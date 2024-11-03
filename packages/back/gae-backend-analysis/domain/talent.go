package domain

type Talent struct {
	Repos        *[]*Repo
	Contributors *[]*Contributor
}

type Repo struct {
	RepoId           int64    `parquet:"name=repo_id, type=INT64"`
	OwnerName        string   `parquet:"name=owner_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	OwnerId          string   `parquet:"name=owner_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Name             string   `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8"`
	CreatedAt        int64    `parquet:"name=created_at, type=INT64"`
	UpdatedAt        int64    `parquet:"name=updated_at, type=INT64"`
	ForksCount       int32    `parquet:"name=forks_count, type=INT32"`
	NetworkCount     int32    `parquet:"name=network_count, type=INT32"`
	OpenIssuesCount  int32    `parquet:"name=open_issues_count, type=INT32"`
	StargazersCount  int32    `parquet:"name=stargazers_count, type=INT32"`
	SubscribersCount int32    `parquet:"name=subscribers_count, type=INT32"`
	WatchersCount    int32    `parquet:"name=watchers_count, type=INT32"`
	ContributorsId   []string `parquet:"name=contributors_id, type=BYTE_ARRAY, repetitiontype=REPEATED, convertedtype=UTF8"`
}

type Contributor struct {
	Login          string `parquet:"name=login, type=BYTE_ARRAY, convertedtype=UTF8"`
	Id             int64  `parquet:"name=id, type=INT64"`
	AvatarURL      string `parquet:"name=avatar_url, type=BYTE_ARRAY, convertedtype=UTF8"`
	URL            string `parquet:"name=url, type=BYTE_ARRAY, convertedtype=UTF8"`
	HTMLURL        string `parquet:"name=html_url, type=BYTE_ARRAY, convertedtype=UTF8"`
	GravatarID     string `parquet:"name=gravatar_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Name           string `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8"`
	Company        string `parquet:"name=company, type=BYTE_ARRAY, convertedtype=UTF8"`
	Blog           string `parquet:"name=blog, type=BYTE_ARRAY, convertedtype=UTF8"`
	Location       string `parquet:"name=location, type=BYTE_ARRAY, convertedtype=UTF8"`
	Email          string `parquet:"name=email, type=BYTE_ARRAY, convertedtype=UTF8"`
	Bio            string `parquet:"name=bio, type=BYTE_ARRAY, convertedtype=UTF8"`
	PublicRepos    int32  `parquet:"name=public_repos, type=INT32"`
	PublicGists    int32  `parquet:"name=public_gists, type=INT32"`
	FollowerCount  int32  `parquet:"name=follower_count, type=INT32"`
	FollowingCount int32  `parquet:"name=following_count, type=INT32"`
	CreatedAt      int64  `parquet:"name=created_at, type=INT64"`
	UpdatedAt      int64  `parquet:"name=updated_at, type=INT64"`
	Type           string `parquet:"name=type, type=BYTE_ARRAY, convertedtype=UTF8"`

	// 嵌套数组字段
	Followers  []*User `parquet:"name=followers, type=LIST"`
	Followings []*User `parquet:"name=followings, type=LIST"`
}

type User struct {
	ID    int64  `parquet:"name=id, type=INT64"`
	Login string `parquet:"name=login, type=BYTE_ARRAY, convertedtype=UTF8"`
	// 根据需要可以添加其他字段
}

type TalentEvent interface {
	ConsumeRepo()
	ConsumeContributors()
}

type TalentRepository interface {
	CleansingDataTemporaryStorageCache(header string, value string, businessId int) error
	// GetAndUpdateCleansingDataShardOffset 当为空时返回nil
	GetAndUpdateCleansingDataShardOffset(businessId int) (bool, int, []string, error)
}

type TalentCron interface {
	AnalyseContributor()
	AnalyseRepo()
}
