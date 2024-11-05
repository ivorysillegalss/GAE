package repository

import (
	"database/sql"
	"gae-backend-web/constant/search"
	"gae-backend-web/domain"
	"gae-backend-web/infrastructure/elasticsearch"
	"gae-backend-web/infrastructure/hive"
	"gae-backend-web/infrastructure/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
)

type userRepository struct {
	hive hive.Client
	es   elasticsearch.Client
}

func NewUserRepository(hc hive.Client) domain.UserRepository {
	return &userRepository{hive: hc}
}

func (u *userRepository) QueryUser(username string) *[]*domain.Contributor {
	//TODO query条件补全
	query, err := u.hive.Query("")
	if err != nil {
		log.GetTextLogger().Warn("error querying user, username is: " + username)
		return nil
	}
	var c []*domain.Contributor
	for query.Next() {
		cv, err := newContributorValue(query)
		if err != nil {
			if err != nil {
				log.GetTextLogger().Warn("error querying user, username is: " + username)
				return nil
			}
		}
		c = append(c, cv)
	}
	return &c
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
