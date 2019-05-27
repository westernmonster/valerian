package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/passport-login/model"
	"valerian/library/ecode"
	"valerian/library/log"
)

var (
	_noLoginIdentity = &model.IdentityInfo{
		Aid:     0,
		Expires: 86400,
	}
)

func (p *Service) GetIdentity(c context.Context, accessToken string) (id *model.IdentityInfo, err error) {
	var (
		t *model.AccessToken
	)

	t, _ = p.d.AccessTokenCache(c, accessToken)

	if t != nil && time.Now().Unix() < t.ExpiresAt {
		id = &model.IdentityInfo{
			Aid:     t.AccountID,
			Expires: 86400,
		}
		return
	}

	if t, err = p.d.GetAccessToken(c, p.d.AuthDB(), accessToken); err != nil {
		return
	}

	if t != nil {
		if time.Now().Unix() < t.ExpiresAt {
			id = &model.IdentityInfo{
				Aid:     t.AccountID,
				Expires: 86400,
			}
		}

		p.addCache(func() {
			p.d.SetAccessTokenCache(context.TODO(), t)
		})

		return
	}

	id = _noLoginIdentity
	err = ecode.NoLogin
	log.For(c).Info(fmt.Sprintf("get identity use access token(%s) not found", accessToken))
	return
}
