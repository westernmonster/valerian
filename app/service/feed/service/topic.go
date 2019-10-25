package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/feed/def"
	"valerian/app/service/feed/model"
	relation "valerian/app/service/relation/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onTopicAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded Unmarshal failed %#v", err))
		return
	}

	var topic *topic.TopicInfo
	if topic, err = p.d.GetTopic(c, info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded GetTopic failed %#v", err))
		if ecode.Cause(err) == ecode.TopicNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded GetFansIDs failed %#v", err))
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
			ActionType: def.ActionTypeCreateTopic,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateTopic,
			ActorID:    topic.Creator.ID,
			ActorType:  def.ActorTypeUser,
			TargetID:   topic.ID,
			TargetType: def.TargetTypeTopic,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onTopicAdded() failed %#v", err))
			return
		}
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
	info := new(def.MsgTopicFollowed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowed Unmarshal failed %#v", err))
		return
	}

	var topic *topic.TopicInfo
	if topic, err = p.d.GetTopic(c, info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowed GetTopic failed %#v", err))
		if ecode.Cause(err) == ecode.TopicNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFollowed GetFansIDs failed %#v", err))
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
			ActionType: def.ActionTypeFollowTopic,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextFollowTopic,
			ActorID:    topic.Creator.ID,
			ActorType:  def.ActorTypeUser,
			TargetID:   topic.ID,
			TargetType: def.TargetTypeTopic,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onTopicFollowed() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onTopicDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgTopicDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onTopicDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelFeedByCond(context.Background(), p.d.DB(), def.TargetTypeTopic, info.TopicID); err != nil {
		log.Errorf("service.onTopicDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
