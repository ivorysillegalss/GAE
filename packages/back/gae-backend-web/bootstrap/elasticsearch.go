package bootstrap

import (
	"gae-backend-web/infrastructure/elasticsearch"
	"log"
)

func NewEsEngine(env *Env) elasticsearch.Client {
	client, err := elasticsearch.NewElasticSearchClient(env.ElasticSearchUrl)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
