package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	identify "valerian/app/service/identify/api/grpc"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

// GetAllAccountsPaged 添加账户
func (p *Service) GetAllAccountsPaged(c context.Context, limit, offset int) (resp *model.AdminAccountsResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	req := &account.AccountsPagedReq{
		Aid:    aid,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	var data *account.AdminAccountsResp
	if data, err = p.d.AllAccountsPaged(c, req); err != nil {
		return
	}

	resp = &model.AdminAccountsResp{
		Items:  make([]*model.AdminAccountItem, 0),
		Paging: &model.Paging{},
	}

	for _, v := range data.Items {
		item := &model.AdminAccountItem{
			ID:             v.ID,
			Email:          v.Email,
			Mobile:         v.Mobile,
			Prefix:         v.Prefix,
			UserName:       v.UserName,
			Gender:         v.Gender,
			Location:       v.Location,
			LocationString: v.LocationString,
			Introduction:   v.Introduction,
			Avatar:         v.Avatar,
			IDCert:         v.IDCert,
			WorkCert:       v.WorkCert,
			IsOrg:          v.IsOrg,
			IsVIP:          v.IsVIP,
			Company:        v.Company,
			Position:       v.Position,
			IsLock:         v.IsLock,
		}

		resp.Items = append(resp.Items, item)
	}

	if resp.Paging.Prev, err = genURL("/api/v1/admin/account/list", url.Values{
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/admin/account/list", url.Values{
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

// AdminAddAccount 添加账户
func (p *Service) AdminAddAccount(c context.Context, arg *model.ArgAdminAddAccount) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	req := &identify.AdminCreateAccountReq{
		Email:    arg.Email,
		Password: arg.Password,
		Prefix:   arg.Prefix,
		Mobile:   arg.Mobile,
		Aid:      aid,
		RemoteIP: metadata.String(c, metadata.RemoteIP),
	}

	if err = p.d.AdminCreateAccount(c, req); err != nil {
		return
	}
	return
}

// AdminUpdateAccount 更新账户信息
func (p *Service) AdminUpdateAccount(c context.Context, arg *model.ArgAdminUpdateProfile) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	req := &account.UpdateProfileReq{Aid: aid, AccountID: arg.AccountID}

	if arg.Gender != nil {
		req.Gender = &account.UpdateProfileReq_GenderValue{*arg.Gender}
	}

	if arg.Avatar != nil {
		req.Avatar = &account.UpdateProfileReq_AvatarValue{*arg.Avatar}
	}

	if arg.Introduction != nil {
		req.Avatar = &account.UpdateProfileReq_AvatarValue{*arg.Avatar}
	}

	if arg.UserName != nil {
		req.UserName = &account.UpdateProfileReq_UserNameValue{*arg.UserName}
	}

	if arg.BirthYear != nil {
		req.BirthYear = &account.UpdateProfileReq_BirthYearValue{*arg.BirthYear}
	}

	if arg.BirthMonth != nil {
		req.BirthMonth = &account.UpdateProfileReq_BirthMonthValue{*arg.BirthMonth}
	}

	if arg.BirthDay != nil {
		req.BirthDay = &account.UpdateProfileReq_BirthDayValue{*arg.BirthDay}
	}

	if arg.Password != nil {
		req.Password = &account.UpdateProfileReq_PasswordValue{*arg.Password}

	}

	if arg.Location != nil {
		req.Location = &account.UpdateProfileReq_LocationValue{*arg.Location}
	}

	return p.d.UpdateProfile(c, req)
}

// AdminLockAccount 锁定账户
func (p *Service) AdminLockAccount(c context.Context, arg *model.ArgAdminLockAccount) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if err = p.d.AccountLock(c, &identify.LockReq{Aid: aid, TargetAccountID: arg.AccountID}); err != nil {
		return
	}
	return
}

// AdminUnlockAccount 解锁账户
func (p *Service) AdminUnlockAccount(c context.Context, arg *model.ArgAdminLockAccount) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if err = p.d.AccountUnlock(c, &identify.LockReq{Aid: aid, TargetAccountID: arg.AccountID}); err != nil {
		return
	}
	return
}

// AdminDeactiveAccount 管理员注销账户
func (p *Service) AdminDeactive(c context.Context, arg *model.ArgAdminDeactiveAccount) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if err = p.d.AdminDeactive(c, &identify.AdminDeactiveReq{Aid: aid, AccountID: arg.AccountID}); err != nil {
		return
	}
	return
}
