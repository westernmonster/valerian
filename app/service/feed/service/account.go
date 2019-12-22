package service

import (
	"context"
	"time"

	"valerian/app/service/feed/def"
	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"

	"github.com/nats-io/stan.go"
)

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
			p.d.SetAccountCache(context.TODO(), account)
		})
	}
	return
}

func (p *Service) onMemberFollowed(m *stan.Msg) {
	var err error

	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)

	info := new(def.MsgMemberFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	var member *model.Account
	if member, err = p.getAccount(c, tx, info.TargetAccountID); err != nil {
		if ecode.IsNotExistEcode(err) {
			// 如果不存在这个账户，直接Ack
			m.Ack()
			return
		}
		PromError("feed: getAccount", "getAccount(), id(%d),error(%+v)", info.TargetAccountID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, tx, info.AccountID); err != nil {
		PromError("feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.AccountID, err)
		return
	}

	if !setting.ActivityFollowMember {
		m.Ack()
		return
	}

	var ids []int64
	if ids, err = p.d.GetFansIDs(c, tx, info.AccountID); err != nil {
		PromError("feed: GetFansIDs", "GetFansIDs(), aid(%d),error(%+v)", info.AccountID, err)
		return
	}

	for _, v := range ids {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeFollowMember,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextFollowMember,
			ActorID:    info.AccountID,
			ActorType:  def.ActorTypeUser,
			TargetID:   member.ID,
			TargetType: def.TargetTypeMember,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			PromError("feed: AddFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		PromError("feed: tx.Rollback", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()
}
