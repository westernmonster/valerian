package dao

import (
	"context"
	"fmt"
	topic "valerian/app/service/topic/api"
	"valerian/library/log"
)

func (p *Dao) GetTopic(c context.Context, id int64) (resp *topic.TopicInfo, err error) {
	if resp, err = p.topicRPC.GetTopicInfo(c, &topic.TopicReq{ID: id}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserTopicsPaged error(%+v), id(%d) ", err, id))
	}
	return
}

func (p *Dao) GetTopicMemberRole(c context.Context, topicID, accountID int64) (resp *topic.MemberRoleReply, err error) {
	if resp, err = p.topicRPC.GetTopicMemberRole(c, &topic.TopicMemberRoleReq{TopicID: topicID, AccountID: accountID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMemberRole error(%+v), topic_id(%d) aid(%d)", err, topicID, accountID))
	}
	return
}
