package dao

import (
	"context"
	"fmt"
	topicFeed "valerian/app/service/topic-feed/api"
	"valerian/library/log"
)

func (p *Dao) GetTopicFeedPaged(c context.Context, topicID int64, limit, offset int) (info *topicFeed.TopicFeedResp, err error) {
	if info, err = p.topicFeedRPC.GetTopicFeedPaged(c, &topicFeed.TopicFeedReq{TopicID: topicID, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicFeedPaged err(%+v)", err))
	}
	return
}
