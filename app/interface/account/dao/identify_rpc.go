package dao

import (
	"context"
	"fmt"

	identify "valerian/app/service/identify/api/grpc"
	"valerian/library/log"
)

func (p *Dao) ForgetPassword(c context.Context, identity, valcode, prefix string, identifyType int32) (info *identify.ForgetPasswordResp, err error) {
	req := &identify.ForgetPasswordReq{Identity: identity, Prefix: prefix, Valcode: valcode, IdentityType: identifyType}
	if info, err = p.identifyRPC.ForgetPassword(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.ForgetPassword err(%+v) req(%s)", err, req))
	}
	return
}

func (p *Dao) ResetPassword(c context.Context, sessionID, password string) (err error) {
	req := &identify.ResetPasswordReq{Password: password, SessionID: sessionID}
	if _, err = p.identifyRPC.ResetPassword(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.ResetPassword err(%+v) req(%s)", err, req))
	}
	return
}
