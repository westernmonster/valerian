package api

import (
	"valerian/app/service/topic/model"

	account "valerian/app/service/account/api"
)

func FromTopic(v *model.Topic, stat *model.TopicStat, acc *account.BaseInfoReply) *TopicInfo {
	reply := &TopicInfo{
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
		TopicHome:       v.TopicHome,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}

	reply.Stat = &TopicStat{
		MemberCount:     int32(stat.MemberCount),
		DiscussionCount: int32(stat.DiscussionCount),
		ArticleCount:    int32(stat.ArticleCount),
	}

	if v.Avatar != nil {
		reply.Avatar = &TopicInfo_AvatarValue{*v.Avatar}
	}

	if v.Bg != nil {
		reply.Bg = &TopicInfo_BgValue{*v.Bg}
	}

	reply.Creator = &Creator{
		ID:       acc.ID,
		UserName: acc.UserName,
		Avatar:   acc.Avatar,
	}

	if acc.Introduction != nil {
		reply.Creator.Introduction = &Creator_IntroductionValue{acc.GetIntroductionValue()}
	}

	return reply
}

func FromTopicMeta(v *model.TopicMeta) *TopicMetaInfo {
	reply := &TopicMetaInfo{
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
