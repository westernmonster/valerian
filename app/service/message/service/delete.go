package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onDiscussionDeleted(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionDeleted Unmarshal failed %#v", err))
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

	var items []*model.Message
	if items, err = p.d.GetMessagesByCond(c, tx, map[string]interface{}{
		"target_id":   info.DiscussionID,
		"target_type": model.TargetTypeDiscussion,
	}); err != nil {
		return
	}

	for _, v := range items {
		if !v.IsRead {
			if err = p.d.IncrMessageStat(c, tx, &model.MessageStat{AccountID: v.AccountID, UnreadCount: -1}); err != nil {
				return
			}
		}

		if err = p.d.DelMessage(c, tx, v.ID); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}
