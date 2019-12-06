package dao

import (
	"context"
	"fmt"
	account "valerian/app/service/account/api"
	"valerian/library/log"
)

func (p *Dao) SetAccountLock(c context.Context, aid int64) (info *account.EmptyStruct, err error) {
	if info, err = p.accountRPC.AccountLock(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SetAccountLock err(%+v)", err))
	}
	return
}

func (p *Dao) SetAccountUnlock(c context.Context, aid int64) (info *account.EmptyStruct, err error) {
	if info, err = p.accountRPC.AccountUnlock(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SetAccountUnlock err(%+v)", err))
	}
	return
}