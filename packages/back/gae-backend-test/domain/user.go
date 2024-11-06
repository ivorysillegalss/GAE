package domain

type Contributor struct {
	Login      string `parquet:"name=login, type=BYTE_ARRAY, convertedtype=UTF8"`
	Id         int64  `parquet:"name=id, type=INT64"`
	AvatarURL  string `parquet:"name=avatar_url, type=BYTE_ARRAY, convertedtype=UTF8"`
	URL        string `parquet:"name=url, type=BYTE_ARRAY, convertedtype=UTF8"`
	HTMLURL    string `parquet:"name=html_url, type=BYTE_ARRAY, convertedtype=UTF8"`
	GravatarID string `parquet:"name=gravatar_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Name       string `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8"`
	Company    string `parquet:"name=company, type=BYTE_ARRAY, convertedtype=UTF8"`
	Blog       string `parquet:"name=blog, type=BYTE_ARRAY, convertedtype=UTF8"`
	Location   string `parquet:"name=location, type=BYTE_ARRAY, convertedtype=UTF8"`
	Email      string `parquet:"name=email, type=BYTE_ARRAY, convertedtype=UTF8"`
	Bio        string `parquet:"name=bio, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type User struct {
	ID    int64  `parquet:"name=id, type=INT64"`
	Login string `parquet:"name=login, type=BYTE_ARRAY, convertedtype=UTF8"`
	// 根据需要可以添加其他字段
}

// UserRelation 用户关系网络相关实体类
type UserRelation struct {
	Username  string
	NetMember *[]*NetMember
}

type NetMember struct {
	Username string
	Relation string
}

type RelationInfo struct {
	Username       string
	RelationMember *[]*RelationMember
}

type RelationMember struct {
	Username  string
	Weight    int
	AvatarUrl string
}

type UserRepository interface {
	QueryUser(username string) *[]*Contributor

	SearchUserByLevel(level string, score int) *[]*RankUser

	SearchUserByNation(nation string, score int) *[]*RankUser

	SearchUserByTech(tech string, score int) *[]*RankUser
}

type UserUsecase interface {
	QueryUserSpecificInfo(username string) *[]*Contributor
	// QueryUserNetwork grpc访问py的api
	QueryUserNetwork(username string) *RelationInfo

	//TODO 打缓存
	QueryByNation(nation string, score int) *[]*RankUser

	QueryByTech(tech string, score int) *[]*RankUser

	QueryByGrade(level string, score int) *[]*RankUser
}
