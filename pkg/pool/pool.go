package pool

import (
	"time"

	"github.com/pg-sharding/spqr/pkg/config"
	"github.com/pg-sharding/spqr/pkg/models/kr"
	"github.com/pg-sharding/spqr/pkg/shard"
	"github.com/pg-sharding/spqr/pkg/tsa"
)

const (
	defaultInstanceConnectionLimit   = 50
	defaultInstanceConnectionRetries = 10
	defaultInstanceConnectionTimeout = time.Second
	defaultKeepAlive                 = time.Second
	defaultTcpUserTimeout            = time.Millisecond * 9500
)

type ConnectionKeeper interface {
	Put(host shard.ShardHostInstance) error
	Discard(sh shard.ShardHostInstance) error
	View() Statistics
}

type Statistics struct {
	DB                string
	Usr               string
	Hostname          string
	RouterName        string
	UsedConnections   int
	IdleConnections   int
	QueueResidualSize int
}

/* dedicated host connection pool */
type Pool interface {
	ConnectionKeeper
	shard.ShardHostIterator

	Connection(clid uint, shardKey kr.ShardKey) (shard.ShardHostInstance, error)
}

/* Host  */
type ShardHostsPool interface {
	ConnectionKeeper
	shard.ShardHostIterator
	PoolIterator

	ID() uint

	ConnectionHost(clid uint, shardKey kr.ShardKey, host config.Host) (shard.ShardHostInstance, error)

	SetRule(rule *config.BackendRule)
}

type MultiShardTSAPool interface {
	ShardHostsPool

	ShardMapping() map[string]*config.Shard
	ConnectionWithTSA(clid uint, key kr.ShardKey, targetSessionAttrs tsa.TSA) (shard.ShardHostInstance, error)
	InstanceHealthChecks() map[string]tsa.CachedCheckResult

	StopCacheWatchdog()
}

type PoolIterator interface {
	ForEachPool(cb func(p Pool) error) error
}

type ConnectionAllocFn func(shardKey kr.ShardKey, host config.Host, rule *config.BackendRule) (shard.ShardHostInstance, error)
