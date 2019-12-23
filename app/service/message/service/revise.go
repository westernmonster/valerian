package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onReviseAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded Unmarshal failed %#v", err))
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

	var revise *model.Revise
	if revise, err = p.getRevise(c, tx, info.ReviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded GetRevise failed %#v", err))
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, tx, revise.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded GetArticle failed %#v", err))
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
		log.For(c).Error(fmt.Sprintf("service.onCommentAdded AddMessage failed %#v", err))
		return
	}

	if err = p.d.IncrMessageStat(c, tx, &model.MessageStat{AccountID: msg.AccountID, UnreadCount: 1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()

	p.addCache(func() {
		if _, err := p.pushSingleUser(context.Background(),
			msg.AccountID,
			msg.ID,
			def.PushMsgTitleReviseAdded,
			def.PushMsgTitleReviseAdded,
			fmt.Sprintf(def.LinkRevise, revise.ID),
		); err != nil {
			log.For(context.Background()).Error(fmt.Sprintf("service.onReviseAdded Push message failed %#v", err))
		}
	})
}
