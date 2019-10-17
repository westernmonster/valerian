package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
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

	var revise *article.ReviseInfo
	if revise, err = p.d.GetRevise(c, info.ReviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded GetRevise failed %#v", err))
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(c, revise.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked GetArticle failed %#v", err))
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  article.Creator.ID,
		ActionType: model.MsgReviseAdded,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextReviseAdded,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   revise.ID,
		TargetType: model.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded AddMessage failed %#v", err))
		return
	}

	m.Ack()

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded GetAccountBaseInfo failed %#v", err))
		return
	}

	var msgID string
	message := fmt.Sprintf("%s%s", account.UserName, model.MsgTextLikeArticle)
	if msgID, err = p.pushSingleUser(c, msg.AccountID, message); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded pushSingleUser failed %#v, msg_id(%s)", err, msgID))
		return
	}

}
