package sys

const (
	// GoRoutinePoolTypesAmount 线程池种类数量
	GoRoutinePoolTypesAmount = 1

	// GzipCompress 压缩方式为Gzip 搭配Json序列化
	GzipCompress = "gzip&json"
	// ProtoBufCompress 序列化方式
	ProtoBufCompress = "protobuf"

	StreamOverSignal = "data: [DONE]"
	StreamPrefix     = "data: "

	// StreamGenerationResponseChannelBuffer 新初始化流式传递rpc数据的channel大小
	StreamGenerationResponseChannelBuffer = 100
)

const (
	UnCleansingUserId = 3
	UnCleansingRepoId = 2
)
