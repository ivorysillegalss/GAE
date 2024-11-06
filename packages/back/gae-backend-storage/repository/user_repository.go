package repository

import (
	"database/sql"
	"gae-backend-storage/constant/search"
	"gae-backend-storage/domain"
	"gae-backend-storage/infrastructure/elasticsearch"
	"gae-backend-storage/infrastructure/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
)

type userRepository struct {
	es elasticsearch.Client
}

func NewUserRepository(es elasticsearch.Client) domain.UserRepository {
	return &userRepository{es: es}
}

func (u *userRepository) QueryUser(username string) *[]*domain.Contributor {
	//TODO
	panic("TODO")
}

// 根据查询结果构建并返回 Contributor 实例
func newContributorValue(rows *sql.Rows) (*domain.Contributor, error) {
	var contributor domain.Contributor
	err := rows.Scan(
		&contributor.Login,
		&contributor.Id,
		&contributor.AvatarURL,
		&contributor.URL,
		&contributor.HTMLURL,
		&contributor.GravatarID,
		&contributor.Name,
		&contributor.Company,
		&contributor.Blog,
		&contributor.Location,
		&contributor.Email,
		&contributor.Bio,
		&contributor.PublicRepos,
		&contributor.PublicGists,
		&contributor.FollowerCount,
		&contributor.FollowingCount,
		&contributor.CreatedAt,
		&contributor.UpdatedAt,
		&contributor.Type,
	)
	if err != nil {
		return nil, err
	}
	return &contributor, nil
}

func (u *userRepository) SearchUserByLevel(level string, score int) *[]*domain.RankUser {
	esClient := u.es.GetClient()
	return searchUser(esClient, "level", level, score)
}

func (u *userRepository) SearchUserByNation(nation string, score int) *[]*domain.RankUser {
	esClient := u.es.GetClient()
	return searchUser(esClient, "nation", nation, score)
}

func (u *userRepository) SearchUserByTech(tech string, score int) *[]*domain.RankUser {
	esClient := u.es.GetClient()
	return searchUser(esClient, "tech", tech, score)
}

func searchUser(es *elastic.Client, esCondition string, condition string, score int) *[]*domain.RankUser {

	opts := elastic.NewTermQuery(esCondition, condition)
	boolQuery := elastic.NewBoolQuery().Filter(opts)
	res, err := es.Search().
		Index(search.RankSearchIndex).
		Query(boolQuery).
		//Sort("_id", true).
		Sort("score", true).
		SearchAfter([]any{score}).
		Size(10).
		Do(ctx)

	if err != nil || len(res.Hits.Hits) == 0 {
		log.GetTextLogger().Warn("can't find target condition  :" + condition)
		return nil
	}
	var rankUser []*domain.RankUser
	for _, v := range res.Hits.Hits {
		var u domain.RankUser
		if v.Source != nil {
			_ = jsoniter.Unmarshal(v.Source, &u)
			rankUser = append(rankUser, &u)
		}
	}
	return &rankUser
}
