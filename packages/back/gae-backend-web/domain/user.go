package domain

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

// UserRelation 用户关系网络相关实体类
type UserRelation struct {
	Username  string
	NetMember *[]*NetMember
}

type NetMember struct {
	Username string
	Relation string
}

type NetInfo struct {
	Relation  string
	Username1 string
	Username2 string
}

type NestedNetInfo struct {
	Username1 string
	Relations []NetInfo
}

type UserRepository interface {
	QueryUser(username string) *[]*Contributor
}

type UserUsecase interface {
	QueryUserSpecificInfo(username string) *[]*Contributor
	// QueryUserNetwork grpc访问py的api
	QueryUserNetwork(username string) *[]*NestedNetInfo

	QueryByNation(username string) *[]*Contributor

	QueryByTech(username string) *[]*Contributor
}
