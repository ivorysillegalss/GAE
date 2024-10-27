package consume

import (
	"context"
	"gae-backend-analysis/bootstrap"
	"gae-backend-analysis/constant/mq"
	"gae-backend-analysis/domain"
	jsoniter "github.com/json-iterator/go"
)

type GenerateEvent struct {
	*baseMessageHandler
	env                  *bootstrap.Env
	channels             *bootstrap.Channels
	generationRepository domain.GenerationRepository
}

func (g GenerateEvent) StreamDataReady(b []byte) error {
	var dataReady domain.StreamGenerationReadyStorageData
	_ = jsoniter.Unmarshal(b, dataReady)
	g.generationRepository.ReadyStreamDataStorage(context.Background(), dataReady)
	return nil
}

func (g GenerateEvent) AsyncStreamStorageDataReady() {
	g.ConsumeMessage(mq.UserChatReadyCallingQueue, g.StreamDataReady)
}

func (g GenerateEvent) PublishStreamReadyStorageData(data *domain.StreamGenerationReadyStorageData) {
	//此方法应用于流信息 发起调用前 所以此时没generateText
	marshal, _ := jsoniter.Marshal(data)
	g.PublishMessage(mq.UserChatReadyCallingQueue, marshal)
}

func NewGenerateEvent(h MessageHandler, e *bootstrap.Env, c *bootstrap.Channels, g domain.GenerationRepository) domain.GenerateEvent {
	messageHandler := h.(*baseMessageHandler)
	chatReadyCalling := &MessageQueueArgs{
		ExchangeName: mq.UserChatReadyCallingExchange,
		QueueName:    mq.UserChatReadyCallingQueue,
		KeyName:      mq.UserChatReadyCallingKey,
	}
	messageHandler.InitMessageQueue(chatReadyCalling)
	return &GenerateEvent{
		baseMessageHandler:   messageHandler,
		env:                  e,
		channels:             c,
		generationRepository: g,
	}
}
