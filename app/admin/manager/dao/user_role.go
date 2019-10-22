package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/admin/manager/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetUserRoles(c context.Context, node sqalx.Node) (items []*model.UserRole, err error) {
	items = make([]*model.UserRole, 0)
	sqlSelect := "SELECT a.* FROM user_roles a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserRoles err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetUserRolesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.UserRole, err error) {
	items = make([]*model.UserRole, 0)
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM user_roles a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserRolesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetUserRoleByID(c context.Context, node sqalx.Node, id string) (item *model.UserRole, err error) {
	item = new(model.UserRole)
	sqlSelect := "SELECT a.* FROM user_roles a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetUserRoleByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetUserRoleByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.UserRole, err error) {
	item = new(model.UserRole)
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM user_roles a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetUserRolesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddUserRole(c context.Context, node sqalx.Node, item *model.UserRole) (err error) {
	sqlInsert := "INSERT INTO user_roles( id,name,perms,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Name, item.Perms, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddUserRoles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateUserRole(c context.Context, node sqalx.Node, item *model.UserRole) (err error) {
	sqlUpdate := "UPDATE user_roles SET name=?,perms=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Name, item.Perms, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateUserRoles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelUserRole(c context.Context, node sqalx.Node, id string) (err error) {
	sqlDelete := "UPDATE user_roles SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelUserRoles err(%+v), item(%+v)", err, id))
		return
	}

	return
}
