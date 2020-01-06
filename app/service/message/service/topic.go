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

	"github.com/davecgh/go-spew/spew"
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

func (p *Service) getTopicInviteRequest(c context.Context, node sqalx.Node, id int64) (item *model.TopicInviteRequest, err error) {
	var req *model.TopicInviteRequest
	if req, err = p.d.GetTopicInviteRequestByID(c, node, id); err != nil {
		return
	} else if req == nil {
		err = ecode.TopicInviteRequestNotExist
		return
	}

	return
}

func (p *Service) onTopicFollowed(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)

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
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

}

func (p *Service) onTopicFollowRequested(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)

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
		fmt.Println(err)
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetTopicFollowRequest", "GetTopicFollowRequest(), id(%d),error(%+v)", info.RequestID, err)
		return
	}

	spew.Dump(req)

	pushMsgs := make([]*model.PushMessage, 0)

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

		pushMsgs = append(pushMsgs, &model.PushMessage{
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
			if _, err := p.pushSingleUser(context.Background(), v); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", v, err)
			}
		}
	})

}

func (p *Service) onTopicFollowApproved(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)

	info := new(def.MsgTopicFollowApproved)
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
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	stat := &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}
	if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

	p.addCache(func() {
		push := &model.PushMessage{
			Aid:     msg.AccountID,
			MsgID:   msg.ID,
			Title:   def.PushMsgTitleTopicFollowApproved,
			Content: def.PushMsgTitleTopicFollowApproved,
			Link:    fmt.Sprintf(def.LinkTopicRequest, req.ID),
		}
		if _, err := p.pushSingleUser(context.Background(), push); err != nil {
			PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
		}
	})

}

func (p *Service) onTopicFollowRejected(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)

	info := new(def.MsgTopicFollowRejected)
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
		PromError("message: GetTopic", "GetTopic(), id(%d),error(%+v)", req.TopicID, err)
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
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	stat := &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}
	if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

	p.addCache(func() {
		push := &model.PushMessage{
			Aid:     msg.AccountID,
			MsgID:   msg.ID,
			Title:   def.PushMsgTitleTopicFollowRejected,
			Content: def.PushMsgTitleTopicFollowRejected,
			Link:    fmt.Sprintf(def.LinkTopicRequest, req.ID),
		}
		if _, err := p.pushSingleUser(context.Background(), push); err != nil {
			PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
		}
	})

}

func (p *Service) onTopicInvite(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)
	info := new(def.MsgTopicInviteSent)
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

	var req *model.TopicInviteRequest
	if req, err = p.getTopicInviteRequest(c, tx, info.InviteID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetTopicInviteRequest", "GetTopicInviteRequest(), id(%d),error(%+v)", info.InviteID, err)
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
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	stat := &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}
	if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
		PromError("message: IncrMessageStat", "IncrMessageStat(), stat(%+v),error(%+v)", stat, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

	p.addCache(func() {
		push := &model.PushMessage{
			Aid:     msg.AccountID,
			MsgID:   msg.ID,
			Title:   def.PushMsgTitleTopicFollowInvited,
			Content: def.PushMsgTitleTopicFollowInvited,
			Link:    fmt.Sprintf(def.LinkTopicInvite, req.ID),
		}
		if _, err := p.pushSingleUser(context.Background(), push); err != nil {
			PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
		}
	})
}
