package consume

import (
	"context"
	"fmt"
	"gae-backend-analysis/bootstrap"
	"gae-backend-analysis/constant/mq"
	"gae-backend-analysis/domain"
	kq "gae-backend-analysis/infrastructure/kafka"
	"gae-backend-analysis/infrastructure/log"
)

type talentEvent struct {
	env              *bootstrap.Env
	kafka            *bootstrap.KafkaConf
	talentRepository domain.TalentRepository
}

func NewTalentEvent(env *bootstrap.Env, conf *bootstrap.KafkaConf, repository domain.TalentRepository) domain.TalentEvent {
	return talentEvent{env: env, kafka: conf, talentRepository: repository}
}

type talentRankHandler struct {
	talentRepository domain.TalentRepository
	businessId       int
}

// Consume 消费具体逻辑
//
//	此处将消息按照id存入list当中 TODO 可拓展为多消费者
func (t *talentRankHandler) Consume(ctx context.Context, key, value string) error {
	log.GetTextLogger().Info("success receive :" + value)
	return t.talentRepository.CleansingDataTemporaryStorageCache(key, value, t.businessId)
}

func (t talentEvent) ConsumeRepo() {
	ta := &talentRankHandler{talentRepository: t.talentRepository, businessId: mq.UnCleansingRepoId}
	fmt.Println(ta)

	conf := *(t.kafka)
	kqConf := conf.ConfMap[mq.UnCleansingRepoId]

	handler := kq.WithHandle(func(ctx context.Context, key, value string) error {
		err := t.talentRepository.CleansingDataTemporaryStorageCache(key, value, mq.UnCleansingRepoId)
		return err
	})

	//queue, err := kq.NewQueue(kqConf, ta)
	queue, err := kq.NewQueue(kqConf, handler)
	if err != nil {
		log.GetTextLogger().Error(fmt.Sprintf("queue get error :" + err.Error()))
	}
	go queue.Start()
}

func (t talentEvent) ConsumeContributors() {
	ta := &talentRankHandler{talentRepository: t.talentRepository, businessId: mq.UnCleansingUserId}
	fmt.Println(ta)

	conf := *(t.kafka)
	kqConf := conf.ConfMap[mq.UnCleansingUserId]

	handler := kq.WithHandle(func(ctx context.Context, key, value string) error {
		err := t.talentRepository.CleansingDataTemporaryStorageCache(key, value, mq.UnCleansingUserId)
		return err
	})

	//queue, err := kq.NewQueue(kqConf, ta)
	queue, err := kq.NewQueue(kqConf, handler)
	if err != nil {
		log.GetTextLogger().Error(fmt.Sprintf("queue get error :" + err.Error()))
	}
	go queue.Start()
}
