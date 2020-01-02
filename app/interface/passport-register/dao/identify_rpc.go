package dao

import (
	"context"
	"fmt"

	identify "valerian/app/service/identify/api/grpc"
	"valerian/library/log"
)

func (p *Dao) EmailRegister(c context.Context, arg *identify.EmailRegisterReq) (info *identify.LoginResp, err error) {
	if info, err = p.identifyRPC.EmailRegister(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.EmailRegister err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) MobileRegister(c context.Context, arg *identify.MobileRegisterReq) (info *identify.LoginResp, err error) {
	if info, err = p.identifyRPC.MobileRegister(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.MobileRegister err(%+v) arg(%+v)", err, arg))
	}
	return
}
