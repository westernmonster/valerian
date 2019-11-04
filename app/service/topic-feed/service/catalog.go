package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/feed/def"
	"valerian/app/service/topic-feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/kamilsk/retry/v4"
	"github.com/kamilsk/retry/v4/strategy"
	"github.com/nats-io/stan.go"
)

func (p *Service) onTopicTaxonomyCatalogAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)

	info := new(def.MsgTopicTaxonomyCatalogAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onTopicTaxonomyCatalogAdded Unmarshal failed %#v", err)
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

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, tx, info.CatalogID); err != nil {
		return
	} else if catalog == nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicTaxonomyCatalogAdded() catalog not exist id(%d)", err, info.CatalogID))
		m.Ack()
		return
	}

	var v *model.Account
	action := func(c context.Context, _ uint) error {
		acc, e := p.getAccount(c, p.d.DB(), info.ActorID)
		if e != nil {
			return e
		}

		v = acc
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    catalog.TopicID,
		ActionType: def.ActionTypeTopicTaxonomyCatalogAdded,
		ActionTime: time.Now().Unix(),
		ActionText: fmt.Sprintf(def.ActionTextTopicTaxonomyCatalogAdded, info.Name),
		ActorID:    info.ActorID,
		ActorType:  def.ActorTypeUser,
		TargetID:   catalog.ID,
		TargetType: def.TargetTypeTopicCatalog,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(c, p.d.DB(), feed); err != nil {
		log.Errorf("service.onTopicTaxonomyCatalogAdded() failed %#v", err)
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	m.Ack()
}

func (p *Service) onTopicTaxonomyCatalogDeleted(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)

	info := new(def.MsgTopicTaxonomyCatalogDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onTopicTaxonomyCatalogDeleted Unmarshal failed %#v", err)
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

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, tx, info.CatalogID); err != nil {
		return
	} else if catalog == nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicTaxonomyCatalogDeleted() catalog not exist id(%d)", err, info.CatalogID))
		m.Ack()
		return
	}

	var v *model.Account
	action := func(c context.Context, _ uint) error {
		acc, e := p.getAccount(c, p.d.DB(), info.ActorID)
		if e != nil {
			return e
		}

		v = acc
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    catalog.TopicID,
		ActionType: def.ActionTypeTopicTaxonomyCatalogDeleted,
		ActionTime: time.Now().Unix(),
		ActionText: fmt.Sprintf(def.ActionTextTopicTaxonomyCatalogDeleted, info.Name),
		ActorID:    info.ActorID,
		ActorType:  def.ActorTypeUser,
		TargetID:   catalog.ID,
		TargetType: def.TargetTypeTopicCatalog,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(c, p.d.DB(), feed); err != nil {
		log.Errorf("service.onTopicTaxonomyCatalogDeleted() failed %#v", err)
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	m.Ack()
}

func (p *Service) onTopicTaxonomyCatalogRenamed(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)

	info := new(def.MsgTopicTaxonomyCatalogRenamed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onTopicTaxonomyCatalogRenamed Unmarshal failed %#v", err)
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

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, tx, info.CatalogID); err != nil {
		return
	} else if catalog == nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicTaxonomyCatalogRenamed() catalog not exist id(%d)", err, info.CatalogID))
		m.Ack()
		return
	}

	var v *model.Account
	action := func(c context.Context, _ uint) error {
		acc, e := p.getAccount(c, p.d.DB(), info.ActorID)
		if e != nil {
			return e
		}

		v = acc
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    catalog.TopicID,
		ActionType: def.ActionTypeTopicTaxonomyCatalogRenamed,
		ActionTime: time.Now().Unix(),
		ActionText: fmt.Sprintf(def.ActionTextTopicTaxonomyCatalogRenamed, info.OldName, info.NewName),
		ActorID:    info.ActorID,
		ActorType:  def.ActorTypeUser,
		TargetID:   catalog.ID,
		TargetType: def.TargetTypeTopicCatalog,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(c, p.d.DB(), feed); err != nil {
		log.Errorf("service.onTopicTaxonomyCatalogRenamed() failed %#v", err)
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	m.Ack()
}

func (p *Service) onTopicTaxonomyCatalogMoved(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)

	info := new(def.MsgTopicTaxonomyCatalogMoved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onTopicTaxonomyCatalogMoved Unmarshal failed %#v", err)
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

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, tx, info.CatalogID); err != nil {
		return
	} else if catalog == nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicTaxonomyCatalogMoved() catalog not exist id(%d)", err, info.CatalogID))
		m.Ack()
		return
	}

	var v *model.Account
	action := func(c context.Context, _ uint) error {
		acc, e := p.getAccount(c, p.d.DB(), info.ActorID)
		if e != nil {
			return e
		}

		v = acc
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    catalog.TopicID,
		ActionType: def.ActionTypeTopicTaxonomyCatalogMoved,
		ActionTime: time.Now().Unix(),
		ActionText: fmt.Sprintf(def.ActionTextTopicTaxonomyCatalogMoved, info.Name),
		ActorID:    info.ActorID,
		ActorType:  def.ActorTypeUser,
		TargetID:   catalog.ID,
		TargetType: def.TargetTypeTopicCatalog,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(c, p.d.DB(), feed); err != nil {
		log.Errorf("service.onTopicTaxonomyCatalogMoved() failed %#v", err)
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	m.Ack()

}

func (p *Service) onTopicUpdated(m *stan.Msg) {
	var err error
	c := context.Background()

	info := new(def.MsgTopicUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onTopicUpdated Unmarshal failed %#v", err)
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

	var t *model.Topic
	action := func(c context.Context, _ uint) error {
		tp, e := p.getTopic(c, p.d.DB(), info.TopicID)
		if e != nil {
			return e
		}

		t = tp
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	var v *model.Account
	action = func(c context.Context, _ uint) error {
		acc, e := p.getAccount(c, p.d.DB(), info.ActorID)
		if e != nil {
			return e
		}

		v = acc
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    info.TopicID,
		ActionType: def.ActionTypeUpdateTopic,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextUpdateTopic,
		ActorID:    info.ActorID,
		ActorType:  def.ActorTypeUser,
		TargetID:   info.TopicID,
		TargetType: def.TargetTypeTopic,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(c, p.d.DB(), feed); err != nil {
		log.Errorf("service.onTopicUpdated() failed %#v", err)
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	m.Ack()

}

func (p *Service) onTopicFollowed(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)

	info := new(def.MsgTopicFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onTopicUpdated Unmarshal failed %#v", err)
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

	var t *model.Topic
	action := func(c context.Context, _ uint) error {
		tp, e := p.getTopic(c, p.d.DB(), info.TopicID)
		if e != nil {
			return e
		}

		t = tp
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	if !t.IsPrivate {
		return
	}

	var v *model.Account
	action = func(c context.Context, _ uint) error {
		acc, e := p.getAccount(c, p.d.DB(), info.ActorID)
		if e != nil {
			return e
		}

		v = acc
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    info.TopicID,
		ActionType: def.ActionTypeFollowTopic,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextFollowTopic,
		ActorID:    info.ActorID,
		ActorType:  def.ActorTypeUser,
		TargetID:   info.TopicID,
		TargetType: def.TargetTypeTopic,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(c, p.d.DB(), feed); err != nil {
		log.Errorf("service.onTopicUpdated() failed %#v", err)
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	m.Ack()

}
