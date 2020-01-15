package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/library/log"
)

func (p *Service) emitFeedBackAccuseSuit(c context.Context, feedbackId int64, actionTime int64) {
	msg := &def.MsgFeedbackAccuseSuit{
		FeedbackID: feedbackId,
		ActionTime: actionTime,
	}
	var data []byte
	var err error
	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("emitFeedBackAccuseSuit.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusFeedBackAccuseSuit, data); err != nil {
		log.For(c).Error(fmt.Sprintf("emitFeedBackAccuseSuit.Publish(), err(%+v)", err))
		return
	}
}

func (p *Service) emitFeedBackAccuseNotSuit(c context.Context, feedbackId, actionTime int64) {
	msg := &def.MsgFeedbackAccuseSuit{
		FeedbackID: feedbackId,
		ActionTime: actionTime,
	}
	var data []byte
	var err error
	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("emitFeedBackAccuseNotSuit.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusFeedBackAccuseNotSuit, data); err != nil {
		log.For(c).Error(fmt.Sprintf("emitFeedBackAccuseNotSuit.Publish(), err(%+v)", err))
		return
	}
}
