package service

import (
	"context"
	"strings"
	api "valerian/app/service/identify/api/grpc"
	"valerian/app/service/identify/model"
	"valerian/library/ecode"
)

// loginAccount 登录
func (p *Service) loginAccount(c context.Context, acc *model.Account, clientID string) (resp *api.LoginResp, err error) {
	accessToken, refreshToken, err := p.grantToken(c, clientID, acc.ID)
	if err != nil {
		return
	}

	resp = &api.LoginResp{
		Aid:          acc.ID,
		Role:         acc.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    _accessExpireSeconds,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	return

}

// EmailLogin 邮件登录
func (p *Service) EmailLogin(ctx context.Context, req *api.EmailLoginReq) (resp *api.LoginResp, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	acc, err := p.d.GetAccountByEmail(ctx, p.d.DB(), req.Email)
	if err != nil {
		return
	}
	if acc == nil {
		err = ecode.UserNotExist
		return
	}
	if acc.IsLock {
		err = ecode.UserDisabled
		return
	}
	if err = p.checkPassword(req.Password, acc.Password, acc.Salt); err != nil {
		return
	}

	return p.loginAccount(ctx, acc, req.ClientID)
}

// MobileLogin 手机登录
func (p *Service) MobileLogin(ctx context.Context, req *api.MobileLoginReq) (resp *api.LoginResp, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	mobile := req.Prefix + req.Mobile
	acc, err := p.d.GetAccountByMobile(ctx, p.d.DB(), mobile)
	if err != nil {
		return
	}
	if acc == nil {
		err = ecode.UserNotExist
		return
	}
	if acc.IsLock {
		err = ecode.UserDisabled
		return
	}
	if err = p.checkPassword(req.Password, acc.Password, acc.Salt); err != nil {
		return
	}

	return p.loginAccount(ctx, acc, req.ClientID)
}

// DigitLogin 验证码登录
func (p *Service) DigitLogin(ctx context.Context, req *api.DigitLoginReq) (resp *api.LoginResp, err error) {
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

	var acc *model.Account
	if acc, err = p.d.GetAccountByMobile(ctx, p.d.DB(), mobile); err != nil {
		return
	} else if acc == nil {
		err = ecode.UserNotExist
		return
	}
	if acc.IsLock {
		err = ecode.UserDisabled
		return
	}

	p.addCache(func() {
		p.d.DelMobileValcodeCache(context.TODO(), model.ValcodeLogin, mobile)
	})

	return p.loginAccount(ctx, acc, req.ClientID)
}

// checkPassword 检验密码
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
