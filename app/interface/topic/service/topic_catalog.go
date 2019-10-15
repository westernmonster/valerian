package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/topic/model"
	article "valerian/app/service/article/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

type dicItem struct {
	Done bool
	Item *model.TopicCatalog
}

func (p *Service) GetCatalogsHierarchy(c context.Context, topicID int64) (items []*model.TopicLevel1Catalog, err error) {
	return p.getCatalogsHierarchy(c, p.d.DB(), topicID)
}

func (p *Service) getCatalogsHierarchy(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicLevel1Catalog, err error) {
	// var (
	// 	addCache = true
	// )

	// if items, err = p.d.TopicCatalogCache(c, topicID); err != nil {
	// 	addCache = false
	// } else if items != nil {
	// 	return
	// }

	if items, err = p.getCatalogHierarchyOfAll(c, node, topicID); err != nil {
		return
	}

	// if addCache {
	// 	p.addCache(func() {
	// 		p.d.SetTopicCatalogCache(context.TODO(), topicID, items)
	// 	})
	// }

	return
}

func (p *Service) GetCatalogTaxonomiesHierarchy(c context.Context, topicID int64) (items []*model.TopicLevel1Catalog, err error) {
	if items, err = p.getCatalogTaxonomyHierarchyOfAll(c, p.d.DB(), topicID); err != nil {
		return
	}

	return
}

func (p *Service) getCatalogHierarchyOfAll(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicLevel1Catalog, err error) {
	items = make([]*model.TopicLevel1Catalog, 0)

	parents, err := p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"topic_id":  topicID,
		"parent_id": 0,
	})
	if err != nil {
		return
	}

	for _, lvl1 := range parents {
		parent := &model.TopicLevel1Catalog{
			ID:       &lvl1.ID,
			Name:     lvl1.Name,
			Seq:      lvl1.Seq,
			Type:     lvl1.Type,
			RefID:    lvl1.RefID,
			Children: make([]*model.TopicLevel2Catalog, 0),
		}

		switch lvl1.Type {
		case model.TopicCatalogArticle:
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, *lvl1.RefID); err != nil {
				return
			}
			parent.Article = p.FromArticle(article)
			continue
		case model.TopicCatalogTestSet:
			continue
		}

		children, eInner := p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
			"topic_id":  topicID,
			"parent_id": lvl1.ID,
		})
		if eInner != nil {
			err = eInner
			return
		}

		for _, lvl2 := range children {
			child := &model.TopicLevel2Catalog{
				ID:       &lvl2.ID,
				Name:     lvl2.Name,
				Seq:      lvl2.Seq,
				Type:     lvl2.Type,
				RefID:    lvl2.RefID,
				Children: make([]*model.TopicChildCatalog, 0),
			}

			switch lvl2.Type {
			case model.TopicCatalogArticle:
				var article *article.ArticleInfo
				if article, err = p.d.GetArticle(c, *lvl2.RefID); err != nil {
					return
				}
				child.Article = p.FromArticle(article)
				continue
			case model.TopicCatalogTestSet:
				continue
			}

			sub, eInner := p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
				"topic_id":  topicID,
				"parent_id": lvl2.ID,
			})
			if eInner != nil {
				err = eInner
				return
			}

			for _, lvl3 := range sub {
				subItem := &model.TopicChildCatalog{
					ID:    &lvl3.ID,
					Name:  lvl3.Name,
					Seq:   lvl3.Seq,
					Type:  lvl3.Type,
					RefID: lvl3.RefID,
				}

				switch lvl3.Type {
				case model.TopicCatalogArticle:
					var article *article.ArticleInfo
					if article, err = p.d.GetArticle(c, *lvl3.RefID); err != nil {
						return
					}
					subItem.Article = p.FromArticle(article)
					continue
				case model.TopicCatalogTestSet:
					continue
				}

				child.Children = append(child.Children, subItem)
			}

			parent.Children = append(parent.Children, child)
		}

		items = append(items, parent)

	}

	return
}

func (p *Service) getCatalogTaxonomyHierarchyOfAll(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicLevel1Catalog, err error) {
	items = make([]*model.TopicLevel1Catalog, 0)

	parents, err := p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"topic_id":  topicID,
		"parent_id": 0,
		"type":      model.TopicCatalogTaxonomy,
	})
	if err != nil {
		return
	}

	for _, lvl1 := range parents {
		parent := &model.TopicLevel1Catalog{
			ID:       &lvl1.ID,
			Name:     lvl1.Name,
			Seq:      lvl1.Seq,
			Type:     lvl1.Type,
			RefID:    lvl1.RefID,
			Children: make([]*model.TopicLevel2Catalog, 0),
		}

		children, eInner := p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
			"topic_id":  topicID,
			"parent_id": lvl1.ID,
			"type":      model.TopicCatalogTaxonomy,
		})
		if eInner != nil {
			err = eInner
			return
		}

		for _, lvl2 := range children {
			child := &model.TopicLevel2Catalog{
				ID:       &lvl2.ID,
				Name:     lvl2.Name,
				Seq:      lvl2.Seq,
				Type:     lvl2.Type,
				RefID:    lvl2.RefID,
				Children: make([]*model.TopicChildCatalog, 0),
			}

			parent.Children = append(parent.Children, child)
		}

		items = append(items, parent)

	}

	return
}

func (p *Service) createCatalog(c context.Context, node sqalx.Node, topicID int64, name string, seq int, rtype string, refID *int64, parentID int64) (id int64, err error) {
	// // 当前为分类，则不允许处于第三级
	if parentID != 0 && rtype == model.TopicCatalogTaxonomy {
		var parent *model.TopicCatalog
		if parent, err = p.d.GetTopicCatalogByCond(c, node, map[string]interface{}{
			"topic_id": topicID,
			"id":       parentID,
		}); err != nil {
			return
		} else if parent == nil {
			err = ecode.TopicCatalogNotExist
			// return
		}

		if parent.Type != model.TopicCatalogTaxonomy {
			err = ecode.InvalidCatalog
			return
		}

		if parent.ParentID != 0 {
			err = ecode.InvalidCatalog
			return
		}
	}

	item := &model.TopicCatalog{
		ID:        gid.NewID(),
		Name:      name,
		Seq:       seq,
		Type:      rtype,
		ParentID:  parentID,
		RefID:     refID,
		TopicID:   topicID,
		Deleted:   false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddTopicCatalog(c, node, item); err != nil {
		return
	}

	return item.ID, nil

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

	// check topic
	if err = p.checkTopic(c, tx, req.TopicID); err != nil {
		return
	}

	// check admin role
	if err = p.checkTopicMemberAdmin(c, tx, req.TopicID, aid); err != nil {
		return
	}

	var dbCatalogs []*model.TopicCatalog
	if dbCatalogs, err = p.d.GetTopicCatalogsByCond(c, tx, map[string]interface{}{"topic_id": req.TopicID, "parent_id": req.ParentID}); err != nil {
		return
	}

	dic := make(map[int64]dicItem)
	for _, v := range dbCatalogs {
		dic[v.ID] = dicItem{
			Done: false,
			Item: v,
		}
	}

	for _, v := range req.Items {
		if v.ID == nil {
			if _, err = p.createCatalog(c, tx, req.TopicID, v.Name, v.Seq, v.Type, v.RefID, req.ParentID); err != nil {
				return
			}

			if v.Type == model.TopicCatalogArticle {
				p.addCache(func() {
					p.onCatalogArticleAdded(c, *v.RefID, req.TopicID, aid, time.Now().Unix())
				})
			}
			continue
		}

		// Deal Move Logic
		var item *model.TopicCatalog
		if item, err = p.d.GetTopicCatalogByID(c, tx, *v.ID); err != nil {
			return
		} else if item == nil {
			return ecode.TopicCatalogNotExist
		}
		if item.ParentID != req.ParentID {
			var parent *model.TopicCatalog
			if parent, err = p.d.GetTopicCatalogByCond(c, tx, map[string]interface{}{"topic_id": req.TopicID, "id": req.ParentID}); err != nil {
				return
			} else if parent == nil {
				err = ecode.TopicCatalogNotExist
				return
			}

			if item.Type == model.TopicCatalogTaxonomy && parent.ParentID != 0 {
				err = ecode.InvalidCatalog
				return
			}
		}

		dic[*v.ID] = dicItem{Done: true}

		item.Name = v.Name
		item.Seq = v.Seq
		item.ParentID = req.ParentID
		item.RefID = v.RefID
		item.UpdatedAt = time.Now().Unix()

		if err = p.d.UpdateTopicCatalog(c, tx, item); err != nil {
			return
		}

	}

	for k, v := range dic {
		if v.Done {
			continue
		}

		if v.Item.Type == model.TopicCatalogTaxonomy {
			var childrenCount int
			if childrenCount, err = p.d.GetTopicCatalogChildrenCount(c, tx, req.TopicID, k); err != nil {
				return
			}
			if childrenCount > 0 {
				err = ecode.MustDeleteChildrenCatalogFirst
				return
			}
		} else if v.Item.IsPrimary == true {
			err = ecode.NeedPrimaryTopic
			return
		}

		if err = p.d.DelTopicCatalog(c, tx, k); err != nil {
			return
		}

		if v.Item.Type == model.TopicCatalogArticle {
			p.addCache(func() {
				p.onCatalogArticleDeleted(c, *v.Item.RefID, req.TopicID, aid, time.Now().Unix())
			})
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicCatalogCache(context.TODO(), req.TopicID)
	})
	return
}
