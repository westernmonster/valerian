package dao

import (
	"context"
	"fmt"
	"valerian/app/admin/account/model"
	certification "valerian/app/service/certification/api"
	"valerian/library/log"
)

func (p *Dao) SetWorkCert(c context.Context, arg *model.ArgWorkCert, managerID int64) (info *certification.EmptyStruct, err error) {

	argRpc := certification.AuditWorkCertReq{AccountID: arg.AccountID,
		ManagerID:   managerID,
		Approve:     arg.Approve,
		AuditResult: arg.AuditResult}
	if info, err = p.certificationRPC.AuditWorkCert(c, &argRpc); err == nil {
		log.For(c).Error(fmt.Sprintf("dao.SetWorkCert err(%+v)", err))
	}
	return
}
