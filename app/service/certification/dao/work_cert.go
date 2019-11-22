package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/certification/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetWorkCertificationsPaged(c context.Context, node sqalx.Node, cond map[string]interface{}, limit, offset int32) (items []*model.WorkCertification, err error) {
	items = make([]*model.WorkCertification, 0)
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
	if val, ok := cond["status"]; ok {
		clause += " AND a.status =?"
		condition = append(condition, val)
	}
	if val, ok := cond["work_pic"]; ok {
		clause += " AND a.work_pic =?"
		condition = append(condition, val)
	}
	if val, ok := cond["other_pic"]; ok {
		clause += " AND a.other_pic =?"
		condition = append(condition, val)
	}
	if val, ok := cond["company"]; ok {
		clause += " AND a.company =?"
		condition = append(condition, val)
	}
	if val, ok := cond["department"]; ok {
		clause += " AND a.department =?"
		condition = append(condition, val)
	}
	if val, ok := cond["position"]; ok {
		clause += " AND a.position =?"
		condition = append(condition, val)
	}
	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at =?"
		condition = append(condition, val)
	}
	if val, ok := cond["audit_result"]; ok {
		clause += " AND a.audit_result =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM work_certifications a WHERE a.deleted=0 %s ORDER BY a.status ASC, a.id DESC LIMIT ?,?", clause)

	condition = append(condition, offset)
	condition = append(condition, limit)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertificationsPaged err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// GetAll get all records
func (p *Dao) GetWorkCertifications(c context.Context, node sqalx.Node) (items []*model.WorkCertification, err error) {
	items = make([]*model.WorkCertification, 0)
	sqlSelect := "SELECT a.* FROM work_certifications a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertifications err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetWorkCertificationsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.WorkCertification, err error) {
	items = make([]*model.WorkCertification, 0)
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
	if val, ok := cond["status"]; ok {
		clause += " AND a.status =?"
		condition = append(condition, val)
	}
	if val, ok := cond["work_pic"]; ok {
		clause += " AND a.work_pic =?"
		condition = append(condition, val)
	}
	if val, ok := cond["other_pic"]; ok {
		clause += " AND a.other_pic =?"
		condition = append(condition, val)
	}
	if val, ok := cond["company"]; ok {
		clause += " AND a.company =?"
		condition = append(condition, val)
	}
	if val, ok := cond["department"]; ok {
		clause += " AND a.department =?"
		condition = append(condition, val)
	}
	if val, ok := cond["position"]; ok {
		clause += " AND a.position =?"
		condition = append(condition, val)
	}
	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at =?"
		condition = append(condition, val)
	}
	if val, ok := cond["audit_result"]; ok {
		clause += " AND a.audit_result =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM work_certifications a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertificationsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetWorkCertificationByID(c context.Context, node sqalx.Node, id int64) (item *model.WorkCertification, err error) {
	item = new(model.WorkCertification)
	sqlSelect := "SELECT a.* FROM work_certifications a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertificationByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetWorkCertificationByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.WorkCertification, err error) {
	item = new(model.WorkCertification)
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
	if val, ok := cond["status"]; ok {
		clause += " AND a.status =?"
		condition = append(condition, val)
	}
	if val, ok := cond["work_pic"]; ok {
		clause += " AND a.work_pic =?"
		condition = append(condition, val)
	}
	if val, ok := cond["other_pic"]; ok {
		clause += " AND a.other_pic =?"
		condition = append(condition, val)
	}
	if val, ok := cond["company"]; ok {
		clause += " AND a.company =?"
		condition = append(condition, val)
	}
	if val, ok := cond["department"]; ok {
		clause += " AND a.department =?"
		condition = append(condition, val)
	}
	if val, ok := cond["position"]; ok {
		clause += " AND a.position =?"
		condition = append(condition, val)
	}
	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at =?"
		condition = append(condition, val)
	}
	if val, ok := cond["audit_result"]; ok {
		clause += " AND a.audit_result =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM work_certifications a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertificationsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddWorkCertification(c context.Context, node sqalx.Node, item *model.WorkCertification) (err error) {
	sqlInsert := "INSERT INTO work_certifications( id,account_id,status,work_pic,other_pic,company,department,position,expires_at,audit_result,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.Status, item.WorkPic, item.OtherPic, item.Company, item.Department, item.Position, item.ExpiresAt, item.AuditResult, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddWorkCertifications err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateWorkCertification(c context.Context, node sqalx.Node, item *model.WorkCertification) (err error) {
	sqlUpdate := "UPDATE work_certifications SET account_id=?,status=?,work_pic=?,other_pic=?,company=?,department=?,position=?,expires_at=?,audit_result=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.Status, item.WorkPic, item.OtherPic, item.Company, item.Department, item.Position, item.ExpiresAt, item.AuditResult, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateWorkCertifications err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelWorkCertification(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE work_certifications SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelWorkCertifications err(%+v), item(%+v)", err, id))
		return
	}

	return
}
