package service

import (
	"context"
	"time"

	"valerian/app/service/account-feed/model"
	account "valerian/app/service/account/api"
	"valerian/app/service/feed/def"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onMemberFollowed(m *stan.Msg) {
	var err error
	info := new(def.MsgMemberFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onMemberFollowed Unmarshal failed %#v", err)
		return
	}

	var member *account.BaseInfoReply
	if member, err = p.d.GetAccountBaseInfo(context.Background(), info.TargetAccountID); err != nil {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.AccountID,
		ActionType: def.ActionTypeFollowMember,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextFollowMember,
		TargetID:   member.ID,
		TargetType: def.TargetTypeMember,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onMemberFollowed() failed %#v", err)
		return
	}

	m.Ack()
}
