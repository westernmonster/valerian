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
func (p *Dao) GetTopicInviteRequests(c context.Context, node sqalx.Node) (items []*model.TopicInviteRequest, err error) {
	items = make([]*model.TopicInviteRequest, 0)
	sqlSelect := "SELECT a.* FROM topic_invite_requests a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicInviteRequests err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetTopicInviteRequestsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicInviteRequest, err error) {
	items = make([]*model.TopicInviteRequest, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["status"]; ok {
		clause += " AND a.status =?"
		condition = append(condition, val)
	}
	if val, ok := cond["from_account_id"]; ok {
		clause += " AND a.from_account_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_invite_requests a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicInviteRequestsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetTopicInviteRequestByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicInviteRequest, err error) {
	item = new(model.TopicInviteRequest)
	sqlSelect := "SELECT a.* FROM topic_invite_requests a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicInviteRequestByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetTopicInviteRequestByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicInviteRequest, err error) {
	item = new(model.TopicInviteRequest)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["status"]; ok {
		clause += " AND a.status =?"
		condition = append(condition, val)
	}
	if val, ok := cond["from_account_id"]; ok {
		clause += " AND a.from_account_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_invite_requests a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicInviteRequestsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTopicInviteRequest(c context.Context, node sqalx.Node, item *model.TopicInviteRequest) (err error) {
	sqlInsert := "INSERT INTO topic_invite_requests( id,account_id,topic_id,status,deleted,created_at,updated_at,from_account_id) VALUES ( ?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.TopicID, item.Status, item.Deleted, item.CreatedAt, item.UpdatedAt, item.FromAccountID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicInviteRequests err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicInviteRequest(c context.Context, node sqalx.Node, item *model.TopicInviteRequest) (err error) {
	sqlUpdate := "UPDATE topic_invite_requests SET account_id=?,topic_id=?,status=?,updated_at=?,from_account_id=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.TopicID, item.Status, item.UpdatedAt, item.FromAccountID, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicInviteRequests err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopicInviteRequest(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_invite_requests SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicInviteRequests err(%+v), item(%+v)", err, id))
		return
	}

	return
}
