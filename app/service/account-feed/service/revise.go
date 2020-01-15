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

// getRevise 获取补充信息
func (p *Service) getRevise(c context.Context, node sqalx.Node, reviseID int64) (item *model.Revise, err error) {
	var addCache = true
	if item, err = p.d.ReviseCache(c, reviseID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetReviseByID(c, p.d.DB(), reviseID); err != nil {
		return
	} else if item == nil {
		err = ecode.ReviseNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetReviseCache(context.Background(), item)
		})
	}
	return
}

// onReviseAdded 当补充新增时
func (p *Service) onReviseAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)

	info := new(def.MsgReviseAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var revise *model.Revise
	if revise, err = p.getRevise(c, p.d.DB(), info.ReviseID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetRevise", "GetRevise(), id(%d),error(%+v)", info.ReviseID, err)
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCreateRevise,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateRevise,
		TargetID:   revise.ID,
		TargetType: def.TargetTypeRevise,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

// onReviseUpdated 当补充更新时
func (p *Service) onReviseUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgReviseUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var revise *model.Revise
	if revise, err = p.getRevise(c, p.d.DB(), info.ReviseID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetRevise", "GetRevise(), id(%d),error(%+v)", info.ReviseID, err)
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeUpdateRevise,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextUpdateRevise,
		TargetID:   revise.ID,
		TargetType: def.TargetTypeRevise,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

// onReviseLiked 当补充被点赞时
func (p *Service) onReviseLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgReviseLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var revise *model.Revise
	if revise, err = p.getRevise(c, p.d.DB(), info.ReviseID); err != nil {
		if ecode.Cause(err) == ecode.ArticleNotExist {
			m.Ack()
		}
		PromError("account-feed: GetRevise", "GetRevise(), id(%d),error(%+v)", info.ReviseID, err)
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
		ActionType: def.ActionTypeLikeRevise,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextLikeRevise,
		TargetID:   revise.ID,
		TargetType: def.TargetTypeRevise,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}
