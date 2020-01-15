package dao

import (
	"context"
	"fmt"

	like "valerian/app/service/like/api"
	"valerian/library/log"
)

func (p *Dao) Like(c context.Context, aid, targetID int64, targetType string) (err error) {
	if _, err = p.likeRPC.Like(c, &like.LikeReq{
		TargetID:   targetID,
		AccountID:  aid,
		TargetType: targetType,
	}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.Like() aid(%d) target_id(%d), target_type(%s)", aid, targetID, targetType))
		return
	}

	return
}

func (p *Dao) CancelLike(c context.Context, aid, targetID int64, targetType string) (err error) {
	if _, err = p.likeRPC.CancelLike(c, &like.LikeReq{
		TargetID:   targetID,
		AccountID:  aid,
		TargetType: targetType,
	}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CancelLike() aid(%d) target_id(%d), target_type(%s)", aid, targetID, targetType))
		return
	}

	return
}

func (p *Dao) Dislike(c context.Context, aid, targetID int64, targetType string) (err error) {
	if _, err = p.likeRPC.Dislike(c, &like.DislikeReq{
		TargetID:   targetID,
		AccountID:  aid,
		TargetType: targetType,
	}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.Dislike() aid(%d) target_id(%d), target_type(%s)", aid, targetID, targetType))
		return
	}

	return
}

func (p *Dao) CancelDislike(c context.Context, aid, targetID int64, targetType string) (err error) {
	if _, err = p.likeRPC.CancelDislike(c, &like.DislikeReq{
		TargetID:   targetID,
		AccountID:  aid,
		TargetType: targetType,
	}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CancelDislike() aid(%d) target_id(%d), target_type(%s)", aid, targetID, targetType))
		return
	}

	return
}
