package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/common/model"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) Like(c context.Context, arg *model.ArgLike) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	switch arg.TargetType {
	case model.TargetTypeArticle:
		if _, err = p.d.GetArticle(c, arg.TargetID); err != nil {
			return
		}
		break
	case model.TargetTypeRevise:
		if _, err = p.d.GetRevise(c, arg.TargetID); err != nil {
			return
		}
		break
	case model.TargetTypeDiscussion:
		if _, err = p.d.GetDiscussion(c, arg.TargetID); err != nil {
			return
		}
		break
	case model.TargetTypeComment:
		if _, err = p.d.GetComment(c, arg.TargetID); err != nil {
			return
		}
		break
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
			p.onCommentLiked(context.Background(), arg.TargetID, aid, time.Now().Unix())
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

	fmt.Printf("CancelLike() aid(%d) target_id(%d), target_type(%s)\n", aid, arg.TargetID, arg.TargetType)

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
