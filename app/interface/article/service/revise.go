package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/article/model"
	article "valerian/app/service/article/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetArticleRevisesPaged(c context.Context, articleID int64, sort string, limit, offset int) (resp *model.ReviseListResp, err error) {
	var data *article.ReviseListResp
	if data, err = p.d.GetArticleRevisesPaged(c, &article.ArgArticleRevisesPaged{ArticleID: articleID, Sort: sort, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		return
	}

	resp = &model.ReviseListResp{
		Paging: &model.Paging{},
		Items:  make([]*model.ReviseItem, 0),
	}

	for _, v := range data.Items {
		item := &model.ReviseItem{
			ID:        v.ID,
			Excerpt:   v.Excerpt,
			ImageUrls: make([]string, 0),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		item.Creator = &model.Creator{
			ID:           v.Creator.ID,
			UserName:     v.Creator.UserName,
			Avatar:       v.Creator.Avatar,
			Introduction: v.Creator.Introduction,
		}

		if v.ImageUrls != nil {
			item.ImageUrls = v.ImageUrls
		}

		resp.Items = append(resp.Items, item)
	}

	if resp.Paging.Prev, err = genURL("/api/v1/article/list/revises", url.Values{
		"article_id": []string{strconv.FormatInt(articleID, 10)},
		"sort":       []string{sort},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/article/list/revises", url.Values{
		"article_id": []string{strconv.FormatInt(articleID, 10)},
		"sort":       []string{sort},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset + limit)},
	}); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	var stat *model.ArticleStat
	if stat, err = p.d.GetArticleStatByID(c, p.d.DB(), articleID); err != nil {
		return
	}

	resp.ReviseCount = int32(stat.ReviseCount)

	return
}

func (p *Service) AddRevise(c context.Context, arg *model.ArgAddRevise) (id int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.ArgAddRevise{Aid: aid, ArticleID: arg.ArticleID, Content: arg.Content, Files: make([]*article.AddReviseFile, 0)}
	for _, v := range arg.Files {
		item.Files = append(item.Files, &article.AddReviseFile{
			FileName: v.FileName,
			FileURL:  v.FileURL,
			Seq:      int32(v.Seq),
		})
	}
	var idResp *article.IDResp
	if idResp, err = p.d.AddRevise(c, item); err != nil {
		return
	}

	id = idResp.ID
	return
}

func (p *Service) UpdateRevise(c context.Context, arg *model.ArgUpdateRevise) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.ArgUpdateRevise{Aid: aid, ID: arg.ID, Content: arg.Content, Files: make([]*article.AddReviseFile, 0)}
	for _, v := range arg.Files {
		item.Files = append(item.Files, &article.AddReviseFile{
			FileName: v.FileName,
			FileURL:  v.FileURL,
			Seq:      int32(v.Seq),
		})
	}
	if err = p.d.UpdateRevise(c, item); err != nil {
		return
	}

	return
}

func (p *Service) DelRevise(c context.Context, id int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.IDReq{Aid: aid, ID: id}
	if err = p.d.DelRevise(c, item); err != nil {
		return
	}

	return
}

func (p *Service) GetRevise(c context.Context, reviseID int64) (resp *model.ReviseDetailResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data *article.ReviseDetail
	if data, err = p.d.GetReviseDetail(c, &article.IDReq{ID: reviseID, Aid: aid}); err != nil {
		return
	}

	resp = &model.ReviseDetailResp{
		ID:        data.ID,
		Title:     data.Title,
		ArticleID: data.ArticleID,
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Files:     make([]*model.ReviseFileResp, 0),
	}

	resp.Creator = &model.Creator{
		ID:           data.Creator.ID,
		UserName:     data.Creator.UserName,
		Avatar:       data.Creator.Avatar,
		Introduction: data.Creator.Introduction,
	}

	if data.Files != nil {
		for _, v := range data.Files {
			resp.Files = append(resp.Files, &model.ReviseFileResp{
				ID:       v.ID,
				FileName: v.FileName,
				FileURL:  v.FileURL,
				Seq:      int(v.Seq),
			})
		}
	}

	resp.DislikeCount = int(data.Stat.DislikeCount)
	resp.LikeCount = int(data.Stat.LikeCount)
	resp.CommentCount = int(data.Stat.CommentCount)

	if resp.Fav, err = p.d.IsFav(c, aid, reviseID, model.TargetTypeRevise); err != nil {
		return
	}

	if resp.Like, err = p.d.IsLike(c, aid, reviseID, model.TargetTypeRevise); err != nil {
		return
	}

	if resp.Dislike, err = p.d.IsDislike(c, aid, reviseID, model.TargetTypeRevise); err != nil {
		return
	}

	if aid == data.Creator.ID {
		resp.CanEdit = true
	}

	return
}
