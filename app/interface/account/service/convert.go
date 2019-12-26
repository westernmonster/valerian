package service

import (
	"valerian/app/interface/account/model"
	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
)

func (p *Service) fromComment(v *comment.CommentInfo) (item *model.TargetComment) {
	item = &model.TargetComment{
		ID:            v.ID,
		Type:          v.TargetType,
		Excerpt:       v.Content,
		CreatedAt:     v.CreatedAt,
		ResourceID:    v.ResourceID,
		ChildrenCount: v.Stat.ChildrenCount,
		LikeCount:     v.Stat.LikeCount,
	}

	return
}

func (p *Service) fromDiscussion(v *discuss.DiscussionInfo) (item *model.TargetDiscuss) {
	item = &model.TargetDiscuss{
		ID:           v.ID,
		Excerpt:      v.Excerpt,
		CommentCount: (v.Stat.CommentCount),
		LikeCount:    (v.Stat.LikeCount),
		DislikeCount: (v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Title:     v.Title,
	}

	if v.ImageUrls == nil {
		item.ImageUrls = make([]string, 0)
	} else {
		item.ImageUrls = v.ImageUrls
	}

	return
}

func (p *Service) fromRevise(v *article.ReviseInfo) (item *model.TargetRevise) {
	item = &model.TargetRevise{
		ID:           v.ID,
		Title:        v.Title,
		Excerpt:      v.Excerpt,
		CommentCount: (v.Stat.CommentCount),
		LikeCount:    (v.Stat.LikeCount),
		DislikeCount: (v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
	if v.ImageUrls == nil {
		item.ImageUrls = make([]string, 0)
	} else {
		item.ImageUrls = v.ImageUrls
	}

	return
}

func (p *Service) fromArticle(v *article.ArticleInfo) (item *model.TargetArticle) {
	item = &model.TargetArticle{
		ID:           v.ID,
		Title:        v.Title,
		Excerpt:      v.Excerpt,
		ChangeDesc:   v.ChangeDesc,
		ReviseCount:  (v.Stat.ReviseCount),
		CommentCount: (v.Stat.CommentCount),
		LikeCount:    (v.Stat.LikeCount),
		DislikeCount: (v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
	if v.ImageUrls == nil {
		item.ImageUrls = make([]string, 0)
	} else {
		item.ImageUrls = v.ImageUrls
	}

	return
}

func (p *Service) fromTopic(v *topic.TopicInfo) (item *model.TargetTopic) {
	item = &model.TargetTopic{
		ID:              v.ID,
		Name:            v.Name,
		Introduction:    v.Introduction,
		MemberCount:     (v.Stat.MemberCount),
		DiscussionCount: (v.Stat.DiscussionCount),
		ArticleCount:    (v.Stat.ArticleCount),
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Avatar:    v.Avatar,
	}

	return
}
