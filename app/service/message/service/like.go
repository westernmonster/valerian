package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"

	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleLiked(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)

	info := new(def.MsgArticleLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("message: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("message: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("message: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	var article *model.Article
	if article, err = p.getArticle(c, tx, info.ArticleID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  article.CreatedBy,
		ActionType: model.MsgLike,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextLikeArticle,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   article.ID,
		TargetType: model.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, tx, msg); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	stat := &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}
	if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

	p.addCache(func() {
		// 强制使用Master库
		c := sqalx.NewContext(context.Background(), true)

		var setting *model.SettingResp
		if setting, err = p.getAccountSetting(c, p.d.DB(), msg.AccountID); err != nil {
			PromError("message: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", msg.AccountID, err)
			return
		}
		if setting.NotifyLike {
			push := &model.PushMessage{
				Aid:     msg.AccountID,
				MsgID:   msg.ID,
				Title:   def.PushMsgTitleArticleLiked,
				Content: def.PushMsgTitleArticleLiked,
				Link:    fmt.Sprintf(def.LinkArticle, article.ID),
			}
			if _, err := p.pushSingleUser(c, push); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
			}
		}
	})

}

func (p *Service) onReviseLiked(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)
	info := new(def.MsgReviseLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("message: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("message: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("message: tx.Rollback", "tx.Rollback(), error(%+v)", err)
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
		PromError("message: GetRevise", "GetRevise(), id(%d),error(%+v)", info.ReviseID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  revise.CreatedBy,
		ActionType: model.MsgLike,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextLikeRevise,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   revise.ID,
		TargetType: model.TargetTypeRevise,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, tx, msg); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	stat := &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}
	if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

	p.addCache(func() {
		// 强制使用Master库
		c := sqalx.NewContext(context.Background(), true)

		var setting *model.SettingResp
		if setting, err = p.getAccountSetting(c, p.d.DB(), msg.AccountID); err != nil {
			PromError("message: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", msg.AccountID, err)
			return
		}
		if setting.NotifyLike {
			push := &model.PushMessage{
				Aid:     msg.AccountID,
				MsgID:   msg.ID,
				Title:   def.PushMsgTitleReviseLiked,
				Content: def.PushMsgTitleReviseLiked,
				Link:    fmt.Sprintf(def.LinkRevise, revise.ID),
			}
			if _, err := p.pushSingleUser(c, push); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
			}
		}
	})

}

func (p *Service) onDiscussionLiked(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)
	info := new(def.MsgDiscussionLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("message: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("message: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("message: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	var discuss *model.Discussion
	if discuss, err = p.getDiscussion(c, tx, info.DiscussionID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetDiscussion", "GetDiscussion(), id(%+v),error(%+v)", info.DiscussionID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  discuss.CreatedBy,
		ActionType: model.MsgLike,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextLikeDiscussion,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   discuss.ID,
		TargetType: model.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, tx, msg); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	stat := &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}
	if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

	p.addCache(func() {
		// 强制使用Master库
		c := sqalx.NewContext(context.Background(), true)

		var setting *model.SettingResp
		if setting, err = p.getAccountSetting(c, p.d.DB(), msg.AccountID); err != nil {
			PromError("message: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", msg.AccountID, err)
			return
		}
		if setting.NotifyLike {
			push := &model.PushMessage{
				Aid:     msg.AccountID,
				MsgID:   msg.ID,
				Title:   def.PushMsgTitleDiscussionLiked,
				Content: def.PushMsgTitleDiscussionLiked,
				Link:    fmt.Sprintf(def.LinkDiscussion, discuss.ID),
			}
			if _, err := p.pushSingleUser(c, push); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
			}
		}
	})
}

func (p *Service) onCommentLiked(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)
	info := new(def.MsgCommentLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("message: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("message: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("message: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	var comment *model.Comment
	if comment, err = p.getComment(c, tx, info.CommentID); err != nil {
		PromError("message: GetComment", "GetComment(), id(%+v),error(%+v)", info.CommentID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  comment.CreatedBy,
		ActionType: model.MsgLike,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextLikeComment,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   comment.ID,
		TargetType: model.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, tx, msg); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	stat := &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}
	if err = p.d.IncrMessageStat(c, tx, stat); err != nil {
		PromError("message: AddMessage", "AddMessage(), message(%+v),error(%+v)", msg, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("message: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()

	p.addCache(func() {
		// 强制使用Master库
		c := sqalx.NewContext(context.Background(), true)

		var setting *model.SettingResp
		if setting, err = p.getAccountSetting(c, p.d.DB(), msg.AccountID); err != nil {
			PromError("message: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", msg.AccountID, err)
			return
		}
		if setting.NotifyLike {
			push := &model.PushMessage{
				Aid:     msg.AccountID,
				MsgID:   msg.ID,
				Title:   def.PushMsgTitleCommentLiked,
				Content: def.PushMsgTitleCommentLiked,
				Link:    fmt.Sprintf(def.LinkComment, comment.OwnerType, comment.OwnerID, comment.ID),
			}
			if _, err := p.pushSingleUser(c, push); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
			}
		}
	})

}
