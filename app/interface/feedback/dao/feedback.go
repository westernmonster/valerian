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
func (p *Dao) GetFeedbacks(c context.Context, node sqalx.Node) (items []*model.Feedback, err error) {
	items = make([]*model.Feedback, 0)
	sqlSelect := "SELECT a.id,a.target_id,a.target_type,a.target_desc,a.feedback_type,a.feedback_desc,a.created_by,a.deleted,a.created_at,a.updated_at,a.verify_status,a.verify_desc FROM feedbacks a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbacks err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetFeedbacksByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Feedback, err error) {
	items = make([]*model.Feedback, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_id"]; ok {
		clause += " AND a.target_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_type"]; ok {
		clause += " AND a.target_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_desc"]; ok {
		clause += " AND a.target_desc =?"
		condition = append(condition, val)
	}
	if val, ok := cond["feedback_type"]; ok {
		clause += " AND a.feedback_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["feedback_desc"]; ok {
		clause += " AND a.feedback_desc =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}
	if val, ok := cond["verify_status"]; ok {
		clause += " AND a.verify_status =?"
		condition = append(condition, val)
	}
	if val, ok := cond["verify_desc"]; ok {
		clause += " AND a.verify_desc =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.target_id,a.target_type,a.target_desc,a.feedback_type,a.feedback_desc,a.created_by,a.deleted,a.created_at,a.updated_at,a.verify_status,a.verify_desc FROM feedbacks a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbacksByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetFeedbackByID(c context.Context, node sqalx.Node, id int64) (item *model.Feedback, err error) {
	item = new(model.Feedback)
	sqlSelect := "SELECT a.id,a.target_id,a.target_type,a.target_desc,a.feedback_type,a.feedback_desc,a.created_by,a.deleted,a.created_at,a.updated_at,a.verify_status,a.verify_desc FROM feedbacks a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbackByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetFeedbackByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Feedback, err error) {
	item = new(model.Feedback)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_id"]; ok {
		clause += " AND a.target_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_type"]; ok {
		clause += " AND a.target_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_desc"]; ok {
		clause += " AND a.target_desc =?"
		condition = append(condition, val)
	}
	if val, ok := cond["feedback_type"]; ok {
		clause += " AND a.feedback_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["feedback_desc"]; ok {
		clause += " AND a.feedback_desc =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}
	if val, ok := cond["verify_status"]; ok {
		clause += " AND a.verify_status =?"
		condition = append(condition, val)
	}
	if val, ok := cond["verify_desc"]; ok {
		clause += " AND a.verify_desc =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.target_id,a.target_type,a.target_desc,a.feedback_type,a.feedback_desc,a.created_by,a.deleted,a.created_at,a.updated_at,a.verify_status,a.verify_desc FROM feedbacks a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbacksByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddFeedback(c context.Context, node sqalx.Node, item *model.Feedback) (err error) {
	sqlInsert := "INSERT INTO feedbacks( id,target_id,target_type,target_desc,feedback_type,feedback_desc,created_by,deleted,created_at,updated_at,verify_status,verify_desc) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TargetID, item.TargetType, item.TargetDesc, item.FeedbackType, item.FeedbackDesc, item.CreatedBy, item.Deleted, item.CreatedAt, item.UpdatedAt, item.VerifyStatus, item.VerifyDesc); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddFeedbacks err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateFeedback(c context.Context, node sqalx.Node, item *model.Feedback) (err error) {
	sqlUpdate := "UPDATE feedbacks SET target_id=?,target_type=?,target_desc=?,feedback_type=?,feedback_desc=?,created_by=?,updated_at=?,verify_status=?,verify_desc=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TargetID, item.TargetType, item.TargetDesc, item.FeedbackType, item.FeedbackDesc, item.CreatedBy, item.UpdatedAt, item.VerifyStatus, item.VerifyDesc, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFeedbacks err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelFeedback(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE feedbacks SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeedbacks err(%+v), item(%+v)", err, id))
		return
	}

	return
}
