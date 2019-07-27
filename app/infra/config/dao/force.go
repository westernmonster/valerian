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
func (p *Dao) GetForces(c context.Context, node sqalx.Node) (items []*model.Force, err error) {
	items = make([]*model.Force, 0)
	sqlSelect := "SELECT a.* FROM forces a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetForces err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetForcesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Force, err error) {
	items = make([]*model.Force, 0)
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
	if val, ok := cond["host_name"]; ok {
		clause += " AND a.host_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["ip"]; ok {
		clause += " AND a.ip =?"
		condition = append(condition, val)
	}
	if val, ok := cond["version"]; ok {
		clause += " AND a.version =?"
		condition = append(condition, val)
	}
	if val, ok := cond["operator"]; ok {
		clause += " AND a.operator =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM forces a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetForcesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetForceByID(c context.Context, node sqalx.Node, id int64) (item *model.Force, err error) {
	item = new(model.Force)
	sqlSelect := "SELECT a.* FROM forces a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetForceByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetForceByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Force, err error) {
	item = new(model.Force)
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
	if val, ok := cond["host_name"]; ok {
		clause += " AND a.host_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["ip"]; ok {
		clause += " AND a.ip =?"
		condition = append(condition, val)
	}
	if val, ok := cond["version"]; ok {
		clause += " AND a.version =?"
		condition = append(condition, val)
	}
	if val, ok := cond["operator"]; ok {
		clause += " AND a.operator =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM forces a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetForcesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddForce(c context.Context, node sqalx.Node, item *model.Force) (err error) {
	sqlInsert := "INSERT INTO forces( id,app_id,host_name,ip,version,operator,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AppID, item.HostName, item.IP, item.Version, item.Operator, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddForces err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateForce(c context.Context, node sqalx.Node, item *model.Force) (err error) {
	sqlUpdate := "UPDATE forces SET app_id=?,host_name=?,ip=?,version=?,operator=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AppID, item.HostName, item.IP, item.Version, item.Operator, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateForces err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelForce(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE forces SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelForces err(%+v), item(%+v)", err, id))
		return
	}

	return
}
