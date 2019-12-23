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

func (p *Service) GetRevise(c context.Context, reviseID int64) (item *model.Revise, err error) {
	if item, err = p.getRevise(c, p.d.DB(), reviseID); err != nil {
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, p.d.DB(), item.ArticleID); err != nil {
		return
	}

	item.Title = article.Title
	return
}

func (p *Service) getRevise(c context.Context, node sqalx.Node, reviseID int64) (item *model.Revise, err error) {
	if item, err = p.d.GetReviseByID(c, p.d.DB(), reviseID); err != nil {
		return
	} else if item == nil {
		err = ecode.ReviseNotExist
		return
	}
	return
}

func (p *Service) onReviseAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgReviseAdded)
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

	var revise *model.Revise
	if revise, err = p.getRevise(c, tx, info.ReviseID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetRevise", "GetRevise(), id(%d),error(%+v)", info.ReviseID, err)
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
			ActionType: def.ActionTypeCreateRevise,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateRevise,
			ActorID:    revise.CreatedBy,
			ActorType:  def.ActorTypeUser,
			TargetID:   revise.ID,
			TargetType: def.TargetTypeRevise,
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

func (p *Service) onReviseUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgReviseUpdated)
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

	var revise *model.Revise
	if revise, err = p.getRevise(c, tx, info.ReviseID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetRevise", "GetRevise(), id(%d),error(%+v)", info.ReviseID, err)
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
			ActionType: def.ActionTypeUpdateRevise,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextUpdateRevise,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   revise.ID,
			TargetType: def.TargetTypeRevise,
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

func (p *Service) onReviseLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgReviseLiked)
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

	var revise *model.Revise
	if revise, err = p.getRevise(c, tx, info.ReviseID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetRevise", "GetRevise(), id(%d),error(%+v)", info.ReviseID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), info.ActorID); err != nil {
		PromError("feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	if !setting.ActivityLike {
		m.Ack()
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
			ActionType: def.ActionTypeLikeRevise,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextLikeRevise,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   revise.ID,
			TargetType: def.TargetTypeRevise,
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
