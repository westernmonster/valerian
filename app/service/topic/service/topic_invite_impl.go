package service

import (
	"context"
	"valerian/app/service/topic/model"
)

func (p *Service) HasInvited(c context.Context, accountID, topicID int64) (hasInvited bool, err error) {
	var req *model.TopicInviteRequest
	if req, err = p.d.GetTopicInviteRequestByCond(c, p.d.DB(), map[string]interface{}{
		"topic_id":   topicID,
		"account_id": accountID,
	}); err != nil {
		return
	} else if req != nil {
		hasInvited = true
	}
	return
}
