package handler

import (
	"context"
	"gae-backend-storage/bootstrap"
	"gae-backend-storage/constant/cache"
	"gae-backend-storage/constant/common"
	"gae-backend-storage/constant/search"
	"gae-backend-storage/constant/sys"
	"gae-backend-storage/infrastructure/mysql"
	"gae-backend-storage/infrastructure/redis"
	jsoniter "github.com/json-iterator/go"
)

type ParquetStorageEngine interface {
	Setup()
}

type storageEngine struct {
	es     *bootstrap.SearchEngine
	client redis.Client
	mysql  mysql.Client
}

func (s *storageEngine) Setup() {
	handler := NewIntermediateHandler()
	readData := handler.ReadData("../../algorithm/result/all_rank.parquet", sys.UnCleansingUserId)
	contributors := *readData.Contributors
	_, _ = s.es.EsEngine.AddDoc(search.RankSearchIndex, contributors)

	hotRankContributors := contributors[:100]
	marshal, _ := jsoniter.Marshal(hotRankContributors)
	_ = s.client.Set(context.Background(), cache.HotRankPhase+common.Infix, marshal)

	s.mysql.Gorm().Create(contributors)

	data := handler.ReadData("../../algorithm/data/repo_output.parquet", sys.UnCleansingRepoId)
	repos := data.Repos

	s.mysql.Gorm().Create(repos)
}

func NewStorageEngine(b *bootstrap.SearchEngine, client redis.Client, client2 mysql.Client) ParquetStorageEngine {
	return &storageEngine{es: b, client: client, mysql: client2}
}
