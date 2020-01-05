package service

import (
	"valerian/app/interface/article/model"
	article "valerian/app/service/article/api"
	"valerian/library/net/http/mars"
)

func (s *Service) AppArticleCachePull(c *mars.Context, arg *model.ArgArticleAppCache) (resp *model.ArticleListCacheResp, err error) {
	var reqItems []*article.IdUpdatedItem
	for _, item := range arg.Items {
		reqItem := article.IdUpdatedItem{
			ID:        item.ID,
			UpdatedAt: item.UpdatedAt,
		}
		reqItems = append(reqItems, &reqItem)
	}
	resp = &model.ArticleListCacheResp{}
	resp.Items = []*model.ArticleResp{}
	for _, argItem := range arg.Items {
		article, _ := s.GetArticle(c, argItem.ID, 0, arg.Include)
		if article != nil && article.UpdatedAt != argItem.UpdatedAt {
			resp.Items = append(resp.Items, article)
		}
	}
	return
}

func (s *Service) AppReviseCachePull(c *mars.Context, arg *model.ArgReviseAppCache) (resp *model.ReviseDetailListCacheResp, err error) {
	var reqItems []*article.IdUpdatedItem
	for _, item := range arg.Items {
		reqItem := article.IdUpdatedItem{
			ID:        item.ID,
			UpdatedAt: item.UpdatedAt,
		}
		reqItems = append(reqItems, &reqItem)
	}
	resp = &model.ReviseDetailListCacheResp{}
	resp.Items = []*model.ReviseDetailResp{}
	for _, item := range arg.Items {
		revise, _ := s.GetRevise(c, item.ID)
		if revise != nil && revise.UpdatedAt != item.UpdatedAt {
			resp.Items = append(resp.Items, revise)
		}
	}
	return
}
