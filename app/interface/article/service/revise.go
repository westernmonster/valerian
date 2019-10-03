package service

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	"valerian/app/interface/article/model"
	account "valerian/app/service/account/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
	"valerian/library/xstr"

	"github.com/PuerkitoBio/goquery"
)

func (p *Service) GetArticleRevisesPaged(c context.Context, articleID int64, offset, limit int) (resp *model.ReviseListResp, err error) {
	var data []*model.Revise
	if data, err = p.d.GetArticleRevisesPaged(c, p.d.DB(), articleID, offset, limit); err != nil {
		return
	}

	resp = &model.ReviseListResp{
		Paging: &model.Paging{},
		Items:  make([]*model.ReviseItem, len(data)),
	}

	for i, v := range data {
		item := &model.ReviseItem{
			ID:        v.ID,
			Excerpt:   xstr.Excerpt(v.ContentText),
			ImageUrls: make([]string, 0),
		}

		var account *account.BaseInfoReply
		if account, err = p.d.GetAccountBaseInfo(c, v.CreatedBy); err != nil {
			return
		}
		item.Creator = &model.Creator{
			ID:       account.ID,
			UserName: account.UserName,
			Avatar:   account.Avatar,
		}
		intro := account.GetIntroductionValue()
		item.Creator.Introduction = &intro

		resp.Items[i] = item
	}

	param := url.Values{}
	param.Set("article_id", strconv.FormatInt(articleID, 10))
	param.Set("limit", strconv.Itoa(limit))
	param.Set("offset", strconv.Itoa(offset-limit))

	if resp.Paging.Prev, err = genURL("/api/v1/article/list/revises", param); err != nil {
		return
	}
	param.Set("offset", strconv.Itoa(offset+limit))
	if resp.Paging.Next, err = genURL("/api/v1/article/list/revises", param); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	return
}

func (p *Service) AddRevise(c context.Context, arg *model.ArgAddRevise) (id int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
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

	item := &model.Revise{
		ID:        gid.NewID(),
		ArticleID: arg.ArticleID,
		Content:   arg.Content,
		CreatedBy: aid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		err = ecode.ParseHTMLFailed
		return
	}
	item.ContentText = doc.Text()

	if err = p.d.AddRevise(c, tx, item); err != nil {
		return
	}

	if err = p.bulkCreateReviseFiles(c, tx, item.ID, arg.Files); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	id = item.ID

	return
}

func (p *Service) UpdateRevise(c context.Context, arg *model.ArgUpdateRevise) (err error) {
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

	var item *model.Revise
	if item, err = p.d.GetReviseByID(c, tx, arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.ReviseNotExist
		return
	}

	item.Content = arg.Content

	if err = p.d.UpdateRevise(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelReviseCache(context.TODO(), arg.ID)
	})

	return
}

func (p *Service) DelRevise(c context.Context, id int64) (err error) {

	p.addCache(func() {
		p.d.DelReviseCache(context.TODO(), id)
	})

	return
}

func (p *Service) GetRevise(c context.Context, reviseID int64) (resp *model.ReviseDetailResp, err error) {
	var data *model.Revise
	if data, err = p.getRevise(c, p.d.DB(), reviseID); err != nil {
		return
	}

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, data.CreatedBy); err != nil {
		return
	}
	resp = &model.ReviseDetailResp{
		ID:        data.ID,
		ArticleID: data.ArticleID,
		Content:   data.Content,
	}

	resp.Creator = &model.Creator{
		ID:       account.ID,
		UserName: account.UserName,
		Avatar:   account.Avatar,
	}
	intro := account.GetIntroductionValue()
	resp.Creator.Introduction = &intro

	if resp.Files, err = p.GetReviseFiles(c, reviseID); err != nil {
		return
	}

	return
}

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
			p.d.SetReviseCache(context.TODO(), item)
		})
	}
	return
}
