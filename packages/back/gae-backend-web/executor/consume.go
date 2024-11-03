package executor

type ConsumeExecutor struct {
}

func (d *ConsumeExecutor) SetupConsume() {
}

func NewConsumeExecutor() *ConsumeExecutor {
	return &ConsumeExecutor{}
}
