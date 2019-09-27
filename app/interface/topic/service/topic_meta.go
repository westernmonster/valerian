package service

import (
	"context"

	"valerian/app/interface/topic/model"
	stopic "valerian/app/service/topic/api"
)

func (p *Service) GetTopicMeta(c context.Context, aid, topicID int64) (meta *model.TopicMeta, err error) {
	var m *stopic.TopicMetaInfo
	if m, err = p.d.GetTopicMeta(c, aid, topicID); err != nil {
		return
	}

	meta = &model.TopicMeta{
		CanFollow:    m.CanFollow,
		CanEdit:      m.CanEdit,
		Fav:          m.Fav,
		CanView:      m.CanView,
		FollowStatus: int(m.FollowStatus),
		IsMember:     m.CanView,
		MemberRole:   m.MemberRole,
	}

	return
}
