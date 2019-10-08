package service

import (
	"context"
	"time"
	"valerian/app/service/account-feed/model"
	discuss "valerian/app/service/discuss/api"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onDiscussionAdded(m *stan.Msg) {
	var err error
	info := new(model.MsgDiscussionAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onDiscussionAdded Unmarshal failed %#v", err)
		return
	}

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(context.Background(), info.DiscussionID); err != nil {
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: model.ActionTypeCreateDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: model.ActionTextCreateDiscussion,
		TargetID:   discuss.ID,
		TargetType: model.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onDiscussionAdded() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionUpdated(m *stan.Msg) {
	var err error
	info := new(model.MsgDiscussionUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onDiscussionUpdated Unmarshal failed %#v", err)
		return
	}

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(context.Background(), info.DiscussionID); err != nil {
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: model.ActionTypeUpdateDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: model.ActionTextUpdateDiscussion,
		TargetID:   discuss.ID,
		TargetType: model.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onDiscussionUpdated() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionDeleted(m *stan.Msg) {
	var err error
	info := new(model.MsgDiscussionDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onDiscussionDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelAccountFeedByCond(context.Background(), p.d.DB(), model.TargetTypeDiscussion, info.DiscussionID); err != nil {
		log.Errorf("service.onDiscussionDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
