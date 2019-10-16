package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"valerian/app/service/search/model"
	"valerian/library/conf/env"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"gopkg.in/olivere/elastic.v6"
)

const topicMapping = `
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
    "topic": {
      "properties": {
        "id": { "type": "long" },
        "name": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart"
        },
        "avatar": { "type": "text" },
        "bg": { "type": "text" },
        "introduction": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart"
        },
        "is_private": { "type": "boolean" },
        "allow_chat": { "type": "boolean" },
        "allow_discuss": { "type": "boolean" },
        "edit_permission": { "type": "text" },
        "view_permission": { "type": "text" },
        "join_permission": { "type": "text" },
        "catalog_view_type": { "type": "text" },
        "creator": {
          "properties": {
            "id": { "type": "long" },
            "user_name": {
              "type": "text",
              "analyzer": "ik_max_word",
              "search_analyzer": "ik_smart"
            },
            "avatar": { "type": "text" },
            "introduction": { "type": "text" }
          }
        },
        "created_at": { "type": "integer" },
        "updated_at": { "type": "integer" }
      }
    }
  }
}
`

func (p *Dao) CreateTopicIndices(c context.Context) (err error) {
	indexName := fmt.Sprintf("%s_topics", env.DeployEnv)

	// Check if index exists
	var indexExist bool
	if indexExist, err = p.esClient.IndexExists(indexName).Do(c); err != nil {
		log.For(c).Error(fmt.Sprintf("check index exist failed, error(%+v)", err))
		return
	}
	if indexExist {
		return
	}

	var createRet *elastic.IndicesCreateResult
	if createRet, err = p.esClient.CreateIndex(indexName).Body(topicMapping).IncludeTypeName(true).Do(c); err != nil {
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

func (p *Dao) PutTopic2ES(c context.Context, item *model.ESTopic) (err error) {
	indexName := fmt.Sprintf("%s_topics", env.DeployEnv)
	var ret *elastic.IndexResponse
	if ret, err = p.esClient.Index().Index(indexName).Type("topic").Id(strconv.FormatInt(item.ID, 10)).BodyJson(item).Do(c); err != nil {
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

func (p *Dao) GetTopics(c context.Context, node sqalx.Node) (items []*model.Topic, err error) {
	items = make([]*model.Topic, 0)
	sqlSelect := "SELECT a.* FROM topics a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopics err(%+v)", err))
		return
	}
	return
}

func (p *Dao) GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error) {
	item = new(model.Topic)
	sqlSelect := "SELECT a.* FROM topics a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicByID err(%+v), id(%+v)", err, id))
	}

	return
}
