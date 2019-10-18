package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/dm/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetTopicMembers(c context.Context, node sqalx.Node) (items []*model.TopicMember, err error) {
	items = make([]*model.TopicMember, 0)
	sqlSelect := "SELECT a.* FROM topic_members a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembers err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetTopicMembersByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicMember, err error) {
	items = make([]*model.TopicMember, 0)
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
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_members a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetTopicMemberByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicMember, err error) {
	item = new(model.TopicMember)
	sqlSelect := "SELECT a.* FROM topic_members a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMemberByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetTopicMemberByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicMember, err error) {
	item = new(model.TopicMember)
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
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_members a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error) {
	sqlInsert := "INSERT INTO topic_members( id,topic_id,account_id,role,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TopicID, item.AccountID, item.Role, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicMembers err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error) {
	sqlUpdate := "UPDATE topic_members SET topic_id=?,account_id=?,role=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicID, item.AccountID, item.Role, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicMembers err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopicMember(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_members SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicMembers err(%+v), item(%+v)", err, id))
		return
	}

	return
}
