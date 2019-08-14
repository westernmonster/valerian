package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetAccountFollowersCount(c context.Context, node sqalx.Node, accountID int64) (count int, err error) {
	countSQL := "SELECT COUNT(1) as count FROM account_followers a WHERE a.deleted=0 AND a.account_id=?"
	if err = node.GetContext(c, &count, countSQL, accountID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFollowersCount error(%+v), account_id(%d)", err, accountID))
	}
	return
}

func (p *Dao) GetAccountFollowersPaged(c context.Context, node sqalx.Node, accountID int64, page, pageSize int) (count int, items []*model.AccountFollower, err error) {
	items = make([]*model.AccountFollower, 0)
	offset := (page - 1) * pageSize

	countSQL := "SELECT COUNT(1) as count FROM account_followers a WHERE a.deleted=0 AND a.account_id=?"
	selectSQL := "SELECT a.* FROM account_followers a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.id DESC limit ?,?"

	if err = node.GetContext(c, &count, countSQL, accountID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFollowersPaged error(%+v), account_id(%d)", err, accountID))
	}

	if err = node.SelectContext(c, &items, selectSQL, accountID, offset, pageSize); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFollowersPaged error(%+v), account_id(%d) page(%d) pageSize(%d)", err, accountID, page, pageSize))
	}
	return
}

// GetAll get all records
func (p *Dao) GetAccountFollowers(c context.Context, node sqalx.Node) (items []*model.AccountFollower, err error) {
	items = make([]*model.AccountFollower, 0)
	sqlSelect := "SELECT a.* FROM account_followers a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFollowers err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetAccountFollowersByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountFollower, err error) {
	items = make([]*model.AccountFollower, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["follower_id"]; ok {
		clause += " AND a.follower_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM account_followers a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFollowersByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetAccountFollowerByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountFollower, err error) {
	item = new(model.AccountFollower)
	sqlSelect := "SELECT a.* FROM account_followers a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFollowerByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetAccountFollowerByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountFollower, err error) {
	item = new(model.AccountFollower)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["follower_id"]; ok {
		clause += " AND a.follower_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM account_followers a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFollowersByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAccountFollower(c context.Context, node sqalx.Node, item *model.AccountFollower) (err error) {
	sqlInsert := "INSERT INTO account_followers( id,account_id,follower_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.FollowerID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountFollowers err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccountFollower(c context.Context, node sqalx.Node, item *model.AccountFollower) (err error) {
	sqlUpdate := "UPDATE account_followers SET account_id=?,follower_id=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.FollowerID, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountFollowers err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelAccountFollower(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE account_followers SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountFollowers err(%+v), item(%+v)", err, id))
		return
	}

	return
}
