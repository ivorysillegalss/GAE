package bootstrap

import (
	"gae-backend-analysis/constant/mq"
	kq "gae-backend-analysis/infrastructure/kafka"
	"github.com/zeromicro/go-zero/core/service"
)

func initKafkaConf(*Env) map[int]kq.KqConf {
	m := new(map[int]kq.KqConf)
	confM := *m
	conf := kq.KqConf{
		ServiceConf: service.ServiceConf{
			Name: mq.UnRankCleansingService,
		},
		Brokers: []string{mq.KafkaDefaultLocalBroker},
		Group:   mq.UnRankCleansingGroup,
		Topic:   mq.UnRankCleansingTopic,
		Offset:  mq.FirstOffset,
		Conns:   1,
	}
	confM[mq.UnRankCleansingServiceId] = conf
	return confM
}

func NewKafkaConf(e *Env) *KafkaConf {
	conf := initKafkaConf(e)
	return &KafkaConf{confMap: conf}
}

type KafkaConf struct {
	confMap map[int]kq.KqConf
}
