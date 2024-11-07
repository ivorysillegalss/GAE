package bootstrap

import (
	kq "gae-backend-web/infrastructure/kafka"
)

func initKafkaConf(*Env) map[int]kq.KqConf {
	m := new(map[int]kq.KqConf)
	confM := *m

	return confM
}

func NewKafkaConf(e *Env) *KafkaConf {
	conf := initKafkaConf(e)
	return &KafkaConf{ConfMap: conf}
}

type KafkaConf struct {
	ConfMap map[int]kq.KqConf
}
