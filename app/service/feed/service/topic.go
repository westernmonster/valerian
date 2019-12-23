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

func (p *Service) getTopic(c context.Context, node sqalx.Node, topicID int64) (item *model.Topic, err error) {
	if item, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if item == nil {
		return nil, ecode.TopicNotExist
	}
	return
}

func (p *Service) onTopicAdded(m *stan.Msg) {
	var err error

	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)

	info := new(def.MsgTopicAdded)
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

	var topic *model.Topic
	if topic, err = p.getTopic(c, tx, info.TopicID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetTopic", "GetTopic(), id(%d),error(%+v)", info.TopicID, err)
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
			ActionType: def.ActionTypeCreateTopic,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateTopic,
			ActorID:    topic.CreatedBy,
			ActorType:  def.ActorTypeUser,
			TargetID:   topic.ID,
			TargetType: def.TargetTypeTopic,
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

func (p *Service) onTopicFollowed(m *stan.Msg) {
	var err error

	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)

	info := new(def.MsgTopicFollowed)
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

	var topic *model.Topic
	if topic, err = p.getTopic(c, tx, info.TopicID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("feed: GetTopic", "GetTopic(), id(%d),error(%+v)", info.TopicID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), info.ActorID); err != nil {
		PromError("feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	if !setting.ActivityFollowTopic {
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
			ActionType: def.ActionTypeFollowTopic,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextFollowTopic,
			ActorID:    topic.CreatedBy,
			ActorType:  def.ActorTypeUser,
			TargetID:   topic.ID,
			TargetType: def.TargetTypeTopic,
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
