package service

import (
	"context"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) GetTopic(c context.Context, topicID int64) (item *model.Topic, err error) {
	return p.getTopic(c, p.d.DB(), topicID)
}

func (p *Service) getTopic(c context.Context, node sqalx.Node, topicID int64) (item *model.Topic, err error) {
	var addCache = true
	if item, err = p.d.TopicCache(c, topicID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if item == nil {
		return nil, ecode.TopicNotExist
	}

	if addCache {
		p.addCache(func() {
			p.d.SetTopicCache(context.TODO(), item)
		})
	}

	return
}

func (p *Service) CheckTopicManager(c context.Context, topicID, aid int64) (err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member == nil {
		err = ecode.NotTopicMember
		return
	} else if member.Role == model.MemberRoleUser {
		err = ecode.NotTopicAdmin
		return
	}

	return nil
}
