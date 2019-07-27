package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/infra/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetApps(c context.Context, node sqalx.Node) (items []*model.App, err error) {
	items = make([]*model.App, 0)
	sqlSelect := "SELECT a.* FROM apps a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetApps err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetAppsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.App, err error) {
	items = make([]*model.App, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name LIKE ?"
		condition = append(condition, "%"+val.(string))
	}
	if val, ok := cond["token"]; ok {
		clause += " AND a.token =?"
		condition = append(condition, val)
	}
	if val, ok := cond["env"]; ok {
		clause += " AND a.env =?"
		condition = append(condition, val)
	}
	if val, ok := cond["zone"]; ok {
		clause += " AND a.zone =?"
		condition = append(condition, val)
	}
	if val, ok := cond["tree_id"]; ok {
		clause += " AND a.tree_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM apps a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAppsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetAppByID(c context.Context, node sqalx.Node, id int64) (item *model.App, err error) {
	item = new(model.App)
	sqlSelect := "SELECT a.* FROM apps a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAppByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetAppByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.App, err error) {
	item = new(model.App)
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
	if val, ok := cond["token"]; ok {
		clause += " AND a.token =?"
		condition = append(condition, val)
	}
	if val, ok := cond["env"]; ok {
		clause += " AND a.env =?"
		condition = append(condition, val)
	}
	if val, ok := cond["zone"]; ok {
		clause += " AND a.zone =?"
		condition = append(condition, val)
	}
	if val, ok := cond["tree_id"]; ok {
		clause += " AND a.tree_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM apps a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAppsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddApp(c context.Context, node sqalx.Node, item *model.App) (err error) {
	sqlInsert := "INSERT INTO apps( id,name,token,env,zone,tree_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Name, item.Token, item.Env, item.Zone, item.TreeID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddApps err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateApp(c context.Context, node sqalx.Node, item *model.App) (err error) {
	sqlUpdate := "UPDATE apps SET name=?,token=?,env=?,zone=?,tree_id=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Name, item.Token, item.Env, item.Zone, item.TreeID, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateApps err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelApp(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE apps SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelApps err(%+v), item(%+v)", err, id))
		return
	}

	return
}
