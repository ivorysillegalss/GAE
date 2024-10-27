package bootstrap

import (
	"gae-backend-analysis/api/controller"
	"gae-backend-analysis/domain"
	"gae-backend-analysis/executor"
	"gae-backend-analysis/infrastructure/mongo"
	"gae-backend-analysis/infrastructure/mysql"
	"gae-backend-analysis/infrastructure/pool"
	"gae-backend-analysis/infrastructure/redis"
)

type Application struct {
	Env          *Env
	Databases    *Databases
	PoolsFactory *PoolsFactory
	Channels     *Channels
	Controllers  *Controllers
	Executor     *Executor
}

type Databases struct {
	Mongo mongo.Client
	Redis redis.Client
	Mysql mysql.Client
}

// PoolsFactory k为pool业务号 v为poll详细配置信息
type PoolsFactory struct {
	Pools map[int]*pool.Pool
}

type Channels struct {
	RpcRes       chan *domain.GenerationResponse
	Stop         chan bool
	StreamRpcRes chan *domain.GenerationResponse
}

type Controllers struct {
	TestController *controller.TestController
}

type Executor struct {
	CronExecutor    *executor.CronExecutor
	ConsumeExecutor *executor.ConsumeExecutor
	DataExecutor    *executor.DataExecutor
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Databases.Mongo)
}

func NewControllers() *Controllers {
	return &Controllers{}
}

func NewExecutors(ce *executor.CronExecutor, cse *executor.ConsumeExecutor, de *executor.DataExecutor) *Executor {
	return &Executor{
		CronExecutor:    ce,
		ConsumeExecutor: cse,
		DataExecutor:    de,
	}
}
