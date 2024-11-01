package elasticsearch

import (
	"context"
	"gae-backend-analysis/constant/dao"
	"github.com/olivere/elastic/v7"
	"log"
)

var ctx = context.Background()
var Url = "http://127.0.0.1:9200"

//var esClient *elastic.Client

var index = "student"

// 定义一些变量，mapping为定制的index字段类型
// number_of_replicas备份数 ， number_of_shards分片数
const mapping = `
{
   "settings":{
      "number_of_shards": 1,
      "number_of_replicas": 0 
   },
}`

// 初始化es连接
//func init() {
//	client, err := elastic.NewClient(
//		elastic.SetURL(Url),
//	)
//	if err != nil {
//		log.Fatal("es 连接失败:", err)
//	}
//	// ping通服务端，并获得服务端的es版本,本实例的es版本为version 7.16.2
//	info, code, err := client.Ping(Url).Do(ctx)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Elasticsearch call code:", code, " version:", info.Version.Number)
//	esClient = client
//}

type Client interface {
	Ping() (int, error)
}

type esClient struct {
	es *elastic.Client
}

func (es *esClient) Ping() (int, error) {
	_, code, err := es.es.Ping(dao.EsUrl).Do(context.Background())
	return code, err
}

func NewElasticSearchClient() (Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(dao.EsUrl),
	)
	if err != nil {
		log.Fatal(err)
	}
	info, code, err := client.Ping(dao.EsUrl).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Elasticsearch call code:", code, " version:", info.Version.Number)
	return &esClient{es: client}, nil
}
