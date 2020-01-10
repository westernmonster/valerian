package dao

import (
	"context"
	"fmt"
	topic "valerian/app/service/topic/api"
	"valerian/library/log"
)

func (p *Dao) GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error) {
	if info, err = p.topicRPC.GetTopicInfo(c, &topic.TopicReq{ID: id}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicInfo err(%+v) id(%d)", err, id))
	}
	return
}

func (p *Dao) GetTopicMeta(c context.Context, aid, topicID int64) (info *topic.TopicMetaInfo, err error) {
	if info, err = p.topicRPC.GetTopicMeta(c, &topic.TopicMetaReq{AccountID: aid, TopicID: topicID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMeta err(%+v)", err))
	}
	return
}

func (p *Dao) GetRecommendTopicsIDs(c context.Context) (resp *topic.IDsResp, err error) {
	if resp, err = p.topicRPC.GetRecommendTopicsIDs(c, &topic.EmptyStruct{}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecommendTopicsIDs err(%+v) ", err))
	}
	return
}

func (p *Dao) GetRecommendAuthTopicsIDs(c context.Context, arg *topic.IDsReq) (resp *topic.IDsResp, err error) {
	if resp, err = p.topicRPC.GetRecommendAuthTopicsIDs(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecommendAuthTopicsIDs err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) GetRecommendMemberIDs(c context.Context, arg *topic.IDsReq) (resp *topic.IDsResp, err error) {
	if resp, err = p.topicRPC.GetRecommendMemberIDs(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecommendMemberIDs err(%+v) arg(%+v)", err, arg))
	}
	return
}
