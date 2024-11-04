package bootstrap

import (
	"gae-backend-web/api/controller"
	"gae-backend-web/executor"
	"gae-backend-web/infrastructure/elasticsearch"
	"gae-backend-web/infrastructure/mongo"
	"gae-backend-web/infrastructure/mysql"
	"gae-backend-web/infrastructure/pool"
	"gae-backend-web/infrastructure/redis"
)

type Application struct {
	Env          *Env
	Databases    *Databases
	PoolsFactory *PoolsFactory
	Channels     *Channels
	Controllers  *Controllers
	Executor     *Executor
	SearchEngine *SearchEngine
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
	Stop chan bool
}

type Controllers struct {
	RankController *controller.RankController
}

type Executor struct {
	CronExecutor    *executor.CronExecutor
	ConsumeExecutor *executor.ConsumeExecutor
}

type SearchEngine struct {
	EsEngine elasticsearch.Client
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Databases.Mongo)
}

func NewControllers(rc *controller.RankController) *Controllers {
	return &Controllers{RankController: rc}
}

func NewExecutors(ce *executor.CronExecutor, cse *executor.ConsumeExecutor) *Executor {
	return &Executor{
		CronExecutor:    ce,
		ConsumeExecutor: cse,
	}
}

func NewSearchEngine(es elasticsearch.Client) *SearchEngine {
	return &SearchEngine{EsEngine: es}
}
