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

type NationInfo struct {
	Nation       string // 用户所属的国家/地区
	NationWeight int64  // 国家/地区的权重，表示重要性或影响力
}

type NestedNetInfo struct {
	Relation   string           // 与中心节点的关系类型，例如 "friend" 或 "colleague"
	Username   string           // 关联用户的用户名
	NationInfo *NationInfo      // 关联用户的国别信息
	Relations  []*NestedNetInfo // 嵌套的下级关系，用于表示下级节点
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
	QueryUserNetwork(username string) *[]*NestedNetInfo

	//TODO 打缓存
	QueryByNation(nation string, score int) *[]*RankUser

	QueryByTech(tech string, score int) *[]*RankUser

	QueryByGrade(level string, score int) *[]*RankUser
}
