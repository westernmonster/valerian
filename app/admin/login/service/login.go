package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"valerian/app/admin/login/model"
	"valerian/library/ecode"
	"valerian/library/log"
)

func (p *Service) EmailLogin(ctx context.Context, req *model.ArgEmailLogin) (resp *model.LoginResp, sid string, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	account, err := p.d.GetAccountByEmail(ctx, p.d.DB(), req.Email)
	if err != nil {
		return
	}
	if account == nil {
		err = ecode.UserNotExist
		return
	}

	if err = p.checkPassword(req.Password, account.Password, account.Salt); err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID: account.ID,
	}

	if resp.Profile, err = p.GetProfile(ctx, account.ID); err != nil {
		return
	}

	if sid, err = p.setSession(ctx, resp.AccountID, account.UserName); err != nil {
		return
	}

	return
}

func (p *Service) MobileLogin(ctx context.Context, req *model.ArgMobileLogin) (resp *model.LoginResp, sid string, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	mobile := req.Prefix + req.Mobile
	account, err := p.d.GetAccountByMobile(ctx, p.d.DB(), mobile)
	if err != nil {
		return
	}
	if account == nil {
		err = ecode.UserNotExist
		return
	}

	if err = p.checkPassword(req.Password, account.Password, account.Salt); err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID: account.ID,
	}

	if resp.Profile, err = p.GetProfile(ctx, account.ID); err != nil {
		return
	}

	if sid, err = p.setSession(ctx, resp.AccountID, account.UserName); err != nil {
		return
	}

	return
}

func (p *Service) DigitLogin(ctx context.Context, req *model.ArgDigitLogin) (resp *model.LoginResp, sid string, err error) {
	mobile := req.Prefix + req.Mobile

	var code string
	if code, err = p.d.MobileValcodeCache(ctx, model.ValcodeLogin, mobile); err != nil {
		return
	} else if code == "" {
		err = ecode.ValcodeExpires
		return
	} else if code != req.Valcode {
		err = ecode.ValcodeWrong
		return
	}

	var account *model.Account
	if account, err = p.d.GetAccountByMobile(ctx, p.d.DB(), mobile); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	resp = &model.LoginResp{
		AccountID: account.ID,
	}

	if resp.Profile, err = p.GetProfile(ctx, account.ID); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelMobileValcodeCache(context.TODO(), model.ValcodeLogin, mobile)
	})

	if sid, err = p.setSession(ctx, resp.AccountID, account.UserName); err != nil {
		return
	}

	return
}

func (p *Service) checkClient(ctx context.Context, clientID string) (err error) {
	client, err := p.d.GetClient(ctx, p.d.AuthDB(), clientID)
	if err != nil {
		return
	}

	if client == nil {
		err = ecode.ClientNotExist
		return
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

func (p *Service) GetProfile(c context.Context, accountID int64) (profile *model.Profile, err error) {
	var item *model.Account
	if item, err = p.getAccountByID(c, accountID); err != nil {
		return
	} else if item == nil {
		err = ecode.UserNotExist
		return
	}

	profile = &model.Profile{
		ID:           item.ID,
		Mobile:       item.Mobile,
		Email:        item.Email,
		Gender:       item.Gender,
		BirthYear:    item.BirthYear,
		BirthMonth:   item.BirthMonth,
		BirthDay:     item.BirthDay,
		Location:     item.Location,
		Introduction: item.Introduction,
		Avatar:       item.Avatar,
		Source:       item.Source,
		IDCert:       bool(item.IDCert),
		WorkCert:     bool(item.WorkCert),
		IsOrg:        bool(item.IsOrg),
		IsVIP:        bool(item.IsVip),
		Role:         item.Role, UserName: item.UserName,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	profile.IP = InetNtoA(item.IP)
	return
}

func (p *Service) getAccountByID(c context.Context, aid int64) (account *model.Account, err error) {
	var needCache = true

	if account, err = p.d.AccountCache(c, aid); err != nil {
		needCache = false
	} else if account != nil {
		return
	}

	if account, err = p.d.GetAccountByID(c, p.d.DB(), aid); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	if needCache {
		p.addCache(func() {
			p.d.SetAccountCache(context.TODO(), account)
		})
	}
	return
}

func (p *Service) setSession(c context.Context, aid int64, uname string) (sid string, err error) {
	si := p.newSession(c)
	si.Set(_sessUIDKey, strconv.FormatInt(aid, 10))
	si.Set(_sessUnameKey, uname)

	if err = p.d.SetSession(c, si); err != nil {
		log.For(c).Error(fmt.Sprintf("p.setSession(%v) error(%v)", si, err))
		return
	}

	sid = si.Sid

	return
}
