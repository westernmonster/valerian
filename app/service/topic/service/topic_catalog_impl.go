package service

import (
	"context"
	"time"

	article "valerian/app/service/article/api"
	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
)

type dicItem struct {
	Done bool
	Item *model.TopicCatalog
}

func (p *Service) getCatalogHierarchyOfAll(c context.Context, node sqalx.Node, topicID int64) (items []*api.TopicRootCatalogInfo, err error) {
	items = make([]*api.TopicRootCatalogInfo, 0)

	parents, err := p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"topic_id":  topicID,
		"parent_id": 0,
	})
	if err != nil {
		return
	}

	for _, lvl1 := range parents {
		parent := &api.TopicRootCatalogInfo{
			ID:        lvl1.ID,
			Name:      lvl1.Name,
			Seq:       lvl1.Seq,
			Type:      lvl1.Type,
			RefID:     lvl1.RefID,
			IsPrimary: bool(lvl1.IsPrimary),
			Children:  make([]*api.TopicParentCatalogInfo, 0),
		}

		switch lvl1.Type {
		case model.TopicCatalogArticle:
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, lvl1.RefID); err != nil {
				return
			}
			parent.Article = p.FromArticle(article)
			if parent.Article.RelationIDs, err = p.d.GetArticleRelationIDs(c, node, article.ID); err != nil {
				return
			}
		case model.TopicCatalogTestSet:
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
			child := &api.TopicParentCatalogInfo{
				ID:        lvl2.ID,
				Name:      lvl2.Name,
				Seq:       lvl2.Seq,
				Type:      lvl2.Type,
				RefID:     lvl2.RefID,
				IsPrimary: bool(lvl2.IsPrimary),
				Children:  make([]*api.TopicChildCatalogInfo, 0),
			}

			switch lvl2.Type {
			case model.TopicCatalogArticle:
				var article *article.ArticleInfo
				if article, err = p.d.GetArticle(c, lvl2.RefID); err != nil {
					return
				}
				child.Article = p.FromArticle(article)
				if child.Article.RelationIDs, err = p.d.GetArticleRelationIDs(c, node, article.ID); err != nil {
					return
				}
			case model.TopicCatalogTestSet:
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
				subItem := &api.TopicChildCatalogInfo{
					ID:        lvl3.ID,
					Name:      lvl3.Name,
					Seq:       lvl3.Seq,
					Type:      lvl3.Type,
					RefID:     lvl3.RefID,
					IsPrimary: bool(lvl3.IsPrimary),
				}

				switch lvl3.Type {
				case model.TopicCatalogArticle:
					var article *article.ArticleInfo
					if article, err = p.d.GetArticle(c, lvl3.RefID); err != nil {
						return
					}
					subItem.Article = p.FromArticle(article)
					if subItem.Article.RelationIDs, err = p.d.GetArticleRelationIDs(c, node, article.ID); err != nil {
						return
					}
				case model.TopicCatalogTestSet:
				}

				child.Children = append(child.Children, subItem)
			}

			parent.Children = append(parent.Children, child)
		}

		items = append(items, parent)

	}

	return
}

func (p *Service) getCatalogTaxonomyHierarchyOfAll(c context.Context, node sqalx.Node, topicID int64) (items []*api.TopicRootCatalogInfo, err error) {
	items = make([]*api.TopicRootCatalogInfo, 0)

	parents, err := p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"topic_id":  topicID,
		"parent_id": 0,
		"type":      model.TopicCatalogTaxonomy,
	})
	if err != nil {
		return
	}

	for _, lvl1 := range parents {
		parent := &api.TopicRootCatalogInfo{
			ID:        lvl1.ID,
			Name:      lvl1.Name,
			Seq:       lvl1.Seq,
			Type:      lvl1.Type,
			RefID:     lvl1.RefID,
			IsPrimary: bool(lvl1.IsPrimary),
			Children:  make([]*api.TopicParentCatalogInfo, 0),
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
			child := &api.TopicParentCatalogInfo{
				ID:        lvl2.ID,
				Name:      lvl2.Name,
				Seq:       lvl2.Seq,
				Type:      lvl2.Type,
				RefID:     lvl2.RefID,
				IsPrimary: bool(lvl2.IsPrimary),
				Children:  make([]*api.TopicChildCatalogInfo, 0),
			}

			parent.Children = append(parent.Children, child)
		}

		items = append(items, parent)

	}

	return
}

func (p *Service) createCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (id int64, err error) {
	// // 当前为分类，则不允许处于第三级
	if item.ParentID != 0 && item.Type == model.TopicCatalogTaxonomy {
		var parent *model.TopicCatalog
		if parent, err = p.d.GetTopicCatalogByCond(c, node, map[string]interface{}{
			"topic_id": item.TopicID,
			"id":       item.ParentID,
		}); err != nil {
			return
		} else if parent == nil {
			err = ecode.TopicCatalogNotExist
			return
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

	if err = p.d.AddTopicCatalog(c, node, item); err != nil {
		return
	}

	return item.ID, nil

}

func (p *Service) getTopicCatalogsMap(c context.Context, node sqalx.Node, topicID, parentID int64) (dic map[int64]dicItem, err error) {
	var dbCatalogs []*model.TopicCatalog
	if dbCatalogs, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{"topic_id": topicID, "parent_id": parentID}); err != nil {
		return
	}

	dic = make(map[int64]dicItem)
	for _, v := range dbCatalogs {
		dic[v.ID] = dicItem{
			Done: false,
			Item: v,
		}
	}

	return
}

func (p *Service) saveCatalogs(c context.Context, node sqalx.Node, aid int64, req *api.ArgSaveCatalogs) (change *model.CatalogChange, err error) {
	change = &model.CatalogChange{
		NewArticles:          make([]*model.ArticleItem, 0),
		DelArticles:          make([]*model.ArticleItem, 0),
		NewTaxonomyItems:     make([]*model.NewTaxonomyItem, 0),
		DelTaxonomyItems:     make([]*model.DelTaxonomyItem, 0),
		RenamedTaxonomyItems: make([]*model.RenamedTaxonomyItem, 0),
		MovedTaxonomyItems:   make([]*model.MovedTaxonomyItem, 0),
	}

	// check topic
	if err = p.checkTopic(c, node, req.TopicID); err != nil {
		return
	}

	var dic map[int64]dicItem
	if dic, err = p.getTopicCatalogsMap(c, node, req.TopicID, req.ParentID); err != nil {
		return
	}

	for _, v := range req.Items {
		if v.ID == nil {
			tc := &model.TopicCatalog{
				ID:        gid.NewID(),
				Name:      v.Name,
				Seq:       v.Seq,
				Type:      v.Type,
				ParentID:  req.ParentID,
				RefID:     v.RefID,
				TopicID:   req.TopicID,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}
			if _, err = p.createCatalog(c, node, tc); err != nil {
				return
			}

			if v.Type == model.TopicCatalogArticle {
				change.NewArticles = append(change.NewArticles, &model.ArticleItem{TopicID: req.TopicID, ArticleID: v.RefID})
			}

			if v.Type == model.TopicCatalogTaxonomy {
				change.NewTaxonomyItems = append(change.NewTaxonomyItems, &model.NewTaxonomyItem{TopicID: req.TopicID, ID: tc.ID, ParentID: req.ParentID, Name: tc.Name})
			}

			continue
		}

		// Deal Move Logic
		var item *model.TopicCatalog
		if item, err = p.d.GetTopicCatalogByID(c, node, v.GetIDValue()); err != nil {
			return
		} else if item == nil {
			err = ecode.TopicCatalogNotExist
			return
		}

		// 如果该条目是从别的地方移动过来的
		if item.ParentID != req.ParentID {
			if req.ParentID != 0 {
				var parent *model.TopicCatalog
				if parent, err = p.d.GetTopicCatalogByCond(c, node, map[string]interface{}{"topic_id": req.TopicID, "id": req.ParentID}); err != nil {
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

			if v.Type == model.TopicCatalogTaxonomy {
				change.MovedTaxonomyItems = append(change.MovedTaxonomyItems, &model.MovedTaxonomyItem{
					TopicID:     item.TopicID,
					ID:          item.ID,
					OldParentID: item.ParentID,
					NewParentID: req.ParentID,
					Name:        v.Name})
			}

		} else {
			if v.Type == model.TopicCatalogTaxonomy && v.Name != item.Name {
				change.RenamedTaxonomyItems = append(change.RenamedTaxonomyItems, &model.RenamedTaxonomyItem{
					TopicID: item.TopicID,
					ID:      item.ID,
					OldName: item.Name,
					NewName: v.Name})
			}
		}

		dic[v.GetIDValue()] = dicItem{Done: true}

		item.Name = v.Name
		item.Seq = v.Seq
		item.ParentID = req.ParentID
		item.RefID = v.RefID
		item.UpdatedAt = time.Now().Unix()

		if err = p.d.UpdateTopicCatalog(c, node, item); err != nil {
			return
		}

	}

	for k, v := range dic {
		if v.Done {
			continue
		}

		if v.Item.Type == model.TopicCatalogTaxonomy {
			var childrenCount int
			if childrenCount, err = p.d.GetTopicCatalogChildrenCount(c, node, req.TopicID, k); err != nil {
				return
			}
			if childrenCount > 0 {
				err = ecode.MustDeleteChildrenCatalogFirst
				return
			}
		}

		if err = p.d.DelTopicCatalog(c, node, k); err != nil {
			return
		}

		if v.Item.Type == model.TopicCatalogArticle {
			change.DelArticles = append(change.DelArticles, &model.ArticleItem{TopicID: req.TopicID, ArticleID: v.Item.RefID})
		}

		if v.Item.Type == model.TopicCatalogTaxonomy {
			change.DelTaxonomyItems = append(change.DelTaxonomyItems, &model.DelTaxonomyItem{
				TopicID:  req.TopicID,
				ID:       v.Item.ID,
				ParentID: v.Item.ParentID,
				Name:     v.Item.Name})
		}
	}
	return
}
