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

func (p *Service) FromArticle(v *article.ArticleInfo) (item *api.TargetArticle) {
	item = &api.TargetArticle{
		ID:           v.ID,
		Title:        v.Title,
		Excerpt:      v.Excerpt,
		ImageUrls:    v.ImageUrls,
		ReviseCount:  (v.Stat.ReviseCount),
		CommentCount: (v.Stat.CommentCount),
		LikeCount:    (v.Stat.LikeCount),
		DislikeCount: (v.Stat.DislikeCount),
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.CreatedAt,
		Creator: &api.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
	}

	return
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
			ID:       lvl1.ID,
			Name:     lvl1.Name,
			Seq:      lvl1.Seq,
			Type:     lvl1.Type,
			RefID:    lvl1.RefID,
			Children: make([]*api.TopicParentCatalogInfo, 0),
		}

		switch lvl1.Type {
		case model.TopicCatalogArticle:
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, lvl1.RefID); err != nil {
				return
			}
			parent.Article = p.FromArticle(article)
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
				ID:       lvl2.ID,
				Name:     lvl2.Name,
				Seq:      lvl2.Seq,
				Type:     lvl2.Type,
				RefID:    lvl2.RefID,
				Children: make([]*api.TopicChildCatalogInfo, 0),
			}

			switch lvl2.Type {
			case model.TopicCatalogArticle:
				var article *article.ArticleInfo
				if article, err = p.d.GetArticle(c, lvl2.RefID); err != nil {
					return
				}
				child.Article = p.FromArticle(article)
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
					ID:    lvl3.ID,
					Name:  lvl3.Name,
					Seq:   lvl3.Seq,
					Type:  lvl3.Type,
					RefID: lvl3.RefID,
				}

				switch lvl3.Type {
				case model.TopicCatalogArticle:
					var article *article.ArticleInfo
					if article, err = p.d.GetArticle(c, lvl3.RefID); err != nil {
						return
					}
					subItem.Article = p.FromArticle(article)
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
			ID:       lvl1.ID,
			Name:     lvl1.Name,
			Seq:      lvl1.Seq,
			Type:     lvl1.Type,
			RefID:    lvl1.RefID,
			Children: make([]*api.TopicParentCatalogInfo, 0),
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
				ID:       lvl2.ID,
				Name:     lvl2.Name,
				Seq:      lvl2.Seq,
				Type:     lvl2.Type,
				RefID:    lvl2.RefID,
				Children: make([]*api.TopicChildCatalogInfo, 0),
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

func (p *Service) saveCatalogs(c context.Context, node sqalx.Node, aid int64, req *api.ArgSaveCatalogs) (delArticles []*model.ArticleItem, newArticles []*model.ArticleItem, err error) {
	newArticles = make([]*model.ArticleItem, 0)
	delArticles = make([]*model.ArticleItem, 0)

	// check topic
	if err = p.checkTopic(c, node, req.TopicID); err != nil {
		return
	}

	// check admin role
	if err = p.checkTopicMemberAdmin(c, node, req.TopicID, aid); err != nil {
		return
	}

	var dbCatalogs []*model.TopicCatalog
	if dbCatalogs, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{"topic_id": req.TopicID, "parent_id": req.ParentID}); err != nil {
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
				newArticles = append(newArticles, &model.ArticleItem{TopicID: req.TopicID, ArticleID: v.RefID})
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
		if item.ParentID != req.ParentID {
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
		} else if v.Item.IsPrimary == true {
			err = ecode.NeedPrimaryTopic
			return
		}

		if err = p.d.DelTopicCatalog(c, node, k); err != nil {
			return
		}

		if v.Item.Type == model.TopicCatalogArticle {
			delArticles = append(delArticles, &model.ArticleItem{TopicID: req.TopicID, ArticleID: v.Item.RefID})
		}
	}
	return
}
