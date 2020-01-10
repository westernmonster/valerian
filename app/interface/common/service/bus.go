package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/library/log"
)

func (p *Service) onCommentLiked(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgCommentLiked{CommentID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onCommentLiked.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusCommentLiked, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onCommentLiked.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onArticleLiked(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgArticleLiked{ArticleID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleLiked.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusArticleLiked, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleLiked.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onReviseLiked(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgReviseLiked{ReviseID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseLiked.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusReviseLiked, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseLiked.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onDiscussionLiked(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgDiscussionLiked{DiscussionID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionLiked.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusDiscussionLiked, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionLiked.Publish(), err(%+v)", err))
		return
	}

	return
}

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
