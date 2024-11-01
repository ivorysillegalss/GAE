package bootstrap

import (
	"context"
	"gae-backend-crawl/executor"
	"gae-backend-crawl/infrastructure/bloom"
	"gae-backend-crawl/infrastructure/pool"
	"gae-backend-crawl/infrastructure/redis"
	"log"
	"time"
)

type Application struct {
	Env           *Env
	Databases     *Databases
	PoolsFactory  *PoolsFactory
	Bloom         bloom.Client
	CrawlExecutor *executor.CrawlExecutor
	KafkaConf     *KafkaConf
}

type Databases struct {
	Redis redis.Client
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

func NewBloomFilter() bloom.Client {
	return bloom.NewBloomClient()
}

// PoolsFactory k为pool业务号 v为poll详细配置信息
type PoolsFactory struct {
	Pools map[int]*pool.Pool
}
