package dao

import (
	"context"
	"fmt"
	identify "valerian/app/service/identify/api/grpc"
	"valerian/library/log"
)

func (p *Dao) SetAccountLock(c context.Context, aid, targetAid int64) (info *identify.EmptyStruct, err error) {
	if info, err = p.identifyRPC.AccountLock(c, &identify.LockReq{Aid: aid, TargetAccountID: targetAid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SetAccountLock err(%+v)", err))
	}
	return
}

func (p *Dao) SetAccountUnlock(c context.Context, aid, targetAid int64) (info *identify.EmptyStruct, err error) {
	if info, err = p.identifyRPC.AccountUnlock(c, &identify.LockReq{Aid: aid, TargetAccountID: targetAid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SetAccountUnlock err(%+v)", err))
	}
	return
}
