package executor

import (
	"gae-backend-analysis/domain"
	"gae-backend-analysis/infrastructure/log"
)

type ConsumeExecutor struct {
	generateEvent domain.GenerateEvent
}

func (d *ConsumeExecutor) SetupConsume() {
	d.generateEvent.AsyncStreamStorageDataReady()
	log.GetTextLogger().Info("AsyncStreamStorageDataReady QUEUE start")

	log.GetTextLogger().Info("ALL-----QUEUE----START-----SUCCESSFULLY")
	//TODO
	//在这里全部启动消费者逻辑
}

func NewConsumeExecutor(g domain.GenerateEvent) *ConsumeExecutor {
	return &ConsumeExecutor{generateEvent: g}
}
