package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"valerian/app/interface/article/model"
	article "valerian/app/service/article/api"
	search "valerian/app/service/search/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
	"valerian/library/xstr"
)

func (p *Service) GetArticleRelations(c context.Context, articleID int64) (items []*model.ArticleRelationResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	return p.getArticleRelations(c, aid, articleID)
}

func (p *Service) getArticleRelations(c context.Context, aid int64, articleID int64) (items []*model.ArticleRelationResp, err error) {
	var data *article.ArticleRelationsResp
	if data, err = p.d.GetArticleRelations(c, &article.IDReq{ID: articleID, Aid: aid}); err != nil {
		return
	}

	items = make([]*model.ArticleRelationResp, 0)
	if data.Items != nil {
		for _, v := range data.Items {
			items = append(items, &model.ArticleRelationResp{
				ID:              v.ID,
				ToTopicID:       v.ToTopicID,
				CatalogFullPath: v.CatalogFullPath,
				Primary:         v.Primary,
				Name:            v.Name,
				Avatar:          v.Avatar,
				Permission:      v.Permission,
				EditPermission:  v.EditPermission,
				Introduction:    v.Introduction,
				MemberCount:     v.Stat.MemberCount,
				ArticleCount:    v.Stat.ArticleCount,
				DiscussionCount: v.Stat.DiscussionCount,
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

func (p *Service) GetUserCanEditArticles(c context.Context, query string, pn, ps int) (resp *model.ArticleListResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var idsResp *article.IDsResp
	if idsResp, err = p.d.GetUserCanEditArticleIDs(c, &article.AidReq{AccountID: aid}); err != nil {
		return
	}

	if idsResp.IDs == nil || len(idsResp.IDs) == 0 {
		resp = &model.ArticleListResp{
			Items:  make([]*model.ArticleItem, 0),
			Paging: &model.Paging{IsEnd: true},
		}
		return
	}

	var data *search.SearchResult
	if data, err = p.d.SearchArticle(c, &search.SearchParam{KW: query, Sort: []string{"updated_at"}, Order: []string{"desc"}, Pn: int32(pn), Ps: int32(ps), IDs: idsResp.IDs}); err != nil {
		err = ecode.SearchArticleFailed
		return
	}

	resp = &model.ArticleListResp{
		Items:  make([]*model.ArticleItem, len(data.Result)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Result {
		t := new(model.ESArticle)
		err = json.Unmarshal(v, t)
		if err != nil {
			return
		}
		item := &model.ArticleItem{
			ID:    t.ID,
			Title: *t.Title,
		}

		item.CreatedAt = *t.CreatedAt
		item.UpdatedAt = *t.UpdatedAt

		if t.ContentText != nil {
			item.Excerpt = xstr.Excerpt(*t.ContentText)
		}

		var stat *model.ArticleStat
		if stat, err = p.d.GetArticleStatByID(c, p.d.DB(), t.ID); err != nil {
			return
		}

		urls, err := p.GetArticleImageUrls(c, t.ID)
		if err != nil {
			return nil, err
		}

		item.ImageUrls = urls

		item.LikeCount = stat.LikeCount
		item.DislikeCount = stat.DislikeCount
		item.ReviseCount = stat.ReviseCount
		item.CommentCount = stat.CommentCount

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/article/list/has_edit_permission", url.Values{
		"query": []string{query},
		"pn":    []string{strconv.Itoa(pn - 1)},
		"ps":    []string{strconv.Itoa(ps)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/article/list/has_edit_permission", url.Values{
		"query": []string{query},
		"pn":    []string{strconv.Itoa(pn + 1)},
		"ps":    []string{strconv.Itoa(ps)},
	}); err != nil {
		return
	}

	if len(resp.Items) < ps {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if pn == 1 {
		resp.Paging.Prev = ""
	}

	return
}

func (p *Service) GetArticleImageUrls(c context.Context, articleID int64) (urls []string, err error) {
	urls = make([]string, 0)
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeArticle,
		"target_id":   articleID,
	}); err != nil {
		return
	}

	for _, v := range imgs {
		urls = append(urls, v.URL)
	}

	return
}
