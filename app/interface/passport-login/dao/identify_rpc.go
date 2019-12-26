package dao

import (
	"context"
	"fmt"

	identify "valerian/app/service/identify/api/grpc"
	"valerian/library/log"
)

func (p *Dao) EmailLogin(c context.Context, arg *identify.EmailLoginReq) (info *identify.LoginResp, err error) {
	if info, err = p.identifyRPC.EmailLogin(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.EmailLogin err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) DigitLogin(c context.Context, arg *identify.DigitLoginReq) (info *identify.LoginResp, err error) {
	if info, err = p.identifyRPC.DigitLogin(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DigitLogin err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) MobileLogin(c context.Context, arg *identify.MobileLoginReq) (info *identify.LoginResp, err error) {
	if info, err = p.identifyRPC.MobileLogin(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.MobileLogin err(%+v) arg(%+v)", err, arg))
	}
	return
}
