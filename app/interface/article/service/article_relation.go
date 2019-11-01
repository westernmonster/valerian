package service

import (
	"context"

	"valerian/app/interface/article/model"
	article "valerian/app/service/article/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetArticleRelations(c context.Context, articleID int64) (items []*model.ArticleRelationResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var data *article.ArticleRelationsResp
	if data, err = p.d.GetArticleRelations(c, &article.IDReq{ID: articleID, Aid: aid}); err != nil {
		return
	}

	items = make([]*model.ArticleRelationResp, 0)
	if data.Items != nil {
		for _, v := range data.Items {
			items = append(items, &model.ArticleRelationResp{
				ID:              v.ID,
				CatalogFullPath: v.CatalogFullPath,
				Primary:         v.Primary,
				Name:            v.Name,
				Avatar:          v.Avatar,
				Permission:      v.Permission,
				EditPermission:  v.EditPermission,
			})
		}
	}

	return
}

func (p *Service) UpdateArticleRelation(c context.Context, arg *model.ArgUpdateArticleRelation) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.ArgUpdateArticleRelation{
		ID:         arg.ID,
		Primary:    arg.Primary,
		Permission: arg.Permission,
		Aid:        aid,
	}

	if err = p.d.UpdateArticleRelation(c, item); err != nil {
		return
	}

	return
}

func (p *Service) SetPrimary(c context.Context, arg *model.ArgSetPrimaryArticleRelation) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.ArgSetPrimaryArticleRelation{
		ID:        arg.ID,
		ArticleID: arg.ArticleID,
		Aid:       aid,
	}

	if err = p.d.SetPrimary(c, item); err != nil {
		return
	}

	return
}

func (p *Service) AddArticleRelation(c context.Context, arg *model.ArgAddArticleRelation) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.ArgAddArticleRelation{
		ArticleID:  arg.ArticleID,
		ParentID:   arg.ParentID,
		TopicID:    arg.TopicID,
		Primary:    arg.Primary,
		Permission: arg.Permission,
		Aid:        aid,
	}

	if err = p.d.AddArticleRelation(c, item); err != nil {
		return
	}

	return
}

func (p *Service) DelArticleRelation(c context.Context, arg *model.ArgDelArticleRelation) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.ArgDelArticleRelation{
		ID:        arg.ID,
		ArticleID: arg.ArticleID,
		Aid:       aid,
	}

	if err = p.d.DelArticleRelation(c, item); err != nil {
		return
	}

	return
}
