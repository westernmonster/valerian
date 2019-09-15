package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by Account ID
func (p *Dao) GetAccountRelationStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountRelationStat, err error) {
	item = new(model.AccountRelationStat)
	sqlSelect := "SELECT a.* FROM account_relation_stats a WHERE a.account_id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, aid); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountRelationStatByID err(%+v), id(%+v)", err, aid))
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAccountRelationStat(c context.Context, node sqalx.Node, item *model.AccountRelationStat) (err error) {
	sqlInsert := "INSERT INTO account_relation_stats( account_id,following,fans,black,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.AccountID, item.Following, item.Fans, item.Black, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountRelationStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccountRelationStat(c context.Context, node sqalx.Node, item *model.AccountRelationStat) (err error) {
	sqlUpdate := "UPDATE account_relation_stats SET following=?,fans=?,black=?,updated_at=? WHERE account_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Following, item.Fans, item.Black, item.UpdatedAt, item.AccountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountRelationStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}
