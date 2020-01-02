package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/identify/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

const (
	_accessExpireSeconds  = 60 * 60 * 24 * 30 // 30 days
	_refreshExpireSeconds = 60 * 60 * 24 * 90 // 90 days
)

// grantToken 生成Token
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

// checkClient 检测ClientID
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

// deleteToken 删除Token
func (p *Service) deleteToken(ctx context.Context, clientID string, accountID int64) (err error) {
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

// deleteAllToken 删除所有Token
func (p *Service) deleteAllToken(ctx context.Context, accountID int64) (err error) {
	tx, err := p.d.AuthDB().Beginx(ctx)
	if err != nil {
		log.For(ctx).Error(fmt.Sprintf("deleteAllToken Beginx err(%v) accountID(%d)", err, accountID))
		return
	}

	defer tx.Rollback()

	var accTokens []*model.AccessToken
	if accTokens, err = p.d.GetAccessTokensByCond(ctx, tx, map[string]interface{}{
		"account_id": accountID,
	}); err != nil {
		return
	}

	var refTokens []*model.RefreshToken
	if refTokens, err = p.d.GetRefreshTokensByCond(ctx, tx, map[string]interface{}{
		"account_id": accountID,
	}); err != nil {
		return
	}

	if _, err = p.d.DelAllAccessToken(ctx, tx, accountID); err != nil {
		return
	}

	if _, err = p.d.DelAllRefreshToken(ctx, tx, accountID); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(ctx).Error(fmt.Sprintf("deleteAllToken Commit err(%v) ", err))
		return
	}

	p.addCache(func() {
		for _, v := range accTokens {
			p.d.DelAccessTokenCache(context.TODO(), v.Token)
		}

		for _, v := range refTokens {
			p.d.DelRefreshTokenCache(context.TODO(), v.Token)
		}
	})

	return
}
