package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) getArticleRelations(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleRelationResp, err error) {
	var data []*model.TopicCatalog
	if data, err = p.d.GetTopicCatalogsByCondition(c, node, map[string]interface{}{
		"type":   model.TopicCatalogArticle,
		"ref_id": articleID,
	}); err != nil {
		return
	}

	items = make([]*model.ArticleRelationResp, 0)
	for _, v := range data {
		item := &model.ArticleRelationResp{
			TopicID:        v.TopicID,
			TopicCatalogID: v.ID,
			Primary:        bool(v.IsPrimary),
		}

		var t *model.TopicResp
		if t, err = p.getTopic(c, node, v.TopicID); err != nil {
			return
		}

		item.TopicName = t.Name
		if item.CatalogFullPath, err = p.getCatalogFullPath(c, node, v); err != nil {
			return
		}

		items = append(items, item)
	}

	return
}

func (p *Service) getCatalogFullPath(c context.Context, node sqalx.Node, articleItem *model.TopicCatalog) (path string, err error) {
	if articleItem.ParentID == 0 {
		path = ""
		return
	}

	var p1 *model.TopicCatalog
	if p1, err = p.d.GetTopicCatalogByID(c, node, articleItem.ParentID); err != nil {
		return
	} else if p1 == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	path = p1.Name
	if p1.ParentID == 0 {
		return
	}

	var p2 *model.TopicCatalog
	if p2, err = p.d.GetTopicCatalogByID(c, node, articleItem.ParentID); err != nil {
		return
	} else if p2 == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	path = p2.Name + "/" + path

	return
}

func (p *Service) bulkCreateArticleCatalogs(c context.Context, node sqalx.Node, articleID int64, title string, relations []*model.AddArticleRelation) (err error) {
	var tx sqalx.Node
	if tx, err = node.Beginx(c); err != nil {
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

	if err = p.checkRelations(relations); err != nil {
		return
	}

	if len(relations) == 0 {
		err = ecode.NeedPrimaryTopic
		return
	}

	for _, v := range relations {
		if err = p.checkCatalog(c, tx, v.TopicID, v.ParentID); err != nil {
			return
		}

		if err = p.checkEditPermission(c, tx, v.TopicID); err != nil {
			return
		}

		var maxSeq int
		if maxSeq, err = p.d.GetTopicCatalogMaxChildrenSeq(c, tx, v.TopicID, v.ParentID); err != nil {
			return
		}

		item := &model.TopicCatalog{
			ID:        gid.NewID(),
			Name:      title,
			Seq:       maxSeq + 1,
			Type:      model.TopicCatalogArticle,
			ParentID:  v.ParentID,
			TopicID:   v.TopicID,
			IsPrimary: types.BitBool(v.Primary),
			RefID:     &articleID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddTopicCatalog(c, tx, item); err != nil {
			return
		}

		p.addCache(func() {
			p.d.DelTopicCatalogCache(context.TODO(), v.TopicID)
		})

	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

func (p *Service) checkRelations(items []*model.AddArticleRelation) (err error) {
	if len(items) == 0 {
		return ecode.NeedPrimaryTopic
	}

	primary := false
	dic := make(map[int64]bool)
	for _, v := range items {
		if primary == true {
			return ecode.OnlyAllowOnePrimaryTopic
		}
		if v.Primary {
			primary = true
		}

		if _, ok := dic[v.TopicID]; ok {
			return ecode.DuplicateTopicID
		}
	}

	return nil
}

func (p *Service) checkCatalog(c context.Context, node sqalx.Node, topicID, parentID int64) (err error) {
	if parentID == 0 {
		return
	}

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, node, parentID); err != nil {
		return
	} else if catalog == nil {
		err = ecode.TopicCatalogNotExist
		return
	} else if catalog.Type != model.TopicCatalogTaxonomy {
		err = ecode.InvalidCatalog
		return
	} else if catalog.TopicID != topicID {
		err = ecode.InvalidCatalog
		return
	}
	return nil
}

func (p *Service) SetPrimary(c context.Context, arg *model.ArgSetPrimaryArticleRelation) (err error) {
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

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, tx, arg.ID); err != nil {
		return
	} else if catalog == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	if err = p.checkEditPermission(c, tx, catalog.TopicID); err != nil {
		return
	}

	if catalog.IsPrimary == false {
		var orgPrimary *model.TopicCatalog
		if orgPrimary, err = p.d.GetTopicCatalogByCondition(c, tx, map[string]interface{}{
			"topic_id":   catalog.TopicID,
			"type":       model.TopicCatalogArticle,
			"ref_id":     arg.ArticleID,
			"is_primary": 1,
		}); err != nil {
			return
		} else if orgPrimary == nil {
			err = ecode.TopicCatalogNotExist
			return
		}

		orgPrimary.IsPrimary = false

		if err = p.d.UpdateTopicCatalog(c, tx, orgPrimary); err != nil {
			return
		}
	}

	catalog.IsPrimary = true
	if err = p.d.UpdateTopicCatalog(c, tx, catalog); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

func (p *Service) AddArticleRelation(c context.Context, arg *model.ArgAddArticleRelation) (err error) {
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

	if err = p.checkEditPermission(c, tx, arg.TopicID); err != nil {
		return
	}

	if arg.Primary {
		var catalog *model.TopicCatalog
		if catalog, err = p.d.GetTopicCatalogByCondition(c, tx, map[string]interface{}{
			"topic_id": arg.TopicID,
			"type":     model.TopicCatalogArticle,
			"ref_id":   arg.ArticleID,
		}); err != nil {
			return
		}

		if catalog != nil && catalog.IsPrimary == true {
			err = ecode.OnlyAllowOnePrimaryTopic
			return
		}
	}

	var maxSeq int
	if maxSeq, err = p.d.GetTopicCatalogMaxChildrenSeq(c, tx, arg.TopicID, arg.ParentID); err != nil {
		return
	}

	item := &model.TopicCatalog{
		ID:        gid.NewID(),
		Name:      article.Title,
		Seq:       maxSeq + 1,
		Type:      model.TopicCatalogArticle,
		ParentID:  arg.ParentID,
		RefID:     &article.ID,
		TopicID:   arg.TopicID,
		IsPrimary: types.BitBool(arg.Primary),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddTopicCatalog(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

func (p *Service) DelArticleRelation(c context.Context, arg *model.ArgDelArticleRelation) (err error) {
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

	// TODO: 删除关联话题需要判断什么权限？
	// if err = p.checkEditPermission(c, tx, arg.TopicID); err != nil {
	// 	return
	// }

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, tx, arg.ID); err != nil {
		return
	} else if catalog == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	if catalog.IsPrimary == true {
		err = ecode.NeedPrimaryTopic
		return
	}

	if err = p.d.DelTopicCatalog(c, tx, catalog.ID); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}
