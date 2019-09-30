package dao

import (
	"context"

	like "valerian/app/service/like/api"
)

func (p *Dao) Like(c context.Context, aid, targetID int64, targetType string) (err error) {
	if _, err = p.likeRPC.Like(c, &like.LikeReq{
		TargetID:   targetID,
		AccountID:  aid,
		TargetType: targetType,
	}); err != nil {
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
		return
	}

	return
}
