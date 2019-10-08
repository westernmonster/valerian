package service

import (
	"context"
	"time"
	"valerian/app/service/account-feed/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onTopicAdded(m *stan.Msg) {
	var err error
	info := new(model.MsgTopicAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onTopicAdded Unmarshal failed %#v", err)
		return
	}

	var topic *topic.TopicInfo
	if topic, err = p.d.GetTopic(context.Background(), info.TopicID); err != nil {
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: model.ActionTypeCreateTopic,
		ActionTime: time.Now().Unix(),
		ActionText: model.ActionTextCreateTopic,
		TargetID:   topic.ID,
		TargetType: model.TargetTypeTopic,
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
	info := new(model.MsgTopicFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onTopicFollowed Unmarshal failed %#v", err)
		return
	}

	var topic *topic.TopicInfo
	if topic, err = p.d.GetTopic(context.Background(), info.TopicID); err != nil {
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: model.ActionTypeFollowTopic,
		ActionTime: time.Now().Unix(),
		ActionText: model.ActionTextFollowTopic,
		TargetID:   topic.ID,
		TargetType: model.TargetTypeTopic,
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
	info := new(model.MsgTopicDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onTopicDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelAccountFeedByCond(context.Background(), p.d.DB(), model.TargetTypeTopic, info.TopicID); err != nil {
		log.Errorf("service.onTopicDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
