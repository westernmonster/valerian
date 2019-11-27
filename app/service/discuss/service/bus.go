package service

import (
	"context"
	"fmt"

	"valerian/app/service/feed/def"
	"valerian/library/log"
)

func (p *Service) onDiscussionAdded(c context.Context, discussionID, aid, actionTime int64) {
	msg := &def.MsgDiscussionAdded{DiscussionID: discussionID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionAdded.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusDiscussionAdded, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionAdded.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onDiscussionUpdated(c context.Context, discussionID, aid, actionTime int64) {
	msg := &def.MsgDiscussionUpdated{DiscussionID: discussionID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionUpdated.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusDiscussionUpdated, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionUpdated.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onDiscussionDeleted(c context.Context, discussionID, aid, actionTime int64) {
	msg := &def.MsgDiscussionDeleted{DiscussionID: discussionID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionDeleted.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusDiscussionDeleted, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionDeleted.Publish(), err(%+v)", err))
		return
	}

	return
}
