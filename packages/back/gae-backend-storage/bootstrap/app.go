package bootstrap

import (
	"gae-backend-storage/executor"
	"gae-backend-storage/handler"
	"gae-backend-storage/infrastructure/elasticsearch"
	"gae-backend-storage/infrastructure/grpc"
	"gae-backend-storage/infrastructure/mysql"
	"gae-backend-storage/infrastructure/pool"
	"gae-backend-storage/infrastructure/redis"
)

type Application struct {
	Env           *Env
	Databases     *Databases
	PoolsFactory  *PoolsFactory
	Channels      *Channels
	Executor      *Executor
	SearchEngine  *SearchEngine
	RpcEngine     *RpcEngine
	StorageEngine *StorageEngine
}

type Databases struct {
	Redis redis.Client
	Mysql mysql.Client
}

// PoolsFactory k为pool业务号 v为poll详细配置信息
type PoolsFactory struct {
	Pools map[int]*pool.Pool
}

type Channels struct {
	Stop chan bool
}

type Executor struct {
	GrpcExecutor *executor.GrpcExecutor
}

type SearchEngine struct {
	EsEngine elasticsearch.Client
}

type StorageEngine struct {
	ParquetStorage handler.ParquetStorageEngine
}

type RpcEngine struct {
	GrpcClient grpc.Client
	//GrpcServer grpc.Server
}

type Engine struct {
}

func NewExecutors(ge *executor.GrpcExecutor) *Executor {
	return &Executor{
		GrpcExecutor: ge,
	}
}

func NewSearchEngine(es elasticsearch.Client) *SearchEngine {
	return &SearchEngine{EsEngine: es}
}

func NewRpcEngine(env *Env) *RpcEngine {
	client := NewGrpc(env)
	//server := grpc.NewGrpcServer(env.GrpcServerNetwork, env.GrpcServerAddress)
	//return &RpcEngine{GrpcClient: client, GrpcServer: server}
	return &RpcEngine{GrpcClient: client}
}

func NewStorageEngine(engine handler.ParquetStorageEngine) *StorageEngine {
	return &StorageEngine{ParquetStorage: engine}
}
