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
func (p *Dao) GetRecommendTopics(c context.Context, node sqalx.Node) (items []*model.RecommendTopic, err error) {
	items = make([]*model.RecommendTopic, 0)
	sqlSelect := "SELECT a.id,a.topic_id,a.deleted,a.created_at,a.updated_at FROM recommend_topics a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecommendTopics err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetRecommendTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.RecommendTopic, err error) {
	items = make([]*model.RecommendTopic, 0)
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

	sqlSelect := fmt.Sprintf("SELECT a.id,a.topic_id,a.deleted,a.created_at,a.updated_at FROM recommend_topics a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecommendTopicsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetRecommendTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.RecommendTopic, err error) {
	item = new(model.RecommendTopic)
	sqlSelect := "SELECT a.id,a.topic_id,a.deleted,a.created_at,a.updated_at FROM recommend_topics a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetRecommendTopicByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetRecommendTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.RecommendTopic, err error) {
	item = new(model.RecommendTopic)
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

	sqlSelect := fmt.Sprintf("SELECT a.id,a.topic_id,a.deleted,a.created_at,a.updated_at FROM recommend_topics a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetRecommendTopicsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddRecommendTopic(c context.Context, node sqalx.Node, item *model.RecommendTopic) (err error) {
	sqlInsert := "INSERT INTO recommend_topics( id,topic_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TopicID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddRecommendTopics err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateRecommendTopic(c context.Context, node sqalx.Node, item *model.RecommendTopic) (err error) {
	sqlUpdate := "UPDATE recommend_topics SET topic_id=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicID, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateRecommendTopics err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelRecommendTopic(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE recommend_topics SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelRecommendTopics err(%+v), item(%+v)", err, id))
		return
	}

	return
}
