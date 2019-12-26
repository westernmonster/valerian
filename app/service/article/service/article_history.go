package service

import (
	"context"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/ecode"
)

// GetArticleHistoriesPaged 获取文章修改历史记录
func (p *Service) GetArticleHistoriesPaged(c context.Context, articleID int64, limit, offset int) (resp *api.ArticleHistoryListResp, err error) {
	var data []*model.ArticleHistory
	if data, err = p.d.GetArticleHistoriesPaged(c, p.d.DB(), articleID, limit, offset); err != nil {
		return
	}

	resp = &api.ArticleHistoryListResp{
		Items: make([]*api.ArticleHistoryItem, len(data)),
	}

	for i, v := range data {
		item := &api.ArticleHistoryItem{
			ID:         v.ID,
			ArticleID:  v.ArticleID,
			Seq:        v.Seq,
			ChangeDesc: v.ChangeDesc,
			UpdatedAt:  v.UpdatedAt,
			CreatedAt:  v.CreatedAt,
		}

		var acc *model.Account
		if acc, err = p.getAccount(c, p.d.DB(), v.UpdatedBy); err != nil {
			return
		}
		item.Updator = &api.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		resp.Items[i] = item
	}

	return
}

// GetArticleHistory 获取指定文章修改历史记录
func (p *Service) GetArticleHistory(c context.Context, articleHistoryID int64) (item *api.ArticleHistoryResp, err error) {
	var v *model.ArticleHistory
	if v, err = p.d.GetArticleHistoryByID(c, p.d.DB(), articleHistoryID); err != nil {
		return
	} else if v == nil {
		err = ecode.ArticleHistoryNotExist
		return
	}

	item = &api.ArticleHistoryResp{
		ID:         v.ID,
		ArticleID:  v.ArticleID,
		Seq:        v.Seq,
		ChangeDesc: v.ChangeDesc,
		Diff:       v.Diff,
		UpdatedAt:  v.UpdatedAt,
		CreatedAt:  v.CreatedAt,
	}

	var acc *model.Account
	if acc, err = p.getAccount(c, p.d.DB(), v.UpdatedBy); err != nil {
		return
	}
	item.Updator = &api.Creator{
		ID:           acc.ID,
		UserName:     acc.UserName,
		Avatar:       acc.Avatar,
		Introduction: acc.Introduction,
	}

	return
}
