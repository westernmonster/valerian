package service

import (
	"context"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
)

func (p *Service) checkViewPermission(c context.Context, aid, topicID int64) (err error) {
	var canView bool
	if canView, err = p.d.CanView(c, &topic.TopicReq{ID: topicID, Aid: aid}); err != nil {
		return
	}

	if !canView {
		err = ecode.NoTopicViewPermission
		return
	}

	return
}
