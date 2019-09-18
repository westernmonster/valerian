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
func (p *Dao) GetTags(c context.Context, node sqalx.Node) (items []*model.Tag, err error) {
	items = make([]*model.Tag, 0)
	sqlSelect := "SELECT a.* FROM tags a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTags err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetTagsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Tag, err error) {
	items = make([]*model.Tag, 0)
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
	if val, ok := cond["config_ids"]; ok {
		clause += " AND a.config_ids =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mark"]; ok {
		clause += " AND a.mark =?"
		condition = append(condition, val)
	}
	if val, ok := cond["force"]; ok {
		clause += " AND a.force =?"
		condition = append(condition, val)
	}
	if val, ok := cond["operator"]; ok {
		clause += " AND a.operator =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM tags a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTagsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetTagByID(c context.Context, node sqalx.Node, id int64) (item *model.Tag, err error) {
	item = new(model.Tag)
	sqlSelect := "SELECT a.* FROM tags a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTagByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetTagByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Tag, err error) {
	item = new(model.Tag)
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
	if val, ok := cond["config_ids"]; ok {
		clause += " AND a.config_ids =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mark"]; ok {
		clause += " AND a.mark =?"
		condition = append(condition, val)
	}
	if val, ok := cond["force"]; ok {
		clause += " AND a.force =?"
		condition = append(condition, val)
	}
	if val, ok := cond["operator"]; ok {
		clause += " AND a.operator =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM tags a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTagsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTag(c context.Context, node sqalx.Node, item *model.Tag) (err error) {
	sqlInsert := "INSERT INTO tags( id,app_id,config_ids,mark,`force`,`operator`,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AppID, item.ConfigIds, item.Mark, item.Force, item.Operator, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTags err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTag(c context.Context, node sqalx.Node, item *model.Tag) (err error) {
	sqlUpdate := "UPDATE tags SET app_id=?,config_ids=?,mark=?,force=?,operator=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AppID, item.ConfigIds, item.Mark, item.Force, item.Operator, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTags err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTag(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE tags SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTags err(%+v), item(%+v)", err, id))
		return
	}

	return
}
