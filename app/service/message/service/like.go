package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/message/model"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked Unmarshal failed %#v", err))
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(c, info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked GetArticle failed %#v", err))
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  article.Creator.ID,
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

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked AddMessage failed %#v", err))
		return
	}

	m.Ack()

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked GetAccountBaseInfo failed %#v", err))
		return
	}

	var msgID string
	message := fmt.Sprintf("%s%s", account.UserName, model.MsgTextLikeArticle)
	if msgID, err = p.pushSingleUser(c, msg.AccountID, message); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked pushSingleUser failed %#v, msg_id(%s)", err, msgID))
		return
	}

}

func (p *Service) onReviseLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseLiked Unmarshal failed %#v", err))
		return
	}

	var article *article.ReviseInfo
	if article, err = p.d.GetRevise(c, info.ReviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseLiked GetRevise failed %#v", err))
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  article.Creator.ID,
		ActionType: model.MsgLike,
		ActionTime: time.Now().Unix(),
		ActionText: model.MsgTextLikeRevise,
		Actors:     strconv.FormatInt(info.ActorID, 10),
		MergeCount: 1,
		ActorType:  model.ActorTypeUser,
		TargetID:   article.ID,
		TargetType: model.TargetTypeRevise,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseLiked AddMessage failed %#v", err))
		return
	}

	m.Ack()

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseLiked GetAccountBaseInfo failed %#v", err))
		return
	}

	var msgID string
	message := fmt.Sprintf("%s%s", account.UserName, model.MsgTextLikeRevise)
	if msgID, err = p.pushSingleUser(c, msg.AccountID, message); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseLiked pushSingleUser failed %#v, msg_id(%s)", err, msgID))
		return
	}

}

func (p *Service) onDiscussionLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked Unmarshal failed %#v", err))
		return
	}

	var discuss *discuss.DiscussionInfo
	if discuss, err = p.d.GetDiscussion(c, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked GetDiscussion failed %#v", err))
		return
	}

	msg := &model.Message{
		ID:         gid.NewID(),
		AccountID:  discuss.Creator.ID,
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

	if err = p.d.AddMessage(c, p.d.DB(), msg); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked AddMessage failed %#v", err))
		return
	}

	m.Ack()

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked GetAccountBaseInfo failed %#v", err))
		return
	}

	var msgID string
	message := fmt.Sprintf("%s%s", account.UserName, model.MsgTextLikeDiscussion)
	if msgID, err = p.pushSingleUser(c, msg.AccountID, message); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionLiked pushSingleUser failed %#v, msg_id(%s)", err, msgID))
		return
	}
}
