package service

import (
	"context"
	"time"
	"valerian/app/service/account-feed/model"
	article "valerian/app/service/article/api"
	"valerian/app/service/feed/def"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleAdded(m *stan.Msg) {
	var err error
	info := new(def.MsgArticleCreated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onArticleAdded Unmarshal failed %#v", err)
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(context.Background(), info.ArticleID); err != nil {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCreateArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateArticle,
		TargetID:   article.ID,
		TargetType: def.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onArticleAdded() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleUpdated(m *stan.Msg) {
	var err error
	info := new(def.MsgArticleUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onArticleUpdated Unmarshal failed %#v", err)
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(context.Background(), info.ArticleID); err != nil {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeUpdateArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextUpdateArticle,
		TargetID:   article.ID,
		TargetType: def.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onArticleUpdated() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleLiked(m *stan.Msg) {
	var err error
	info := new(def.MsgArticleLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onArticleLiked Unmarshal failed %#v", err)
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(context.Background(), info.ArticleID); err != nil {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeLikeArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextLikeArticle,
		TargetID:   article.ID,
		TargetType: def.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onArticleLiked() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleFavd(m *stan.Msg) {
	var err error
	info := new(def.MsgArticleFaved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onArticleFavd Unmarshal failed %#v", err)
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(context.Background(), info.ArticleID); err != nil {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeFavArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextFavArticle,
		TargetID:   article.ID,
		TargetType: def.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onArticleFavd() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgArticleDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onArticleDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelAccountFeedByCond(context.Background(), p.d.DB(), def.TargetTypeArticle, info.ArticleID); err != nil {
		log.Errorf("service.onArticleDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
