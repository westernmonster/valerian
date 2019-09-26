package api

import "valerian/app/service/topic/model"

func FromTopic(v *model.Topic) *TopicInfo {
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
		CreatedBy:       v.CreatedBy,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}

	if v.Avatar != nil {
		reply.Avatar = &TopicInfo_AvatarValue{*v.Avatar}
	}

	if v.Bg != nil {
		reply.Bg = &TopicInfo_BgValue{*v.Bg}
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
