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

func (p *Service) onReviseAdded(m *stan.Msg) {
	var err error
	// 强制使用Master库
	c := sqalx.NewContext(context.Background(), true)

	info := new(def.MsgReviseAdded)
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

	var article *model.Article
	if article, err = p.getArticle(c, tx, revise.ArticleID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("message: GetArticle", "GetArticle(), id(%d),error(%+v)", revise.ArticleID, err)
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  article.CreatedBy,
		ActionType: model.MsgReviseAdded,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextReviseAdded,
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
		push := &model.PushMessage{
			Aid:     msg.AccountID,
			MsgID:   msg.ID,
			Title:   def.PushMsgTitleReviseAdded,
			Content: def.PushMsgTitleReviseAdded,
			Link:    fmt.Sprintf(def.LinkRevise, revise.ID),
		}
		if _, err := p.pushSingleUser(context.Background(), push); err != nil {
			PromError("message: pushSingleUser", "pushSingleUser(), push(%+v),error(%+v)", push, err)
		}
	})
}
