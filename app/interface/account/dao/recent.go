package dao

import (
	"context"
	"fmt"
	recent "valerian/app/service/recent/api"
	"valerian/library/log"
)

func (p *Dao) GetRecentPubsPaged(c context.Context, aid int64, targetType string, limit, offset int) (info *recent.RecentPubsResp, err error) {
	if info, err = p.recentRPC.GetRecentPubsPaged(c, &recent.RecentPubsReq{AccountID: aid, TargetType: targetType, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecentPubsPaged(), err(%+v), aid(%d), limit(%d), offset(%d)", err, aid, limit, offset))
	}
	return
}

func (p *Dao) GetRecentViewsPaged(c context.Context, aid int64, targetType string, limit, offset int) (info *recent.RecentViewsResp, err error) {
	if info, err = p.recentRPC.GetRecentViewsPaged(c, &recent.RecentViewsReq{AccountID: aid, TargetType: targetType, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecentViewsPaged(), err(%+v), aid(%d), limit(%d), offset(%d)", err, aid, limit, offset))
	}
	return
}
