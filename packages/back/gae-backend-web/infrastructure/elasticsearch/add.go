package elasticsearch

import "gae-backend-web/infrastructure/log"

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
