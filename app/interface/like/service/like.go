package service

import (
	"context"
	"time"

	"valerian/app/interface/like/model"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) Like(c context.Context, arg *model.ArgLike) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.Like(c, aid, arg.TargetID, arg.TargetType); err != nil {
		return
	}

	p.addCache(func() {
		switch arg.TargetType {
		case model.TargetTypeArticle:
			p.onArticleLiked(context.Background(), arg.TargetID, aid, time.Now().Unix())
			break
		case model.TargetTypeRevise:
			p.onReviseLiked(context.Background(), arg.TargetID, aid, time.Now().Unix())
			break
		case model.TargetTypeDiscussion:
			p.onDiscussionLiked(context.Background(), arg.TargetID, aid, time.Now().Unix())
			break

		case model.TargetTypeComment:
			// p.onTopicLiked(context.Background(), arg.TargetID, aid, fav.CreatedAt)
			break

		}
	})

	return
}

func (p *Service) CancelLike(c context.Context, arg *model.ArgCancelLike) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.CancelLike(c, aid, arg.TargetID, arg.TargetType); err != nil {
		return
	}

	return
}

func (p *Service) Dislike(c context.Context, arg *model.ArgDislike) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.Dislike(c, aid, arg.TargetID, arg.TargetType); err != nil {
		return
	}

	return
}

func (p *Service) CancelDislike(c context.Context, arg *model.ArgCancelDislike) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.CancelDislike(c, aid, arg.TargetID, arg.TargetType); err != nil {
		return
	}

	return
}
