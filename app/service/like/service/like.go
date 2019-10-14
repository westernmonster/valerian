package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/like/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) IsLike(c context.Context, aid int64, targetID int64, targetType string) (isLike bool, err error) {
	var fav *model.Like
	if fav, err = p.d.GetLikeByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		isLike = true
		return
	}
	return
}

func (p *Service) Like(c context.Context, aid, targetID int64, targetType string) (err error) {
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var fav *model.Like
	if fav, err = p.d.GetLikeByCond(c, tx, map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		return
	}

	fav = &model.Like{
		ID:         gid.NewID(),
		AccountID:  aid,
		TargetID:   targetID,
		TargetType: targetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddLike(c, tx, fav); err != nil {
		return
	}

	switch targetType {
	case model.TargetTypeArticle:
		if err = p.d.IncrArticleStat(c, tx, &model.ArticleStat{LikeCount: 1}); err != nil {
			return
		}
		break
	case model.TargetTypeDiscussion:
		if err = p.d.IncrDiscussionStat(c, tx, &model.DiscussionStat{LikeCount: 1}); err != nil {
			return
		}
		break
	case model.TargetTypeRevise:
		if err = p.d.IncrReviseStat(c, tx, &model.ReviseStat{LikeCount: 1}); err != nil {
			return
		}
		break
	case model.TargetTypeComment:
		if err = p.d.IncrCommentStat(c, tx, &model.CommentStat{LikeCount: 1}); err != nil {
			return
		}
		break
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	return
}

func (p *Service) CancelLike(c context.Context, aid, targetID int64, targetType string) (err error) {
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var fav *model.Like
	if fav, err = p.d.GetLikeByCond(c, tx, map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav == nil {
		return
	}

	if err = p.d.DelLike(c, tx, fav.ID); err != nil {
		return
	}

	switch targetType {
	case model.TargetTypeArticle:
		if err = p.d.IncrArticleStat(c, tx, &model.ArticleStat{LikeCount: -1}); err != nil {
			return
		}
		break
	case model.TargetTypeDiscussion:
		if err = p.d.IncrDiscussionStat(c, tx, &model.DiscussionStat{LikeCount: -1}); err != nil {
			return
		}
		break
	case model.TargetTypeRevise:
		if err = p.d.IncrReviseStat(c, tx, &model.ReviseStat{LikeCount: -1}); err != nil {
			return
		}
		break
	case model.TargetTypeComment:
		if err = p.d.IncrCommentStat(c, tx, &model.CommentStat{LikeCount: -1}); err != nil {
			return
		}
		break
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	return
}
