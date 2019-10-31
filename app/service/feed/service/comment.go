package service

import (
	"context"
	"fmt"
	"time"

	comment "valerian/app/service/comment/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/feed/model"
	relation "valerian/app/service/relation/api"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/kamilsk/retry/v4"
	"github.com/kamilsk/retry/v4/strategy"
	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleCommented Unmarshal failed %#v", err))
		return
	}

	var cmt *comment.CommentInfo
	action := func(c context.Context, _ uint) error {
		ct, e := p.d.GetComment(c, info.CommentID)
		if e != nil {
			log.For(c).Error(fmt.Sprintf("service.onArticleCommented GetArticle failed %#v", e))
			return e
		}

		cmt = ct
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	var fansResp *relation.IDsResp
	action = func(c context.Context, _ uint) error {
		res, e := p.d.GetFansIDs(c, info.ActorID)
		if e != nil {
			log.For(c).Error(fmt.Sprintf("service.onArticleCommented GetFansIDs failed %#v", e))
			return e
		}

		fansResp = res
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

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

	for _, v := range fansResp.IDs {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeCommentArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCommentArticle,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   cmt.ID,
			TargetType: def.TargetTypeComment,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onArticleCommented() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onReviseCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseCommented Unmarshal failed %#v", err))
		return
	}

	var cmt *comment.CommentInfo
	action := func(c context.Context, _ uint) error {
		ct, e := p.d.GetComment(c, info.CommentID)
		if e != nil {
			log.For(c).Error(fmt.Sprintf("service.onReviseCommented GetRevise failed %#v", e))
			return e
		}

		cmt = ct
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	var fansResp *relation.IDsResp
	action = func(c context.Context, _ uint) error {
		res, e := p.d.GetFansIDs(c, info.ActorID)
		if e != nil {
			log.For(c).Error(fmt.Sprintf("service.onReviseCommented GetFansIDs failed %#v", e))
			return e
		}

		fansResp = res
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

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

	for _, v := range fansResp.IDs {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeCommentRevise,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCommentRevise,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   info.CommentID,
			TargetType: def.TargetTypeComment,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onReviseCommented() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionCommented Unmarshal failed %#v", err))
		return
	}

	var cmt *comment.CommentInfo
	action := func(c context.Context, _ uint) error {
		ct, e := p.d.GetComment(c, info.CommentID)
		if e != nil {
			log.For(c).Error(fmt.Sprintf("service.onDiscussionCommented GetDiscussion failed %#v", e))
			return e
		}

		cmt = ct
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	var fansResp *relation.IDsResp
	action = func(c context.Context, _ uint) error {
		res, e := p.d.GetFansIDs(c, info.ActorID)
		if e != nil {
			log.For(c).Error(fmt.Sprintf("service.onDiscussionCommented GetFansIDs failed %#v", e))
			return e
		}

		fansResp = res
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

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

	for _, v := range fansResp.IDs {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeCommentDiscussion,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCommentDiscussion,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   info.CommentID,
			TargetType: def.TargetTypeComment,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onDiscussionCommented() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}
