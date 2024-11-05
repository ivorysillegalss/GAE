package bootstrap

import (
	"context"
	"gae-backend-analysis/executor"
	"gae-backend-analysis/infrastructure/pool"
	"gae-backend-analysis/infrastructure/redis"
	"log"
	"time"
)

type Application struct {
	Env          *Env
	Databases    *Databases
	PoolsFactory *PoolsFactory
	Executor     *Executor
}

type Databases struct {
	Redis redis.Client
}

// PoolsFactory k为pool业务号 v为poll详细配置信息
type PoolsFactory struct {
	Pools map[int]*pool.Pool
}

type Executor struct {
	CronExecutor    *executor.CronExecutor
	ConsumeExecutor *executor.ConsumeExecutor
	DataExecutor    *executor.DataExecutor
}

func NewDatabases(env *Env) *Databases {
	return &Databases{
		Redis: NewRedisDatabase(env),
	}
}

func NewRedisDatabase(env *Env) redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbAddr := env.RedisAddr
	dbPassword := env.RedisPassword

	client, err := redis.NewRedisClient(redis.NewRedisApplication(dbAddr, dbPassword))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func NewExecutors(ce *executor.CronExecutor, cse *executor.ConsumeExecutor, de *executor.DataExecutor) *Executor {
	return &Executor{
		CronExecutor:    ce,
		ConsumeExecutor: cse,
		DataExecutor:    de,
	}
}
