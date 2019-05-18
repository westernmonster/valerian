package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"valerian/app/interface/passport-login/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/models"
)

const (
	_accessExpireSeconds  = 2592000     // 30 days
	_refreshExpireSeconds = 2592000 * 2 // 60 days
)

func (p *Service) EmailLogin(ctx context.Context, req *model.ArgEmailLogin) (loginResult *model.LoginResp, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	account, err := p.d.GetAccountByEmail(ctx, p.d.Node(), req.Email)
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

	return
}

func (p *Service) grantToken(ctx context.Context, clientID string, accountID int64) (accessToken *model.OauthAccessToken, refreshToken *model.OauthRefreshToken, err error) {
	tx, err := p.d.Node().Beginx(ctx)
	if err != nil {
		p.logger.For(ctx).Error(fmt.Sprintf("grantAccessToken Beginx err(%v) clientID(%s) accountID(%d)", err, clientID, accountID))
		return
	}

	defer tx.Rollback()

	now := time.Now()
	accessTokenStr := generateAccess(accountID, now, p.c.DC.Num)
	refreshTokenStr := generateRefresh(accountID, now, p.c.DC.Num)

	if _, err = p.d.DelExpiredAccessToken(ctx, tx, clientID, accountID, now.Unix()); err != nil {
		return
	}

	accessToken = &model.OauthAccessToken{
		ID:        gid.NewID(),
		ClientID:  clientID,
		AccountID: accountID,
		Token:     accessTokenStr,
		ExpiresAt: now.Unix() + _accessExpireSeconds,
		Scope:     "",
		Deleted:   false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if _, err = p.d.AddAccessToken(ctx, tx, accessToken); err != nil {
		return
	}

	refreshToken = &model.OauthRefreshToken{
		ID:        gid.NewID(),
		ClientID:  clientID,
		AccountID: accountID,
		Token:     refreshTokenStr,
		ExpiresAt: now.Unix() + _refreshExpireSeconds,
		Scope:     "",
		Deleted:   false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if _, err = p.d.AddRefreshToken(ctx, tx, refreshToken); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		p.logger.For(ctx).Error(fmt.Sprintf("grantAccessToken Commit err(%v) ", err))
		return
	}

	return
}

func (p *Service) MobileLogin(ctx context.Context, req *model.ArgMobileLogin) (loginResult *model.LoginResp, err error) {
	return
}

func (p *Service) DigitLogin(ctx context.Context, req *model.ArgMobileLogin) (loginResult *model.LoginResp, err error) {
	return
}

func (p *Service) checkClient(ctx context.Context, clientID string) (err error) {
	client, err := p.d.GetClient(ctx, p.d.Node(), clientID)
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
