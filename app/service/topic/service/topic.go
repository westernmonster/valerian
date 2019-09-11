package service

import (
	"context"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) checkTopic(c context.Context, node sqalx.Node, topicID int64) (err error) {
	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if t == nil {
		return ecode.TopicNotExist
	}

	return
}
