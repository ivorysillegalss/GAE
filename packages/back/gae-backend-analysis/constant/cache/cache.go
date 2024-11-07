package cache

const (
	// TemporaryUnRankCleansingRepoShardList  此处数据全部代指仓库数据
	//指目前活跃的切片的编号
	TemporaryUnRankCleansingRepoShardList     = "UnRankCleansingRepoList"
	TemporaryUnRankCleansingRepoShardMaxValue = 15000
	TemporaryUnRankCleansingRepoData          = "UnRankCleansingRepo"
	TemporaryUnRankCleansingRepoShard         = "shard"
	TemporaryUnRankCleansingRepoReadyShard    = "UnRankCleansingReadyRepoShardList"
)

const (
	// TemporaryUnRankCleansingContributorShardList 代指贡献者的数据
	//指目前活跃的切片的编号
	TemporaryUnRankCleansingContributorShardList     = "UnRankCleansingContributorList"
	TemporaryUnRankCleansingContributorShardMaxValue = 15000
	TemporaryUnRankCleansingContributorData          = "UnRankCleansingContributor"
	TemporaryUnRankCleansingContributorShard         = "shard"
	TemporaryUnRankCleansingContributorReadyShard    = "UnRankCleansingReadyContributorShardList"
)
