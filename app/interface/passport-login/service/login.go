package service

import (
	"context"
	"strings"

	"valerian/app/interface/passport-login/model"
	"valerian/library/ecode"
	"valerian/models"
)

func (p *Service) EmailLogin(ctx context.Context, req *model.ArgEmailLogin) (loginResult *model.LoginResp, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	account, err := p.d.GetAccountByEmail(ctx, p.d.DB(), req.Email)
	if err != nil {
		return
	}
	if account != nil {
		err = ecode.UserNotExist
		return
	}

	if err = p.checkPassword(req.Password, account.Password, account.Salt); err != nil {
		return
	}

	accessToken, refreshToken, err := p.grantToken(ctx, req.ClientID, account.ID)
	if err != nil {
		return
	}

	loginResult = &model.LoginResp{
		AccountID:    account.ID,
		Role:         account.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    models.ExpiresIn,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	p.addCache(func() {
		p.d.SetAccessTokenCache(context.TODO(), accessToken)
	})

	return
}

func (p *Service) MobileLogin(ctx context.Context, req *model.ArgMobileLogin) (loginResult *model.LoginResp, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	account, err := p.d.GetAccountByMobile(ctx, p.d.DB(), req.Mobile)
	if err != nil {
		return
	}
	if account != nil {
		err = ecode.UserNotExist
		return
	}

	if err = p.checkPassword(req.Password, account.Password, account.Salt); err != nil {
		return
	}

	accessToken, refreshToken, err := p.grantToken(ctx, req.ClientID, account.ID)
	if err != nil {
		return
	}

	loginResult = &model.LoginResp{
		AccountID:    account.ID,
		Role:         account.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    models.ExpiresIn,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	p.addCache(func() {
		p.d.SetAccessTokenCache(context.TODO(), accessToken)
	})

	return
}

func (p *Service) DigitLogin(ctx context.Context, req *model.ArgMobileLogin) (loginResult *model.LoginResp, err error) {
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
