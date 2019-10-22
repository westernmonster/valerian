package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"valerian/app/admin/manager/model"
	"valerian/library/ecode"
	"valerian/library/log"
)

func (p *Service) Auth(c context.Context, reqSID string) (sid string, uid int64, uname string, err error) {
	si := p.session(c, reqSID)
	if si.Get(_sessUIDKey) == nil {
		err = ecode.Unauthorized
		return
	}

	sid = si.Sid
	uidStr := si.Get(_sessUIDKey).(string)
	if uid, err = strconv.ParseInt(uidStr, 10, 64); err != nil {
		return
	}

	var u *model.User
	if u, err = p.d.GetUserByID(c, p.d.DB(), uid); err != nil {
		return
	} else if u == nil {
		err = ecode.AdminNotExist
		return
	}
	uname = si.Get(_sessUnameKey).(string)

	return
}

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

	si := p.newSession(c)
	si.Set(_sessUIDKey, strconv.FormatInt(u.ID, 10))
	si.Set(_sessUnameKey, u.UserName)

	if err = p.d.SetSession(c, si); err != nil {
		log.For(c).Error(fmt.Sprintf("s.SetSession(%v) error(%v)", si, err))
		err = nil
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
