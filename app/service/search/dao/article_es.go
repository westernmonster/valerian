package dao

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"valerian/app/service/search/model"
	"valerian/library/conf/env"
	"valerian/library/log"
	"valerian/library/sync/errgroup"

	"gopkg.in/olivere/elastic.v6"
)

const articleMapping = `
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
    "article": {
      "properties": {
	   "suggest": { "type": "completion" },
        "id": { "type": "long" },
        "title": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart",
		  "copy_to":["suggest"]
        },
        "content": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart"
        },
        "content_text": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart",
		  "copy_to":["suggest"]
        },
        "disable_revise": { "type": "boolean" },
        "disable_comment": { "type": "boolean" },
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

func (p *Dao) CreateArticleIndices(c context.Context) (err error) {
	indexName := fmt.Sprintf("%s_articles", env.DeployEnv)

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
	if createRet, err = p.esClient.CreateIndex(indexName).Body(articleMapping).IncludeTypeName(true).Do(c); err != nil {
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

func (p *Dao) BulkArticle2ES(c context.Context, items []*model.ESArticle) (err error) {
	indexName := fmt.Sprintf("%s_articles", env.DeployEnv)
	docsc := make(chan *model.ESArticle)
	g, ctx := errgroup.WithContext(c)
	g.Go(func() error {
		defer close(docsc)
		for _, v := range items {
			select {
			case docsc <- v:
			case <-c.Done():
				return ctx.Err()
			}
		}
		return nil
	})
	g.Go(func() error {
		bulk := p.esClient.Bulk().Index(indexName).Type("article")
		for d := range docsc {
			// Enqueue the document
			bulk.Add(elastic.NewBulkIndexRequest().Id(strconv.FormatInt(d.ID, 10)).Doc(d))
			if bulk.NumberOfActions() >= 20 {
				// Commit
				res, err := bulk.Do(ctx)
				if err != nil {
					return err
				}
				if res.Errors {
					// Look up the failed documents with res.Failed(), and e.g. recommit
					return errors.New("bulk commit failed")
				}
				// "bulk" is reset after Do, so you can reuse it
			}

			select {
			default:
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		// Commit the final batch before exiting
		if bulk.NumberOfActions() > 0 {
			_, err = bulk.Do(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})

	// Wait until all goroutines are finished
	if err = g.Wait(); err != nil {
		return
	}

	return
}

func (p *Dao) PutArticle2ES(c context.Context, item *model.ESArticle) (err error) {
	indexName := fmt.Sprintf("%s_articles", env.DeployEnv)
	var ret *elastic.IndexResponse
	if ret, err = p.esClient.Index().Index(indexName).Type("article").Id(strconv.FormatInt(item.ID, 10)).BodyJson(item).Do(c); err != nil {
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

func (p *Dao) DelESArticle(c context.Context, id int64) (err error) {
	indexName := fmt.Sprintf("%s_articles", env.DeployEnv)
	var ret *elastic.DeleteResponse
	if ret, err = p.esClient.Delete().Index(indexName).Type("article").Id(strconv.FormatInt(id, 10)).Do(c); err != nil {
		log.For(c).Error(fmt.Sprintf("delete doc failed, error(%+v)", err))
		return
	}
	if ret == nil {
		msg := fmt.Sprintf("expected delete response to be != nil, index_name(%s),id(%d) ", indexName, id)
		log.For(c).Error(msg)
		err = errors.New(msg)
		return
	}
	return
}
