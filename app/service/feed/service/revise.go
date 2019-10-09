package service

import (
	"context"
	"fmt"
	"time"

	article "valerian/app/service/article/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/feed/model"
	relation "valerian/app/service/relation/api"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onReviseAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded Unmarshal failed %#v", err))
		return
	}

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(c, info.ReviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded GetRevise failed %#v", err))
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded GetFansIDs failed %#v", err))
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
			ActionType: def.ActionTypeCreateRevise,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateRevise,
			ActorID:    article.Creator.ID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeRevise,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onReviseAdded() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onReviseUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseUpdated Unmarshal failed %#v", err))
		return
	}

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(c, info.ReviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseUpdated GetRevise failed %#v", err))
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseUpdated GetFansIDs failed %#v", err))
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
			ActionType: def.ActionTypeUpdateRevise,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextUpdateRevise,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeRevise,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onReviseUpdated() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onReviseLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseLiked Unmarshal failed %#v", err))
		return
	}

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(c, info.ReviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseLiked GetRevise failed %#v", err))
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseLiked GetFansIDs failed %#v", err))
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
			ActionType: def.ActionTypeLikeRevise,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextLikeRevise,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeRevise,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onReviseLiked() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onReviseFaved(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseFaved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseFaved Unmarshal failed %#v", err))
		return
	}

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(c, info.ReviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseFaved GetRevise failed %#v", err))
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseFaved GetFansIDs failed %#v", err))
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
			ActionType: def.ActionTypeFavRevise,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextFavRevise,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeRevise,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onReviseFaved() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onReviseDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgReviseDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onReviseDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelFeedByCond(context.Background(), p.d.DB(), def.TargetTypeRevise, info.ReviseID); err != nil {
		log.Errorf("service.onReviseDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
