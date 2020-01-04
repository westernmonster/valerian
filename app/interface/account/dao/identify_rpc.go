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

func (p *Dao) Deactive(c context.Context, arg *identify.DeactiveReq) (err error) {
	if _, err = p.identifyRPC.Deactive(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.Deactive err(%+v) req(%s)", err, arg))
	}
	return
}

func (p *Dao) AdminCreateAccount(c context.Context, arg *identify.AdminCreateAccountReq) (err error) {
	if _, err = p.identifyRPC.AdminCreateAccount(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AdminCreateAccount err(%+v) req(%s)", err, arg))
	}
	return
}

func (p *Dao) AccountLock(c context.Context, arg *identify.LockReq) (err error) {
	if _, err = p.identifyRPC.AccountLock(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AccountLock err(%+v) req(%s)", err, arg))
	}
	return
}

func (p *Dao) AccountUnlock(c context.Context, arg *identify.LockReq) (err error) {
	if _, err = p.identifyRPC.AccountUnlock(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AccountUnlock err(%+v) req(%s)", err, arg))
	}
	return
}
