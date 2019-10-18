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

func (p *Dao) GetBelongsTopicIDs(c context.Context, aid int64) (resp *topic.IDsResp, err error) {
	if resp, err = p.topicRPC.GetBelongsTopicIDs(c, &topic.AidReq{AccountID: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetBelongsTopicIDs error(%+v), aid(%d) ", err, aid))
	}

	return
}

func (p *Dao) GetTopicMemberIDs(c context.Context, topicID int64) (resp *topic.IDsResp, err error) {
	if resp, err = p.topicRPC.GetTopicMemberIDs(c, &topic.TopicReq{ID: topicID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMemberIDs error(%+v), topic_id(%d) ", err, topicID))
	}

	return
}
