package elasticsearch

import (
	"context"
	"gae-backend-analysis/constant/dao"
	"gae-backend-analysis/infrastructure/log"
	"github.com/olivere/elastic/v7"
)

var ctx = context.Background()

// TODO 根据Talent搜索所需的元素，改造mapping,此处为示例
const mapping = `
{
   "settings":{
      "number_of_shards": 1,
      "number_of_replicas": 0 
   },
}`

type Client interface {
	Ping() (int, error)
	AddDoc(index string, data any) (bool, error)
}

func NewElasticSearchClient() (Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(dao.EsUrl),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.GetTextLogger().Fatal(err.Error())
	}
	info, code, err := client.Ping(dao.EsUrl).Do(context.Background())
	if err != nil {
		log.GetTextLogger().Fatal("error is :%v", err)
	}
	log.GetTextLogger().Info("Elasticsearch call code:", code, " version:", info.Version.Number)

	initEsIndex(client)

	return &esClient{es: client}, nil
}

func initEsIndex(client *elastic.Client) {
	//TODO 加入TalentRank所需的元素，构建搜索部分资源的表
	//addIndex(client,"","")

}

type esClient struct {
	es *elastic.Client
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

func (es *esClient) AddDoc(index string, data any) (bool, error) {
	exist, err := checkIndex(es.es, index)
	if err != nil {
		log.GetTextLogger().Fatal("create index error,error is: %v", err.Error())
		return false, err
	}

	if !exist {
		log.GetTextLogger().Fatal("cannot find target index: %v", err.Error())
	}
	_, err = es.es.Index().Index(index).BodyJson(data).Do(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (es *esClient) Ping() (int, error) {
	_, code, err := es.es.Ping(dao.EsUrl).Do(context.Background())
	return code, err
}
