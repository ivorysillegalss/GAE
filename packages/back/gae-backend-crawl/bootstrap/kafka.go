package bootstrap

import (
	"gae-backend-crawl/constant/mq"
	kq "gae-backend-crawl/infrastructure/kafka"
	"github.com/zeromicro/go-zero/core/service"
)

func initKafkaConf(*Env) *kq.KqConf {
	return &kq.KqConf{
		ServiceConf: service.ServiceConf{
			Name: "gaeMessageConsumerService",
		},
		Brokers: []string{mq.KafkaDefaultLocalBroker},
		Group:   mq.UnRankCleansingGroup,
		Topic:   mq.UnRankCleansingTopic,
		Offset:  mq.FirstOffset,
		Conns:   1,
	}
}

func NewKafkaConf(e *Env) *kq.KqConf {
	conf := initKafkaConf(e)
	return conf
}
