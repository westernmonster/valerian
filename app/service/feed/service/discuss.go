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

func (p *Service) getDiscussion(c context.Context, node sqalx.Node, articleID int64) (item *model.Discussion, err error) {
	var addCache = true
	if item, err = p.d.DiscussionCache(c, articleID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), articleID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetDiscussionCache(context.TODO(), item)
		})
	}
	return
}

func (p *Service) onDiscussionAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgDiscussionAdded)
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

	var discuss *model.Discussion
	if discuss, err = p.getDiscussion(c, tx, info.DiscussionID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetDiscussion", "GetDiscussion(), id(%d),error(%+v)", info.DiscussionID, err)
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
			ActionType: def.ActionTypeCreateDiscussion,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateDiscussion,
			ActorID:    discuss.CreatedBy,
			ActorType:  def.ActorTypeUser,
			TargetID:   discuss.ID,
			TargetType: def.TargetTypeDiscussion,
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

func (p *Service) onDiscussionUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgDiscussionUpdated)
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

	var discuss *model.Discussion
	if discuss, err = p.getDiscussion(c, tx, info.DiscussionID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetDiscussion", "GetDiscussion(), id(%d),error(%+v)", info.DiscussionID, err)
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
			ActionType: def.ActionTypeUpdateDiscussion,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextUpdateDiscussion,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   discuss.ID,
			TargetType: def.TargetTypeDiscussion,
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

func (p *Service) onDiscussionLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgDiscussionLiked)
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

	var discuss *model.Discussion
	if discuss, err = p.getDiscussion(c, tx, info.DiscussionID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetDiscussion", "GetDiscussion(), id(%d),error(%+v)", info.DiscussionID, err)
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
			ActionType: def.ActionTypeLikeDiscussion,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextLikeDiscussion,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   discuss.ID,
			TargetType: def.TargetTypeDiscussion,
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
