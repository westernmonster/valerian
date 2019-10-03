package dao

import (
	"context"
	"fmt"
	feed "valerian/app/service/feed/api"
	"valerian/library/log"
)

func (p *Dao) GetTopicFeedPaged(c context.Context, topicID int64, limit, offset int) (info *feed.TopicFeedResp, err error) {
	if info, err = p.feedRPC.GetTopicFeedPaged(c, &feed.TopicFeedReq{TopicID: topicID, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicFeedPaged err(%+v)", err))
	}
	return
}
