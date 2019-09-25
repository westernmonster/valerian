package dao

import (
	"context"
	feed "valerian/app/service/feed/api"
)

func (p *Dao) GetTopicFeedPaged(c context.Context, topicID int64, limit, offset int) (info *feed.TopicFeedResp, err error) {
	return p.feedRPC.GetTopicFeedPaged(c, &feed.TopicFeedReq{TopicID: topicID, Limit: int32(limit), Offset: int32(offset)})
}
