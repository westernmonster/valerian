package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/passport-login/model"
	"valerian/library/ecode"
	"valerian/library/gid"
)

const (
	_accessExpireSeconds  = 86400       // 1 days
	_refreshExpireSeconds = 2592000 * 2 // 60 days
)

func (p *Service) grantToken(ctx context.Context, clientID string, accountID int64) (accessToken *model.AccessToken, refreshToken *model.RefreshToken, err error) {
	tx, err := p.d.AuthDB().Beginx(ctx)
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

	accessToken = &model.AccessToken{
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

	refreshToken = &model.RefreshToken{
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

func (p *Service) checkClient(ctx context.Context, clientID string) (err error) {
	client, err := p.d.GetClient(ctx, p.d.DB(), clientID)
	if err != nil {
		return
	}

	if client == nil {
		err = ecode.ClientNotExist
		return
	}
	return
}
