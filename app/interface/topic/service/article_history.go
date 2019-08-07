package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
)

func (p *Service) GetArticleHistoriesResp(c context.Context, node sqalx.Node, articleVersionID int64) (items []*model.ArticleHistoryResp, err error) {
	var addCache = true

	if items, err = p.d.ArticleHistoryCache(c, articleVersionID); err != nil {
		addCache = false
	} else if items != nil {
		for _, v := range items {
			if v.Updator, err = p.getBasicAccountResp(c, node, v.UpdatedBy); err != nil {
				return
			}
		}
		return
	}

	var data []*model.ArticleHistory
	if data, err = p.d.GetArticleHistories(c, node, articleVersionID); err != nil {
		return
	}

	items = make([]*model.ArticleHistoryResp, 0)
	for _, v := range data {
		item := &model.ArticleHistoryResp{
			ID:               v.ID,
			ArticleVersionID: v.ArticleVersionID,
			Seq:              v.Seq,
			ChangeID:         v.ChangeID,
			Description:      v.Description,
			Content:          &v.Content,
			ContentText:      &v.ContentText,
			Diff:             &v.Diff,
			UpdatedBy:        v.UpdatedBy,
			UpdatedAt:        v.UpdatedAt,
			CreatedAt:        v.CreatedAt,
		}

		items = append(items, item)
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleHistoryCache(context.TODO(), articleVersionID, items)
		})
	}

	return
}
