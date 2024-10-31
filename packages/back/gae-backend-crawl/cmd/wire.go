//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"gae-backend-crawl/bootstrap"
	"github.com/google/wire"
)

var appSet = wire.NewSet(
	bootstrap.NewEnv,
	bootstrap.NewDatabases,
	bootstrap.NewPoolFactory,
	bootstrap.NewKafkaConf,
	bootstrap.NewBloomFilter,

	wire.Struct(new(bootstrap.Application), "*"),
)

// InitializeApp init application.
func InitializeApp() (*bootstrap.Application, error) {
	wire.Build(appSet)
	return &bootstrap.Application{}, nil
}
