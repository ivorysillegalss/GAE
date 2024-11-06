package mq

const (
	UnRankCleansingService   = "gaeMessageConsumerService"
	UnRankCleansingServiceId = 1
	KafkaDefaultLocalBroker  = "localhost:9092"

	FirstOffset          = "first"
	UnRankCleansingGroup = "gaeMessageGroup"
	UnRankCleansingTopic = "gaeMessageTopic"
)

const (
	MaxGoroutine = 10

	UnCleansingRepoId    = 2
	UnCleansingRepoGroup = "gaeUnCleansingRepoGroup"
	UnCleansingRepoTopic = "gaeUnCleansingRepoTopic"

	UnCleansingUserId    = 3
	UnCleansingUserGroup = "gaeUnCleansingUserGroup"
	UnCleansingUserTopic = "gaeUnCleansingUserTopic"
)
