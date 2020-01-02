package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"

	"github.com/nats-io/stan.go"
)

func (p *Service) onMemberFollowed(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgMemberFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("message: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("message: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("message: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  info.TargetAccountID,
		ActionType: model.MsgFollowed,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextFollowed,
		Actors:     strconv.FormatInt(info.AccountID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   info.AccountID,
		TargetType: model.TargetTypeMember,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, tx, msg); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	stat := &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}
	if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

	p.addCache(func() {
		c := context.Background()
		// 强制使用Master库
		c = sqalx.NewContext(c, true)
		var setting *model.SettingResp
		if setting, err = p.getAccountSetting(c, p.d.DB(), msg.AccountID); err != nil {
			PromError("message: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", msg.AccountID, err)
			return
		}

		push := &model.PushMessage{
			Aid:     msg.AccountID,
			MsgID:   msg.ID,
			Title:   def.PushMsgTitleFollowed,
			Content: def.PushMsgTitleFollowed,
			Link:    fmt.Sprintf(def.LinkUser, info.TargetAccountID),
		}
		if setting.NotifyNewFans {
			if _, err := p.pushSingleUser(c, push); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
			}
		}

	})

}
