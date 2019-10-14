package service

import (
	"context"
	"time"
	"valerian/app/service/account-feed/model"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/feed/def"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onDiscussionAdded(m *stan.Msg) {
	var err error
	info := new(def.MsgDiscussionAdded)
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
		ActionType: def.ActionTypeCreateDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateDiscussion,
		TargetID:   discuss.ID,
		TargetType: def.TargetTypeDiscussion,
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
	info := new(def.MsgDiscussionUpdated)
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
		ActionType: def.ActionTypeUpdateDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextUpdateDiscussion,
		TargetID:   discuss.ID,
		TargetType: def.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onDiscussionUpdated() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionLiked(m *stan.Msg) {
	var err error
	info := new(def.MsgDiscussionLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onDiscussionLiked Unmarshal failed %#v", err)
		return
	}

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(context.Background(), info.DiscussionID); err != nil {
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeLikeDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextLikeDiscussion,
		TargetID:   discuss.ID,
		TargetType: def.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onDiscussionLiked() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionFaved(m *stan.Msg) {
	var err error
	info := new(def.MsgDiscussionFaved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onDiscussionFaved Unmarshal failed %#v", err)
		return
	}

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(context.Background(), info.DiscussionID); err != nil {
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeFavDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextFavDiscussion,
		TargetID:   discuss.ID,
		TargetType: def.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onDiscussionFaved() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgDiscussionDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onDiscussionDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelAccountFeedByCond(context.Background(), p.d.DB(), def.TargetTypeDiscussion, info.DiscussionID); err != nil {
		log.Errorf("service.onDiscussionDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
