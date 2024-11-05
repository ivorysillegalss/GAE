//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"gae-backend-analysis/bootstrap"
	"gae-backend-analysis/consume"
	"gae-backend-analysis/cron"
	"gae-backend-analysis/executor"
	"gae-backend-analysis/repository"
	"github.com/google/wire"
)

var appSet = wire.NewSet(
	bootstrap.NewEnv,

	bootstrap.NewPoolFactory,

	bootstrap.NewExecutors,
	bootstrap.NewKafkaConf,
	bootstrap.NewDatabases,
	bootstrap.NewRedisDatabase,

	repository.NewTalentRepository,

	consume.NewTalentEvent,

	cron.NewTalentCron,

	executor.NewCronExecutor,
	executor.NewConsumeExecutor,
	executor.NewDataExecutor,

	wire.Struct(new(bootstrap.Application), "*"),
)

// InitializeApp init application.
func InitializeApp() (*bootstrap.Application, error) {
	wire.Build(appSet)
	return &bootstrap.Application{}, nil
}
