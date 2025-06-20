package rrelation

import (
	protos "github.com/pg-sharding/spqr/pkg/protos"
	spqrparser "github.com/pg-sharding/spqr/yacc/console"
)

type ReferenceRelation struct {
	TableName             string
	SchemaVersion         uint64
	ColumnSequenceMapping map[string]string
	ShardId               []string
}

type AutoIncrementEntry struct {
	Column string
	Start  uint64
}

func ReferenceRelationEntriesFromDB(inEntries []*spqrparser.AutoIncrementEntry) []*AutoIncrementEntry {
	var ret []*AutoIncrementEntry

	for _, e := range inEntries {
		ret = append(ret, &AutoIncrementEntry{
			Column: e.Column,
			Start:  e.Start,
		})
	}

	return ret
}

func AutoIncrementEntriesToProto(inEntries []*AutoIncrementEntry) []*protos.AutoIncrementEntry {
	var ret []*protos.AutoIncrementEntry

	for _, e := range inEntries {
		ret = append(ret, &protos.AutoIncrementEntry{
			ColName:    e.Column,
			StartValue: e.Start,
		})
	}

	return ret
}

func AutoIncrementEntriesFromProto(inEntries []*protos.AutoIncrementEntry) []*AutoIncrementEntry {
	var ret []*AutoIncrementEntry

	for _, e := range inEntries {
		ret = append(ret, &AutoIncrementEntry{
			Column: e.ColName,
			Start:  e.StartValue,
		})
	}

	return ret
}

func RefRelationFromProto(p *protos.ReferenceRelation) *ReferenceRelation {
	return &ReferenceRelation{
		TableName:             p.Name,
		SchemaVersion:         p.SchemaVersion,
		ColumnSequenceMapping: p.SequenceColumns,
		ShardId:               p.ShardIds,
	}
}

func RefRelationToProto(p *ReferenceRelation) *protos.ReferenceRelation {
	return &protos.ReferenceRelation{
		Name:            p.TableName,
		SchemaVersion:   p.SchemaVersion,
		SequenceColumns: p.ColumnSequenceMapping,
		ShardIds:        p.ShardId,
	}
}
