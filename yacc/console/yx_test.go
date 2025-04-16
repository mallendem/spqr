package spqrparser_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	spqrparser "github.com/pg-sharding/spqr/yacc/console"
)

func TestSimpleTrace(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	/*  */
	for _, tt := range []tcase{
		{
			query: "START TRACE ALL MESSAGES",
			exp: &spqrparser.TraceStmt{
				All: true,
			},
			err: nil,
		},

		{
			query: "START TRACE CLIENT 129191;",
			exp: &spqrparser.TraceStmt{
				Client: 129191,
			},
			err: nil,
		},

		{
			query: "STOP TRACE MESSAGES",
			exp:   &spqrparser.StopTraceStmt{},
			err:   nil,
		},
	} {
		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestSimpleShow(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	/* POOLS STATS LISTS SERVERS CLIENTS DATABASES BACKEND_CONNECTIONS */
	for _, tt := range []tcase{
		{
			query: "SHOW version",
			exp: &spqrparser.Show{
				Cmd:     spqrparser.VersionStr,
				Where:   spqrparser.WhereClauseEmpty{},
				GroupBy: spqrparser.GroupByClauseEmpty{},
			},
			err: nil,
		},
		/* case insensitive */
		{
			query: "ShOw versIon",
			exp: &spqrparser.Show{
				Cmd:     spqrparser.VersionStr,
				Where:   spqrparser.WhereClauseEmpty{},
				GroupBy: spqrparser.GroupByClauseEmpty{},
			},
			err: nil,
		},

		{
			query: "ShOw pools",
			exp: &spqrparser.Show{
				Cmd:     spqrparser.PoolsStr,
				Where:   spqrparser.WhereClauseEmpty{},
				GroupBy: spqrparser.GroupByClauseEmpty{},
			},
			err: nil,
		},

		{
			query: "ShOw instance",
			exp: &spqrparser.Show{
				Cmd:     spqrparser.InstanceStr,
				Where:   spqrparser.WhereClauseEmpty{},
				GroupBy: spqrparser.GroupByClauseEmpty{},
			},
			err: nil,
		},
		{
			query: "ShOw clients",
			exp: &spqrparser.Show{
				Cmd:     spqrparser.ClientsStr,
				Where:   spqrparser.WhereClauseEmpty{},
				GroupBy: spqrparser.GroupByClauseEmpty{},
			},
			err: nil,
		},
		{
			query: "ShOw DATABASES",
			exp: &spqrparser.Show{
				Cmd:     spqrparser.DatabasesStr,
				Where:   spqrparser.WhereClauseEmpty{},
				GroupBy: spqrparser.GroupByClauseEmpty{},
			},
			err: nil,
		},
		{
			query: "ShOw BACKEND_CONNECTIONS",
			exp: &spqrparser.Show{
				Cmd:     spqrparser.BackendConnectionsStr,
				Where:   spqrparser.WhereClauseEmpty{},
				GroupBy: spqrparser.GroupByClauseEmpty{},
			},
			err: nil,
		},
		{
			query: "kill client 824636929312;",
			exp: &spqrparser.Kill{
				Cmd:    spqrparser.ClientStr,
				Target: 824636929312,
			},
			err: nil,
		},
	} {
		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestSimpleWhere(t *testing.T) {

	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SHOW clients where user = 'usr1';",
			exp: &spqrparser.Show{
				Cmd: spqrparser.ClientsStr,
				Where: spqrparser.WhereClauseLeaf{
					Op:     "=",
					ColRef: spqrparser.ColumnRef{ColName: "user"},
					Value:  "usr1",
				},
				GroupBy: spqrparser.GroupByClauseEmpty{},
			},
			err: nil,
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestNestedWhere(t *testing.T) {

	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SHOW clients where user = 'usr1' or dbname = 'db1';",
			exp: &spqrparser.Show{
				Cmd: spqrparser.ClientsStr,
				Where: spqrparser.WhereClauseOp{
					Op: "OR",
					Left: spqrparser.WhereClauseLeaf{
						Op:     "=",
						ColRef: spqrparser.ColumnRef{ColName: "user"},
						Value:  "usr1",
					},
					Right: spqrparser.WhereClauseLeaf{
						Op:     "=",
						ColRef: spqrparser.ColumnRef{ColName: "dbname"},
						Value:  "db1",
					},
				},
				GroupBy: spqrparser.GroupByClauseEmpty{},
			},
			err: nil,
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestGroupBy(t *testing.T) {

	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SHOW backend_connections GROUP BY hostname;",
			exp: &spqrparser.Show{
				Cmd:     spqrparser.BackendConnectionsStr,
				Where:   spqrparser.WhereClauseEmpty{},
				GroupBy: spqrparser.GroupBy{Col: spqrparser.ColumnRef{ColName: "hostname"}},
			},
			err: nil,
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestRedistribute(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "REDISTRIBUTE KEY RANGE kr1 TO sh2 BATCH SIZE 500",
			exp: &spqrparser.RedistributeKeyRange{
				KeyRangeID:  "kr1",
				DestShardID: "sh2",
				BatchSize:   500,
				Check:       true,
				Apply:       true,
			},
			err: nil,
		},
		{
			query: "REDISTRIBUTE KEY RANGE kr1 TO sh2 BATCH SIZE -1",
			exp:   nil,
			err:   fmt.Errorf("syntax error"),
		},
		{
			query: "REDISTRIBUTE KEY RANGE kr1 TO sh2",
			exp: &spqrparser.RedistributeKeyRange{
				KeyRangeID:  "kr1",
				DestShardID: "sh2",
				BatchSize:   -1,
				Check:       true,
				Apply:       true,
			},
			err: nil,
		},
		{
			query: "REDISTRIBUTE KEY RANGE kr1 TO sh2 CHECK",
			exp: &spqrparser.RedistributeKeyRange{
				KeyRangeID:  "kr1",
				DestShardID: "sh2",
				BatchSize:   -1,
				Check:       true,
				Apply:       false,
			},
			err: nil,
		},
		{
			query: "REDISTRIBUTE KEY RANGE kr1 TO sh2 APPLY",
			exp: &spqrparser.RedistributeKeyRange{
				KeyRangeID:  "kr1",
				DestShardID: "sh2",
				BatchSize:   -1,
				Check:       false,
				Apply:       true,
			},
			err: nil,
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		if err != nil {
			assert.EqualError(err, tt.err.Error())
		} else {
			assert.NoError(err, "query %s", tt.query)
		}

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestKeyRange(t *testing.T) {

	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "CREATE KEY RANGE krid1 FROM 1 ROUTE TO sh1 FOR DISTRIBUTION ds1;",
			exp: &spqrparser.Create{
				Element: &spqrparser.KeyRangeDefinition{
					ShardID:      "sh1",
					KeyRangeID:   "krid1",
					Distribution: "ds1",
					LowerBound: &spqrparser.KeyRangeBound{
						Pivots: [][]byte{
							{2, 0, 0, 0, 0, 0, 0, 0},
						},
					},
				},
			},
			err: nil,
		},

		{
			query: "CREATE KEY RANGE krid2 FROM 88888888-8888-8888-8888-888888888889 ROUTE TO sh2 FOR DISTRIBUTION ds1;",
			exp: &spqrparser.Create{
				Element: &spqrparser.KeyRangeDefinition{
					ShardID:      "sh2",
					KeyRangeID:   "krid2",
					Distribution: "ds1",
					LowerBound: &spqrparser.KeyRangeBound{
						Pivots: [][]byte{
							[]byte("88888888-8888-8888-8888-888888888889"),
						},
					},
				},
			},
			err: nil,
		},

		{
			query: `
			CREATE KEY RANGE krid1 FROM 0, 'a' ROUTE TO sh1 FOR DISTRIBUTION ds1;`,

			exp: &spqrparser.Create{
				Element: &spqrparser.KeyRangeDefinition{
					ShardID:      "sh1",
					KeyRangeID:   "krid1",
					Distribution: "ds1",
					LowerBound: &spqrparser.KeyRangeBound{
						Pivots: [][]byte{
							{0, 0, 0, 0, 0, 0, 0, 0},
							[]byte("a"),
						},
					},
				},
			},
			err: nil,
		},

		{
			query: "CREATE KEY RANGE krid1 FROM 1 TO 10 ROUTE TO sh1 FOR DISTRIBUTION ds1;",
			exp:   nil,
			err:   fmt.Errorf("syntax error"),
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		if tt.err == nil {
			assert.NoError(err, "query %s", tt.query)
		} else {
			assert.Error(err, "query %s", tt.query)
		}

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestShardingRule(t *testing.T) {

	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "CREATE SHARDING RULE rule1 COLUMNS id FOR DISTRIBUTION ds1;",
			exp: &spqrparser.Create{
				Element: &spqrparser.ShardingRuleDefinition{
					ID:           "rule1",
					TableName:    "",
					Distribution: "ds1",
					Entries: []spqrparser.ShardingRuleEntry{
						{
							Column: "id",
						},
					},
				},
			},
			err: nil,
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestSplitKeyRange(t *testing.T) {

	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SPLIT KEY RANGE krid3 FROM krid1 BY 5;",
			exp: &spqrparser.SplitKeyRange{
				Border: &spqrparser.KeyRangeBound{
					Pivots: [][]byte{
						{10, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				KeyRangeFromID: "krid1",
				KeyRangeID:     "krid3",
			},
			err: nil,
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestAttachTable(t *testing.T) {

	assert := assert.New(t)

	tmp, err := spqrparser.Parse("ATTACH TABLE t TO DISTRIBUTION ds1;")
	assert.Error(err)
	assert.Equal(nil, tmp, "query %s", "ATTACH TABLE t TO DISTRIBUTION ds1;")
}

func TestAlter(t *testing.T) {

	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "ALTER DISTRIBUTION ds1 ATTACH RELATION t DISTRIBUTION KEY id;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AttachRelation{
						Relations: []*spqrparser.DistributedRelation{
							&spqrparser.DistributedRelation{
								Name: "t",
								DistributionKey: []spqrparser.DistributionKeyEntry{
									{
										Column: "id",
									},
								},
							},
						},
						Distribution: &spqrparser.DistributionSelector{ID: "ds1"},
					},
				},
			},
			err: nil,
		},
		{
			query: "ALTER DISTRIBUTION ds1 ATTACH RELATION t DISTRIBUTION KEY id1, id2;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AttachRelation{
						Relations: []*spqrparser.DistributedRelation{
							&spqrparser.DistributedRelation{
								Name: "t",
								DistributionKey: []spqrparser.DistributionKeyEntry{
									{
										Column: "id1",
									},
									{
										Column: "id2",
									},
								},
							},
						},
						Distribution: &spqrparser.DistributionSelector{ID: "ds1"},
					},
				},
			},
			err: nil,
		},
		{
			query: "ALTER DISTRIBUTION ds1 ATTACH RELATION t DISTRIBUTION KEY id1, id2 HASH FUNCTION murmur;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AttachRelation{
						Relations: []*spqrparser.DistributedRelation{
							&spqrparser.DistributedRelation{
								Name: "t",
								DistributionKey: []spqrparser.DistributionKeyEntry{
									{
										Column: "id1",
									},
									{
										Column:       "id2",
										HashFunction: "murmur",
									},
								},
							},
						},
						Distribution: &spqrparser.DistributionSelector{ID: "ds1"},
					},
				},
			},
			err: nil,
		},

		{
			query: `
		ALTER DISTRIBUTION 
			ds1 
		ATTACH
			RELATION t DISTRIBUTION KEY id1, id2 HASH FUNCTION murmur
			RELATION t2 DISTRIBUTION KEY xd1, xd2 HASH FUNCTION city
			`,
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AttachRelation{
						Relations: []*spqrparser.DistributedRelation{
							{
								Name: "t",
								DistributionKey: []spqrparser.DistributionKeyEntry{
									{
										Column: "id1",
									},
									{
										Column:       "id2",
										HashFunction: "murmur",
									},
								},
							},
							{
								Name: "t2",
								DistributionKey: []spqrparser.DistributionKeyEntry{
									{
										Column: "xd1",
									},
									{
										Column:       "xd2",
										HashFunction: "city",
									},
								},
							},
						},
						Distribution: &spqrparser.DistributionSelector{ID: "ds1"},
					},
				},
			},
			err: nil,
		},

		{
			query: "ALTER DISTRIBUTION ds1 DETACH RELATION t;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.DetachRelation{
						RelationName: "t",
						Distribution: &spqrparser.DistributionSelector{ID: "ds1"},
					},
				},
			},
			err: nil,
		},

		{
			query: "ALTER REPLICATED DISTRIBUTION ATTACH RELATION t;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AttachRelation{
						Distribution: &spqrparser.DistributionSelector{
							ID:         "REPLICATED",
							Replicated: true,
						},
						Relations: []*spqrparser.DistributedRelation{
							{
								Name:               "t",
								ReplicatedRelation: true,
							},
						},
					},
				},
			},
			err: nil,
		},

		{
			query: "ALTER REPLICATED DISTRIBUTION ATTACH RELATION t AUTO INCREMENT id1;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AttachRelation{
						Distribution: &spqrparser.DistributionSelector{
							ID:         "REPLICATED",
							Replicated: true,
						},
						Relations: []*spqrparser.DistributedRelation{
							{
								Name:                 "t",
								ReplicatedRelation:   true,
								AutoIncrementColumns: []string{"id1"},
							},
						},
					},
				},
			},
		},
		{
			query: "ALTER DISTRIBUTION ds1 ATTACH RELATION t DISTRIBUTION KEY id1 AUTO INCREMENT id1, id2;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AttachRelation{
						Relations: []*spqrparser.DistributedRelation{
							{
								Name: "t",
								DistributionKey: []spqrparser.DistributionKeyEntry{
									{
										Column: "id1",
									},
								},
								AutoIncrementColumns: []string{"id1", "id2"},
							},
						},
						Distribution: &spqrparser.DistributionSelector{ID: "ds1"},
					},
				},
			},
		},
		{
			query: "ALTER DISTRIBUTION ds1 ATTACH RELATION t DISTRIBUTION KEY id SCHEMA test;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AttachRelation{
						Relations: []*spqrparser.DistributedRelation{
							{
								Name:       "t",
								SchemaName: "test",
								DistributionKey: []spqrparser.DistributionKeyEntry{
									{
										Column: "id",
									},
								},
							},
						},
						Distribution: &spqrparser.DistributionSelector{ID: "ds1"},
					},
				},
			},
			err: nil,
		},
		{
			query: "ALTER DISTRIBUTION ds1 ALTER RELATION t DISTRIBUTION KEY id;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AlterRelation{
						Relation: &spqrparser.DistributedRelation{
							Name: "t",
							DistributionKey: []spqrparser.DistributionKeyEntry{
								{
									Column: "id",
								},
							},
						},
						Distribution: &spqrparser.DistributionSelector{ID: "ds1"},
					},
				},
			},
			err: nil,
		},
		{
			query: "ALTER DISTRIBUTION ds1 ALTER RELATION t DISTRIBUTION KEY id SCHEMA test;",
			exp: &spqrparser.Alter{
				Element: &spqrparser.AlterDistribution{
					Element: &spqrparser.AlterRelation{
						Relation: &spqrparser.DistributedRelation{
							Name:       "t",
							SchemaName: "test",
							DistributionKey: []spqrparser.DistributionKeyEntry{
								{
									Column: "id",
								},
							},
						},
						Distribution: &spqrparser.DistributionSelector{ID: "ds1"},
					},
				},
			},
			err: nil,
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestDistribution(t *testing.T) {

	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "CREATE DISTRIBUTION db1 COLUMN TYPES integer;",
			exp: &spqrparser.Create{
				Element: &spqrparser.DistributionDefinition{
					ID: "db1",
					ColTypes: []string{
						"integer",
					},
				},
			},
			err: nil,
		},
		{
			query: "CREATE REPLICATED DISTRIBUTION",
			exp: &spqrparser.Create{
				Element: &spqrparser.DistributionDefinition{
					Replicated: true,
				},
			},
			err: nil,
		},
		{
			query: "CREATE REFERENCE TABLE xtab",
			exp: &spqrparser.Create{
				Element: &spqrparser.ReferenceRelationDefinition{
					TableName: "xtab",
				},
			},
			err: nil,
		},
		{
			query: "CREATE REFERENCE TABLE xtab AUTO INCREMENT id",
			exp: &spqrparser.Create{
				Element: &spqrparser.ReferenceRelationDefinition{
					TableName:            "xtab",
					AutoIncrementColumns: []string{"id"},
				},
			},
			err: nil,
		},
		{
			query: "CREATE DISTRIBUTION db1 COLUMN TYPES varchar hash;",
			exp: &spqrparser.Create{
				Element: &spqrparser.DistributionDefinition{
					ID: "db1",
					ColTypes: []string{
						"varchar hashed",
					},
				},
			},
			err: nil,
		},
		{
			query: "CREATE DISTRIBUTION db1 COLUMN TYPES varchar, varchar;",
			exp: &spqrparser.Create{
				Element: &spqrparser.DistributionDefinition{
					ID: "db1",
					ColTypes: []string{
						"varchar",
						"varchar",
					},
				},
			},
			err: nil,
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestShard(t *testing.T) {

	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "CREATE SHARD sh1 WITH HOSTS localhost:6432;",
			exp: &spqrparser.Create{
				Element: &spqrparser.ShardDefinition{
					Id:    "sh1",
					Hosts: []string{"localhost:6432"},
				},
			},
			err: nil,
		},
		{
			query: "CREATE SHARD sh1 WITH HOSTS localhost:6432, other_hosts:6432;",
			exp: &spqrparser.Create{
				Element: &spqrparser.ShardDefinition{
					Id: "sh1",
					Hosts: []string{
						"localhost:6432",
						"other_hosts:6432",
					},
				},
			},
			err: nil,
		},
		{
			query: "DROP SHARD sh1;",
			exp: &spqrparser.Drop{
				Element: &spqrparser.ShardSelector{
					ID: "sh1",
				},
			},
			err: nil,
		},
	} {

		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestRefresh(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   spqrparser.Statement
		err   error
	}

	/*  */
	for _, tt := range []tcase{
		{
			query: "INVALIDATE CACHE",
			exp:   &spqrparser.InvalidateCache{},
			err:   nil,
		},
	} {
		tmp, err := spqrparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}
