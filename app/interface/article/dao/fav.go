package dao

import (
	"context"
	"fmt"
	fav "valerian/app/service/fav/api"
	"valerian/library/log"
)

func (p *Dao) IsFav(c context.Context, aid, targetID int64, targetType string) (isFav bool, err error) {
	var resp *fav.FavInfo
	if resp, err = p.favRPC.IsFav(c, &fav.FavReq{AccountID: aid, TargetID: targetID, TargetType: targetType}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsFav(), err(%+v), aid(%d), target_id(%d), target_type(%s)", err, aid, targetID, targetType))
		return
	}

	return resp.IsFav, nil
}
