package dao

import (
	"context"
	"fmt"
	recent "valerian/app/service/recent/api"
	"valerian/library/log"
)

func (p *Dao) GetRecentPubsPaged(c context.Context, accountID int64, limit, offset int) (info *recent.RecentPubsResp, err error) {
	if info, err = p.recentRPC.GetRecentPubsPaged(c, &recent.RecentPubsReq{AccountID: accountID, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecentPubsPaged(), err(%+v), aid(%d), limit(%d), offset(%d)", err, accountID, limit, offset))
	}
	return
}

func (p *Dao) GetRecentViewsPaged(c context.Context, accountID int64, limit, offset int) (info *recent.RecentViewsResp, err error) {
	if info, err = p.recentRPC.GetRecentViewsPaged(c, &recent.RecentViewsReq{AccountID: accountID, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecentViewsPaged(), err(%+v), aid(%d), limit(%d), offset(%d)", err, accountID, limit, offset))
	}
	return
}
