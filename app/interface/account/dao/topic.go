package dao

import (
	"context"
	"fmt"
	topic "valerian/app/service/topic/api"
	"valerian/library/log"
)

func (p *Dao) GetUserTopicsPaged(c context.Context, aid int64, limit, offset int) (resp *topic.UserTopicsResp, err error) {
	if resp, err = p.topicRPC.GetUserTopicsPaged(c, &topic.UserTopicsReq{AccountID: aid, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserTopicsPaged error(%+v), aid(%d) limit(%d) offset(%d)", err, aid, limit, offset))
	}

	return
}

func (p *Dao) GetTopic(c context.Context, id int64) (resp *topic.TopicInfo, err error) {
	if resp, err = p.topicRPC.GetTopicInfo(c, &topic.TopicReq{ID: id}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserTopicsPaged error(%+v), id(%d) ", err, id))
	}

	return
}

func (p *Dao) GetFollowedTopicIDsPaged(c context.Context, aid int64, limit, offset int) (resp *topic.IDsResp, err error) {
	if resp, err = p.topicRPC.GetFollowedTopicIDsPaged(c, &topic.UserTopicsReq{AccountID: aid, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFollowedTopicIDsPaged error(%+v), aid(%d) limit(%d) offset(%d)", err, aid, limit, offset))
	}

	return
}

func (p *Dao) GetManageTopicIDsPaged(c context.Context, aid int64, limit, offset int) (resp *topic.IDsResp, err error) {
	if resp, err = p.topicRPC.GetManageTopicIDsPaged(c, &topic.UserTopicsReq{AccountID: aid, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetManageTopicIDsPaged error(%+v), aid(%d) limit(%d) offset(%d)", err, aid, limit, offset))
	}

	return
}
