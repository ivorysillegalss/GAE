//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"gae-backend-web/api/controller"
	"gae-backend-web/bootstrap"
	"gae-backend-web/consume"
	"gae-backend-web/executor"
	"gae-backend-web/internal/tokenutil"
	"gae-backend-web/repository"
	"gae-backend-web/usecase"
	"github.com/google/wire"
)

var appSet = wire.NewSet(
	bootstrap.NewEnv,
	tokenutil.NewTokenUtil,
	bootstrap.NewDatabases,
	bootstrap.NewRedisDatabase,
	bootstrap.NewMysqlDatabase,
	bootstrap.NewMongoDatabase,
	bootstrap.NewPoolFactory,
	bootstrap.NewChannel,
	bootstrap.NewControllers,
	bootstrap.NewExecutors,
	bootstrap.NewKafkaConf,
	bootstrap.NewEsEngine,
	bootstrap.NewSearchEngine,
	bootstrap.NewHiveDBConnection,

	usecase.NewRankUsecase,
	usecase.NewUsecase,

	repository.NewRankRepository,
	repository.NewUserRepository,

	controller.NewRankController,

	consume.NewTalentEvent,

	executor.NewCronExecutor,
	executor.NewConsumeExecutor,

	wire.Struct(new(bootstrap.Application), "*"),
)

// InitializeApp init application.
func InitializeApp() (*bootstrap.Application, error) {
	wire.Build(appSet)
	return &bootstrap.Application{}, nil
}
