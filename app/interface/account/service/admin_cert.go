package service

import (
	"context"
	"net/url"
	"strconv"
	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
)

// GetWorkCertsPaged 分页获取所有认证信息
func (p *Service) GetWorkCertsPaged(c context.Context, status int, limit, offset int) (resp *model.WorkCertsPagedResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	arg := &account.WorkCertPagedReq{Aid: aid, Status: int32(status), Limit: int32(limit), Offset: int32(offset)}

	var ret *account.WorkCertPagedResp
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

		var id *account.IDCertInfo
		if id, err = p.d.GetIDCert(c, v.AccountID); err != nil {
			return
		}

		item.IDName = id.Name

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/admin/certification/workcert/list", url.Values{
		"status": []string{strconv.Itoa(status)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/admin/certification/workcert/list", url.Values{
		"status": []string{strconv.Itoa(status)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset + limit)},
	}); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	return
}

// AuditWorkCert 审核工作认证
func (s *Service) AuditWorkCert(c *mars.Context, arg *model.ArgAuditWorkCert) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	req := &account.AuditWorkCertReq{
		AccountID:   arg.AccountID,
		Approve:     arg.Approve,
		AuditResult: arg.AuditResult,
		Aid:         aid,
	}
	if err = s.d.AuditWorkCert(c, req); err != nil {
		return err
	}
	return
}

// GetWorkCertHistoriesPaged 分页获取用户工作认证历史记录
func (p *Service) GetWorkCertHistoriesPaged(c context.Context, accountID int64, limit, offset int) (resp *model.WorkCertHistoriesResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	arg := &account.WorkCertHistoriesPagedReq{Aid: aid, AccountID: accountID, Limit: int32(limit), Offset: int32(offset)}

	var ret *account.WorkCertHistoriesPagedResp
	if ret, err = p.d.GetWorkCertHistoriesPaged(c, arg); err != nil {
		return
	}

	resp = &model.WorkCertHistoriesResp{
		Paging: &model.Paging{},
		Items:  make([]*model.WorkCertHistoryItem, len(ret.Items)),
	}

	for i, v := range ret.Items {
		item := &model.WorkCertHistoryItem{
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

		var m *account.BaseInfoReply
		if m, err = p.d.GetAccountBaseInfo(c, v.ManagerID); err != nil {
			return
		}

		item.Manager = &model.Member{
			ID:           m.ID,
			UserName:     m.UserName,
			Avatar:       m.Avatar,
			Introduction: m.Introduction,
		}

		var id *account.IDCertInfo
		if id, err = p.d.GetIDCert(c, v.AccountID); err != nil {
			return
		}

		item.IDName = id.Name

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/admin/certification/workcert/history/list", url.Values{
		"account_id": []string{strconv.FormatInt(accountID, 10)},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/admin/certification/workcert/history/list", url.Values{
		"account_id": []string{strconv.FormatInt(accountID, 10)},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset + limit)},
	}); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	return
}
