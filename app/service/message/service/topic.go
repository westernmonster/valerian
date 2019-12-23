package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) getTopic(c context.Context, node sqalx.Node, topicID int64) (item *model.Topic, err error) {
	if item, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if item == nil {
		return nil, ecode.TopicNotExist
	}
	return
}

func (p *Service) getTopicFollowRequest(c context.Context, node sqalx.Node, id int64) (item *model.TopicFollowRequest, err error) {
	var req *model.TopicFollowRequest
	if req, err = p.d.GetTopicFollowRequestByID(c, node, id); err != nil {
		return
	} else if req == nil {
		err = ecode.TopicFollowRequestNotExist
		return
	}

	return
}

func (p *Service) onTopicFollowed(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("message: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("message: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("message: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	var topic *model.Topic
	if topic, err = p.getTopic(c, tx, info.TopicID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetTopic", "GetTopic(), id(%d),error(%+v)", info.TopicID, err)
		return
	}

	var admins []*model.TopicMember
	if admins, err = p.d.GetAdminTopicMembers(c, tx, info.TopicID); err != nil {
		PromError("message: GetAdminTopicMembers", "GetAdminTopicMembers(), topic_id(%d),error(%+v)", info.TopicID, err)
		return
	}

	for _, v := range admins {
		msg := &model.Message{
			ID:         gid.NewID(),
			AccountID:  v.AccountID,
			ActionType: model.MsgJoined,
			ActionTime: time.Now().Unix(),
			ActionText: model.MsgTextJoined,
			Actors:     strconv.FormatInt(info.ActorID, 10),
			MergeCount: 1,
			ActorType:  model.ActorTypeUser,
			TargetID:   topic.ID,
			TargetType: model.TargetTypeTopic,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddMessage(c, tx, msg); err != nil {
			PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
			return
		}

		stat := &model.MessageStat{AccountID: v.AccountID, UnreadCount: 1}
		if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
			PromError("message: IncrMessageStat", "IncrMessageStat(), stat(%+v),error(%+v)", stat, err)
			return
		}

	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()

}

func (p *Service) onTopicFollowRequested(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicFollowRequested)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("message: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("message: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("message: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	var admins []*model.TopicMember
	if admins, err = p.d.GetAdminTopicMembers(c, tx, info.TopicID); err != nil {
		PromError("message: GetAdminTopicMembers", "GetAdminTopicMembers(), topic_id(%d),error(%+v)", info.TopicID, err)
		return
	}

	var req *model.TopicFollowRequest
	if req, err = p.getTopicFollowRequest(c, tx, info.RequestID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetTopicFollowRequest", "GetTopicFollowRequest(), id(%d),error(%+v)", info.RequestID, err)
		return
	}

	type PushMessge struct {
		Aid     int64
		MsgID   int64
		Title   string
		Content string
		Link    string
	}

	pushMsgs := make([]*PushMessge, 0)

	for _, v := range admins {

		msg := &model.Message{
			ID:         gid.NewID(),
			AccountID:  v.AccountID,
			ActionType: model.MsgApply,
			ActionTime: time.Now().Unix(),
			ActionText: model.MsgTextApply,
			Actors:     strconv.FormatInt(info.ActorID, 10),
			MergeCount: 1,
			ActorType:  model.ActorTypeUser,
			TargetID:   req.ID,
			CreatedAt:  time.Now().Unix(),
			TargetType: model.TargetTypeTopicFollowRequest,
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddMessage(c, tx, msg); err != nil {
			PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
			return
		}

		stat := &model.MessageStat{AccountID: v.AccountID, UnreadCount: 1}
		if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
			PromError("message: IncrMessageStat", "IncrMessageStat(), stat(%+v),error(%+v)", stat, err)
			return
		}

		pushMsgs = append(pushMsgs, &PushMessge{
			MsgID:   msg.ID,
			Title:   def.PushMsgTitleTopicFollowRequested,
			Content: def.PushMsgTitleTopicFollowRequested,
			Link:    fmt.Sprintf(def.LinkTopicRequest, req.ID),
			Aid:     v.AccountID,
		})

	}

	if err = tx.Commit(); err != nil {
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

	p.addCache(func() {
		for _, v := range pushMsgs {
			if _, err := p.pushSingleUser(context.Background(), v.Aid, v.MsgID, v.Title, v.Content, v.Link); err != nil {
				log.For(context.Background()).Error(fmt.Sprintf("service.onTopicFollowRequested Push message failed %#v", err))
			}
		}
	})

}

func (p *Service) onTopicFollowApproved(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicFollowApproved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowApproved Unmarshal failed %#v", err))
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var req *model.TopicFollowRequest
	if req, err = p.getTopicFollowRequest(c, tx, info.RequestID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetTopicFollowRequest", "GetTopicFollowRequest(), id(%d),error(%+v)", info.RequestID, err)
		return
	}

	var topic *model.Topic
	if topic, err = p.getTopic(c, tx, req.TopicID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetTopic", "GetTopic(), id(%d),error(%+v)", info.TopicID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  req.AccountID,
		ActionType: model.MsgApplyApproved,
		ActionTime: time.Now().Unix(),
		ActionText: fmt.Sprintf(model.MsgTextApplyApproved, topic.Name),
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   req.ID,
		CreatedAt:  time.Now().Unix(),
		TargetType: model.TargetTypeTopicFollowRequest,
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, tx, msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowApproved AddMessage failed %#v", err))
		return
	}

	if err = p.d.IncrMessageStat(c, tx, &model.MessageStat{AccountID: req.AccountID, UnreadCount: 1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()

	p.addCache(func() {
		if _, err := p.pushSingleUser(context.Background(),
			msg.AccountID,
			msg.ID,
			def.PushMsgTitleTopicFollowApproved,
			def.PushMsgTitleTopicFollowApproved,
			fmt.Sprintf(def.LinkTopicRequest, req.ID),
		); err != nil {
			log.For(context.Background()).Error(fmt.Sprintf("service.onTopicFollowApproved Push message failed %#v", err))
		}
	})

}

func (p *Service) onTopicFollowRejected(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicFollowRejected)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowRejected Unmarshal failed %#v", err))
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var req *model.TopicFollowRequest
	if req, err = p.d.GetTopicFollowRequestByID(c, tx, info.RequestID); err != nil {
		return
	} else if req == nil {
		err = ecode.TopicFollowRequestNotExist
		return
	}

	var topic *model.Topic
	if topic, err = p.getTopic(c, tx, req.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowRejected GetTopic failed %#v", err))
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  req.AccountID,
		ActionType: model.MsgApplyRejected,
		ActionTime: time.Now().Unix(),
		ActionText: fmt.Sprintf(model.MsgTextApplyRejected, topic.Name, req.RejectReason),
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   req.ID,
		CreatedAt:  time.Now().Unix(),
		TargetType: model.TargetTypeTopicFollowRequest,
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, tx, msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowRejected AddMessage failed %#v", err))
		return
	}

	if err = p.d.IncrMessageStat(c, tx, &model.MessageStat{AccountID: req.AccountID, UnreadCount: 1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()

	p.addCache(func() {
		if _, err := p.pushSingleUser(context.Background(),
			msg.AccountID,
			msg.ID,
			def.PushMsgTitleTopicFollowRejected,
			def.PushMsgTitleTopicFollowRejected,
			fmt.Sprintf(def.LinkTopicRequest, req.ID),
		); err != nil {
			log.For(context.Background()).Error(fmt.Sprintf("service.onTopicFollowRejected Push message failed %#v", err))
		}
	})

}

func (p *Service) onTopicInvite(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicInviteSent)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowed Unmarshal failed %#v", err))
		return
	}
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var req *model.TopicInviteRequest
	if req, err = p.d.GetTopicInviteRequestByID(c, tx, info.InviteID); err != nil {
		return
	} else if req == nil {
		err = ecode.TopicInviteRequestNotExist
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  req.AccountID,
		ActionType: model.MsgInvite,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextInvite,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   req.ID,
		CreatedAt:  time.Now().Unix(),
		TargetType: model.TargetTypeTopicInviteRequest,
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, tx, msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowed AddMessage failed %#v", err))
		return
	}

	if err = p.d.IncrMessageStat(c, tx, &model.MessageStat{AccountID: req.AccountID, UnreadCount: 1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()

	p.addCache(func() {
		if _, err := p.pushSingleUser(context.Background(),
			msg.AccountID,
			msg.ID,
			def.PushMsgTitleTopicFollowInvited,
			def.PushMsgTitleTopicFollowInvited,
			fmt.Sprintf(def.LinkTopicInvite, req.ID),
		); err != nil {
			log.For(context.Background()).Error(fmt.Sprintf("service.onTopicInvite Push message failed %#v", err))
		}
	})
}
