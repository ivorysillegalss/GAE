package repository

import (
	"context"
	"gae-backend-storage/constant/cache"
	"gae-backend-storage/constant/common"
	"gae-backend-storage/constant/search"
	"gae-backend-storage/domain"
	"gae-backend-storage/infrastructure/elasticsearch"
	"gae-backend-storage/infrastructure/log"
	"gae-backend-storage/infrastructure/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
)

type rankRepository struct {
	rcl redis.Client
	es  elasticsearch.Client
}

func NewRankRepository(rcl redis.Client, es elasticsearch.Client) domain.RankRepository {
	return &rankRepository{rcl: rcl, es: es}
}

var ctx = context.Background()

func (r *rankRepository) CacheGetRankPhase() *[]string {
	v, _ := r.rcl.LRangeAll(ctx, cache.HotRankPhase)
	return &v
}

func (r *rankRepository) CacheGetHotRank(page int, phase string) *[]*domain.RankUser {
	var hotUsers []*domain.RankUser
	value, _ := r.rcl.ZRange(ctx, cache.HotRankPhase+common.Infix+phase, page, page+cache.HotRankPerPage)
	for _, v := range value {
		var u domain.RankUser
		_ = jsoniter.Unmarshal([]byte(v), u)
		hotUsers = append(hotUsers, &u)
	}
	return &hotUsers
}

func (r *rankRepository) SearchUserRank(username string) *[]*domain.RankUser {

	//TODO es字段存入,与读取出的大小写匹配问题
	esClient := r.es.GetClient()
	opts := elastic.NewTermQuery("login", username)
	targetQuery := elastic.NewBoolQuery().Filter(opts)
	searchResult, err := esClient.
		Search().
		Index(search.RankSearchIndex).
		Query(targetQuery).
		Sort("score", true).
		Size(10).
		Do(ctx)

	if err != nil || len(searchResult.Hits.Hits) == 0 {
		log.GetTextLogger().Warn("can't find target user, username is:" + username)
		return nil
	}

	sortValue := searchResult.Hits.Hits[0].Sort[0]

	elastic.NewBoolQuery().Filter(opts).Filter(elastic.NewRangeQuery("score").Gte(sortValue))
	res, err := esClient.
		Search().
		Index(search.RankSearchIndex).
		Sort("score", true).
		Size(10).
		Do(ctx)

	if err != nil {
		log.GetTextLogger().Warn("can't find after user, username is:" + username)
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
