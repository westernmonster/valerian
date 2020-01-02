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

// getDiscussion 获取讨论信息
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

// onDiscussionAdded 当讨论新增时
func (p *Service) onDiscussionAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)
	info := new(def.MsgDiscussionAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var discuss *model.Discussion
	if discuss, err = p.getDiscussion(c, p.d.DB(), info.DiscussionID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetDiscussion", "GetDiscussion(), id(%d),error(%+v)", info.DiscussionID, err)
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCreateDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateDiscussion,
		TargetID:   discuss.ID,
		TargetType: def.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

// onDiscussionUpdated 当讨论更新时
func (p *Service) onDiscussionUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)
	info := new(def.MsgDiscussionUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var discuss *model.Discussion
	if discuss, err = p.getDiscussion(c, p.d.DB(), info.DiscussionID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetDiscussion", "GetDiscussion(), id(%d),error(%+v)", info.DiscussionID, err)
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeUpdateDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextUpdateDiscussion,
		TargetID:   discuss.ID,
		TargetType: def.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

// onDiscussionLiked 当讨论被点赞时
func (p *Service) onDiscussionLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)
	info := new(def.MsgDiscussionLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var discuss *model.Discussion
	if discuss, err = p.getDiscussion(c, p.d.DB(), info.DiscussionID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetDiscussion", "GetDiscussion(), id(%d),error(%+v)", info.DiscussionID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), info.ActorID); err != nil {
		PromError("account-feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	if !setting.ActivityLike {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeLikeDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextLikeDiscussion,
		TargetID:   discuss.ID,
		TargetType: def.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}
