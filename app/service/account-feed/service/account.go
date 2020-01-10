package service

import (
	"context"
	"time"

	"valerian/app/service/account-feed/model"
	"valerian/app/service/feed/def"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"

	"github.com/nats-io/stan.go"
)

// getAccount 获取用户信息
func (p *Service) getAccount(c context.Context, node sqalx.Node, aid int64) (account *model.Account, err error) {
	var needCache = true

	if account, err = p.d.AccountCache(c, aid); err != nil {
		needCache = false
	} else if account != nil {
		return
	}

	if account, err = p.d.GetAccountByID(c, node, aid); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	if needCache {
		p.addCache(func() {
			p.d.SetAccountCache(context.Background(), account)
		})
	}
	return
}

// onMemberFollowed 关注用户时
func (p *Service) onMemberFollowed(m *stan.Msg) {
	var err error

	c := context.Background()
	// 强制使用master库
	c = sqalx.NewContext(c, true)

	info := new(def.MsgMemberFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		m.Ack()
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var member *model.Account
	if member, err = p.getAccount(c, p.d.DB(), info.TargetAccountID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: getAccount", "getAccount(), id(%d),error(%+v)", info.TargetAccountID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), info.AccountID); err != nil {
		PromError("account-feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.AccountID, err)
		return
	}

	if !setting.ActivityFollowMember {
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

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}
