package service

import (
	"context"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) getFeedback(c context.Context, node sqalx.Node, id int64) (item *model.Feedback, err error) {
	if item, err = p.d.GetFeedbackByID(c, node, id); err != nil {
		return
	} else if item == nil {
		return nil, ecode.FeedbackNotExist
	}
	return
}

//func (p *Service) onFeedBackAccuseSuit(m *stan.Msg) {
//	var err error
//	c := context.Background()
//	info := new(def.MsgFeedbackAccuseSuit)
//	if err = info.Unmarshal(m.Data); err != nil {
//		log.For(c).Error(fmt.Sprintf("service.onFeedBackAccuseSuit Unmarshal failed %#v", err))
//		return
//	}

//	var tx sqalx.Node
//	if tx, err = p.d.DB().Beginx(c); err != nil {
//		PromError("feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
//		return
//	}

//	defer func() {
//		if err != nil {
//			if err1 := tx.Rollback(); err1 != nil {
//				PromError("feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
//			}
//			return
//		}
//	}()

//	var feedback *model.Feedback
//	if feedback, err = p.getFeedback(c, tx, info.FeedbackID); err != nil {
//		if ecode.IsNotExistEcode(err) {
//			m.Ack()
//			return
//		}
//		PromError("message: GetFeedback", "GetFeedback(), id(%d),error(%+v)", info.FeedbackID, err)
//		return
//	}

//	msgTargetType, targetDesc := convertFBTargetType(feedback.TargetType)
//	// 发送给举报人
//	msg := &model.Message{
//		ID:         gid.NewID(),
//		AccountID:  feedback.CreatedBy,
//		ActionType: model.MsgFeedbackAccuseSuit,
//		ActionTime: time.Now().Unix(),
//		ActionText: fmt.Sprintf(model.MsgTextFeedBackAccuseSuitToReporter, targetDesc+":"+feedback.FeedbackDesc, feedback.VerifyDesc),
//		MergeCount: 1,
//		ActorType:  model.ActorTypeUser,
//		Actors:     strconv.FormatInt(feedback.CreatedBy, 10),
//		TargetID:   feedback.TargetID,
//		TargetType: msgTargetType,
//		CreatedAt:  time.Now().Unix(),
//		UpdatedAt:  time.Now().Unix(),
//	}
//	if err = p.d.AddMessage(c, tx, msg); err != nil {
//		PromError("feed: AddMessage", "AddMessage(), feed(%+v),error(%+v)", msg, err)
//		return
//	}

//	if err = tx.Commit(); err != nil {
//		PromError("feed: tx.Commit", "tx.Commit(), error(%+v)", err)
//		return
//	}

//	m.Ack()

//	p.addCache(func() {
//		if _, err := p.pushSingleUser(context.Background(),
//			msg.AccountID,
//			msg.ID,
//			fmt.Sprintf(def.PushMsgFeedBackAccuseSuitToReporter, targetDesc+":"+feedback.FeedbackDesc, feedback.VerifyDesc),
//			fmt.Sprintf(def.PushMsgFeedBackAccuseSuitToReporter, targetDesc+":"+feedback.FeedbackDesc, feedback.VerifyDesc),
//			"",
//		); err != nil {
//			log.For(context.Background()).Error(fmt.Sprintf("service.onCommentReplied Push message failed %#v", err))
//		}

//		p.sendToBeReported(context.Background(), feedback)
//	})

//}

//// 推送给被举报人
//func (p *Service) sendToBeReported(c context.Context, fb *model.Feedback) {
//	c = sqalx.NewContext(c, true)
//	var accountId int64
//	msgTargetType, _ := convertFBTargetType(fb.TargetType)
//	if msgTargetType == model.TargetTypeTopic {
//		if resp, err := p.getTopic(c, p.d.DB(), fb.TargetID); err != nil {
//			return
//		} else {
//			accountId = resp.CreatedBy
//		}
//	} else if msgTargetType == model.TargetTypeMember {
//		accountId = fb.TargetID
//	} else if msgTargetType == model.TargetTypeDiscussion {
//		if info, err := p.d.GetDiscussion(c, fb.TargetID); err != nil {
//			return
//		} else {
//			accountId = info.Creator.ID
//		}
//	} else if msgTargetType == model.TargetTypeComment {
//		if comment, err := p.d.GetComment(c, fb.TargetID, false); err != nil {
//			return
//		} else {
//			accountId = comment.Creator.ID
//		}

//	} else if msgTargetType == model.TargetTypeRevise {
//		if revise, err := p.d.GetRevise(c, fb.TargetID, false); err != nil {
//			return
//		} else {
//			accountId = revise.Creator.ID
//		}
//	} else if msgTargetType == model.TargetTypeArticle {
//		if article, err := p.d.GetArticle(c, fb.TargetID, false); err != nil {
//			return
//		} else {
//			accountId = article.Creator.ID
//		}
//	} else {
//		return
//	}

//	p.publishToBeReported(c, accountId, fb)

//	// 文章类型需要发给相关编辑者,主题主理人
//	if msgTargetType == model.TargetTypeArticle {

//	}
//}

//func (p *Service) publishToBeReported(c context.Context, accountId int64, fb *model.Feedback) {
//	var err error
//	msgTargetType, targetDesc := convertFBTargetType(fb.TargetType)
//	// 发送给被举报人
//	msg := &model.Message{
//		ID:         gid.NewID(),
//		AccountID:  accountId,
//		ActionType: model.MsgFeedbackAccuseSuit,
//		ActionTime: time.Now().Unix(),
//		ActionText: fmt.Sprintf(model.MsgTextFeedBackAccuseSuitToAuthor, targetDesc+":"+fb.FeedbackDesc, fb.VerifyDesc),
//		MergeCount: 1,
//		ActorType:  model.ActorTypeUser,
//		Actors:     strconv.FormatInt(fb.CreatedBy, 10),
//		TargetID:   fb.TargetID,
//		TargetType: msgTargetType,
//		CreatedAt:  time.Now().Unix(),
//		UpdatedAt:  time.Now().Unix(),
//	}
//	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
//		log.For(c).Error(fmt.Sprintf("service.onCommentReplied AddMessage failed %#v", err))
//		return
//	}
//	p.addCache(func() {
//		if _, err := p.pushSingleUser(context.Background(),
//			msg.AccountID,
//			msg.ID,
//			fmt.Sprintf(def.PushMsgFeedBackAccuseSuitToAuthor, targetDesc+":"+fb.FeedbackDesc, fb.VerifyDesc),
//			fmt.Sprintf(def.PushMsgFeedBackAccuseSuitToAuthor, targetDesc+":"+fb.FeedbackDesc, fb.VerifyDesc),
//			"",
//		); err != nil {
//			log.For(context.Background()).Error(fmt.Sprintf("service.onCommentReplied Push message failed %#v", err))
//		}
//	})
//}

//func convertFBTargetType(feedBackTargetType int32) (msgTargetType, name string) {

//	switch feedBackTargetType {
//	// 类型 参看interface-feedback
//	//TargetTypeMember   = int32(2)
//	case 2:
//		return model.TargetTypeMember, "用户"
//	//TargetTypeTopic    = int32(3)
//	case 3:
//		return model.TargetTypeTopic, "主题"
//		// TargetTypeArticle  = int32(4)
//	case 4:
//		return model.TargetTypeArticle, "文章"
//		//TargetTypeDiscuss = int32(5)
//	case 5:
//		return model.TargetTypeDiscussion, "讨论"
//		//TargetTypeRevise = int32(6)
//	case 6:
//		return model.TargetTypeRevise, "修正"
//		//TargetTypeComment = int32(7)
//	case 7:
//		return model.TargetTypeComment, "评论"
//	default:
//		return "", ""
//	}
//}

//func (p *Service) onFeedBackAccuseNotSuit(m *stan.Msg) {
//	var err error
//	c := context.Background()
//	info := new(def.MsgFeedbackAccuseNotSuit)
//	if err = info.Unmarshal(m.Data); err != nil {
//		log.For(c).Error(fmt.Sprintf("service.onFeedBackAccuseNotSuit Unmarshal failed %#v", err))
//		return
//	}
//	var fb *model.Feedback
//	action := func(c context.Context, _ uint) error {
//		feedback, err := p.d.GetFeedbackByID(c, p.d.DB(), info.FeedbackID)
//		if err != nil {
//			return err
//		}

//		fb = feedback
//		return nil
//	}

//	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
//		log.For(c).Error(fmt.Sprintf("service.onFeedBackAccuseNotSuit GetComment failed %#v", err))
//		m.Ack()
//		return
//	}
//	msgTargetType, targetDesc := convertFBTargetType(fb.TargetType)
//	actionText := fmt.Sprintf(model.MsgTextFeedBackAccuseNotSuit, targetDesc+":"+fb.FeedbackDesc, fb.VerifyDesc)
//	// 发送给举报人
//	msg := &model.Message{
//		ID:         gid.NewID(),
//		AccountID:  fb.CreatedBy,
//		ActionType: model.MsgFeedbackAccuseNotSuit,
//		ActionTime: time.Now().Unix(),
//		ActionText: actionText,
//		MergeCount: 1,
//		ActorType:  model.ActorTypeUser,
//		Actors:     strconv.FormatInt(fb.CreatedBy, 10),
//		TargetID:   fb.TargetID,
//		TargetType: msgTargetType,
//		CreatedAt:  time.Now().Unix(),
//		UpdatedAt:  time.Now().Unix(),
//	}
//	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
//		log.For(c).Error(fmt.Sprintf("service.onFeedBackAccuseNotSuit AddMessage failed %#v", err))
//		return
//	}
//	p.addCache(func() {
//		if _, err := p.pushSingleUser(context.Background(),
//			msg.AccountID,
//			msg.ID,
//			fmt.Sprintf(def.PushMsgFeedBackAccuseNotSuit, targetDesc+":"+fb.FeedbackDesc, fb.VerifyDesc),
//			fmt.Sprintf(def.PushMsgFeedBackAccuseNotSuit, targetDesc+":"+fb.FeedbackDesc, fb.VerifyDesc),
//			"",
//		); err != nil {
//			log.For(context.Background()).Error(fmt.Sprintf("service.onFeedBackAccuseNotSuit Push message failed %#v", err))
//		}
//	})
//	m.Ack()
//}
