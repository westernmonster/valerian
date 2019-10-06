package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

func (p *Dao) GetFansPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.AccountFans, err error) {
	items = make([]*model.AccountFans, 0)

	sql := "SELECT a.* FROM account_fans a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.id DESC limit ?,?"

	log.For(c).Info(fmt.Sprintf("dao.GetFansPaged  account_id(%d) limit(%d) offset(%d)", aid, limit, offset))
	if err = node.SelectContext(c, &items, sql, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFansPaged error(%+v), account_id(%d) limit(%d) offset(%d)", err, aid, limit, offset))
	}
	return
}

func (p *Dao) SetFans(c context.Context, node sqalx.Node, attr uint32, aid, fid int64) (err error) {
	sqlUpdate := "UPDATE account_fans SET attribute=?,updated_at=? WHERE account_id=? and target_account_id=? AND deleted=0"

	_, err = node.ExecContext(c, sqlUpdate, attr, time.Now().Unix(), aid, fid)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SetFans err(%+v), attribute (%d) aid(%d) fid(%d)", err, attr, aid, fid))
		return
	}

	return
}

func (p *Dao) GetFansIDs(c context.Context, node sqalx.Node, aid int64) (items []int64, err error) {
	// items: = make([]*model.AccountFans, 0)
	items = make([]int64, 0)
	sqlSelect := "SELECT a.target_account_id FROM account_fans a WHERE a.deleted=0 AND a.account_id =?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFansIDs err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetFansIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

// GetAll get all records
func (p *Dao) GetFansList(c context.Context, node sqalx.Node) (items []*model.AccountFans, err error) {
	items = make([]*model.AccountFans, 0)
	sqlSelect := "SELECT a.* FROM account_fans a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFans err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetFansListByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountFans, err error) {
	items = make([]*model.AccountFans, 0)
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
	if val, ok := cond["target_account_id"]; ok {
		clause += " AND a.target_account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["attribute"]; ok {
		clause += " AND a.attribute =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM account_fans a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFansByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetFansByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountFans, err error) {
	item = new(model.AccountFans)
	sqlSelect := "SELECT a.* FROM account_fans a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFanByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetFansByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountFans, err error) {
	item = new(model.AccountFans)
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
	if val, ok := cond["target_account_id"]; ok {
		clause += " AND a.target_account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["attribute"]; ok {
		clause += " AND a.attribute =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM account_fans a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFansByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddFans(c context.Context, node sqalx.Node, item *model.AccountFans) (err error) {
	sqlInsert := "INSERT INTO account_fans( id,account_id,target_account_id,attribute,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.TargetAccountID, item.Attribute, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddFans err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateFans(c context.Context, node sqalx.Node, item *model.AccountFans) (err error) {
	sqlUpdate := "UPDATE account_fans SET account_id=?,target_account_id=?,attribute=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.TargetAccountID, item.Attribute, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFans err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelFans(c context.Context, node sqalx.Node, aid, fid int64) (err error) {
	sqlDelete := "UPDATE account_fans SET deleted=1 WHERE deleted=0 AND account_id=? and target_account_id=? "

	if _, err = node.ExecContext(c, sqlDelete, aid, fid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFans err(%+v), aid(%+v) fid(%d)", err, aid, fid))
		return
	}

	return
}
