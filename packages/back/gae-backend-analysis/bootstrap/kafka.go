package bootstrap

import (
	"gae-backend-analysis/constant/mq"
	kq "gae-backend-analysis/infrastructure/kafka"
	"github.com/zeromicro/go-zero/core/service"
)

func initKafkaConf(*Env) map[int]kq.KqConf {
	m := new(map[int]kq.KqConf)
	confM := *m
	//TODO modified map size
	confM = make(map[int]kq.KqConf, 10)
	conf := kq.KqConf{
		ServiceConf: service.ServiceConf{
			Name: mq.UnRankCleansingService,
		},
		Brokers:    []string{mq.KafkaDefaultLocalBroker},
		Group:      mq.UnRankCleansingGroup,
		Topic:      mq.UnRankCleansingTopic,
		Offset:     mq.FirstOffset,
		Conns:      1,
		Processors: 1,
		Consumers:  1,
	}
	confM[mq.UnRankCleansingServiceId] = conf

	// 为 UnCleansingRepo 配置
	UnCleansingRepoGroup := kq.KqConf{
		ServiceConf: service.ServiceConf{
			Name: "gaeUnCleansingRepoService",
		},
		Brokers:    []string{mq.KafkaDefaultLocalBroker},
		Group:      mq.UnCleansingRepoGroup,
		Topic:      mq.UnCleansingRepoTopic,
		Offset:     mq.FirstOffset,
		Conns:      1,
		Processors: 1,
		Consumers:  1,
	}

	confM[mq.UnCleansingRepoId] = UnCleansingRepoGroup

	// 为 UnCleansingUser 配置
	UnCleansingUserGroup := kq.KqConf{
		ServiceConf: service.ServiceConf{
			Name: "gaeUnCleansingUserService",
		},
		Brokers:    []string{mq.KafkaDefaultLocalBroker},
		Group:      mq.UnCleansingUserGroup,
		Topic:      mq.UnCleansingUserTopic,
		Offset:     mq.FirstOffset,
		Conns:      1,
		Processors: 1,
		Consumers:  1,
	}

	confM[mq.UnCleansingUserId] = UnCleansingUserGroup

	return confM
}

func NewKafkaConf(e *Env) *KafkaConf {
	conf := initKafkaConf(e)
	return &KafkaConf{ConfMap: conf}
}

type KafkaConf struct {
	ConfMap map[int]kq.KqConf
}
