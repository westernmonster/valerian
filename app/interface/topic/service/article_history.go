package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
)

func (p *Service) getArticleHistoriesResp(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleHistoryResp, err error) {
	var addCache = true

	if items, err = p.d.ArticleHistoryCache(c, articleID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	var data []*model.ArticleHistory
	if data, err = p.d.GetArticleHistories(c, node, articleID); err != nil {
		return
	}

	items = make([]*model.ArticleHistoryResp, 0)
	for _, v := range data {
		item := &model.ArticleHistoryResp{
			ID: v.ID,
			// ArticleID:   v.ArticleID,
			Seq:         v.Seq,
			CreatedAt:   v.CreatedAt,
			Description: v.Description,
			Content:     &v.Content,
			ContentText: &v.ContentText,
			Diff:        &v.Diff,
		}

		if acc, e := p.getAccountByID(c, node, v.UpdatedBy); e != nil {
			return nil, e
		} else {
			item.Updator = &model.BasicAccountResp{
				Avatar:    acc.Avatar,
				UserName:  acc.UserName,
				AccountID: acc.ID,
			}
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
