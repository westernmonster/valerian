package dao

import (
	"context"
	"fmt"
	"valerian/app/service/certification/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Insert insert a new record
func (p *Dao) AddWorkCertHistory(c context.Context, node sqalx.Node, item *model.WorkCertHistory) (err error) {
	sqlInsert := "INSERT INTO work_cert_history( id,account_id,status,work_pic,other_pic,company,department,position," +
		"expires_at,audit_result,manager_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.Status, item.WorkPic, item.OtherPic,
		item.Company, item.Department, item.Position, item.ExpiresAt, item.AuditResult, item.ManagerID,
		item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddWorkCertHistory err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// GetAll get all records
func (p *Dao) ListWorkCertHistoryByAccount(c context.Context, node sqalx.Node) (items []*model.WorkCertHistory, err error) {
	items = make([]*model.WorkCertHistory, 0)
	sqlSelect := "SELECT a.id,a.account_id,a.status,a.work_pic,a.other_pic,a.company,a.department,a.position," +
		"a.expires_at,a.audit_result,a.manager_id,a.deleted,a.created_at,a.updated_at FROM work_cert_history a " +
		"WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertHistory err(%+v)", err))
		return
	}
	return
}