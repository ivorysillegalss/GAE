package domain

type Talent struct {
	Repos        *[]Repo
	Contributors *[]Contributor
}

type Repo struct {
	RepoId           int64
	OwnerName        string
	OwnerId          string
	Name             string
	CreatedAt        int64
	UpdatedAt        int64
	ForksCount       int
	NetworkCount     int
	OpenIssuesCount  int
	StargazersCount  int
	SubscribersCount int
	WatchersCount    int
	ContributorsId   *[]string
}

type Contributor struct {
	Login         string
	Id            int64
	AvatarURL     *string
	URL           string
	Contributions int
}

type TalentEvent interface {
	ConsumeRepo()
	ConsumeContributors()
}

type TalentRepository interface {
	CleansingDataTemporaryStorageCache(header string, value string, businessId int) error
	GetAndUpdateCleansingDataShardOffset(businessId int) (int, []string, error)
}

type TalentCron interface {
	AnalyseContributor()
	AnalyseRepo()
}
