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
	Placement int `parquet:"name=placement, type=INT64"`
	Login     string
	Id        int64
	AvatarURL string
	HTMLURL   string
	Location  string
	Company   string
	//TODO 存储分数TBD
	Score                  int
	Level                  string
	LocationClassification float64
	Tech                   string
}

type RankEntity struct {
	Tech   []string
	Nation []string
	Level  []string
}

func RandRankEntity() *RankEntity {
	return &RankEntity{
		Tech:   []string{"c++", "c", "java"},
		Nation: []string{"china", "usa"},
		Level:  []string{"a", "b", "c"},
	}
}
