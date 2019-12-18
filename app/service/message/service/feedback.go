package service

import (
	"context"
	"fmt"
	"github.com/kamilsk/retry/v4"
	"github.com/kamilsk/retry/v4/strategy"
	"github.com/nats-io/stan.go"
	"time"
	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) onFeedBackAccuseSuit(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgFeedbackAccuseSuit)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onFeedBackAccuseSuit Unmarshal failed %#v", err))
		return
	}
	var fb *model.Feedback
	action := func(c context.Context, _ uint) error {
		feedback, err := p.d.GetFeedbackByID(c, p.d.DB(), info.FeedbackID)
		if err != nil {
			return err
		}

		fb = feedback
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentReplied GetComment failed %#v", err))
		m.Ack()
		return
	}

	msgTargetType, targetDesc := convertFBTargetType(fb.TargetType)
	// 发送给举报人
	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  fb.CreatedBy,
		ActionType: model.MsgFeedbackAccuseSuit,
		ActionTime: time.Now().Unix(),
		ActionText: fmt.Sprintf(model.MsgTextFeedBackAccuseSuitToReporter, targetDesc+":"+fb.TargetDesc, fb.VerifyDesc),
		MergeCount: 1,
		TargetID:   fb.TargetID,
		TargetType: msgTargetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}
	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentReplied AddMessage failed %#v", err))
		return
	}
	p.addCache(func() {
		if _, err := p.pushSingleUser(context.Background(),
			msg.AccountID,
			msg.ID,
			fmt.Sprintf(def.PushMsgFeedBackAccuseSuitToReporter, targetDesc+":"+fb.TargetDesc, fb.VerifyDesc),
			fmt.Sprintf(def.PushMsgFeedBackAccuseSuitToReporter, targetDesc+":"+fb.TargetDesc, fb.VerifyDesc),
			"",
		); err != nil {
			log.For(context.Background()).Error(fmt.Sprintf("service.onCommentReplied Push message failed %#v", err))
		}
	})

	p.sendToBeReported(c, fb)
	m.Ack()
}

// 推送给被举报人
func (p *Service) sendToBeReported(c context.Context, fb *model.Feedback) {
	var accountId int64
	msgTargetType, _ := convertFBTargetType(fb.TargetType)
	if msgTargetType == model.TargetTypeTopic {
		if resp, err := p.d.GetTopic(c, fb.TargetID); err != nil {
			return
		} else {
			accountId = resp.Creator.ID
		}
	} else if msgTargetType == model.TargetTypeMember {
		accountId = fb.TargetID
	} else if msgTargetType == model.TargetTypeDiscussion {
		if info, err := p.d.GetDiscussion(c, fb.TargetID); err != nil {
			return
		} else {
			accountId = info.Creator.ID
		}
	} else if msgTargetType == model.TargetTypeComment {
		if comment, err := p.d.GetComment(c, fb.TargetID, false); err != nil {
			return
		} else {
			accountId = comment.Creator.ID
		}

	} else if msgTargetType == model.TargetTypeRevise {
		if revise, err := p.d.GetRevise(c, fb.TargetID, false); err != nil {
			return
		} else {
			accountId = revise.Creator.ID
		}
	} else if msgTargetType == model.TargetTypeArticle {
		if article, err := p.d.GetArticle(c, fb.TargetID, false); err != nil {
			return
		} else {
			accountId = article.Creator.ID
		}
	} else {
		return
	}

	p.publishToBeReported(c, accountId, fb)

	// 文章类型需要发给相关编辑者,主题主理人
	if msgTargetType == model.TargetTypeArticle {

	}
}

func (p *Service) publishToBeReported(c context.Context, accountId int64, fb *model.Feedback) {
	var err error
	msgTargetType, targetDesc := convertFBTargetType(fb.TargetType)
	// 发送给被举报人
	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  accountId,
		ActionType: model.MsgFeedbackAccuseSuit,
		ActionTime: time.Now().Unix(),
		ActionText: fmt.Sprintf(model.MsgTextFeedBackAccuseSuitToAuthor, targetDesc+":"+fb.TargetDesc, fb.VerifyDesc),
		MergeCount: 1,
		TargetID:   fb.TargetID,
		TargetType: msgTargetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}
	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentReplied AddMessage failed %#v", err))
		return
	}
	p.addCache(func() {
		if _, err := p.pushSingleUser(context.Background(),
			msg.AccountID,
			msg.ID,
			fmt.Sprintf(def.PushMsgFeedBackAccuseSuitToAuthor, targetDesc+":"+fb.TargetDesc, fb.VerifyDesc),
			fmt.Sprintf(def.PushMsgFeedBackAccuseSuitToAuthor, targetDesc+":"+fb.TargetDesc, fb.VerifyDesc),
			"",
		); err != nil {
			log.For(context.Background()).Error(fmt.Sprintf("service.onCommentReplied Push message failed %#v", err))
		}
	})
}

func convertFBTargetType(feedBackTargetType int32) (msgTargetType, name string) {

	switch feedBackTargetType {
	// 类型 参看interface-feedback
	//TargetTypeMember   = int32(2)
	case 2:
		return model.TargetTypeMember, "用户"
	//TargetTypeTopic    = int32(3)
	case 3:
		return model.TargetTypeTopic, "主题"
		// TargetTypeArticle  = int32(4)
	case 4:
		return model.TargetTypeArticle, "文章"
		//TargetTypeDiscuss = int32(5)
	case 5:
		return model.TargetTypeDiscussion, "讨论"
		//TargetTypeRevise = int32(6)
	case 6:
		return model.TargetTypeRevise, "修正"
		//TargetTypeComment = int32(7)
	case 7:
		return model.TargetTypeComment, "评论"
	default:
		return "", ""
	}
}

func (p *Service) onFeedBackAccuseNotSuit(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgFeedbackAccuseNotSuit)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onFeedBackAccuseNotSuit Unmarshal failed %#v", err))
		return
	}

	var fb *model.Feedback
	action := func(c context.Context, _ uint) error {
		feedback, err := p.d.GetFeedbackByID(c, p.d.DB(), info.FeedbackID)
		if err != nil {
			return err
		}

		fb = feedback
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentReplied GetComment failed %#v", err))
		m.Ack()
		return
	}

	msgTargetType, targetDesc := convertFBTargetType(fb.TargetType)
	// 发送给举报人
	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  fb.CreatedBy,
		ActionType: model.MsgFeedbackAccuseNotSuit,
		ActionTime: time.Now().Unix(),
		ActionText: fmt.Sprintf(model.MsgTextFeedBackAccuseNotSuit, targetDesc+":"+fb.TargetDesc, fb.VerifyDesc),
		MergeCount: 1,
		TargetID:   fb.TargetID,
		TargetType: msgTargetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}
	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCommentReplied AddMessage failed %#v", err))
		return
	}
	p.addCache(func() {
		if _, err := p.pushSingleUser(context.Background(),
			msg.AccountID,
			msg.ID,
			fmt.Sprintf(def.PushMsgFeedBackAccuseNotSuit, targetDesc+":"+fb.TargetDesc, fb.VerifyDesc),
			fmt.Sprintf(def.PushMsgFeedBackAccuseNotSuit, targetDesc+":"+fb.TargetDesc, fb.VerifyDesc),
			"",
		); err != nil {
			log.For(context.Background()).Error(fmt.Sprintf("service.onCommentReplied Push message failed %#v", err))
		}
	})
	m.Ack()
}
