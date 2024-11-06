package domain

type Rank struct {
}

type RankRepository interface {
	CacheGetRankPhase() *[]string
	CacheGetHotRank(page int, phase string) *[]*RankUser
	SearchUserRank(username string) *[]*RankUser
}

type RankUser struct {
	Placement              int    `parquet:"name=placement, type=INT64"`
	Login                  string `parquet:"name=login, type=BYTE_ARRAY, convertedtype=UTF8"`
	Id                     int64  `parquet:"name=id, type=INT64"`
	AvatarURL              string `parquet:"name=avatar_url, type=BYTE_ARRAY, convertedtype=UTF8"`
	HTMLURL                string `parquet:"name=html_url, type=BYTE_ARRAY, convertedtype=UTF8"`
	Location               string `parquet:"name=location, type=BYTE_ARRAY, convertedtype=UTF8"`
	Company                string `parquet:"name=company, type=BYTE_ARRAY, convertedtype=UTF8"`
	Score                  int32  `parquet:"name=score, type=INT32"`
	Level                  string `parquet:"name=level, type=BYTE_ARRAY, convertedtype=UTF8"`
	LocationClassification int32  `parquet:"name=location_classification, type=INT32"`
	Tech                   string `parquet:"name=tech, type=BYTE_ARRAY, convertedtype=UTF8"`
}
