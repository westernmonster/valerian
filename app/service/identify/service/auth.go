package service

import (
	"context"
	"time"

	api "valerian/app/service/identify/api/grpc"
	"valerian/app/service/identify/model"
	"valerian/library/ecode"
)

// GetTokenInfo 通过Token获取登录信息
func (p *Service) GetTokenInfo(c context.Context, token string) (r *api.AuthReply, err error) {
	var t *model.AccessToken

	if t, err = p.getAccessToken(c, token); err != nil {
		return
	}

	if t == nil || time.Now().Unix() > t.ExpiresAt {
		r = _noLogin
		return
	}

	r = &api.AuthReply{
		Login:   true,
		Aid:     t.AccountID,
		Expires: t.ExpiresAt,
	}
	return
}

// RenewToken 重新生成Token
func (p *Service) RenewToken(c context.Context, refreshToken, clientID string) (r *model.TokenResp, err error) {
	var t *model.RefreshToken
	if t, err = p.d.GetRefreshToken(c, p.d.AuthDB(), refreshToken); err != nil {
		return
	} else if t == nil {
		err = ecode.RefreshTokenNotExistOrExpired
		return
	} else if t.ClientID != clientID {
		err = ecode.RefreshTokenNotExistOrExpired
		return
	} else if time.Now().Unix() > t.ExpiresAt {
		err = ecode.RefreshTokenNotExistOrExpired
		return
	}

	accToken, refToken, err := p.grantToken(c, clientID, t.AccountID)
	if err != nil {
		return
	}

	r = &model.TokenResp{
		AccountID:    t.AccountID,
		AccessToken:  accToken.Token,
		RefreshToken: refToken.Token,
		ExpiresIn:    _accessExpireSeconds,
	}

	return
}

// getAccessToken 获取AccessToken
func (p *Service) getAccessToken(c context.Context, token string) (t *model.AccessToken, err error) {
	addCache := true
	if t, err = p.d.AccessTokenCache(c, token); err != nil {
		addCache = false
	} else if t != nil {
		return
	}

	if t, err = p.d.GetAccessToken(c, p.d.AuthDB(), token); err != nil {
		return
	} else if t == nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetAccessTokenCache(context.Background(), t)
		})
	}
	return
}
