package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetAuthTopics(c context.Context, node sqalx.Node) (items []*model.AuthTopic, err error) {
	items = make([]*model.AuthTopic, 0)
	sqlSelect := "SELECT a.* FROM auth_topics a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopics err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetAuthTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AuthTopic, err error) {
	items = make([]*model.AuthTopic, 0)
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
	if val, ok := cond["to_topic_id"]; ok {
		clause += " AND a.to_topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["permission"]; ok {
		clause += " AND a.permission =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM auth_topics a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopicsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetAuthTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.AuthTopic, err error) {
	item = new(model.AuthTopic)
	sqlSelect := "SELECT a.* FROM auth_topics a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopicByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetAuthTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AuthTopic, err error) {
	item = new(model.AuthTopic)
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
	if val, ok := cond["to_topic_id"]; ok {
		clause += " AND a.to_topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["permission"]; ok {
		clause += " AND a.permission =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM auth_topics a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopicsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error) {
	sqlInsert := "INSERT INTO auth_topics( id,topic_id,to_topic_id,permission,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TopicID, item.ToTopicID, item.Permission, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAuthTopics err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error) {
	sqlUpdate := "UPDATE auth_topics SET topic_id=?,to_topic_id=?,permission=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicID, item.ToTopicID, item.Permission, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAuthTopics err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelAuthTopic(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE auth_topics SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAuthTopics err(%+v), item(%+v)", err, id))
		return
	}

	return
}
