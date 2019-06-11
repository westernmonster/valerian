package service

import (
	"context"
	"time"
	"valerian/app/interface/passport-auth/model"
	"valerian/library/ecode"
)

func (p *Service) Auth(c context.Context, token string) (r *model.AuthReply, err error) {
	var t *model.AccessToken

	if t, err = p.getAccessToken(c, token); err != nil {
		return
	}

	if t == nil || time.Now().Unix() > t.ExpiresAt {
		r = _noLogin
		return
	}

	r = &model.AuthReply{
		Login:     true,
		Aid:       t.AccountID,
		ExpiresAt: t.ExpiresAt,
	}
	return
}

func (p *Service) RenewToken(c context.Context, arg *model.ArgRenewToken) (r *model.TokenResp, err error) {
	var t *model.RefreshToken
	if t, err = p.d.GetRefreshToken(c, p.d.AuthDB(), arg.RefreshToken); err != nil {
		return
	} else if t == nil {
		err = ecode.RefreshTokenNotExistOrExpired
		return
	} else if t.ClientID != arg.ClientID {
		err = ecode.RefreshTokenNotExistOrExpired
		return
	} else if time.Now().Unix() > t.ExpiresAt {
		err = ecode.RefreshTokenNotExistOrExpired
		return
	}

	accessToken, refreshToken, err := p.grantToken(c, arg.ClientID, t.AccountID)
	if err != nil {
		return
	}

	r = &model.TokenResp{
		AccountID:    t.AccountID,
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
		ExpiresIn:    _accessExpireSeconds,
	}

	return
}

func (p *Service) getAccessToken(c context.Context, token string) (t *model.AccessToken, err error) {
	needCache := true
	if t, err = p.d.AccessTokenCache(c, token); err != nil {
		needCache = false
	} else if t != nil {
		return
	}

	if t, err = p.d.GetAccessToken(c, p.d.AuthDB(), token); err != nil {
		return
	}

	if needCache && t != nil {
		p.addCache(func() {
			p.d.SetAccessTokenCache(context.TODO(), t)
		})
	}
	return
}
