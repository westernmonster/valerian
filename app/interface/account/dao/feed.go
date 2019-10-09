package dao

import (
	"context"
	"fmt"
	accountFeed "valerian/app/service/account-feed/api"
	"valerian/library/log"
)

func (p *Dao) GetAccountFeedPaged(c context.Context, accountID int64, limit, offset int) (info *accountFeed.AccountFeedResp, err error) {
	if info, err = p.accountFeedRPC.GetAccountFeedPaged(c, &accountFeed.AccountFeedReq{AccountID: accountID, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFeedPaged(), err(%+v), aid(%d), limit(%d), offset(%d)", err, accountID, limit, offset))
	}
	return
}
