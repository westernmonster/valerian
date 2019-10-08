package service

import (
	"context"
	"time"
	"valerian/app/service/account-feed/model"
	article "valerian/app/service/article/api"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onReviseAdded(m *stan.Msg) {
	var err error
	info := new(model.MsgReviseAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onReviseAdded Unmarshal failed %#v", err)
		return
	}

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(context.Background(), info.ReviseID); err != nil {
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: model.ActionTypeCreateRevise,
		ActionTime: time.Now().Unix(),
		ActionText: model.ActionTextCreateRevise,
		TargetID:   article.ID,
		TargetType: model.TargetTypeRevise,
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
	info := new(model.MsgReviseUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onReviseUpdated Unmarshal failed %#v", err)
		return
	}

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(context.Background(), info.ReviseID); err != nil {
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: model.ActionTypeUpdateRevise,
		ActionTime: time.Now().Unix(),
		ActionText: model.ActionTextUpdateRevise,
		TargetID:   article.ID,
		TargetType: model.TargetTypeRevise,
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
	info := new(model.MsgReviseDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onReviseDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelAccountFeedByCond(context.Background(), p.d.DB(), model.TargetTypeRevise, info.ReviseID); err != nil {
		log.Errorf("service.onReviseDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
