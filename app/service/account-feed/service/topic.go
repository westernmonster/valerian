package service

import (
	"context"
	"time"
	"valerian/app/service/account-feed/model"
	"valerian/app/service/feed/def"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"

	"github.com/nats-io/stan.go"
)

func (p *Service) getTopic(c context.Context, node sqalx.Node, topicID int64) (item *model.Topic, err error) {
	var addCache = true
	if item, err = p.d.TopicCache(c, topicID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if item == nil {
		return nil, ecode.TopicNotExist
	}

	if addCache {
		p.addCache(func() {
			p.d.SetTopicCache(context.TODO(), item)
		})
	}

	return
}

func (p *Service) onTopicAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgTopicAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var topic *model.Topic
	if topic, err = p.getTopic(c, p.d.DB(), info.TopicID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetTopic", "GetTopic(), id(%d),error(%+v)", info.TopicID, err)
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCreateTopic,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateTopic,
		TargetID:   topic.ID,
		TargetType: def.TargetTypeTopic,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

func (p *Service) onTopicFollowed(m *stan.Msg) {
	var err error
	info := new(def.MsgTopicFollowed)
	c := context.Background()
	// 强制使用Master库
	c = sqalx.NewContext(c, true)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var topic *model.Topic
	if topic, err = p.getTopic(c, p.d.DB(), info.TopicID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetTopic", "GetTopic(), id(%d),error(%+v)", info.TopicID, err)
		return
	}

	// var v *model.Account
	if _, err := p.getAccount(c, p.d.DB(), info.ActorID); err != nil {
		PromError("account-feed: GetAccount", "GetAccount(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), info.ActorID); err != nil {
		PromError("account-feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	if !setting.ActivityFollowTopic {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeFollowTopic,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextFollowTopic,
		TargetID:   topic.ID,
		TargetType: def.TargetTypeTopic,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}
