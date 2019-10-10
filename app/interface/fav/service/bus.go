package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/library/log"
)

func (p *Service) onTopicFaved(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicFaved{TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFaved.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicFaved, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFaved.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onArticleFaved(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgArticleFaved{ArticleID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleFaved.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusArticleFaved, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleFaved.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onReviseFaved(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgReviseFaved{ReviseID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseFaved.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusReviseFaved, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseFaved.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onDiscussionFaved(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgDiscussionFaved{DiscussionID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionFaved.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusDiscussionFaved, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionFaved.Publish(), err(%+v)", err))
		return
	}

	return
}
