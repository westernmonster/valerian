package dao

import (
	"context"
	stopic "valerian/app/service/topic/api"
)

func (p *Dao) GetTopicMeta(c context.Context, aid, topicID int64) (info *stopic.TopicMetaInfo, err error) {
	return p.topicRPC.GetTopicMeta(c, &stopic.TopicMetaReq{AccountID: aid, TopicID: topicID})
}
