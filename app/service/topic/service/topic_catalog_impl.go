package service

import (
	"context"
	"fmt"
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

func (p *Service) getTopicArticleInfos(c context.Context, node sqalx.Node, topicID int64) (items map[int64]*article.ArticleInfo, err error) {
	var articleIDs []int64
	if articleIDs, err = p.d.GetTopicArticleIDs(c, node, topicID); err != nil {
		return
	}

	var resp *article.ArticleInfosResp
	if resp, err = p.d.GetArticleInfos(c, articleIDs); err != nil {
		return
	}

	items = resp.Items
	return
}

// getCatalogHierarchyOfAll 按树形层级获取话题目录信息
func (p *Service) getCatalogHierarchyOfAll(c context.Context, node sqalx.Node, topicID int64) (items []*api.TopicRootCatalogInfo, err error) {
	items = make([]*api.TopicRootCatalogInfo, 0)

	deadline, _ := c.Deadline()
	fmt.Println(deadline)

	var articleInfos map[int64]*article.ArticleInfo
	if articleInfos, err = p.getTopicArticleInfos(c, node, topicID); err != nil {
		return
	}

	fmt.Println("获取文章信息成功")

	var parents []*model.TopicCatalog
	if parents, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"topic_id":  topicID,
		"parent_id": 0,
	}); err != nil {
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
			if val, ok := articleInfos[parent.RefID]; ok {
				parent.Name = val.Title
				parent.Article = p.fromArticle(val)
				if parent.Article.RelationIDs, err = p.d.GetArticleRelationIDs(c, node, val.ID); err != nil {
					return
				}
			} else {
				// 如果未获取到文章信息，则跳过
				continue
			}
			break
		case model.TopicCatalogTopic:
			var t *model.Topic
			if t, err = p.getTopic(c, p.d.DB(), lvl1.RefID); err != nil {
				if ecode.IsNotExistEcode(err) {
					// 如果未获取到话题信息，则跳过
					continue
				}
				return
			}
			var topicStat *model.TopicStat
			if topicStat, err = p.GetTopicStat(c, lvl1.RefID); err != nil {
				return
			}
			parent.Name = t.Name
			parent.Topic = modelCopyToPbTopic(t, topicStat)
			break
		case model.TopicCatalogTestSet:
		}

		var children []*model.TopicCatalog
		if children, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
			"topic_id":  topicID,
			"parent_id": lvl1.ID,
		}); err != nil {
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
				if val, ok := articleInfos[lvl2.RefID]; ok {
					child.Name = val.Title
					child.Article = p.fromArticle(val)
					if child.Article.RelationIDs, err = p.d.GetArticleRelationIDs(c, node, val.ID); err != nil {
						return
					}
				} else {
					// 如果未获取到文章信息，则跳过
					continue
				}
			case model.TopicCatalogTopic:
				var t *model.Topic
				if t, err = p.getTopic(c, p.d.DB(), lvl2.RefID); err != nil {
					if ecode.IsNotExistEcode(err) {
						// 如果未获取到话题信息，则跳过
						continue
					}
					return
				}
				var topicStat *model.TopicStat
				if topicStat, err = p.GetTopicStat(c, lvl1.RefID); err != nil {
					return
				}
				child.Name = t.Name
				parent.Topic = modelCopyToPbTopic(t, topicStat)
			case model.TopicCatalogTestSet:
			}

			var sub []*model.TopicCatalog
			if sub, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
				"topic_id":  topicID,
				"parent_id": lvl2.ID,
			}); err != nil {
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
					if val, ok := articleInfos[lvl3.RefID]; ok {
						subItem.Name = val.Title
						subItem.Article = p.fromArticle(val)
						if subItem.Article.RelationIDs, err = p.d.GetArticleRelationIDs(c, node, val.ID); err != nil {
							return
						}
					} else {
						// 如果未获取到文章信息，则跳过
						continue
					}
				case model.TopicCatalogTopic:
					var t *model.Topic
					if t, err = p.getTopic(c, p.d.DB(), lvl3.RefID); err != nil {
						if ecode.IsNotExistEcode(err) {
							// 如果未获取到话题信息，则跳过
							continue
						}
						return
					}
					var topicStat *model.TopicStat
					if topicStat, err = p.GetTopicStat(c, lvl1.RefID); err != nil {
						return
					}
					subItem.Name = t.Name
					subItem.Topic = modelCopyToPbTopic(t, topicStat)
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

func modelCopyToPbTopic(topic *model.Topic, topicStat *model.TopicStat) (item *api.TopicInfo) {
	item = &api.TopicInfo{
		ID:              topic.ID,
		Name:            topic.Name,
		Avatar:          topic.Avatar,
		Bg:              topic.Bg,
		Introduction:    topic.Introduction,
		AllowDiscuss:    bool(topic.AllowDiscuss),
		AllowChat:       bool(topic.AllowChat),
		IsPrivate:       bool(topic.IsPrivate),
		ViewPermission:  topic.ViewPermission,
		EditPermission:  topic.EditPermission,
		JoinPermission:  topic.JoinPermission,
		CatalogViewType: topic.CatalogViewType,
		TopicHome:       topic.TopicHome,
		Creator: &api.Creator{
			ID: topic.CreatedBy,
		},
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.CreatedAt,
	}
	if topicStat != nil {
		item.Stat = &api.TopicStat{
			MemberCount:     topicStat.MemberCount,
			ArticleCount:    topicStat.ArticleCount,
			DiscussionCount: topicStat.DiscussionCount,
		}
	}

	return
}

// getCatalogTaxonomyHierarchyOfAll 按树形层级获取话题目录（只包含分类）
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

// createCatalog 创建目录条目
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

// getTopicCatalogsMap 获取话题目录字典
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

// saveCatalogs 保存话题目录
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

	// 如果已经存在则报错
	//for _, dbVal := range dic {
	//	for _, reqVal := range req.Items {
	//		if dbVal.Item.TopicID == req.TopicID && dbVal.Item.RefID == reqVal.RefID &&
	//			dbVal.Item.Type == reqVal.Type && dbVal.Item.ParentID == req.ParentID &&
	//			dbVal.Item.ID <= 0 {
	//			return nil, ecode.Error(ecode.RequestErr, "添加的文章已经存在。")
	//		}
	//	}
	//}

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

			if v.Type == model.TopicCatalogArticle {
				var article *article.ArticleInfo
				if article, err = p.d.GetArticle(c, v.RefID); err != nil {
					return
				}

				tc.Name = article.Title
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
