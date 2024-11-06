package bootstrap

import (
	"context"
	"fmt"
	"gae-backend-storage/infrastructure/mysql"
	"gae-backend-storage/infrastructure/redis"
	"log"
	"time"
)

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

func NewMysqlDatabase(env *Env) mysql.Client {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.MysqlUser, env.MysqlPassword, env.MysqlHost, env.MysqlPort, env.MysqlDB)

	client, err := mysql.NewMysqlClient(dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func NewDatabases(env *Env) *Databases {
	return &Databases{
		Redis: NewRedisDatabase(env),
		Mysql: NewMysqlDatabase(env),
	}
}
