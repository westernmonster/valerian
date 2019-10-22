package service

import (
	"context"
	"strings"
	"valerian/app/admin/manager/model"
	"valerian/library/ecode"
)

func (p *Service) Login(c context.Context, arg *model.ArgAdminLogin) (resp *model.LoginResp, err error) {
	var u *model.User
	if u, err = p.d.GetUserByCond(c, p.d.DB(), map[string]interface{}{
		"user_name": arg.UserName,
	}); err != nil {
		return
	} else if u == nil {
		err = ecode.AdminNotExist
		return
	}

	if err = p.checkPassword(arg.Password, u.Password, u.Salt); err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID: u.ID,
		Role:      u.Role,
		Profile: &model.Profile{
			ID:        u.ID,
			UserName:  u.UserName,
			Email:     u.Email,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
	}

	return
}

func (p *Service) checkPassword(password, dbPassword, dbSalt string) (err error) {
	passwordHash, err := hashPassword(password, dbSalt)
	if err != nil {
		return
	}

	if !strings.EqualFold(dbPassword, passwordHash) {
		err = ecode.PasswordErr
		return
	}
	return
}
