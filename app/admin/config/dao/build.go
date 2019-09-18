package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetBuilds(c context.Context, node sqalx.Node) (items []*model.Build, err error) {
	items = make([]*model.Build, 0)
	sqlSelect := "SELECT a.* FROM builds a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetBuilds err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetBuildsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Build, err error) {
	items = make([]*model.Build, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["app_id"]; ok {
		clause += " AND a.app_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["tag_id"]; ok {
		clause += " AND a.tag_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["operator"]; ok {
		clause += " AND a.operator =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM builds a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetBuildsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetBuildByID(c context.Context, node sqalx.Node, id int64) (item *model.Build, err error) {
	item = new(model.Build)
	sqlSelect := "SELECT a.* FROM builds a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetBuildByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetBuildByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Build, err error) {
	item = new(model.Build)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["app_id"]; ok {
		clause += " AND a.app_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["tag_id"]; ok {
		clause += " AND a.tag_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["operator"]; ok {
		clause += " AND a.operator =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM builds a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetBuildsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddBuild(c context.Context, node sqalx.Node, item *model.Build) (err error) {
	sqlInsert := "INSERT INTO builds( id,app_id,name,tag_id,operator,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AppID, item.Name, item.TagID, item.Operator, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddBuilds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateBuild(c context.Context, node sqalx.Node, item *model.Build) (err error) {
	sqlUpdate := "UPDATE builds SET app_id=?,name=?,tag_id=?,operator=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AppID, item.Name, item.TagID, item.Operator, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateBuilds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelBuild(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE builds SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelBuilds err(%+v), item(%+v)", err, id))
		return
	}

	return
}
