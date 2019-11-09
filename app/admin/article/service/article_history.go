package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/admin/article/model"
	article "valerian/app/service/article/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetArticleHistoriesResp(c context.Context, articleID int64, limit, offset int) (resp *model.ArticleHistoryListResp, err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var data *article.ArticleHistoryListResp
	if data, err = p.d.GetArticleHistoriesPaged(c, &article.ArgArticleHistoriesPaged{
		Aid:       aid,
		ArticleID: articleID,
		Limit:     int32(limit),
		Offset:    int32(offset),
	}); err != nil {
		return
	}

	resp = &model.ArticleHistoryListResp{
		Items:  make([]*model.ArticleHistoryItem, 0),
		Paging: &model.Paging{},
	}

	if data.Items != nil {
		for _, v := range data.Items {
			item := &model.ArticleHistoryItem{
				ID:         v.ID,
				ArticleID:  v.ArticleID,
				Seq:        int(v.Seq),
				ChangeDesc: v.ChangeDesc,
				UpdatedAt:  v.UpdatedAt,
				CreatedAt:  v.CreatedAt,
			}

			item.Updator = &model.Creator{
				ID:           v.Updator.ID,
				UserName:     v.Updator.UserName,
				Avatar:       v.Updator.Avatar,
				Introduction: v.Updator.Introduction,
			}

			resp.Items = append(resp.Items, item)
		}
	}

	param := url.Values{}
	param.Set("article_id", strconv.FormatInt(articleID, 10))
	param.Set("limit", strconv.Itoa(limit))
	param.Set("offset", strconv.Itoa(offset-limit))

	if resp.Paging.Prev, err = genURL("/api/v1/article/list/histories", param); err != nil {
		return
	}
	param.Set("offset", strconv.Itoa(offset+limit))
	if resp.Paging.Next, err = genURL("/api/v1/article/list/histories", param); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	return
}

func (p *Service) GetArticleHistoryResp(c context.Context, articleHistoryID int64) (item *model.ArticleHistoryResp, err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var v *article.ArticleHistoryResp
	if v, err = p.d.GetArticleHistory(c, &article.IDReq{ID: articleHistoryID, Aid: aid}); err != nil {
		return
	} else if v == nil {
		err = ecode.ArticleHistoryNotExist
		return
	}

	item = &model.ArticleHistoryResp{
		ID:         v.ID,
		ArticleID:  v.ArticleID,
		Seq:        int(v.Seq),
		ChangeDesc: v.ChangeDesc,
		Diff:       v.Diff,
		UpdatedAt:  v.UpdatedAt,
		CreatedAt:  v.CreatedAt,
	}

	item.Updator = &model.Creator{
		ID:           v.Updator.ID,
		UserName:     v.Updator.UserName,
		Avatar:       v.Updator.Avatar,
		Introduction: v.Updator.Introduction,
	}

	return
}
