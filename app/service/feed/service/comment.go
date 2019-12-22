package service

import (
	"context"
	"time"

	"valerian/app/service/feed/def"
	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"

	"github.com/nats-io/stan.go"
)

func (p *Service) getComment(c context.Context, node sqalx.Node, commentID int64) (item *model.Comment, err error) {
	if item, err = p.d.GetCommentByID(c, p.d.DB(), commentID); err != nil {
		return
	} else if item == nil {
		err = ecode.CommentNotExist
		return
	}
	return
}

func (p *Service) onArticleCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	var cmt *model.Comment
	if cmt, err = p.getComment(c, tx, info.CommentID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	var ids []int64
	if ids, err = p.d.GetFansIDs(c, tx, info.ActorID); err != nil {
		PromError("feed: GetFansIDs", "GetFansIDs(), aid(%d),error(%+v)", info.ActorID, err)
		return
	}

	for _, v := range ids {
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
			PromError("feed: AddFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		PromError("feed: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()
}

func (p *Service) onReviseCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgReviseCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	if _, err = p.getComment(c, tx, info.CommentID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	var ids []int64
	if ids, err = p.d.GetFansIDs(c, tx, info.ActorID); err != nil {
		PromError("feed: GetFansIDs", "GetFansIDs(), aid(%d),error(%+v)", info.ActorID, err)
		return
	}

	for _, v := range ids {
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
			PromError("feed: AddFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		PromError("feed: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgDiscussionCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	if _, err = p.getComment(c, tx, info.CommentID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	var ids []int64
	if ids, err = p.d.GetFansIDs(c, tx, info.ActorID); err != nil {
		PromError("feed: GetFansIDs", "GetFansIDs(), aid(%d),error(%+v)", info.ActorID, err)
		return
	}

	for _, v := range ids {
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
			PromError("feed: AddFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		PromError("feed: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()
}
