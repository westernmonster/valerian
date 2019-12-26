package service

import (
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
)

func (p *Service) fromArticle(v *article.ArticleInfo) (item *api.TargetArticle) {
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

func fromTopic(v *model.Topic, stat *model.TopicStat, acc *account.BaseInfoReply) *api.TopicInfo {
	reply := &api.TopicInfo{
		ID:              v.ID,
		Name:            v.Name,
		Introduction:    v.Introduction,
		AllowDiscuss:    bool(v.AllowDiscuss),
		AllowChat:       bool(v.AllowChat),
		IsPrivate:       bool(v.IsPrivate),
		ViewPermission:  v.ViewPermission,
		EditPermission:  v.EditPermission,
		JoinPermission:  v.JoinPermission,
		CatalogViewType: v.CatalogViewType,
		Avatar:          v.Avatar,
		Bg:              v.Bg,
		TopicHome:       v.TopicHome,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}

	reply.Stat = &api.TopicStat{
		MemberCount:     int32(stat.MemberCount),
		DiscussionCount: int32(stat.DiscussionCount),
		ArticleCount:    int32(stat.ArticleCount),
	}

	reply.Creator = &api.Creator{
		ID:           acc.ID,
		UserName:     acc.UserName,
		Avatar:       acc.Avatar,
		Introduction: acc.Introduction,
	}

	return reply
}

func fromTopicMeta(v *model.TopicMeta) *api.TopicMetaInfo {
	reply := &api.TopicMetaInfo{
		CanFollow:    v.CanFollow,
		CanEdit:      v.CanEdit,
		Fav:          v.Fav,
		CanView:      v.CanView,
		FollowStatus: int32(v.FollowStatus),
		IsMember:     v.IsMember,
		MemberRole:   v.MemberRole,
	}

	return reply
}
