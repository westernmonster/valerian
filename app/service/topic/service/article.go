package service

import (
	article "valerian/app/service/article/api"
	"valerian/app/service/topic/api"
)

func (p *Service) FromArticle(v *article.ArticleInfo) (item *api.TargetArticle) {
	item = &api.TargetArticle{
		ID:           v.ID,
		Title:        v.Title,
		Excerpt:      v.Excerpt,
		ChangeDesc:   v.ChangeDesc,
		ImageUrls:    v.ImageUrls,
		ReviseCount:  (v.Stat.ReviseCount),
		CommentCount: (v.Stat.CommentCount),
		LikeCount:    (v.Stat.LikeCount),
		DislikeCount: (v.Stat.DislikeCount),
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.CreatedAt,
		Creator: &api.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
	}

	return
}
