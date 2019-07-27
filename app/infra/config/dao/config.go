package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"valerian/app/infra/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func Int64Array2StringArray(req []int64) (resp []string) {
	resp = make([]string, 0)

	for _, v := range req {
		resp = append(resp, strconv.FormatInt(v, 10))
	}

	return
}

// GetAll get all records
func (p *Dao) GetConfigs(c context.Context, node sqalx.Node) (items []*model.Config, err error) {
	items = make([]*model.Config, 0)
	sqlSelect := "SELECT a.* FROM configs a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetConfigs err(%+v)", err))
		return
	}
	return
}

func (p *Dao) GetConfigsByIDs(c context.Context, node sqalx.Node, ids []int64) (items []*model.Config, err error) {
	items = make([]*model.Config, 0)
	strIDs := Int64Array2StringArray(ids)

	sqlSelect := "SELECT a.* FROM configs a WHERE a.deleted=0 AND a.state=? AND a.id IN(?)"

	if err = node.SelectContext(c, &items, sqlSelect, model.ConfigEnd, strings.Join(strIDs, ",")); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetConfigsByIDs err(%+v), ids(%+v)", err, strIDs))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetConfigsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Config, err error) {
	items = make([]*model.Config, 0)
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
	if val, ok := cond["comment"]; ok {
		clause += " AND a.comment =?"
		condition = append(condition, val)
	}
	if val, ok := cond["from"]; ok {
		clause += " AND a.from =?"
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM configs a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetConfigsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetConfigByID(c context.Context, node sqalx.Node, id int64) (item *model.Config, err error) {
	item = new(model.Config)
	sqlSelect := "SELECT a.* FROM configs a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetConfigByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetConfigByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Config, err error) {
	item = new(model.Config)
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
	if val, ok := cond["comment"]; ok {
		clause += " AND a.comment =?"
		condition = append(condition, val)
	}
	if val, ok := cond["from"]; ok {
		clause += " AND a.from =?"
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM configs a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetConfigsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddConfig(c context.Context, node sqalx.Node, item *model.Config) (err error) {
	sqlInsert := "INSERT INTO configs( id,app_id,name,comment,from,state,mark,operator,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AppID, item.Name, item.Comment, item.From, item.State, item.Mark, item.Operator, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddConfigs err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateConfig(c context.Context, node sqalx.Node, item *model.Config) (err error) {
	sqlUpdate := "UPDATE configs SET app_id=?,name=?,comment=?,from=?,state=?,mark=?,operator=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AppID, item.Name, item.Comment, item.From, item.State, item.Mark, item.Operator, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateConfigs err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelConfig(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE configs SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelConfigs err(%+v), item(%+v)", err, id))
		return
	}

	return
}
