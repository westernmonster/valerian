package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/metadata"
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

func (p *Service) SaveCatalogs(c context.Context, req *model.ArgSaveTopicCatalog) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

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

	var newArticles, delArticles []*model.ArticleItem
	if delArticles, newArticles, err = p.saveCatalogs(c, tx, aid, req); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		for _, v := range newArticles {
			p.onCatalogArticleAdded(c, v.ArticleID, v.TopicID, aid, time.Now().Unix())
		}

		for _, v := range delArticles {
			p.onCatalogArticleDeleted(c, v.ArticleID, v.TopicID, aid, time.Now().Unix())
		}
	})
	return
}
