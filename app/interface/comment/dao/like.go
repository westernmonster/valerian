package dao

import (
	"context"
	"fmt"
	like "valerian/app/service/like/api"
	"valerian/library/log"
)

func (p *Dao) IsLike(c context.Context, aid, targetID int64, targetType string) (isLike bool, err error) {
	var resp *like.LikeInfo
	if resp, err = p.likeRPC.IsLike(c, &like.LikeReq{AccountID: aid, TargetID: targetID, TargetType: targetType}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsLike(), err(%+v), aid(%d), target_id(%d), target_type(%s)", err, aid, targetID, targetType))
		return
	}

	return resp.IsLike, nil
}

func (p *Dao) IsDislike(c context.Context, aid, targetID int64, targetType string) (isLike bool, err error) {
	var resp *like.DislikeInfo
	if resp, err = p.likeRPC.IsDislike(c, &like.DislikeReq{AccountID: aid, TargetID: targetID, TargetType: targetType}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsDislike(), err(%+v), aid(%d), target_id(%d), target_type(%s)", err, aid, targetID, targetType))
		return
	}

	return resp.IsDislike, nil
}
