package dao

import (
	"context"
	"fmt"
	recent "valerian/app/service/recent/api"
	"valerian/library/log"
)

func (p *Dao) GetRecentViewsPaged(c context.Context, aid int64, targetType string, limit, offset int) (resp *recent.RecentViewsResp, err error) {
	if resp, err = p.recentRPC.GetRecentViewsPaged(c, &recent.RecentViewsReq{AccountID: aid, TargetType: targetType, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecentViewsPaged error(%+v), aid(%d) targetType(%s), limit(%d), offset(%d)", err, aid, targetType, limit, offset))
	}

	return
}
