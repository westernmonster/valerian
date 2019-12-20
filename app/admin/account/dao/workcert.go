package dao

import (
	"context"
	"fmt"
	"valerian/app/admin/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

// GetAllByCondition get records by condition
func (p *Dao) GetWorkCertificationsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}, limit, offset int) (items []*model.WorkCertification, err error) {
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

	condition = append(condition, offset)
	condition = append(condition, limit)

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.status,a.work_pic,a.other_pic,a.company,a.department,"+
		"a.position,a.expires_at,a.audit_result,a.deleted,a.created_at,a.updated_at FROM work_certifications a "+
		"WHERE a.deleted=0 %s ORDER BY a.id DESC limit ?,?", clause)
	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertificationsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

func (p *Dao) GetWorkCertHistorysByAccount(c *mars.Context, node sqalx.Node, cond map[string]interface{}, limit, offset int) (items []*model.WorkCertHistory, err error) {

	items = make([]*model.WorkCertHistory, 0)
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
	if val, ok := cond["manager_id"]; ok {
		clause += " AND a.manager_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.status,a.work_pic,a.other_pic,a.company,a.department,a.position,a.expires_at,a.audit_result,a.manager_id,a.deleted,a.created_at,a.updated_at FROM work_cert_history a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertHistoryByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}
