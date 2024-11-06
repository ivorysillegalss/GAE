package domain

type RankUsecase interface {
	GetHotRank(page int, phase string) *[]*RankUser
	GetHotRankPhase() *[]string
	SearchUserRank(username string) *[]*RankUser
	GetRankEntity() *RankEntity
}

type RankRepository interface {
	CacheGetRankPhase() *[]string
	CacheGetHotRank(page int, phase string) *[]*RankUser
	SearchUserRank(username string) *[]*RankUser
}

type RankEntity struct {
	Tech   []string
	Nation []string
	Level  []string
}

type RankUser struct {
	//排行值
	Placement int
	Login     string
	Id        int64
	AvatarURL string
	HTMLURL   string
	Location  string
	Company   string
	//TODO 存储分数TBD
	Score                  int
	Level                  string
	LocationClassification string
	Tech                   string
}
