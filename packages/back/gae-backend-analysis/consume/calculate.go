package consume

import (
	"context"
	"fmt"
	"gae-backend-analysis/bootstrap"
	"gae-backend-analysis/domain"
	kq "gae-backend-analysis/infrastructure/kafka"
	"gae-backend-analysis/infrastructure/log"
)

type talentEvent struct {
	env              *bootstrap.Env
	kqConf           *kq.KqConf
	talentRepository domain.TalentRepository
}

func NewTalentEvent(env *bootstrap.Env, conf *kq.KqConf) domain.TalentEvent {
	return talentEvent{env: env, kqConf: conf}
}

type talentRankHandler struct {
	talentRepository domain.TalentRepository
}

// Consume 消费具体逻辑
//
//	此处将消息按照id存入list当中 TODO 可拓展为多消费者
func (t *talentRankHandler) Consume(ctx context.Context, key, value string) error {
	return t.talentRepository.CleansingDataTemporaryStorageCache(key, value)
}

func (t talentEvent) ConsumeTalent() {
	ta := &talentRankHandler{talentRepository: t.talentRepository}
	queue, err := kq.NewQueue(*t.kqConf, ta)
	if err != nil {
		log.GetTextLogger().Error(fmt.Sprintf("queue get error :" + err.Error()))
	}
	queue.Start()
}
