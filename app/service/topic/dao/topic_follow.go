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
func (p *Dao) GetTopicFollowRequests(c context.Context, node sqalx.Node) (items []*model.TopicFollowRequest, err error) {
	items = make([]*model.TopicFollowRequest, 0)
	sqlSelect := "SELECT a.id,a.account_id,a.topic_id,a.status,a.deleted,a.created_at,a.updated_at,a.reason,a.allow_view_cert,a.reject_reason FROM topic_follow_requests a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicFollowRequests err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetTopicFollowRequestsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicFollowRequest, err error) {
	items = make([]*model.TopicFollowRequest, 0)
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
	if val, ok := cond["reason"]; ok {
		clause += " AND a.reason =?"
		condition = append(condition, val)
	}
	if val, ok := cond["allow_view_cert"]; ok {
		clause += " AND a.allow_view_cert =?"
		condition = append(condition, val)
	}
	if val, ok := cond["reject_reason"]; ok {
		clause += " AND a.reject_reason =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.topic_id,a.status,a.deleted,a.created_at,a.updated_at,a.reason,a.allow_view_cert,a.reject_reason FROM topic_follow_requests a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicFollowRequestsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetTopicFollowRequestByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicFollowRequest, err error) {
	item = new(model.TopicFollowRequest)
	sqlSelect := "SELECT a.id,a.account_id,a.topic_id,a.status,a.deleted,a.created_at,a.updated_at,a.reason,a.allow_view_cert,a.reject_reason FROM topic_follow_requests a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicFollowRequestByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetTopicFollowRequestByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicFollowRequest, err error) {
	item = new(model.TopicFollowRequest)
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
	if val, ok := cond["reason"]; ok {
		clause += " AND a.reason =?"
		condition = append(condition, val)
	}
	if val, ok := cond["allow_view_cert"]; ok {
		clause += " AND a.allow_view_cert =?"
		condition = append(condition, val)
	}
	if val, ok := cond["reject_reason"]; ok {
		clause += " AND a.reject_reason =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.topic_id,a.status,a.deleted,a.created_at,a.updated_at,a.reason,a.allow_view_cert,a.reject_reason FROM topic_follow_requests a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicFollowRequestsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error) {
	sqlInsert := "INSERT INTO topic_follow_requests( id,account_id,topic_id,status,deleted,created_at,updated_at,reason,allow_view_cert,reject_reason) VALUES ( ?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.TopicID, item.Status, item.Deleted, item.CreatedAt, item.UpdatedAt, item.Reason, item.AllowViewCert, item.RejectReason); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicFollowRequests err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error) {
	sqlUpdate := "UPDATE topic_follow_requests SET account_id=?,topic_id=?,status=?,updated_at=?,reason=?,allow_view_cert=?,reject_reason=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.TopicID, item.Status, item.UpdatedAt, item.Reason, item.AllowViewCert, item.RejectReason, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicFollowRequests err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopicFollowRequest(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_follow_requests SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicFollowRequests err(%+v), item(%+v)", err, id))
		return
	}

	return
}
