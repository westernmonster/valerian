package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/admin/feedback/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Update update a exist record
func (p *Dao) UpdateFeedbackVerify(c context.Context, node sqalx.Node, feedbackID int64, verifyStatus int32, verifyDesc string) (err error) {
	sqlUpdate := "UPDATE feedbacks SET verify_status=?,verify_desc=? WHERE id=?"
	_, err = node.ExecContext(c, sqlUpdate, verifyStatus, verifyDesc, feedbackID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFeedbackVerify err(%+v)", err))
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
func (p *Dao) GetFeedbacksByCondPaged(c context.Context, node sqalx.Node, cond map[string]interface{}, limit, offset int) (items []*model.Feedback, err error) {
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

	condition = append(condition, offset)
	condition = append(condition, limit)

	sqlSelect := fmt.Sprintf("SELECT a.id,a.target_id,a.target_type,a.target_desc,a.feedback_type,a.feedback_desc,"+
		"a.created_by,a.deleted,a.created_at,a.updated_at,a.verify_status,a.verify_desc FROM feedbacks a "+
		"WHERE a.deleted=0 %s ORDER BY a.id DESC  limit ?,?", clause)
	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbacksByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

func (p *Dao) GetReportPaged(c context.Context, node sqalx.Node, cond map[string]interface{}, limit, offset int) (items []*model.Feedback, err error) {
	items = make([]*model.Feedback, 0)
	condition := make([]interface{}, 0)
	clause := ""

	// 非举报类型
	notAccuseTypeList := []int{11, 12, 13}

	for _, notAccuseType := range notAccuseTypeList {
		clause += " AND a.feedback_type <>? "
		condition = append(condition, notAccuseType)
	}

	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}

	condition = append(condition, offset)
	condition = append(condition, limit)

	sqlSelect := fmt.Sprintf("SELECT a.id,a.target_id,a.target_type,a.target_desc,a.feedback_type,a.feedback_desc,"+
		"a.created_by,a.deleted,a.created_at,a.updated_at,a.verify_status,a.verify_desc FROM feedbacks a "+
		"WHERE a.deleted=0 %s "+
		"ORDER BY a.created_at DESC  limit ?,?", clause)
	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetReportByCondPaged err(%+v), condition(%+v)", err, condition))
		return
	}
	return
}

func (p *Dao) GetBeReportedPaged(c context.Context, node sqalx.Node, cond map[string]interface{}, limit, offset int) (items []*model.Feedback, err error) {
	items = make([]*model.Feedback, 0)
	condition := make([]interface{}, 0)
	clause := ""

	// 非举报类型
	notAccuseTypeList := []int{11, 12, 13}

	for _, notAccuseType := range notAccuseTypeList {
		clause += " AND a.feedback_type <>? "
		condition = append(condition, notAccuseType)
	}

	if val, ok := cond["target_user_id"]; ok {
		clause += " AND a.target_user_id =?"
		condition = append(condition, val)
	}

	condition = append(condition, offset)
	condition = append(condition, limit)

	sqlSelect := fmt.Sprintf("SELECT a.id,a.target_id,a.target_type,a.target_desc,a.feedback_type,a.feedback_desc,"+
		"a.created_by,a.deleted,a.created_at,a.updated_at,a.verify_status,a.verify_desc FROM feedbacks a "+
		"WHERE a.deleted=0 %s "+
		"ORDER BY a.created_at DESC  limit ?,?", clause)
	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetBeReportedByCondPaged err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}
