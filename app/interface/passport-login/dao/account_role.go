package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/passport-login/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetAccountRoles(c context.Context, node sqalx.Node) (items []*model.AccountRole, err error) {
	items = make([]*model.AccountRole, 0)
	sqlSelect := "SELECT a.id,a.name,a.perms,a.deleted,a.created_at,a.updated_at FROM account_roles a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountRoles err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetAccountRolesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountRole, err error) {
	items = make([]*model.AccountRole, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["perms"]; ok {
		clause += " AND a.perms =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.perms,a.deleted,a.created_at,a.updated_at FROM account_roles a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountRolesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetAccountRoleByID(c context.Context, node sqalx.Node, id string) (item *model.AccountRole, err error) {
	item = new(model.AccountRole)
	sqlSelect := "SELECT a.id,a.name,a.perms,a.deleted,a.created_at,a.updated_at FROM account_roles a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountRoleByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetAccountRoleByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountRole, err error) {
	item = new(model.AccountRole)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["perms"]; ok {
		clause += " AND a.perms =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.perms,a.deleted,a.created_at,a.updated_at FROM account_roles a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountRolesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAccountRole(c context.Context, node sqalx.Node, item *model.AccountRole) (err error) {
	sqlInsert := "INSERT INTO account_roles( id,name,perms,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Name, item.Perms, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountRoles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccountRole(c context.Context, node sqalx.Node, item *model.AccountRole) (err error) {
	sqlUpdate := "UPDATE account_roles SET name=?,perms=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Name, item.Perms, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountRoles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelAccountRole(c context.Context, node sqalx.Node, id string) (err error) {
	sqlDelete := "UPDATE account_roles SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountRoles err(%+v), item(%+v)", err, id))
		return
	}

	return
}
