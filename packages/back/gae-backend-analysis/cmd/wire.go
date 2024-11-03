//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"gae-backend-analysis/api/controller"
	"gae-backend-analysis/bootstrap"
	"gae-backend-analysis/consume"
	"gae-backend-analysis/cron"
	"gae-backend-analysis/executor"
	"gae-backend-analysis/internal/tokenutil"
	"gae-backend-analysis/repository"
	"gae-backend-analysis/task"
	"gae-backend-analysis/usecase"
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
	bootstrap.NewRabbitConnection,
	bootstrap.NewControllers,
	bootstrap.NewExecutors,
	bootstrap.NewKafkaConf,
	bootstrap.NewEsEngine,
	bootstrap.NewSearchEngine,

	repository.NewGenerationRepository,
	repository.NewChatRepository,
	repository.NewBotRepository,
	repository.NewTalentRepository,

	consume.NewTalentEvent,
	consume.NewStorageEvent,
	consume.NewGenerateEvent,
	consume.NewMessageHandler,

	cron.NewGenerationCron,
	cron.NewTalentCron,

	executor.NewCronExecutor,
	executor.NewConsumeExecutor,
	executor.NewDataExecutor,

	usecase.NewChatUseCase,

	task.NewAskChatTask,
	task.NewChatTitleTask,
	task.NewConvertTask,

	controller.NewTestController,

	wire.Struct(new(bootstrap.Application), "*"),
)

// InitializeApp init application.
func InitializeApp() (*bootstrap.Application, error) {
	wire.Build(appSet)
	return &bootstrap.Application{}, nil
}
