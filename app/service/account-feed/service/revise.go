package service

import (
	"context"
	"time"
	"valerian/app/service/account-feed/model"
	article "valerian/app/service/article/api"
	"valerian/app/service/feed/def"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onReviseAdded(m *stan.Msg) {
	var err error
	info := new(def.MsgReviseAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onReviseAdded Unmarshal failed %#v", err)
		return
	}

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(context.Background(), info.ReviseID); err != nil {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCreateRevise,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateRevise,
		TargetID:   article.ID,
		TargetType: def.TargetTypeRevise,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onReviseAdded() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onReviseUpdated(m *stan.Msg) {
	var err error
	info := new(def.MsgReviseUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onReviseUpdated Unmarshal failed %#v", err)
		return
	}

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(context.Background(), info.ReviseID); err != nil {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeUpdateRevise,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextUpdateRevise,
		TargetID:   article.ID,
		TargetType: def.TargetTypeRevise,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onReviseUpdated() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onReviseDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgReviseDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onReviseDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelAccountFeedByCond(context.Background(), p.d.DB(), def.TargetTypeRevise, info.ReviseID); err != nil {
		log.Errorf("service.onReviseDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
