package service

import (
	"context"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/xstr"
)

func (p *Service) GetReviseStat(c context.Context, reviseID int64) (item *model.ReviseStat, err error) {
	return p.d.GetReviseStatByID(c, p.d.DB(), reviseID)
}

func (p *Service) GetRevise(c context.Context, reviseID int64) (item *model.Revise, err error) {
	if item, err = p.getRevise(c, p.d.DB(), reviseID); err != nil {
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, p.d.DB(), item.ArticleID); err != nil {
		return
	}

	item.Title = article.Title
	return
}

func (p *Service) getRevise(c context.Context, node sqalx.Node, reviseID int64) (item *model.Revise, err error) {
	var addCache = true
	if item, err = p.d.ReviseCache(c, reviseID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetReviseByID(c, p.d.DB(), reviseID); err != nil {
		return
	} else if item == nil {
		err = ecode.ReviseNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetReviseCache(context.TODO(), item)
		})
	}
	return
}

func (p *Service) GetReviseImageUrls(c context.Context, reviseID int64) (urls []string, err error) {
	urls = make([]string, 0)
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeRevise,
		"target_id":   reviseID,
	}); err != nil {
		return
	}

	for _, v := range imgs {
		urls = append(urls, v.URL)
	}

	return
}

func (p *Service) DelRevise(c context.Context, aid int64, reviseID int64) (err error) {
	return
}

func (p *Service) GetReviseInfo(c context.Context, req *api.IDReq) (item *api.ReviseInfo, err error) {
	revise, err := p.GetRevise(c, req.ID)
	if err != nil {
		return nil, err
	}
	article, err := p.GetArticle(c, revise.ArticleID)
	if err != nil {
		return nil, err
	}

	stat, err := p.GetReviseStat(c, req.ID)
	if err != nil {
		return nil, err
	}

	urls, err := p.GetReviseImageUrls(c, req.ID)
	if err != nil {
		return nil, err
	}

	m, err := p.getAccount(c, p.d.DB(), article.CreatedBy)
	if err != nil {
		return nil, err
	}

	resp := &api.ReviseInfo{
		ID:        revise.ID,
		Title:     article.Title,
		Excerpt:   xstr.Excerpt(revise.ContentText),
		CreatedAt: revise.CreatedAt,
		UpdatedAt: revise.UpdatedAt,
		ImageUrls: urls,
		Stat: &api.ReviseStat{
			CommentCount: int32(stat.CommentCount),
			LikeCount:    int32(stat.LikeCount),
			DislikeCount: int32(stat.DislikeCount),
		},
		Creator: &api.Creator{
			ID:           m.ID,
			UserName:     m.UserName,
			Avatar:       m.Avatar,
			Introduction: m.Introduction,
		},
		ArticleID: revise.ArticleID,
	}

	inc := includeParam(req.Include)

	if inc["content"] {
		resp.Content = article.Content
	}

	if inc["content_text"] {
		resp.ContentText = article.ContentText
	}

	return resp, nil

}
