package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetFansPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.AccountRelation, err error) {
	items = make([]*model.AccountRelation, 0)

	sql := "SELECT a.* FROM account_relations a WHERE a.deleted=0 AND a.following_id=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sql, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFansPaged error(%+v), following id(%d) limit(%d) offset(%d)", err, aid, limit, offset))
	}
	return
}

func (p *Dao) GetFollowingsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.AccountRelation, err error) {
	items = make([]*model.AccountRelation, 0)

	sql := "SELECT a.* FROM account_relations a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sql, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFansPaged error(%+v), account id(%d) limit(%d) offset(%d)", err, aid, limit, offset))
	}
	return
}

// GetAll get all records
func (p *Dao) GetAccountRelations(c context.Context, node sqalx.Node) (items []*model.AccountRelation, err error) {
	items = make([]*model.AccountRelation, 0)
	sqlSelect := "SELECT a.* FROM account_relations a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountRelations err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetAccountRelationsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountRelation, err error) {
	items = make([]*model.AccountRelation, 0)
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
	if val, ok := cond["following_id"]; ok {
		clause += " AND a.following_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["attribute"]; ok {
		clause += " AND a.attribute =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM account_relations a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountRelationsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetAccountRelationByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountRelation, err error) {
	item = new(model.AccountRelation)
	sqlSelect := "SELECT a.* FROM account_relations a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountRelationByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetAccountRelationByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountRelation, err error) {
	item = new(model.AccountRelation)
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
	if val, ok := cond["following_id"]; ok {
		clause += " AND a.following_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["attribute"]; ok {
		clause += " AND a.attribute =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM account_relations a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountRelationsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAccountRelation(c context.Context, node sqalx.Node, item *model.AccountRelation) (err error) {
	sqlInsert := "INSERT INTO account_relations( id,account_id,following_id,attribute,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.FollowingID, item.Attribute, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountRelations err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccountRelation(c context.Context, node sqalx.Node, item *model.AccountRelation) (err error) {
	sqlUpdate := "UPDATE account_relations SET account_id=?,following_id=?,attribute=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.FollowingID, item.Attribute, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountRelations err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelAccountRelation(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE account_relations SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountRelations err(%+v), item(%+v)", err, id))
		return
	}

	return
}
