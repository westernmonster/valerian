package dao

import (
	"context"
	"fmt"

	fav "valerian/app/service/fav/api"
	"valerian/library/log"
)

func (p *Dao) GetUserFavsPaged(c context.Context, aid int64, targetType string, limit, offset int32) (resp *fav.FavsResp, err error) {
	if resp, err = p.favRPC.GetUserFavsPaged(c, &fav.UserFavsReq{AccountID: aid, TargetType: targetType, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserFavsPaged, error(%+v), aid(%d), target_type(%s), limit(%d), offset(%d)`", err, aid, targetType, limit, offset))
	}
	return
}
