package consume

import (
	"context"
	"gae-backend-web/bootstrap"
	"gae-backend-web/domain"
)

type talentEvent struct {
	env   *bootstrap.Env
	kafka *bootstrap.KafkaConf
}

func NewTalentEvent(env *bootstrap.Env, conf *bootstrap.KafkaConf) domain.TalentEvent {
	return talentEvent{env: env, kafka: conf}
}

type talentRankHandler struct {
	businessId int
}

func (t *talentRankHandler) Consume(ctx context.Context, key, value string) error {
	return nil
}
