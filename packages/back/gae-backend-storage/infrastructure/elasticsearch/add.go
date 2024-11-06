package elasticsearch

import "gae-backend-storage/infrastructure/log"

func (es *EsClient) AddDoc(index string, data any) (bool, error) {
	exist, err := checkIndex(es.Es, index)
	if err != nil {
		log.GetTextLogger().Fatal("create index error,error is: %v", err.Error())
		return false, err
	}

	if !exist {
		log.GetTextLogger().Fatal("cannot find target index: %v", err.Error())
	}
	_, err = es.Es.Index().Index(index).BodyJson(data).Do(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
