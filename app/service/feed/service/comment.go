package service

import (
	"context"
	"fmt"
	"time"

	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/feed/model"
	relation "valerian/app/service/relation/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

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

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(c, info.CommentID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleCommented GetArticle failed %#v", err))
		if ecode.Cause(err) == ecode.ArticleNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleCommented GetFansIDs failed %#v", err))
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
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
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

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(c, info.CommentID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseCommented GetRevise failed %#v", err))
		if ecode.Cause(err) == ecode.ReviseNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseCommented GetFansIDs failed %#v", err))
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
			TargetID:   article.ID,
			TargetType: def.TargetTypeRevise,
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

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(c, info.CommentID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionCommented GetDiscussion failed %#v", err))
		if ecode.Cause(err) == ecode.DiscussionNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionCommented GetFansIDs failed %#v", err))
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
			TargetID:   discuss.ID,
			TargetType: def.TargetTypeDiscussion,
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
