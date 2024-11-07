package bootstrap

import (
	"gae-backend-analysis/constant/mq"
	kq "gae-backend-analysis/infrastructure/kafka"
	"github.com/zeromicro/go-zero/core/service"
)

var env *Env

func NewKafkaConf(e *Env) *KafkaConf {
	conf := *initKafkaConf(e)
	return &KafkaConf{ConfMap: conf}
}

type KafkaConf struct {
	ConfMap map[int]kq.KqConf
}

func initKafkaConf(e *Env) *map[int]kq.KqConf {
	env = e
	confMap := new(map[int]kq.KqConf)
	m := *confMap
	m = make(map[int]kq.KqConf)

	// 为 UnCleansingRepo 配置
	UnCleansingRepoGroup := newConsumerConf(mq.UnCleansingRepoGroup, mq.UnCleansingRepoTopic, mq.UnCleansingRepoServiceName)
	m[mq.UnCleansingRepoId] = *UnCleansingRepoGroup

	UnCleansingUserGroup := newConsumerConf(mq.UnCleansingUserGroup, mq.UnCleansingUserTopic, mq.UnCleansingUserServiceName)
	// 为 UnCleansingUser 配置

	m[mq.UnCleansingUserId] = *UnCleansingUserGroup

	return &m
}

func newConsumerConf(group string, topic string, opts ...any) *kq.KqConf {
	return &kq.KqConf{
		ServiceConf: service.ServiceConf{
			Name: opts[0].(string),
		},
		Brokers:    []string{env.KafkaBroker},
		Group:      group,
		Topic:      topic,
		Offset:     mq.FirstOffset,
		Conns:      env.KafkaConns,
		Consumers:  env.KafkaConsumers,
		Processors: env.KafkaProcessors,
	}
}
