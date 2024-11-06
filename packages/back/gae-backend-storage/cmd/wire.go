//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"gae-backend-storage/bootstrap"
	"gae-backend-storage/executor"
	"gae-backend-storage/handler"
	"gae-backend-storage/repository"
	"github.com/google/wire"
)

var appSet = wire.NewSet(
	bootstrap.NewEnv,
	bootstrap.NewDatabases,
	bootstrap.NewRedisDatabase,
	bootstrap.NewMysqlDatabase,
	bootstrap.NewPoolFactory,
	bootstrap.NewChannel,

	bootstrap.NewRpcEngine,

	bootstrap.NewEsEngine,
	bootstrap.NewSearchEngine,
	executor.NewGrpcExecutor,
	bootstrap.NewExecutors,

	handler.NewStorageEngine,
	bootstrap.NewStorageEngine,

	repository.NewRankRepository,
	repository.NewUserRepository,

	wire.Struct(new(bootstrap.Application), "*"),
)

// InitializeApp init application.
func InitializeApp() (*bootstrap.Application, error) {
	wire.Build(appSet)
	return &bootstrap.Application{}, nil
}
