package service

import (
	"context"
	"time"

	"valerian/app/interface/article/model"
	article "valerian/app/service/article/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) AddArticle(c context.Context, arg *model.ArgAddArticle) (id int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.ArgAddArticle{
		Aid:            aid,
		Title:          arg.Title,
		Content:        arg.Content,
		DisableRevise:  arg.DisableRevise,
		DisableComment: arg.DisableComment,
		Files:          make([]*article.ArgArticleFile, 0),
		Relations:      make([]*article.ArgArticleRelation, 0),
	}

	if arg.Files != nil {
		for _, v := range arg.Files {
			item.Files = append(item.Files, &article.ArgArticleFile{
				FileName: v.FileName,
				FileURL:  v.FileURL,
				Seq:      int32(v.Seq),
			})
		}
	}

	if arg.Relations != nil {
		for _, v := range arg.Relations {
			item.Relations = append(item.Relations, &article.ArgArticleRelation{
				ParentID:   v.ParentID,
				TopicID:    v.TopicID,
				Primary:    v.Primary,
				Permission: v.Permission,
			})
		}
	}

	var resp *article.IDResp
	if resp, err = p.d.AddArticle(c, item); err != nil {
		return
	}

	id = resp.ID
	return
}

func (p *Service) UpdateArticle(c context.Context, arg *model.ArgUpdateArticle) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.ArgUpdateArticle{
		ID:         arg.ID,
		Content:    arg.Content,
		ChangeDesc: arg.ChangeDesc,
		Aid:        aid,
	}

	if arg.Title != nil {
		item.Title = &article.ArgUpdateArticle_TitleValue{*arg.Title}
	}

	if arg.DisableRevise != nil {
		item.DisableRevise = &article.ArgUpdateArticle_DisableReviseValue{*arg.DisableRevise}
	}

	if arg.DisableComment != nil {
		item.DisableComment = &article.ArgUpdateArticle_DisableCommentValue{*arg.DisableComment}
	}

	if err = p.d.UpdateArticle(c, item); err != nil {
		return
	}

	return
}

func (p *Service) DelArticle(c context.Context, id int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.DelArticle(c, &article.IDReq{Aid: aid, ID: id}); err != nil {
		return
	}

	return
}

func (p *Service) GetArticle(c context.Context, id int64, include string) (item *model.ArticleResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	inc := includeParam(include)
	var data *article.ArticleDetail
	if data, err = p.d.GetArticleDetail(c, &article.IDReq{Aid: aid, ID: id, Include: include}); err != nil {
		return
	}

	item = &model.ArticleResp{
		ID:        data.ID,
		Title:     data.Title,
		Content:   data.Content,
		CreatedBy: data.Creator.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Files:     make([]*model.ArticleFileResp, 0),
		Relations: make([]*model.ArticleRelationResp, 0),
	}

	item.Creator = &model.Creator{
		ID:           data.Creator.ID,
		UserName:     data.Creator.UserName,
		Avatar:       data.Creator.Avatar,
		Introduction: data.Creator.Introduction,
	}

	if data.LastHistory != nil {
		item.Updator = &model.Creator{
			ID:           data.LastHistory.Updator.ID,
			UserName:     data.LastHistory.Updator.UserName,
			Avatar:       data.LastHistory.Updator.Avatar,
			Introduction: data.LastHistory.Updator.Introduction,
		}

		item.ChangeDesc = data.LastHistory.ChangeDesc
	}

	if inc["files"] {
		if item.Files, err = p.GetArticleFiles(c, id); err != nil {
			return
		}
	}

	if inc["relations"] {
		if item.Relations, err = p.GetArticleRelations(c, id); err != nil {
			return
		}
	}

	if inc["meta"] {
		if item.ArticleMeta, err = p.getArticleMeta(c, data); err != nil {
			return
		}
	}

	p.addCache(func() {
		p.onArticleViewed(context.Background(), id, aid, time.Now().Unix())
	})

	return
}

func (p *Service) getArticleMeta(c context.Context, data *article.ArticleDetail) (item *model.ArticleMeta, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	item = new(model.ArticleMeta)

	if item.CanEdit, err = p.d.CanEdit(c, &article.IDReq{Aid: aid, ID: data.ID}); err != nil {
		return
	}

	if item.Like, err = p.d.IsLike(c, aid, data.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if item.Fav, err = p.d.IsFav(c, aid, data.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if item.Dislike, err = p.d.IsDislike(c, aid, data.ID, model.TargetTypeArticle); err != nil {
		return
	}

	item.LikeCount = data.Stat.LikeCount
	item.DislikeCount = data.Stat.DislikeCount
	item.ReviseCount = data.Stat.ReviseCount
	item.CommentCount = data.Stat.CommentCount

	return
}
