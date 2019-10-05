package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onTopicAdded(m *stan.Msg) {
	var err error
	info := new(model.MsgTopicAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onTopicAdded Unmarshal failed %#v", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(context.Background()); err != nil {
		log.Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	if err = tx.Commit(); err != nil {
		log.Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onTopicDeleted(m *stan.Msg) {
	m.Ack()
}

func (p *Service) onTopicFollowed(m *stan.Msg) {
	m.Ack()
}

func (p *Service) onTopicLeaved(m *stan.Msg) {
	m.Ack()
}

func (p *Service) onTopicInviteSent(m *stan.Msg) {
	m.Ack()
}

func (p *Service) onTopicFollowRequested(m *stan.Msg) {
	m.Ack()
}

func (p *Service) onTopicFollowRejected(m *stan.Msg) {
	m.Ack()
}

func (p *Service) onTopicFollowApproved(m *stan.Msg) {
	m.Ack()
}
