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
