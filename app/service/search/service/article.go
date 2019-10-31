package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/app/service/search/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleCreated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded Unmarshal failed %#v", err))
		return
	}

	var v *model.Article
	if v, err = p.d.GetArticleByID(c, p.d.DB(), info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded GetArticleByID failed %#v", err))
		return
	} else if v == nil {
		m.Ack()
		return
	}

	changeDesc := ""
	var h *model.ArticleHistory
	if h, err = p.d.GetLastArticleHistory(c, p.d.DB(), v.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded GetArticleByID failed %#v", err))
		return
	} else if h != nil {
		changeDesc = h.ChangeDesc
	}

	item := &model.ESArticle{
		ID:          v.ID,
		Title:       &v.Title,
		Content:     &v.Content,
		ContentText: &v.ContentText,
		ChangeDesc:  &changeDesc,
		CreatedAt:   &v.CreatedAt,
		UpdatedAt:   &v.UpdatedAt,
	}
	disableRevise := bool(v.DisableRevise)
	item.DisableRevise = &disableRevise
	disableComment := bool(v.DisableComment)
	item.DisableComment = &disableComment

	var acc *model.Account
	if acc, err = p.d.GetAccountByID(c, p.d.DB(), v.CreatedBy); err != nil {
		return
	} else if acc == nil {
		m.Ack()
		return
	}

	item.Creator = &model.ESCreator{
		ID:           acc.ID,
		UserName:     &acc.UserName,
		Avatar:       &acc.Avatar,
		Introduction: &acc.Introduction,
	}

	if err = p.d.PutArticle2ES(c, item); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onArticleUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleUpdated Unmarshal failed %#v", err))
		return
	}

	var v *model.Article
	if v, err = p.d.GetArticleByID(c, p.d.DB(), info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded GetArticleByID failed %#v", err))
		return
	} else if v == nil {
		m.Ack()
		return
	}

	changeDesc := ""
	var h *model.ArticleHistory
	if h, err = p.d.GetLastArticleHistory(c, p.d.DB(), v.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded GetArticleByID failed %#v", err))
		return
	} else if h != nil {
		changeDesc = h.ChangeDesc
	}

	item := &model.ESArticle{
		ID:          v.ID,
		Title:       &v.Title,
		Content:     &v.Content,
		ContentText: &v.ContentText,
		ChangeDesc:  &changeDesc,
		CreatedAt:   &v.CreatedAt,
		UpdatedAt:   &v.UpdatedAt,
	}
	disableRevise := bool(v.DisableRevise)
	item.DisableRevise = &disableRevise
	disableComment := bool(v.DisableComment)
	item.DisableComment = &disableComment

	var acc *model.Account
	if acc, err = p.d.GetAccountByID(c, p.d.DB(), v.CreatedBy); err != nil {
		return
	} else if acc == nil {
		m.Ack()
		return
	}

	item.Creator = &model.ESCreator{
		ID:           acc.ID,
		UserName:     &acc.UserName,
		Avatar:       &acc.Avatar,
		Introduction: &acc.Introduction,
	}

	if err = p.d.PutArticle2ES(c, item); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onArticleDeleted(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleDeleted Unmarshal failed %#v", err))
		return
	}

	if err = p.d.DelESArticle(c, info.ArticleID); err != nil {
		return
	}

	m.Ack()
}
