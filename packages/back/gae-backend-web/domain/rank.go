package domain

type Rank struct {
}

type RankUsecase interface {
	GetHotRank(page int, phase string) *[]*RankUser
	GetHotRankPhase() *[]string
	SearchUserRank(username string) *[]*RankUser
}

type RankRepository interface {
	CacheGetRankPhase() *[]string
	CacheGetHotRank(page int, phase string) *[]*RankUser
	SearchUserRank(username string) *[]*RankUser
}

type RankUser struct {
	Login     string
	Id        int64
	AvatarURL string
	HTMLURL   string
	Location  string
	Company   string
	//TODO 存储分数TBD
	Score int
	Level string
}
