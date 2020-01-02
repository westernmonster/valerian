package dao

import (
	"context"
	"fmt"

	"valerian/app/interface/passport-auth/model"
	identify "valerian/app/service/identify/api/grpc"
	"valerian/library/log"
)

func (p *Dao) RenewToken(c context.Context, refreshToken, clientID string) (r *model.TokenResp, err error) {
	data, err := p.identifyRPC.RenewToken(c, &identify.RenewTokenReq{RefreshToken: refreshToken, ClientId: clientID})
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.RenewToken(), err(%+v), refresh_token(%s), client_id(%s)", err, refreshToken, clientID))
		return
	}

	r = &model.TokenResp{
		AccountID:    data.Aid,
		AccessToken:  data.AccessToken,
		ExpiresIn:    int(data.ExpiresIn),
		TokenType:    data.TokenType,
		RefreshToken: data.RefreshToken,
	}

	return
}

func (p *Dao) Logout(c context.Context, aid int64, clientID string) (err error) {
	if _, err = p.identifyRPC.Logout(c, &identify.LogoutReq{Aid: aid, ClientId: clientID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.Logout(), err(%+v), aid(%d), client_id(%s)", err, aid, clientID))
		return
	}

	return
}
