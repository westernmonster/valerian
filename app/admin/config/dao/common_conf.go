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
func (p *Dao) GetCommonConfigs(c context.Context, node sqalx.Node) (items []*model.CommonConfig, err error) {
	items = make([]*model.CommonConfig, 0)
	sqlSelect := "SELECT a.id,a.name,a.comment,a.state,a.mark,a.operator,a.deleted,a.created_at,a.updated_at FROM common_configs a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCommonConfigs err(%+v)", err))
		return
	}
	return
}

func (p *Dao) GetCommonConfigsByCondPaged(c context.Context, node sqalx.Node, cond map[string]interface{}, page, pageSize int32) (count int32, items []*model.CommonConfig, err error) {
	items = make([]*model.CommonConfig, 0)
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
	if val, ok := cond["comment"]; ok {
		clause += " AND a.comment =?"
		condition = append(condition, val)
	}
	if val, ok := cond["state"]; ok {
		clause += " AND a.state =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mark"]; ok {
		clause += " AND a.mark =?"
		condition = append(condition, val)
	}
	if val, ok := cond["operator"]; ok {
		clause += " AND a.operator =?"
		condition = append(condition, val)
	}

	sqlCount := fmt.Sprintf("SELECT COUNT(1) AS count FROM common_configs a WHERE a.deleted=0 %s", clause)

	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.comment,a.state,a.mark,a.operator,a.deleted,a.created_at,a.updated_at  FROM common_configs a WHERE a.deleted=0 %s ORDER BY a.id DESC LIMIT ?,?", clause)

	if err = node.GetContext(c, &count, sqlCount, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCommonConfigsByCondPaged error(%+v), condition(%+v)", err, condition))
	}

	offset := (page - 1) * pageSize
	condition = append(condition, offset)
	condition = append(condition, pageSize)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCommonConfigsByCondPaged err(%+v), condition(%+v)", err, condition))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetCommonConfigsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.CommonConfig, err error) {
	items = make([]*model.CommonConfig, 0)
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
	if val, ok := cond["comment"]; ok {
		clause += " AND a.comment =?"
		condition = append(condition, val)
	}
	if val, ok := cond["state"]; ok {
		clause += " AND a.state =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mark"]; ok {
		clause += " AND a.mark =?"
		condition = append(condition, val)
	}
	if val, ok := cond["operator"]; ok {
		clause += " AND a.operator =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.comment,a.state,a.mark,a.operator,a.deleted,a.created_at,a.updated_at FROM common_configs a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCommonConfigsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetCommonConfigByID(c context.Context, node sqalx.Node, id int64) (item *model.CommonConfig, err error) {
	item = new(model.CommonConfig)
	sqlSelect := "SELECT a.id,a.name,a.comment,a.state,a.mark,a.operator,a.deleted,a.created_at,a.updated_at FROM common_configs a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetCommonConfigByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetCommonConfigByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.CommonConfig, err error) {
	item = new(model.CommonConfig)
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
	if val, ok := cond["comment"]; ok {
		clause += " AND a.comment =?"
		condition = append(condition, val)
	}
	if val, ok := cond["state"]; ok {
		clause += " AND a.state =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mark"]; ok {
		clause += " AND a.mark =?"
		condition = append(condition, val)
	}
	if val, ok := cond["operator"]; ok {
		clause += " AND a.operator =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.comment,a.state,a.mark,a.operator,a.deleted,a.created_at,a.updated_at FROM common_configs a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetCommonConfigsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddCommonConfig(c context.Context, node sqalx.Node, item *model.CommonConfig) (err error) {
	sqlInsert := "INSERT INTO common_configs( id,name,comment,state,mark,operator,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Name, item.Comment, item.State, item.Mark, item.Operator, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddCommonConfigs err(%+v), item(%+v)", err, item))
		return
	}
	return
}

// Update update a exist record
func (p *Dao) UpdateCommonConfig(c context.Context, node sqalx.Node, item *model.CommonConfig) (err error) {
	sqlUpdate := "UPDATE common_configs SET name=?,comment=?,state=?,mark=?,operator=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Name, item.Comment, item.State, item.Mark, item.Operator, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateCommonConfigs err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelCommonConfig(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE common_configs SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelCommonConfigs err(%+v), item(%+v)", err, id))
		return
	}

	return
}
