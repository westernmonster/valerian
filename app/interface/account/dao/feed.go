package dao

import (
	"context"
	feed "valerian/app/service/feed/api"
)

func (p *Dao) GetAccountFeedPaged(c context.Context, accountID int64, limit, offset int) (info *feed.AccountFeedResp, err error) {
	return p.feedRPC.GetAccountFeedPaged(c, &feed.AccountFeedReq{AccountID: accountID, Limit: int32(limit), Offset: int32(offset)})
}
