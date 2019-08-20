package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/feedback/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetFeedbackTypes(c context.Context, node sqalx.Node) (items []*model.FeedbackType, err error) {
	items = make([]*model.FeedbackType, 0)
	sqlSelect := "SELECT a.* FROM feedback_types a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbackTypes err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetFeedbackTypesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.FeedbackType, err error) {
	items = make([]*model.FeedbackType, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["type"]; ok {
		clause += " AND a.type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM feedback_types a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbackTypesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetFeedbackTypeByID(c context.Context, node sqalx.Node, id int) (item *model.FeedbackType, err error) {
	item = new(model.FeedbackType)
	sqlSelect := "SELECT a.* FROM feedback_types a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbackTypeByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetFeedbackTypeByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.FeedbackType, err error) {
	item = new(model.FeedbackType)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["type"]; ok {
		clause += " AND a.type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM feedback_types a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbackTypesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddFeedbackType(c context.Context, node sqalx.Node, item *model.FeedbackType) (err error) {
	sqlInsert := "INSERT INTO feedback_types( id,type,name,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Type, item.Name, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddFeedbackTypes err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateFeedbackType(c context.Context, node sqalx.Node, item *model.FeedbackType) (err error) {
	sqlUpdate := "UPDATE feedback_types SET type=?,name=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Type, item.Name, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFeedbackTypes err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelFeedbackType(c context.Context, node sqalx.Node, id int) (err error) {
	sqlDelete := "UPDATE feedback_types SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeedbackTypes err(%+v), item(%+v)", err, id))
		return
	}

	return
}
