package bootstrap

import (
	"gae-backend-analysis/infrastructure/elasticsearch"
	"log"
)

func NewEsEngine(env *Env) elasticsearch.Client {
	client, err := elasticsearch.NewElasticSearchClient()
	if err != nil {
		log.Fatal(err)
	}
	return client
}
