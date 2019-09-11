package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetDiscussCategories(c context.Context, node sqalx.Node) (items []*model.DiscussCategory, err error) {
	items = make([]*model.DiscussCategory, 0)
	sqlSelect := "SELECT a.* FROM discuss_categories a WHERE a.deleted=0 ORDER BY a.seq "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussCategories err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetDiscussCategoriesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.DiscussCategory, err error) {
	items = make([]*model.DiscussCategory, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM discuss_categories a WHERE a.deleted=0 %s ORDER BY a.seq", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussCategoriesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetDiscussCategoryByID(c context.Context, node sqalx.Node, id int64) (item *model.DiscussCategory, err error) {
	item = new(model.DiscussCategory)
	sqlSelect := "SELECT a.* FROM discuss_categories a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussCategoryByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetDiscussCategoryByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.DiscussCategory, err error) {
	item = new(model.DiscussCategory)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM discuss_categories a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussCategoriesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error) {
	sqlInsert := "INSERT INTO discuss_categories( id,topic_id,seq,name,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TopicID, item.Seq, item.Name, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddDiscussCategories err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error) {
	sqlUpdate := "UPDATE discuss_categories SET topic_id=?,seq=?,name=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicID, item.Seq, item.Name, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateDiscussCategories err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelDiscussCategory(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE discuss_categories SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelDiscussCategories err(%+v), item(%+v)", err, id))
		return
	}

	return
}
