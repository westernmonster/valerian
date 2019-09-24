package dao

import (
	"context"
	topic "valerian/app/service/topic/api"
)

func (p *Dao) GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error) {
	return p.topicRPC.GetTopicInfo(c, &topic.TopicReq{ID: id})
}

func (p *Dao) CheckTopicManager(c context.Context, topicID, accountID int64) (err error) {
	_, err = p.topicRPC.CheckTopicManager(c, &topic.CheckTopicManagerReq{TopicID: topicID, AccountID: accountID})
	return
}
