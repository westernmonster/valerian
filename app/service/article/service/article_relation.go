package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/log"
)

// GetArticleRelations 文章关联话题信息
func (p *Service) GetArticleRelations(c context.Context, articleID int64) (items []*api.ArticleRelationResp, err error) {
	return p.getArticleRelations(c, p.d.DB(), articleID)
}

// UpdateArticleRelation 更新文章关联话题信息
func (p *Service) UpdateArticleRelation(c context.Context, arg *api.ArgUpdateArticleRelation) (err error) {
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

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, tx, arg.ID); err != nil {
		return
	} else if catalog == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	if err = p.checkEditPermission(c, tx, arg.Aid, catalog.RefID); err != nil {
		return
	}

	catalog.Permission = arg.Permission
	catalog.IsPrimary = types.BitBool(arg.Primary)

	if err = p.d.UpdateTopicCatalog(c, tx, catalog); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicCatalogCache(context.Background(), catalog.TopicID)
	})

	return
}

// AddArticleRelation 添加文章关联话题信息
func (p *Service) AddArticleRelation(c context.Context, arg *api.ArgAddArticleRelation) (err error) {
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

	var article *model.Article
	if article, err = p.d.GetArticleByID(c, tx, arg.ArticleID); err != nil {
		return
	} else if article == nil {
		return ecode.ArticleNotExist
	}

	if err = p.checkEditPermission(c, tx, arg.Aid, arg.ArticleID); err != nil {
		return
	}

	if _, err = p.addArticleRelation(c, tx, article.ID, article.Title, &api.ArgArticleRelation{
		TopicID:    arg.TopicID,
		Permission: arg.Permission,
		ParentID:   arg.ParentID,
		Primary:    arg.Primary,
	}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicCatalogCache(context.Background(), arg.TopicID)
	})

	return
}

// DelArticleRelation 删除文章关联话题信息
func (p *Service) DelArticleRelation(c context.Context, arg *api.ArgDelArticleRelation) (err error) {
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

	var article *model.Article
	if article, err = p.d.GetArticleByID(c, tx, arg.ArticleID); err != nil {
		return
	} else if article == nil {
		return ecode.ArticleNotExist
	}

	if err = p.checkEditPermission(c, tx, arg.Aid, arg.ArticleID); err != nil {
		return
	}

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, tx, arg.ID); err != nil {
		return
	} else if catalog == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	var existItems []*model.TopicCatalog
	if existItems, err = p.d.GetTopicCatalogsByCond(c, tx, map[string]interface{}{
		"ref_id": arg.ArticleID,
		"type":   model.TopicCatalogArticle,
	}); err != nil {
		return
	}
	if len(existItems) == 1 {
		err = ecode.NeedArticleRelation
		return
	}

	if err = p.d.DelTopicCatalog(c, tx, catalog.ID); err != nil {
		return
	}

	if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{TopicID: catalog.TopicID, ArticleCount: -1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicCatalogCache(context.Background(), catalog.TopicID)
		p.onCatalogArticleDeleted(context.Background(), arg.ArticleID, catalog.TopicID, arg.Aid, time.Now().Unix())
	})

	return
}
