package service

import (
	"context"
	"fmt"
	"time"

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

func (p *Service) onDiscussionAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionAdded Unmarshal failed %#v", err))
		return
	}

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(c, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionAdded GetDiscussion failed %#v", err))
		if ecode.Cause(err) == ecode.DiscussionNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionAdded GetFansIDs failed %#v", err))
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
			ActionType: def.ActionTypeCreateDiscussion,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateDiscussion,
			ActorID:    discuss.Creator.ID,
			ActorType:  def.ActorTypeUser,
			TargetID:   discuss.ID,
			TargetType: def.TargetTypeDiscussion,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onDiscussionAdded() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated Unmarshal failed %#v", err))
		return
	}

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(c, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated GetDiscussion failed %#v", err))
		if ecode.Cause(err) == ecode.DiscussionNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated GetFansIDs failed %#v", err))
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
			log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked Unmarshal failed %#v", err))
		return
	}

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(c, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked GetDiscussion failed %#v", err))
		if ecode.Cause(err) == ecode.DiscussionNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked GetFansIDs failed %#v", err))
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
			log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionFaved(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionFaved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionFaved Unmarshal failed %#v", err))
		return
	}

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(c, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionFaved GetDiscussion failed %#v", err))
		if ecode.Cause(err) == ecode.DiscussionNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionFaved GetFansIDs failed %#v", err))
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
			ActionType: def.ActionTypeFavDiscussion,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextFavDiscussion,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   discuss.ID,
			TargetType: def.TargetTypeDiscussion,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onDiscussionFaved() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgDiscussionDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onDiscussionDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelFeedByCond(context.Background(), p.d.DB(), def.TargetTypeDiscussion, info.DiscussionID); err != nil {
		log.Errorf("service.onDiscussionDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
