package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/library/log"
)

// func (p *Service) onTopicLiked(c context.Context, topicID, aid, actionTime int64) {
// 	msg := &def.MsgTopicLiked{TopicID: topicID, ActorID: aid, ActionTime: actionTime}

// 	var data []byte
// 	var err error

// 	if data, err = msg.Marshal(); err != nil {
// 		log.For(c).Error(fmt.Sprintf("onTopicLiked.Marshal(), err(%+v)", err))
// 		return
// 	}

// 	if err = p.mq.Publish(def.BusTopicLiked, data); err != nil {
// 		log.For(c).Error(fmt.Sprintf("onTopicLiked.Publish(), err(%+v)", err))
// 		return
// 	}

// 	return
// }

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
