package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/library/log"
)

func (p *Service) onArticleCommented(c context.Context, commentID, aid, actionTime int64) {
	msg := &def.MsgArticleCommented{CommentID: commentID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleCommented.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusArticleCommented, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleCommented.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onReviseCommented(c context.Context, commentID, aid, actionTime int64) {
	msg := &def.MsgReviseCommented{CommentID: commentID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseCommented.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusReviseCommented, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseCommented.Publish(), err(%+v)", err))
		return
	}

	return

}

func (p *Service) onDiscussionCommented(c context.Context, commentID, aid, actionTime int64) {
	msg := &def.MsgDiscussionCommented{CommentID: commentID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionCommented.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusDiscussionCommented, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onDiscussionCommented.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onCommentReplied(c context.Context, commentID, replyCommentID, aid, actionTime int64) {
	msg := &def.MsgCommentReplied{CommentID: commentID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onCommentReplied.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusCommentReplied, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onCommentReplied.Publish(), err(%+v)", err))
		return
	}

	return
}
