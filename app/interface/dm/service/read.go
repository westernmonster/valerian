package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/dm/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) MarkRead(c context.Context, id int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
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

	var msg *model.Message
	if msg, err = p.d.GetMessageByID(c, tx, id); err != nil {
		return
	} else if msg == nil {
		err = ecode.MessageNotExist
		return
	}

	msg.IsRead = types.BitBool(true)
	if err = p.d.UpdateMessage(c, tx, msg); err != nil {
		return
	}

	if err = p.d.IncrMessageStat(c, tx, &model.MessageStat{AccountID: aid, UnreadCount: -1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return

}

func (p *Service) MarkAllRead(c context.Context) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
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

	if err = p.d.MarkAllRead(c, tx, aid); err != nil {
		return
	}

	var stat *model.MessageStat
	if stat, err = p.d.GetMessageStatForUpdate(c, tx, aid); err != nil {
		return
	}

	stat.UnreadCount = 0
	stat.UpdatedAt = time.Now().Unix()
	if err = p.d.UpdateMessageStat(c, tx, stat); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return

}
