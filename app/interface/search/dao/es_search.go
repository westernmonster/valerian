package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"valerian/app/interface/search/model"

	"gopkg.in/olivere/elastic.v6"
)

// searchResult get result from ES.
func (d *Dao) searchResult(c context.Context, indexName string, query elastic.Query, bsp *model.BasicSearchParams) (res *model.SearchResult, err error) {
	res = &model.SearchResult{Debug: ""}
	if bsp.Debug {
		var src interface{}
		if src, err = query.Source(); err == nil {
			var data []byte
			if data, err = json.Marshal(src); err == nil {
				res = &model.SearchResult{Debug: string(data)}
			}
		}
	}
	// multi sort
	sorterSlice := []elastic.Sorter{}
	if bsp.KW != "" {
		sorterSlice = append(sorterSlice, elastic.NewScoreSort().Desc())
	}
	for i, d := range bsp.Sort {
		if len(bsp.Order) < i+1 {
			if bsp.Order[0] == "desc" {
				sorterSlice = append(sorterSlice, elastic.NewFieldSort(d).Desc())
			} else {
				sorterSlice = append(sorterSlice, elastic.NewFieldSort(d).Asc())
			}
		} else {
			if bsp.Order[i] == "desc" {
				sorterSlice = append(sorterSlice, elastic.NewFieldSort(d).Desc())
			} else {
				sorterSlice = append(sorterSlice, elastic.NewFieldSort(d).Asc())
			}
		}
	}
	fsc := elastic.NewFetchSourceContext(true).Include(bsp.Source...)
	searchResult, err := d.esClient.
		Search().Index(indexName).
		Query(query).
		SortBy(sorterSlice...).
		From((bsp.Pn - 1) * bsp.Ps).
		Size(bsp.Ps).
		Pretty(true).
		FetchSourceContext(fsc).
		Do(c)
	if err != nil {
		PromError(c, fmt.Sprintf("es:执行查询失败"), "%v", err)
		res = &model.SearchResult{Debug: res.Debug + "es:执行查询失败", Page: &model.Page{}}
		return
	}
	var data []json.RawMessage
	for _, hit := range searchResult.Hits.Hits {
		var t json.RawMessage
		e := json.Unmarshal(*hit.Source, &t)
		if e != nil {
			PromError(c, fmt.Sprintf("索引有脏数据"), "s.dao.SearchArchiveCheck(%d,%d) error(%v) ", bsp.Pn*bsp.Ps, bsp.Ps, e)
			continue
		}
		data = append(data, t)
	}
	res = &model.SearchResult{
		Order:  strings.Join(bsp.Order, ","),
		Sort:   strings.Join(bsp.Sort, ","),
		Result: data,
		Debug:  res.Debug,
		Page: &model.Page{
			Pn:    bsp.Pn,
			Ps:    bsp.Ps,
			Total: searchResult.Hits.TotalHits,
		},
	}
	return
}
