package dao

import (
	"context"
	fav "valerian/app/service/fav/api"
)

func (p *Dao) IsFav(c context.Context, aid, targetID int64, targetType string) (isFav bool, err error) {
	var resp *fav.FavInfo
	if resp, err = p.favRPC.IsFav(c, &fav.FavReq{
		AccountID:  aid,
		TargetID:   targetID,
		TargetType: targetType,
	}); err != nil {
		return
	}

	isFav = resp.IsFav

	return
}
