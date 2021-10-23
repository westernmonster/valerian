package service

import (
	"context"
	"time"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
)

// checkArticleRelations 检测文章关联话题请求
func (p *Service) checkArticleRelations(c context.Context, node sqalx.Node, aid int64, items []*api.ArgArticleRelation) (err error) {
	dic := make(map[int64]bool)
	for _, v := range items {
		if _, err = p.getTopic(c, node, v.TopicID); err != nil {
			return
		}

		var canEdit bool
		if canEdit, err = p.d.IsCanEditTopic(c, node, aid, v.TopicID); err != nil {
			return
		} else if !canEdit {
			err = ecode.NoTopicEditPermission
			return
		}

		if _, ok := dic[v.TopicID]; ok {
			return ecode.AuthTopicDuplicate
		}
	}

	return nil
}

// checkTopicCatalogTaxonomy 检测话题目录分类
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

// bulkCreateArticleRelations 批量创建文章关联话题信息
func (p *Service) bulkCreateArticleRelations(c context.Context, node sqalx.Node, aid, articleID int64, title string, relations []*api.ArgArticleRelation) (ids []int64, err error) {
	ids = make([]int64, 0)
	if err = p.checkArticleRelations(c, node, aid, relations); err != nil {
		return
	}

	for _, v := range relations {
		var id int64
		if id, err = p.addArticleRelation(c, node, articleID, title, v); err != nil {
			return
		}
		ids = append(ids, id)
	}

	return
}

// addArticleRelation 添加文章关联话题信息
func (p *Service) addArticleRelation(c context.Context, node sqalx.Node, articleID int64, title string, item *api.ArgArticleRelation) (id int64, err error) {
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

	var maxSeq int
	if maxSeq, err = p.d.GetTopicCatalogMaxChildrenSeq(c, node, item.TopicID, item.ParentID); err != nil {
		return
	}

	d := &model.TopicCatalog{
		ID:         gid.NewID(),
		Name:       title,
		Seq:        int32(maxSeq + 1),
		Type:       model.TopicCatalogArticle,
		ParentID:   item.ParentID,
		RefID:      articleID,
		TopicID:    item.TopicID,
		IsPrimary:  types.BitBool(item.Primary),
		Permission: item.Permission,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicCatalog(c, node, d); err != nil {
		return
	}

	if err = p.d.IncrTopicStat(c, node, &model.TopicStat{TopicID: item.TopicID, ArticleCount: 1}); err != nil {
		return
	}

	id = d.ID

	return
}

// getCatalogFullPath 获取文章在类目中的路径
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
	if p2, err = p.d.GetTopicCatalogByID(c, node, p1.ParentID); err != nil {
		return
	} else if p2 == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	path = p2.Name + "/" + path

	return
}

// getArticleRelations 文章关联话题信息
func (p *Service) getArticleRelations(c context.Context, node sqalx.Node, articleID int64) (items []*api.ArticleRelationResp, err error) {
	var data []*model.TopicCatalog
	if data, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"type":   model.TopicCatalogArticle,
		"ref_id": articleID,
	}); err != nil {
		return
	}

	items = make([]*api.ArticleRelationResp, 0)
	for _, v := range data {
		item := &api.ArticleRelationResp{
			ID:         v.ID,
			ToTopicID:  v.TopicID,
			Primary:    bool(v.IsPrimary),
			Permission: v.Permission,
		}

		var t *model.Topic
		if t, err = p.getTopic(c, p.d.DB(), v.TopicID); err != nil {
			return
		}

		item.Name = t.Name
		item.Avatar = t.Avatar
		item.Introduction = t.Introduction
		item.EditPermission = t.EditPermission

		var stat *model.TopicStat
		if stat, err = p.d.GetTopicStatByID(c, p.d.DB(), v.TopicID); err != nil {
			return
		}

		item.Stat = &api.TopicStat{
			MemberCount:     stat.MemberCount,
			ArticleCount:    stat.ArticleCount,
			DiscussionCount: stat.DiscussionCount,
		}

		if item.CatalogFullPath, err = p.getCatalogFullPath(c, node, v); err != nil {
			return
		}

		items = append(items, item)
	}

	return
}