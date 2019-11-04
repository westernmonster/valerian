package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"github.com/davecgh/go-spew/spew"
)

//  GetCatalogsHierarchy 按层级获取类目
func (p *Service) GetCatalogsHierarchy(c context.Context, topicID int64) (items []*api.TopicRootCatalogInfo, err error) {
	return p.getCatalogHierarchyOfAll(c, p.d.DB(), topicID)
}

//  GetCatalogsHierarchy 按层级获取类目分类
func (p *Service) GetCatalogTaxonomiesHierarchy(c context.Context, topicID int64) (items []*api.TopicRootCatalogInfo, err error) {
	if items, err = p.getCatalogTaxonomyHierarchyOfAll(c, p.d.DB(), topicID); err != nil {
		return
	}

	return
}

func (p *Service) SaveCatalogs(c context.Context, req *api.ArgSaveCatalogs) (err error) {
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var change *model.CatalogChange
	if change, err = p.saveCatalogs(c, tx, req.Aid, req); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		for _, v := range change.NewArticles {
			p.onCatalogArticleAdded(c, v.ArticleID, v.TopicID, req.Aid, time.Now().Unix())
		}

		for _, v := range change.DelArticles {
			p.onCatalogArticleDeleted(c, v.ArticleID, v.TopicID, req.Aid, time.Now().Unix())
		}

		for _, v := range change.NewTaxonomyItems {
			p.onTopicTaxonomyCatalogAdded(c, v, req.Aid, time.Now().Unix())
		}

		for _, v := range change.RenamedTaxonomyItems {
			p.onTopicTaxonomyCatalogRenamed(c, v, req.Aid, time.Now().Unix())
		}

		for _, v := range change.MovedTaxonomyItems {
			p.onTopicTaxonomyCatalogMoved(c, v, req.Aid, time.Now().Unix())
		}

		for _, v := range change.DelTaxonomyItems {
			spew.Dump(v)
			p.onTopicTaxonomyCatalogDeleted(c, v, req.Aid, time.Now().Unix())
		}
	})
	return
}

func (p *Service) HasTaxonomy(c context.Context, topicID int64) (hasTaxonomy bool, err error) {
	if hasTaxonomy, err = p.d.HasTaxonomy(c, p.d.DB(), topicID); err != nil {
		return
	}

	return
}
