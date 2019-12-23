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
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)
	info := new(def.MsgArticleCommented)
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
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, tx, comment.OwnerID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetArticle", "GetArticle(), id(%d),error(%+v)", comment.OwnerID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  article.CreatedBy,
		ActionType: model.MsgArticleCommented,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextArticleCommented,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   comment.ID,
		TargetType: model.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
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

		if setting.NotifyComment {
			push := &model.PushMessage{
				Aid:     msg.AccountID,
				MsgID:   msg.ID,
				Title:   def.PushMsgTitleArticleCommented,
				Content: def.PushMsgTitleArticleCommented,
				Link:    fmt.Sprintf(def.LinkComment, model.TargetTypeArticle, comment.OwnerID, comment.ID),
			}
			if _, err := p.pushSingleUser(c, push); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
			}
		}
	})

}

func (p *Service) onReviseCommented(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)
	info := new(def.MsgReviseCommented)
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
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	var revise *model.Revise
	if revise, err = p.getRevise(c, tx, comment.OwnerID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetRevise", "GetRevise(), id(%d),error(%+v)", comment.OwnerID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  revise.CreatedBy,
		ActionType: model.MsgReviseCommented,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextReviseCommented,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   comment.ID,
		TargetType: model.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
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
		if setting.NotifyComment {
			push := &model.PushMessage{
				Aid:     msg.AccountID,
				MsgID:   msg.ID,
				Title:   def.PushMsgTitleReviseCommented,
				Content: def.PushMsgTitleReviseCommented,
				Link:    fmt.Sprintf(def.LinkComment, model.TargetTypeRevise, comment.OwnerID, comment.ID),
			}
			if _, err := p.pushSingleUser(c, push); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
			}
		}
	})
}

func (p *Service) onDiscussionCommented(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)
	info := new(def.MsgDiscussionCommented)
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
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	var discussion *model.Discussion
	if discussion, err = p.getDiscussion(c, tx, comment.OwnerID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetDiscussion", "GetDiscussion(), id(%d),error(%+v)", comment.OwnerID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  discussion.CreatedBy,
		ActionType: model.MsgDiscussionCommented,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextDiscussionCommented,
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
		if setting.NotifyComment {
			push := &model.PushMessage{
				Aid:     msg.AccountID,
				MsgID:   msg.ID,
				Title:   def.PushMsgTitleDiscussionCommented,
				Content: def.PushMsgTitleDiscussionCommented,
				Link:    fmt.Sprintf(def.LinkComment, model.TargetTypeRevise, comment.OwnerID, comment.ID),
			}

			if _, err := p.pushSingleUser(c, push); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
			}
		}
	})
}

func (p *Service) onCommentReplied(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)

	info := new(def.MsgCommentReplied)
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
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetComment", "GetComment(), id(%d),error(%+v)", info.CommentID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  comment.CreatedBy,
		ActionType: model.MsgCommentReplied,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextCommentReplied,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   comment.ID,
		TargetType: model.TargetTypeComment,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
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

		if setting.NotifyComment {
			push := &model.PushMessage{
				Aid:     msg.AccountID,
				MsgID:   msg.ID,
				Title:   def.PushMsgTitleCommentReplied,
				Content: def.PushMsgTitleCommentReplied,
				Link:    fmt.Sprintf(def.LinkSubComment, comment.OwnerType, comment.OwnerID, comment.ResourceID, comment.ID),
			}
			if _, err := p.pushSingleUser(c, push); err != nil {
				PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
			}
		}
	})

}
