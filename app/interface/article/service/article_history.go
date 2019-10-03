package service

import (
	"context"
	"valerian/app/interface/article/model"
	account "valerian/app/service/account/api"
	"valerian/library/ecode"
)

func (p *Service) GetArticleHistoriesResp(c context.Context, articleID int64, offset, limit int) (resp *model.ArticleHistoryListResp, err error) {
	var data []*model.ArticleHistory
	if data, err = p.d.GetArticleHistoriesPaged(c, p.d.DB(), articleID, offset, limit); err != nil {
		return
	}

	resp = &model.ArticleHistoryListResp{
		Items:  make([]*model.ArticleHistoryItem, len(data)),
		Paging: &model.Paging{},
	}

	for i, v := range data {
		item := &model.ArticleHistoryItem{
			ID:         v.ID,
			ArticleID:  v.ArticleID,
			Seq:        v.Seq,
			ChangeDesc: v.ChangeDesc,
			UpdatedAt:  v.UpdatedAt,
			CreatedAt:  v.CreatedAt,
		}

		var account *account.BaseInfoReply
		if account, err = p.d.GetAccountBaseInfo(c, v.UpdatedBy); err != nil {
			return
		}
		item.Updator = &model.Creator{
			ID:       account.ID,
			UserName: account.UserName,
			Avatar:   account.Avatar,
		}
		intro := account.GetIntroductionValue()
		item.Updator.Introduction = &intro

		resp.Items[i] = item
	}

	return
}

func (p *Service) GetArticleHistoryResp(c context.Context, articleHistoryID int64) (item *model.ArticleHistoryResp, err error) {
	var v *model.ArticleHistory
	if v, err = p.d.GetArticleHistoryByID(c, p.d.DB(), articleHistoryID); err != nil {
		return
	} else if v == nil {
		err = ecode.ArticleHistoryNotExist
		return
	}

	item = &model.ArticleHistoryResp{
		ID:         v.ID,
		ArticleID:  v.ArticleID,
		Seq:        v.Seq,
		ChangeDesc: v.ChangeDesc,
		Diff:       &v.Diff,
		UpdatedAt:  v.UpdatedAt,
		CreatedAt:  v.CreatedAt,
	}

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, v.UpdatedBy); err != nil {
		return
	}
	item.Updator = &model.Creator{
		ID:       account.ID,
		UserName: account.UserName,
		Avatar:   account.Avatar,
	}
	intro := account.GetIntroductionValue()
	item.Updator.Introduction = &intro

	return
}
