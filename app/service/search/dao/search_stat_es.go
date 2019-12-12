package dao

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"strconv"
	"valerian/app/service/search/model"
	"valerian/library/conf/env"
	"valerian/library/log"
)

const searchStatMapping = `
{
  "settings": {
    "analysis": {
      "analyzer": {
        "my_analyzer": {
          "tokenizer": "ik_max_word",
          "char_filter": ["html_strip"]
        }
      }
    }
  },
  "mappings": {
    "search_stat": {
      "properties": {
        "id": { "type": "long" },
        "enterpoint": { "type": "text" },
        "keywords": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart"
        },
        "hits": { "type": "long" },
        "created_by": { "type": "long" },
        "created_at": { "type": "integer" },
        "updated_at": { "type": "integer" }
      }
    }
  }
}
`

func (p *Dao) CreateSearchStatIndices(c context.Context) (err error) {

	indexName := fmt.Sprintf("%s_search_stat", env.DeployEnv)

	// Check if index exists
	var indexExist bool
	if indexExist, err = p.esClient.IndexExists(indexName).Do(c); err != nil {
		log.For(c).Error(fmt.Sprintf("check index exist failed, error(%+v)", err))
		return
	}
	if indexExist {
		var deleteRet *elastic.IndicesDeleteResponse
		if deleteRet, err = p.esClient.DeleteIndex(indexName).Do(c); err != nil {
			log.For(c).Error(fmt.Sprintf("delete index failed, error(%+v)", err))
			return
		}

		if !deleteRet.Acknowledged {
			msg := fmt.Sprintf("expected DeleteIndex.Acknowledged %v; got %v", true, deleteRet.Acknowledged)
			log.For(c).Error(msg)
			err = errors.New(msg)
			return
		}
	}

	var createRet *elastic.IndicesCreateResult
	if createRet, err = p.esClient.CreateIndex(indexName).Body(searchStatMapping).IncludeTypeName(true).Do(c); err != nil {
		log.For(c).Error(fmt.Sprintf("create index failed, error(%+v)", err))
		return
	}

	if !createRet.Acknowledged {
		msg := fmt.Sprintf("expected IndicesCreateResult.Acknowledged %v; got %v", true, createRet.Acknowledged)
		log.For(c).Error(msg)
		err = errors.New(msg)
		return
	}

	return
}

func (p *Dao) PutSearchStat2ES(c context.Context, item *model.ESSearchStat) (err error) {
	indexName := fmt.Sprintf("%s_search_stat", env.DeployEnv)
	var ret *elastic.IndexResponse
	if ret, err = p.esClient.Index().Index(indexName).Type("search_stat").Id(strconv.FormatInt(item.ID, 10)).BodyJson(item).Do(c); err != nil {
		log.For(c).Error(fmt.Sprintf("index doc failed, error(%+v)", err))
		return
	}
	if ret == nil {
		msg := fmt.Sprintf("expected index response to be != nil, index_name(%s),doc(%+v) ", indexName, item)
		log.For(c).Error(msg)
		err = errors.New(msg)
		return
	}
	return
}
