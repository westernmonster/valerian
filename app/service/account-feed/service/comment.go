package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/account-feed/model"
	"valerian/app/service/feed/def"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) getComment(c context.Context, node sqalx.Node, commentID int64) (item *model.Comment, err error) {
	if item, err = p.d.GetCommentByID(c, p.d.DB(), commentID); err != nil {
		return
	} else if item == nil {
		err = ecode.CommentNotExist
		return
	}
	return
}

func (p *Service) onArticleCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	if _, err = p.getComment(c, p.d.DB(), info.CommentID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), info.ActorID); err != nil {
		PromError("account-feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	if !setting.ActivityComment {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCommentArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTypeCommentArticle,
		TargetID:   info.CommentID,
		TargetType: def.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

func (p *Service) onReviseCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	if _, err = p.getComment(c, p.d.DB(), info.CommentID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), info.ActorID); err != nil {
		PromError("account-feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	if !setting.ActivityComment {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCommentRevise,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCommentRevise,
		TargetID:   info.CommentID,
		TargetType: def.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionCommented(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleCommented)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleCommented Unmarshal failed %#v", err))
		return
	}

	if _, err = p.getComment(c, p.d.DB(), info.CommentID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), info.ActorID); err != nil {
		PromError("account-feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	if !setting.ActivityComment {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCommentDiscussion,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCommentDiscussion,
		TargetID:   info.CommentID,
		TargetType: def.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

func (p *Service) onCommentLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)

	info := new(def.MsgCommentLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	if _, err = p.getComment(c, p.d.DB(), info.CommentID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("account-feed: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
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
		ActionType: def.ActionTypeLikeComment,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextLikeComment,
		TargetID:   info.CommentID,
		TargetType: def.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}
