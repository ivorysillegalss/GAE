package elasticsearch

import (
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

// 查询数据
func query(es *elastic.Client, index string, field string, filter elastic.Query, sort string, page int, limit int) (*elastic.SearchResult, error) {
	// 分页数据处理
	isAsc := true
	if sort != "" {
		sortSlice := strings.Split(sort, " ")
		sort = sortSlice[0]
		if sortSlice[1] == "desc" {
			isAsc = false
		}
	}
	// 查询位置处理
	if page <= 1 {
		page = 1
	}

	fsc := elastic.NewFetchSourceContext(true)
	// 返回字段处理
	if field != "" {
		fieldSlice := strings.Split(field, ",")
		if len(fieldSlice) > 0 {
			for _, v := range fieldSlice {
				fsc.Include(v)
			}
		}
	}

	// 开始查询位置
	fromStart := (page - 1) * limit
	res, err := es.Search().
		Index(index).
		FetchSourceContext(fsc).
		Query(filter).
		Sort(sort, isAsc).
		From(fromStart).
		Size(limit).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil

}

// QueryDoc index:"index"      索引
// field:"name,age"   要查询的字段
// filter:*TermQuery  查询规则
// sort:"age asc"     排序规则
//
//page:0             页数
//limit:10           条目数量
func (es *esClient) QueryDoc(index string, field string, filter elastic.Query, sort string, page int, limit int) (interface{}, error) {
	res, err := query(es.es, index, field, filter, sort, page, limit)
	strD, _ := json.Marshal(res)
	if err != nil {
		fmt.Println("失败:", err)
		return nil, nil
	}
	fmt.Println("执行完成")
	return string(strD), nil
}
