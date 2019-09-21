package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetFollowingsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.AccountFollowing, err error) {
	items = make([]*model.AccountFollowing, 0)

	sql := "SELECT a.* FROM account_followings a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sql, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFollowingsPaged error(%+v), aid(%d) limit(%d) offset(%d)", err, aid, limit, offset))
	}
	return
}

func (p *Dao) SetFollowing(c context.Context, node sqalx.Node, attr uint32, aid, fid int64) (err error) {
	sqlUpdate := "UPDATE account_followings SET attribute=?,updated_at=? WHERE account_id=? and following_id=? AND deleted=0"

	_, err = node.ExecContext(c, sqlUpdate, attr, time.Now().Unix(), aid, fid)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SetFollowing err(%+v), attribute (%d) aid(%d) fid(%d)", err, attr, aid, fid))
		return
	}

	return
}

// GetAll get all records
func (p *Dao) GetFollowings(c context.Context, node sqalx.Node) (items []*model.AccountFollowing, err error) {
	items = make([]*model.AccountFollowing, 0)
	sqlSelect := "SELECT a.* FROM account_followings a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFollowings err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetFollowingsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountFollowing, err error) {
	items = make([]*model.AccountFollowing, 0)
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM account_followings a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFollowingsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetFollowingByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountFollowing, err error) {
	item = new(model.AccountFollowing)
	sqlSelect := "SELECT a.* FROM account_followings a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFollowingByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetFollowingByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountFollowing, err error) {
	item = new(model.AccountFollowing)
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM account_followings a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFollowingsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddFollowing(c context.Context, node sqalx.Node, item *model.AccountFollowing) (err error) {
	sqlInsert := "INSERT INTO account_followings( id,account_id,target_account_id,attribute,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.TargetAccountID, item.Attribute, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddFollowing err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateFollowing(c context.Context, node sqalx.Node, item *model.AccountFollowing) (err error) {
	sqlUpdate := "UPDATE account_followings SET account_id=?,target_account_id=?,attribute=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.TargetAccountID, item.Attribute, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFollowing err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelFollowing(c context.Context, node sqalx.Node, aid, fid int64) (err error) {
	sqlDelete := "UPDATE account_followings SET deleted=1 WHERE deleted=0 AND account_id=? AND target_account_id=?"

	if _, err = node.ExecContext(c, sqlDelete, aid, fid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFollowing err(%+v), aid(%+v) fid(%+v)", err, aid, fid))
		return
	}

	return
}
