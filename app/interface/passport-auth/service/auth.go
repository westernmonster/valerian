package service

import (
	"context"
	"time"
	"valerian/app/interface/passport-login/model"
)

func (p *Service) GetTokenInfo(c context.Context, token string) (r *model.AuthReply, err error) {
	var t *model.AccessToken

	if t, err = p.getAccessToken(c, token); err != nil {
		return
	}

	if r == nil || time.Now().Unix() > t.ExpiresAt {
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

func (p *Service) getAccessToken(c context.Context, token string) (t *model.AccessToken, err error) {
	cache := true
	if t, err = p.d.AccessTokenCache(c, token); err != nil {
		cache = false
	} else if t != nil {
		return
	}

	t, err = p.d.GetAccessToken(c, p.d.DB(), token)
	if err != nil || !cache {
		return
	}
	if t != nil {
		p.addCache(func() {
			p.d.SetAccessTokenCache(context.TODO(), t)
		})
		return
	}
	return
}