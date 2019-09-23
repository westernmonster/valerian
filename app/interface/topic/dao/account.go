package dao

import (
	"context"
	account "valerian/app/service/account/api"
)

func (p *Dao) GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error) {
	return p.accountRPC.BasicInfo(c, &account.AidReq{Aid: aid})
}
