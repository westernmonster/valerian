package service

import (
	"valerian/app/interface/article/model"
	article "valerian/app/service/article/api"
	"valerian/library/database/sqlx/types"
	"valerian/library/net/http/mars"
)

func (s *Service) AppArticleCachePull(c *mars.Context, arg *model.ArgArticleAppCache) (resp []model.Article, err error) {
	var reqItems []*article.IdUpdatedItem
	for _, item := range arg.Items {
		reqItem := article.IdUpdatedItem{
			ID:        item.ID,
			UpdatedAt: item.UpdatedAt,
		}
		reqItems = append(reqItems, &reqItem)
	}
	if results, err := s.d.PullArticleAppCache(c, &article.IdUpdatedReq{Items: reqItems}); err != nil {
		return nil, err
	} else {
		for _, result := range results.Items {
			article := model.Article{
				ID:             result.ID,
				Title:          result.Title,
				Content:        result.Content,
				ContentText:    result.ContentText,
				DisableRevise:  types.BitBool(result.DisableRevise),
				DisableComment: types.BitBool(result.DisableComment),
				CreatedBy:      result.CreatedBy,
				CreatedAt:      result.CreatedAt,
				UpdatedAt:      result.UpdatedAt,
			}
			resp = append(resp, article)
		}
		return resp, err
	}
}

func (s *Service) AppReviseCachePull(c *mars.Context, arg *model.ArgReviseAppCache) (resp []model.Revise, err error) {
	var reqItems []*article.IdUpdatedItem
	for _, item := range arg.Items {
		reqItem := article.IdUpdatedItem{
			ID:        item.ID,
			UpdatedAt: item.UpdatedAt,
		}
		reqItems = append(reqItems, &reqItem)
	}
	if results, err := s.d.PullReviseAppCache(c, &article.IdUpdatedReq{Items: reqItems}); err != nil {
		return nil, err
	} else {
		for _, result := range results.Items {
			revise := model.Revise{
				ID:          result.ID,
				ArticleID:   result.ArticleID,
				Title:       result.Title,
				Content:     result.Content,
				ContentText: result.ContentText,

				CreatedAt: result.CreatedAt,
				UpdatedAt: result.UpdatedAt,
			}
			if result.Creator != nil {
				revise.CreatedBy = result.Creator.ID
			}
			resp = append(resp, revise)
		}
		return resp, err
	}

}
