package dao

import (
	"context"
	account "valerian/app/service/account/api"
)

func (p *Dao) GetAccountStat(c context.Context, aid int64) (stat *account.AccountStatInfo, err error) {
	return p.accountRPC.AccountStat(c, &account.AidReq{Aid: aid})
}
