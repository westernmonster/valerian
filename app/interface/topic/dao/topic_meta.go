package dao

import (
	"context"
	"fmt"

	stopic "valerian/app/service/topic/api"
	"valerian/library/log"
)

func (p *Dao) GetTopicMeta(c context.Context, aid, topicID int64) (info *stopic.TopicMetaInfo, err error) {
	if info, err = p.topicRPC.GetTopicMeta(c, &stopic.TopicMetaReq{AccountID: aid, TopicID: topicID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMeta err(%+v)", err))
	}
	return
}
