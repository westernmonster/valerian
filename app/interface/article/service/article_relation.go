package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/article/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) GetArticleRelations(c context.Context, articleID int64) (items []*model.ArticleRelationResp, err error) {
	return p.getArticleRelations(c, p.d.DB(), articleID)
}

func (p *Service) getArticleRelations(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleRelationResp, err error) {
	var data []*model.TopicCatalog
	if data, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"type":   model.TopicCatalogArticle,
		"ref_id": articleID,
	}); err != nil {
		return
	}

	items = make([]*model.ArticleRelationResp, 0)
	for _, v := range data {
		item := &model.ArticleRelationResp{
			ID:         v.ID,
			ToTopicID:  v.TopicID,
			Primary:    bool(v.IsPrimary),
			Permission: *v.Permission,
		}

		var t *topic.TopicInfo
		if t, err = p.d.GetTopic(c, v.TopicID); err != nil {
			return
		}

		item.Name = t.Name
		avatar := t.GetAvatarValue()
		item.Avatar = &avatar

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

func (p *Service) checkArticleRelations(items []*model.AddArticleRelation) (err error) {
	if len(items) == 0 {
		return ecode.NeedPrimaryTopic
	}

	primary := false
	dic := make(map[int64]bool)
	for _, v := range items {
		if primary == true && v.Primary {
			return ecode.OnlyAllowOnePrimaryTopic
		}
		if v.Primary {
			primary = true
		}

		if _, ok := dic[v.TopicID]; ok {
			return ecode.AuthTopicDuplicate
		}
	}

	return nil
}

func (p *Service) checkTopicCatalogTaxonomy(c context.Context, node sqalx.Node, topicID, parentID int64) (err error) {
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
	}
	return nil
}

func (p *Service) bulkCreateArticleRelations(c context.Context, node sqalx.Node, articleID int64, title string, relations []*model.AddArticleRelation) (err error) {
	if err = p.checkArticleRelations(relations); err != nil {
		return
	}

	for _, v := range relations {
		if err = p.addArticleRelation(c, node, articleID, title, v); err != nil {
			return
		}
	}

	return
}

func (p *Service) addArticleRelation(c context.Context, node sqalx.Node, articleID int64, title string, item *model.AddArticleRelation) (err error) {
	var checkExist *model.TopicCatalog
	if checkExist, err = p.d.GetTopicCatalogByCond(c, node, map[string]interface{}{
		"topic_id": item.TopicID,
		"ref_id":   articleID,
		"type":     model.TopicCatalogArticle,
	}); err != nil {
		return
	} else if checkExist != nil {
		err = ecode.AuthTopicExist
		return
	}

	if err = p.checkTopicCatalogTaxonomy(c, node, item.TopicID, item.ParentID); err != nil {
		return
	}

	if item.Primary {
		var catalog *model.TopicCatalog
		if catalog, err = p.d.GetTopicCatalogByCond(c, node, map[string]interface{}{
			"ref_id":  articleID,
			"type":    model.TopicCatalogArticle,
			"primary": 1,
		}); err != nil {
			return
		}

		if catalog != nil && catalog.IsPrimary == true {
			err = ecode.OnlyAllowOnePrimaryTopic
			return
		}
	}

	var maxSeq int
	if maxSeq, err = p.d.GetTopicCatalogMaxChildrenSeq(c, node, item.TopicID, item.ParentID); err != nil {
		return
	}

	d := &model.TopicCatalog{
		ID:         gid.NewID(),
		Name:       title,
		Seq:        maxSeq + 1,
		Type:       model.TopicCatalogArticle,
		ParentID:   item.ParentID,
		RefID:      &articleID,
		TopicID:    item.TopicID,
		IsPrimary:  types.BitBool(item.Primary),
		Permission: &item.Permission,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicCatalog(c, node, d); err != nil {
		return
	}

	if err = p.d.AddTopicStat(c, node, &model.TopicStat{TopicID: item.TopicID, ArticleCount: 1}); err != nil {
		return
	}

	return
}

func (p *Service) UpdateArticleRelation(c context.Context, arg *model.ArgUpdateArticleRelation) (err error) {
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

	if arg.Primary {
		var cata *model.TopicCatalog
		if cata, err = p.d.GetTopicCatalogByCond(c, tx, map[string]interface{}{
			"ref_id":  catalog.RefID,
			"type":    model.TopicCatalogArticle,
			"primary": 1,
		}); err != nil {
			return
		} else if cata != nil {
			if cata.ID != arg.ID {
				err = ecode.OnlyAllowOnePrimaryTopic
				return
			}
		}

	}

	catalog.Permission = &arg.Permission
	catalog.IsPrimary = types.BitBool(arg.Primary)

	if err = p.d.UpdateTopicCatalog(c, tx, catalog); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicCatalogCache(context.TODO(), catalog.TopicID)
	})

	return
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

	var orgPrimary *model.TopicCatalog
	if orgPrimary, err = p.d.GetTopicCatalogByCond(c, tx, map[string]interface{}{
		"type":       model.TopicCatalogArticle,
		"ref_id":     arg.ArticleID,
		"is_primary": 1,
	}); err != nil {
		return
	} else if orgPrimary == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	if orgPrimary.ID == arg.ID {
		return
	}

	orgPrimary.IsPrimary = false
	if err = p.d.UpdateTopicCatalog(c, tx, orgPrimary); err != nil {
		return
	}

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, tx, arg.ID); err != nil {
		return
	} else if catalog == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	catalog.IsPrimary = true
	if err = p.d.UpdateTopicCatalog(c, tx, catalog); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicCatalogCache(context.TODO(), catalog.TopicID)
	})

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

	if err = p.addArticleRelation(c, tx, article.ID, article.Title, &model.AddArticleRelation{
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
		p.d.DelTopicCatalogCache(context.TODO(), arg.TopicID)
	})

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

	if err = p.d.AddTopicStat(c, tx, &model.TopicStat{TopicID: catalog.TopicID, ArticleCount: -1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicCatalogCache(context.TODO(), catalog.TopicID)
	})

	return
}
