package service

import (
	"context"

	"valerian/app/interface/certification/model"
	account "valerian/app/service/account/api"
	certification "valerian/app/service/certification/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetWorkCertsPaged(c context.Context, status *int32, limit, offset int32) (resp *model.WorkCertsPagedResp, err error) {

	arg := &certification.WorkCertPagedReq{Limit: limit, Offset: offset}
	if status != nil {
		arg.Status = &certification.WorkCertPagedReq_StatusValue{*status}
	}

	var ret *certification.WorkCertPagedResp
	if ret, err = p.d.GetWorkCertsPaged(c, arg); err != nil {
		return
	}

	resp = &model.WorkCertsPagedResp{
		Paging: &model.Paging{},
		Items:  make([]*model.WorkCertItem, len(ret.Items)),
	}

	for i, v := range ret.Items {
		item := &model.WorkCertItem{
			ID:         v.ID,
			Status:     v.Status,
			WorkPic:    v.WorkPic,
			OtherPic:   v.OtherPic,
			Company:    v.Company,
			Department: v.Department,
			Position:   v.Position,
			ExpiresAt:  v.ExpiresAt,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, v.AccountID); err != nil {
			return
		}

		item.Member = &model.Member{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		var id *certification.IDCertInfo
		if id, err = p.d.GetIDCert(c, v.AccountID); err != nil {
			return
		}

		item.IDName = id.Name

		resp.Items[i] = item
	}

	return
}

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
