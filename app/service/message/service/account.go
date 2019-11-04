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
	"valerian/library/jpush"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onMemberFollowed(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgMemberFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onMemberFollowed Unmarshal failed %#v", err))
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
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
		log.For(c).Error(fmt.Sprintf("service.onCommentAdded AddMessage failed %#v", err))
		return
	}

	if err = p.d.IncrMessageStat(c, tx, &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()

	p.addCache(func() {
		if _, err := p.pushSingleUser(context.Background(), msg.AccountID, &jpush.Message{
			Title:   def.PushMsgTitleFollowed,
			Content: def.PushMsgTitleFollowed,
			Extras: map[string]interface{}{
				"id":   strconv.FormatInt(msg.ID, 10),
				"type": "link",
				"url":  fmt.Sprintf(def.LinkUser, info.TargetAccountID),
			},
		}); err != nil {
			log.For(context.Background()).Error(fmt.Sprintf("service.onMemberFollowed Push message failed %#v", err))
		}
	})

}
