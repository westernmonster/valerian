package service

import (
	"context"

	"valerian/app/interface/certification/model"
	certification "valerian/app/service/certification/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

// RequestWorkCert 提交工作认证信息
func (p *Service) RequestWorkCert(c context.Context, arg *model.ArgWorkCert) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &certification.WorkCertReq{
		AccountID:  aid,
		WorkPic:    arg.WorkPic,
		OtherPic:   arg.OtherPic,
		Company:    arg.Company,
		Department: arg.Department,
		Position:   arg.Position,
		ExpiresAt:  arg.ExpiresAt,
	}

	if err = p.d.RequestWorkCert(c, item); err != nil {
		return
	}

	return
}

// GetWorkCert 获取当前用户工作认证信息
func (p *Service) GetWorkCert(c context.Context) (resp *model.WorkCertResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var v *certification.WorkCertInfo
	if v, err = p.d.GetWorkCert(c, aid); err != nil {
		return
	}

	var i *certification.IDCertInfo
	if i, err = p.d.GetIDCert(c, aid); err != nil {
		return
	}

	resp = &model.WorkCertResp{
		AccountID:  v.AccountID,
		IDName:     i.Name,
		Status:     int(v.Status),
		WorkPic:    v.WorkPic,
		OtherPic:   v.OtherPic,
		Company:    v.Company,
		Department: v.Department,
		Position:   v.Position,
		ExpiresAt:  v.ExpiresAt,
		CreatedAt:  v.CreatedAt,
		AuditAt:    v.UpdatedAt,
		Result:     v.AuditResult,
	}

	return

}
