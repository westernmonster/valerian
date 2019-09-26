package service

import (
	"context"
	"valerian/app/interface/article/model"
	account "valerian/app/service/account/api"
	"valerian/library/ecode"
)

func (p *Service) GetArticleHistoriesResp(c context.Context, articleID int64) (items []*model.ArticleHistoryResp, err error) {
	var node = p.d.DB()
	var addCache = true

	if items, err = p.d.ArticleHistoryCache(c, articleID); err != nil {
		addCache = false
	} else if items != nil {
		for _, v := range items {
			var account *account.BaseInfoReply
			if account, err = p.d.GetAccountBaseInfo(c, v.UpdatedBy); err != nil {
				return
			}
			v.Updator = &model.Creator{
				ID:       account.ID,
				UserName: account.UserName,
				Avatar:   account.Avatar,
			}
			intro := account.GetIntroductionValue()
			v.Updator.Introduction = &intro
		}
		return
	}

	var data []*model.ArticleHistory
	if data, err = p.d.GetArticleHistoriesByCond(c, node, map[string]interface{}{"article_id": articleID}); err != nil {
		return
	}

	items = make([]*model.ArticleHistoryResp, 0)
	for _, v := range data {
		item := &model.ArticleHistoryResp{
			ID:         v.ID,
			ArticleID:  v.ArticleID,
			Seq:        v.Seq,
			ChangeDesc: v.ChangeDesc,
			// Content:     &v.Content,
			// ContentText: &v.ContentText,
			// Diff:      &v.Diff,
			UpdatedBy: v.UpdatedBy,
			UpdatedAt: v.UpdatedAt,
			CreatedAt: v.CreatedAt,
		}

		items = append(items, item)
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleHistoryCache(context.TODO(), articleID, items)
		})
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
		ID:          v.ID,
		ArticleID:   v.ArticleID,
		Seq:         v.Seq,
		ChangeDesc:  v.ChangeDesc,
		Content:     &v.Content,
		ContentText: &v.ContentText,
		Diff:        &v.Diff,
		UpdatedBy:   v.UpdatedBy,
		UpdatedAt:   v.UpdatedAt,
		CreatedAt:   v.CreatedAt,
	}

	return
}
