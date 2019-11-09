package dao

import (
	"context"
	topic "valerian/app/service/topic/api"
)

func (p *Dao) GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error) {
	return p.topicRPC.GetTopicInfo(c, &topic.TopicReq{ID: id})
}

func (p *Dao) GetTopicMemberRole(c context.Context, topicID, accountID int64) (resp *topic.MemberRoleReply, err error) {
	return p.topicRPC.GetTopicMemberRole(c, &topic.TopicMemberRoleReq{TopicID: topicID, AccountID: accountID})
}
