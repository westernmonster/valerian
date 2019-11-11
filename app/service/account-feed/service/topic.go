package service

import (
	"context"
	"time"
	"valerian/app/service/account-feed/model"
	account "valerian/app/service/account/api"
	"valerian/app/service/feed/def"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/kamilsk/retry/v4"
	"github.com/kamilsk/retry/v4/strategy"
	"github.com/nats-io/stan.go"
)

func (p *Service) onTopicAdded(m *stan.Msg) {
	var err error
	info := new(def.MsgTopicAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onTopicAdded Unmarshal failed %#v", err)
		return
	}

	var topic *topic.TopicInfo
	if topic, err = p.d.GetTopic(context.Background(), info.TopicID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
		}
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCreateTopic,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateTopic,
		TargetID:   topic.ID,
		TargetType: def.TargetTypeTopic,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onTopicAdded() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onTopicFollowed(m *stan.Msg) {
	var err error
	info := new(def.MsgTopicFollowed)
	c := context.Background()
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onTopicFollowed Unmarshal failed %#v", err)
		return
	}

	var topic *topic.TopicInfo
	if topic, err = p.d.GetTopic(context.Background(), info.TopicID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
		}
		return
	}

	var v *account.BaseInfoReply
	action := func(c context.Context, _ uint) error {
		acc, e := p.d.GetAccountBaseInfo(c, info.ActorID)
		if e != nil {
			return e
		}

		v = acc
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeFollowTopic,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextFollowTopic,
		TargetID:   topic.ID,
		TargetType: def.TargetTypeTopic,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onTopicFollowed() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onTopicDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgTopicDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onTopicDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelAccountFeedByCond(context.Background(), p.d.DB(), def.TargetTypeTopic, info.TopicID); err != nil {
		log.Errorf("service.onTopicDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
