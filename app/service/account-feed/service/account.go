package service

import (
	"context"
	"time"

	"valerian/app/service/account-feed/model"
	account "valerian/app/service/account/api"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onMemberFollowed(m *stan.Msg) {
	var err error
	info := new(model.MsgMemberFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onMemberFollowed Unmarshal failed %#v", err)
		return
	}

	var member *account.BaseInfoReply
	if member, err = p.d.GetAccountBaseInfo(context.Background(), info.TargetAccountID); err != nil {
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.AccountID,
		ActionType: model.ActionTypeFollowMember,
		ActionTime: time.Now().Unix(),
		ActionText: model.ActionTextFollowMember,
		TargetID:   member.ID,
		TargetType: model.TargetTypeMember,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onMemberFollowed() failed %#v", err)
		return
	}

	m.Ack()
}
