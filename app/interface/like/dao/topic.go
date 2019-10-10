package dao

import (
	"context"
	topic "valerian/app/service/topic/api"
)

func (p *Dao) GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error) {
	return p.topicRPC.GetTopicInfo(c, &topic.TopicReq{ID: id})
}
