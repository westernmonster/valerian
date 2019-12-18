package dao

import (
	"context"
	"fmt"
	"valerian/app/admin/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
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

	condition = append(condition, limit)
	condition = append(condition, offset)

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.status,a.work_pic,a.other_pic,a.company,a.department," +
		"a.position,a.expires_at,a.audit_result,a.deleted,a.created_at,a.updated_at FROM work_certifications a " +
		"WHERE a.deleted=0 %s ORDER BY a.id DESC limit ?,?", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertificationsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}
