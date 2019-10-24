package service

import (
	"context"

	"valerian/app/interface/passport-auth/model"
	"valerian/library/net/metadata"
)

func (p *Service) RenewToken(c context.Context, arg *model.ArgRenewToken) (r *model.TokenResp, err error) {
	if r, err = p.d.RenewToken(c, arg.RefreshToken, arg.ClientID); err != nil {
		return
	}
	return
}

func (p *Service) Logout(c context.Context, arg *model.ArgLogout) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		return
	}

	if err = p.d.Logout(c, aid, arg.ClientID); err != nil {
		return
	}
	return
}
