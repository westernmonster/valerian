package service

import (
	"context"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/log"
)

func (p *Service) onTopicAdded(c context.Context, id int64) {
	msg := &model.MsgTopicAdded{ID: id}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicAdded.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusTopicAdded, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicAdded.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicDeleted(c context.Context, id int64) {
	msg := &model.MsgTopicDeleted{ID: id}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicDeleted.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusTopicDeleted, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicDeleted.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicFollowed(c context.Context, aid, topicID int64) {
	msg := &model.MsgTopicFollowed{AccountID: aid, TopicID: topicID}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowed.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusTopicFollowed, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowed.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicLeaved(c context.Context, aid, topicID int64) {
	msg := &model.MsgTopicLeaved{AccountID: aid, TopicID: topicID}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicLeaved.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusTopicLeaved, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicLeaved.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicFollowRequested(c context.Context, id int64) {
	msg := &model.MsgTopicFollowRequested{ID: id}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowRequested.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusTopicFollowRequested, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowRequested.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicFollowRejected(c context.Context, id int64) {
	msg := &model.MsgTopicFollowRejected{ID: id}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowRejected.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusTopicFollowRejected, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowRejected.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicFollowApproved(c context.Context, id int64) {
	msg := &model.MsgTopicFollowApproved{ID: id}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowApproved.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusTopicFollowApproved, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowApproved.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicInviteSent(c context.Context, id int64) {
	msg := &model.MsgTopicInviteSent{ID: id}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicInviteSent.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusTopicInviteSent, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicInviteSent.Publish(), err(%+v)", err))
		return
	}

	return
}
