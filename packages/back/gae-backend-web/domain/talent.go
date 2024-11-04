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

type TalentEvent interface {
}

type TalentRepository interface {
}

type TalentCron interface {
}
