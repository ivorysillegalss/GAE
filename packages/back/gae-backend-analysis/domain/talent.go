package domain

type Talent struct {
	//TODO 补充更多个人信息
	username int
	nickname int
	email    string
	location string
	company  string

	bio               string
	Stars             int
	Forks             int
	PullRequest       int
	RepositoryAmounts int
}

type TalentEvent interface {
	ConsumeTalent()
}

type TalentRepository interface {
	CleansingDataTemporaryStorageCache(header string, value string) error
	GetAndUpdateCleansingDataShardOffset() (int, []string, error)
}

type TalentCron interface {
	AnalyseTalent()
}
