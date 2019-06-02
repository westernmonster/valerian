package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/passport-register/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

const (
	_accessExpireSeconds  = 60 * 60 * 24 * 30 // 30 days
	_refreshExpireSeconds = 60 * 60 * 24 * 90 // 90 days
)

func (p *Service) grantToken(ctx context.Context, clientID string, accountID int64) (accessToken *model.AccessToken, refreshToken *model.RefreshToken, err error) {

	now := time.Now()
	accessTokenStr := generateAccess(accountID, now, p.c.DC.Num)
	refreshTokenStr := generateRefresh(accountID, now, p.c.DC.Num)

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

	var accessTokens []string

	tx, err := p.d.AuthDB().Beginx(ctx)
	if err != nil {
		log.For(ctx).Error(fmt.Sprintf("grantAccessToken Beginx err(%v) clientID(%s) accountID(%d)", err, clientID, accountID))
		return
	}

	defer tx.Rollback()

	if accessTokens, err = p.d.GetClientAccessTokens(ctx, tx, accountID, clientID); err != nil {
		return
	}

	if _, err = p.d.DelExpiredAccessToken(ctx, tx, clientID, accountID); err != nil {
		return
	}

	if _, err = p.d.DelExpiredRefreshToken(ctx, tx, clientID, accountID); err != nil {
		return
	}

	if _, err = p.d.AddAccessToken(ctx, tx, accessToken); err != nil {
		return
	}

	if _, err = p.d.AddRefreshToken(ctx, tx, refreshToken); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(ctx).Error(fmt.Sprintf("grantAccessToken Commit err(%v) ", err))
		return
	}

	p.addCache(func() {
		for _, v := range accessTokens {
			p.d.DelAccessTokenCache(context.TODO(), v)
		}
	})

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
