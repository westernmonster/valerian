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
func (p *Dao) GetUsers(c context.Context, node sqalx.Node) (items []*model.User, err error) {
	items = make([]*model.User, 0)
	sqlSelect := "SELECT a.* FROM users a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUsers err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetUsersByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.User, err error) {
	items = make([]*model.User, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["user_name"]; ok {
		clause += " AND a.user_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["nick_name"]; ok {
		clause += " AND a.nick_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["email"]; ok {
		clause += " AND a.email =?"
		condition = append(condition, val)
	}
	if val, ok := cond["password"]; ok {
		clause += " AND a.password =?"
		condition = append(condition, val)
	}
	if val, ok := cond["salt"]; ok {
		clause += " AND a.salt =?"
		condition = append(condition, val)
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =?"
		condition = append(condition, val)
	}
	if val, ok := cond["state"]; ok {
		clause += " AND a.state =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM users a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUsersByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetUserByID(c context.Context, node sqalx.Node, id int64) (item *model.User, err error) {
	item = new(model.User)
	sqlSelect := "SELECT a.* FROM users a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetUserByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetUserByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.User, err error) {
	item = new(model.User)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["user_name"]; ok {
		clause += " AND a.user_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["nick_name"]; ok {
		clause += " AND a.nick_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["email"]; ok {
		clause += " AND a.email =?"
		condition = append(condition, val)
	}
	if val, ok := cond["password"]; ok {
		clause += " AND a.password =?"
		condition = append(condition, val)
	}
	if val, ok := cond["salt"]; ok {
		clause += " AND a.salt =?"
		condition = append(condition, val)
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =?"
		condition = append(condition, val)
	}
	if val, ok := cond["state"]; ok {
		clause += " AND a.state =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM users a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetUsersByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddUser(c context.Context, node sqalx.Node, item *model.User) (err error) {
	sqlInsert := "INSERT INTO users( id,user_name,nick_name,email,password,salt,role,state,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.UserName, item.NickName, item.Email, item.Password, item.Salt, item.Role, item.State, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddUsers err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateUser(c context.Context, node sqalx.Node, item *model.User) (err error) {
	sqlUpdate := "UPDATE users SET user_name=?,nick_name=?,email=?,password=?,salt=?,role=?,state=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.UserName, item.NickName, item.Email, item.Password, item.Salt, item.Role, item.State, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateUsers err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelUser(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE users SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelUsers err(%+v), item(%+v)", err, id))
		return
	}

	return
}
