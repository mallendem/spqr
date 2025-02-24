package rmeta

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pg-sharding/spqr/pkg/meta"
	"github.com/pg-sharding/spqr/pkg/models/distributions"
	"github.com/pg-sharding/spqr/pkg/models/kr"
	"github.com/pg-sharding/spqr/pkg/models/spqrerror"
	"github.com/pg-sharding/spqr/pkg/session"
	"github.com/pg-sharding/spqr/pkg/spqrlog"
	"github.com/pg-sharding/spqr/qdb"
	"github.com/pg-sharding/spqr/router/plan"
	"github.com/pg-sharding/spqr/router/rerrors"
	"github.com/pg-sharding/spqr/router/rfqn"
	"github.com/pg-sharding/spqr/router/routehint"

	"github.com/pg-sharding/lyx/lyx"
)

type RoutingMetadataContext struct {
	// this maps table names to its query-defined restrictions
	// All columns in query should be considered in context of its table,
	// to distinguish composite join/select queries routing schemas
	//
	// For example,
	// SELECT * FROM a join b WHERE a.c1 = <val> and b.c2 = <val>
	// and
	// SELECT * FROM a join b WHERE a.c1 = <val> and a.c2 = <val>
	// can be routed with different rules
	Rels      map[rfqn.RelationFQN]struct{}
	Exprs     map[rfqn.RelationFQN]map[string][]interface{}
	ParamRefs map[rfqn.RelationFQN]map[string][]int

	// cached CTE names
	CteNames map[string]struct{}

	// needed to parse
	// SELECT * FROM t1 a where a.i = 1
	// rarg:{range_var:{relname:"t2" inh:true relpersistence:"p" alias:{aliasname:"b"}
	TableAliases map[string]rfqn.RelationFQN

	SPH session.SessionParamsHolder

	Mgr meta.EntityMgr

	Distributions map[rfqn.RelationFQN]*distributions.Distribution
}

func NewRoutingMetadataContext(sph session.SessionParamsHolder, mgr meta.EntityMgr) *RoutingMetadataContext {
	return &RoutingMetadataContext{
		Rels:          map[rfqn.RelationFQN]struct{}{},
		CteNames:      map[string]struct{}{},
		TableAliases:  map[string]rfqn.RelationFQN{},
		Exprs:         map[rfqn.RelationFQN]map[string][]interface{}{},
		ParamRefs:     map[rfqn.RelationFQN]map[string][]int{},
		Distributions: map[rfqn.RelationFQN]*distributions.Distribution{},
		SPH:           sph,
		Mgr:           mgr,
	}
}

var CatalogDistribution = distributions.Distribution{
	Relations: nil,
	Id:        distributions.REPLICATED,
	ColTypes:  nil,
}

func (rm *RoutingMetadataContext) GetRelationDistribution(ctx context.Context, resolvedRelation rfqn.RelationFQN) (*distributions.Distribution, error) {
	if res, ok := rm.Distributions[resolvedRelation]; ok {
		return res, nil
	}

	if len(resolvedRelation.RelationName) >= 3 && resolvedRelation.RelationName[0:3] == "pg_" {
		return &CatalogDistribution, nil
	}

	if resolvedRelation.SchemaName == "information_schema" {
		return &CatalogDistribution, nil
	}

	ds, err := rm.Mgr.GetRelationDistribution(ctx, resolvedRelation.RelationName)

	if err != nil {
		return nil, err
	}

	rm.Distributions[resolvedRelation] = ds
	return ds, nil
}

func (rm *RoutingMetadataContext) RFQNIsCTE(resolvedRelation rfqn.RelationFQN) bool {
	_, ok := rm.CteNames[resolvedRelation.RelationName]
	return len(resolvedRelation.SchemaName) == 0 && ok
}

// TODO : unit tests
func (rm *RoutingMetadataContext) RecordConstExpr(resolvedRelation rfqn.RelationFQN, colname string, expr interface{}) error {
	if rm.RFQNIsCTE(resolvedRelation) {
		// CTE, skip
		return nil
	}
	rm.Rels[resolvedRelation] = struct{}{}
	if _, ok := rm.Exprs[resolvedRelation]; !ok {
		rm.Exprs[resolvedRelation] = map[string][]interface{}{}
	}
	if _, ok := rm.Exprs[resolvedRelation][colname]; !ok {
		rm.Exprs[resolvedRelation][colname] = make([]interface{}, 0)
	}
	rm.Exprs[resolvedRelation][colname] = append(rm.Exprs[resolvedRelation][colname], expr)
	return nil
}

func (routingMeta *RoutingMetadataContext) RecordParamRefExpr(resolvedRelation rfqn.RelationFQN, colname string, ind int) error {
	if routingMeta.RFQNIsCTE(resolvedRelation) {
		// CTE, skip
		return nil
	}

	if routingMeta.Distributions[resolvedRelation].Id == distributions.REPLICATED {
		// reference relation, skip
		return nil
	}

	routingMeta.Rels[resolvedRelation] = struct{}{}
	if _, ok := routingMeta.ParamRefs[resolvedRelation]; !ok {
		routingMeta.ParamRefs[resolvedRelation] = map[string][]int{}
	}
	if _, ok := routingMeta.ParamRefs[resolvedRelation][colname]; !ok {
		routingMeta.ParamRefs[resolvedRelation][colname] = make([]int, 0)
	}
	routingMeta.ParamRefs[resolvedRelation][colname] = append(routingMeta.ParamRefs[resolvedRelation][colname], ind)
	return nil
}

// TODO : unit tests
func (rm *RoutingMetadataContext) ResolveRelationByAlias(alias string) (rfqn.RelationFQN, error) {
	if _, ok := rm.Rels[rfqn.RelationFQN{RelationName: alias}]; ok {
		return rfqn.RelationFQN{RelationName: alias}, nil
	}
	if resolvedRelation, ok := rm.TableAliases[alias]; ok {
		// TBD: postpone routing from here to root of parsing tree
		return resolvedRelation, nil
	} else {
		// TBD: postpone routing from here to root of parsing tree
		if len(rm.Rels) != 1 {
			// ambiguity in column aliasing
			return rfqn.RelationFQN{}, rerrors.ErrComplexQuery
		}
		for tbl := range rm.Rels {
			resolvedRelation = tbl
		}
		return resolvedRelation, nil
	}
}

// TODO : unit tests
func (rm *RoutingMetadataContext) DeparseKeyWithRangesInternal(_ context.Context, key []interface{}, krs []*kr.KeyRange) (*kr.ShardKey, error) {
	spqrlog.Zero.Debug().
		Interface("key", key[0]).
		Int("key-ranges-count", len(krs)).
		Msg("checking key with key ranges")

	var matchedKrkey *kr.KeyRange = nil

	for _, krkey := range krs {
		if kr.CmpRangesLessEqual(krkey.LowerBound, key, krkey.ColumnTypes) &&
			(matchedKrkey == nil || kr.CmpRangesLessEqual(matchedKrkey.LowerBound, krkey.LowerBound, krkey.ColumnTypes)) {
			matchedKrkey = krkey
		}
	}

	if matchedKrkey != nil {
		if err := rm.Mgr.ShareKeyRange(matchedKrkey.ID); err != nil {
			return nil, err
		}
		return &kr.ShardKey{Name: matchedKrkey.ShardID}, nil
	}
	spqrlog.Zero.Debug().Msg("failed to match key with ranges")

	return nil, fmt.Errorf("failed to match key with ranges")
}

func (rm *RoutingMetadataContext) ResolveRouteHint() (routehint.RouteHint, error) {
	if rm.SPH.ScatterQuery() {
		return &routehint.ScatterRouteHint{}, nil
	}
	if val := rm.SPH.ShardingKey(); val != "" {
		spqrlog.Zero.Debug().Str("sharding key", val).Msg("checking hint key")

		dsId := rm.SPH.Distribution()
		if dsId == "" {
			return nil, spqrerror.New(spqrerror.SPQR_NO_DISTRIBUTION, "sharding key in comment without distribution")
		}

		ctx := context.TODO()
		krs, err := rm.Mgr.ListKeyRanges(ctx, dsId)
		if err != nil {
			return nil, err
		}

		distrib, err := rm.Mgr.GetDistribution(ctx, dsId)
		if err != nil {
			return nil, err
		}

		// TODO: fix this
		compositeKey, err := kr.KeyRangeBoundFromStrings(distrib.ColTypes, []string{val})

		if err != nil {
			return nil, err
		}

		ds, err := rm.DeparseKeyWithRangesInternal(ctx, compositeKey, krs)
		if err != nil {
			return nil, err
		}
		return &routehint.TargetRouteHint{
			State: plan.ShardDispatchPlan{
				ExecTarget: ds,
			},
		}, nil
	}

	return &routehint.EmptyRouteHint{}, nil
}

func (rm *RoutingMetadataContext) GetDistributionKeyOffsetType(resolvedRelation rfqn.RelationFQN, colname string) (int, string) {
	/* do not process non-distributed relations or columns not from relation distribution key */

	ds, err := rm.GetRelationDistribution(context.TODO(), resolvedRelation)
	if err != nil {
		return -1, ""
	} else if ds.Id == distributions.REPLICATED {
		return -1, ""
	}
	// TODO: optimize
	relation, exists := ds.Relations[resolvedRelation.RelationName]
	if !exists {
		return -1, ""
	}
	for ind, c := range relation.DistributionKey {
		if c.Column == colname {
			return ind, ds.ColTypes[ind]
		}
	}
	return -1, ""
}

func (rm *RoutingMetadataContext) ProcessSingleExpr(resolvedRelation rfqn.RelationFQN, tp string, colname string, expr lyx.Node) error {
	switch right := expr.(type) {
	case *lyx.ParamRef:
		return rm.RecordParamRefExpr(resolvedRelation, colname, right.Number-1)
	case *lyx.AExprSConst:
		switch tp {
		case qdb.ColumnTypeVarcharDeprecated:
			fallthrough
		case qdb.ColumnTypeVarcharHashed:
			fallthrough
		case qdb.ColumnTypeVarchar:
			return rm.RecordConstExpr(resolvedRelation, colname, right.Value)
		case qdb.ColumnTypeInteger:
			num, err := strconv.ParseInt(right.Value, 10, 64)
			if err != nil {
				return err
			}
			return rm.RecordConstExpr(resolvedRelation, colname, num)
		case qdb.ColumnTypeUinteger:
			num, err := strconv.ParseUint(right.Value, 10, 64)
			if err != nil {
				return err
			}
			return rm.RecordConstExpr(resolvedRelation, colname, num)
		default:
			return fmt.Errorf("incorrect key-offset type for AExprSConst expression: %s", tp)
		}
	case *lyx.AExprIConst:
		switch tp {
		case qdb.ColumnTypeVarcharDeprecated:
			fallthrough
		case qdb.ColumnTypeVarcharHashed:
			fallthrough
		case qdb.ColumnTypeVarchar:
			return fmt.Errorf("varchar type is not supported for AExprIConst expression")
		case qdb.ColumnTypeInteger:
			return rm.RecordConstExpr(resolvedRelation, colname, int64(right.Value))
		case qdb.ColumnTypeUinteger:
			return rm.RecordConstExpr(resolvedRelation, colname, uint64(right.Value))
		default:
			return fmt.Errorf("incorrect key-offset type for AExprIConst expression: %s", tp)
		}
	default:
		return fmt.Errorf("expression is not const")
	}
}
