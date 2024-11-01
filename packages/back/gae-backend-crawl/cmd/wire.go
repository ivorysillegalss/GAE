//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"gae-backend-crawl/bootstrap"
	"gae-backend-crawl/crawl"
	"gae-backend-crawl/executor"
	"github.com/google/wire"
)

var appSet = wire.NewSet(
	bootstrap.NewEnv,
	bootstrap.NewKafkaConf,
	bootstrap.NewDatabases,
	bootstrap.NewPoolFactory,
	bootstrap.NewBloomFilter,
	crawl.NewRepoCrawl,
	executor.NewCrawlExecutor,

	wire.Struct(new(bootstrap.Application), "*"),
)

// InitializeApp init application.
func InitializeApp() (*bootstrap.Application, error) {
	wire.Build(appSet)
	return &bootstrap.Application{}, nil
}
