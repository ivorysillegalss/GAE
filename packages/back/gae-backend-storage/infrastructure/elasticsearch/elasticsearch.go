package elasticsearch

import (
	"context"
	"gae-backend-storage/constant/search"

	"gae-backend-storage/infrastructure/log"
	"github.com/olivere/elastic/v7"
)

var ctx = context.Background()

type Client interface {
	Ping(esUrl string) (int, error)
	AddDoc(index string, data any) (bool, error)
	QueryDoc(index string, field string, filter elastic.Query, sort string, page int, limit int) (interface{}, error)
	GetClient() *elastic.Client
}

func NewElasticSearchClient(esUrl string) (Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(esUrl),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.GetTextLogger().Fatal(err.Error())
	}
	info, code, err := client.Ping(esUrl).Do(context.Background())
	if err != nil {
		log.GetTextLogger().Fatal("error is :%v", err)
	}
	log.GetTextLogger().Info("Elasticsearch call code:", code, " version:", info.Version.Number)

	initEsIndex(client)

	return &EsClient{Es: client}, nil
}

func initEsIndex(client *elastic.Client) {
	addIndex(client, search.RankSearchIndex, search.RankIndexMapping)
}

type EsClient struct {
	Es *elastic.Client
}

func (es *EsClient) GetClient() *elastic.Client {
	return es.Es
}

func addIndex(es *elastic.Client, index string, mapping string) bool {
	exists, _ := checkIndex(es, index)
	//不存在对应的索引
	if exists {
		return false
	} else {
		//不存在对应的索引
		_, err := es.CreateIndex(index).BodyString(mapping).Do(ctx)
		log.GetTextLogger().Error("error creating index ,error is :", err.Error())
		return true
	}
}

// 查看对应的索引是否存在 如果存在则跳过 不存在则重新创建
func checkIndex(es *elastic.Client, index string) (bool, error) {
	exists, err := es.IndexExists(index).Do(ctx)
	if err != nil {
		log.GetTextLogger().Error("error is:", err.Error())
		return false, nil
	}
	return exists, nil

}

func (es *EsClient) Ping(esUrl string) (int, error) {
	_, code, err := es.Es.Ping(esUrl).Do(context.Background())
	return code, err
}
