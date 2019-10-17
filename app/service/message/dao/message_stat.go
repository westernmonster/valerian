package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by ID
func (p *Dao) GetMessageStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.MessageStat, err error) {
	item = new(model.MessageStat)
	sqlSelect := "SELECT a.* FROM message_stats a WHERE a.account_id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, aid); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetMessageStatByID err(%+v), aid(%+v)", err, aid))
	}

	return
}

// Insert insert a new record
func (p *Dao) AddMessageStat(c context.Context, node sqalx.Node, item *model.MessageStat) (err error) {
	sqlInsert := "INSERT INTO message_stats( account_id,unread_count,created_at,updated_at) VALUES ( ?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.AccountID, item.UnreadCount, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddMessageStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateMessageStat(c context.Context, node sqalx.Node, item *model.MessageStat) (err error) {
	sqlUpdate := "UPDATE message_stats SET unread_count=?,updated_at=? WHERE account_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.UnreadCount, item.UpdatedAt, item.AccountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateMessageStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) IncrMessageStat(c context.Context, node sqalx.Node, item *model.MessageStat) (err error) {
	sqlUpdate := "UPDATE message_stats SET unread_count=unread_count+?,updated_at=? WHERE account_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.UnreadCount, item.UpdatedAt, item.AccountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrMessageStat err(%+v), item(%+v)", err, item))
		return
	}

	return
}
