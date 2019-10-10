package dao

import (
	"context"
	"fmt"
	feed "valerian/app/service/feed/api"
	"valerian/library/log"
)

func (p *Dao) GetFeedPaged(c context.Context, accountID int64, limit, offset int) (info *feed.FeedResp, err error) {
	if info, err = p.feedRPC.GetFeedPaged(c, &feed.FeedReq{AccountID: accountID, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeedPaged err(%+v)", err))
	}
	return
}
