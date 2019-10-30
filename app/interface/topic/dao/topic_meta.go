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

func (p *Dao) CreateTopic(c context.Context, arg *stopic.ArgCreateTopic) (id int64, err error) {
	var idRet *stopic.IDResp
	if idRet, err = p.topicRPC.CreateTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CreateTopic err(%+v) arg(%+v)", err, arg))
		return
	}
	return idRet.ID, nil
}

func (p *Dao) UpdateTopic(c context.Context, arg *stopic.ArgUpdateTopic) (err error) {
	if _, err = p.topicRPC.UpdateTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopic err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) DelTopic(c context.Context, arg *stopic.IDReq) (err error) {
	if _, err = p.topicRPC.DelTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopic err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) GetTopicResp(c context.Context, arg *stopic.IDReq) (resp *stopic.TopicResp, err error) {
	if resp, err = p.topicRPC.GetTopicResp(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicResp err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) ChangeOwner(c context.Context, arg *stopic.ArgChangeOwner) (err error) {
	if _, err = p.topicRPC.ChangeOwner(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.ChangeOwner err(%+v) arg(%+v)", err, arg))
	}
	return
}
