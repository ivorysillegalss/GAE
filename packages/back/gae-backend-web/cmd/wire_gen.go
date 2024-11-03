// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gae-backend-web/bootstrap"
	"gae-backend-web/consume"
	"gae-backend-web/executor"
	"gae-backend-web/internal/tokenutil"
	"github.com/google/wire"
)

// Injectors from wire.go:

// InitializeApp init application.
func InitializeApp() (*bootstrap.Application, error) {
	env := bootstrap.NewEnv()
	databases := bootstrap.NewDatabases(env)
	poolsFactory := bootstrap.NewPoolFactory()
	channels := bootstrap.NewChannel()
	controllers := bootstrap.NewControllers()
	cronExecutor := executor.NewCronExecutor()
	consumeExecutor := executor.NewConsumeExecutor()
	bootstrapExecutor := bootstrap.NewExecutors(cronExecutor, consumeExecutor)
	client := bootstrap.NewEsEngine(env)
	searchEngine := bootstrap.NewSearchEngine(client)
	application := &bootstrap.Application{
		Env:          env,
		Databases:    databases,
		PoolsFactory: poolsFactory,
		Channels:     channels,
		Controllers:  controllers,
		Executor:     bootstrapExecutor,
		SearchEngine: searchEngine,
	}
	return application, nil
}

// wire.go:

var appSet = wire.NewSet(bootstrap.NewEnv, tokenutil.NewTokenUtil, bootstrap.NewDatabases, bootstrap.NewRedisDatabase, bootstrap.NewMysqlDatabase, bootstrap.NewMongoDatabase, bootstrap.NewPoolFactory, bootstrap.NewChannel, bootstrap.NewControllers, bootstrap.NewExecutors, bootstrap.NewKafkaConf, bootstrap.NewEsEngine, bootstrap.NewSearchEngine, consume.NewTalentEvent, executor.NewCronExecutor, executor.NewConsumeExecutor, wire.Struct(new(bootstrap.Application), "*"))
