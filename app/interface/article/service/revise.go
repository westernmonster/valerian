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

func (p *Service) GetReviseImageUrls(c context.Context, reviseID int64) (urls []string, err error) {
	urls = make([]string, 0)
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeRevise,
		"target_id":   reviseID,
	}); err != nil {
		return
	}

	for _, v := range imgs {
		urls = append(urls, v.URL)
	}

	return
}

func (p *Service) GetArticleRevisesPaged(c context.Context, articleID int64, sort string, limit, offset int) (resp *model.ReviseListResp, err error) {
	var data []*model.Revise
	if data, err = p.d.GetArticleRevisesPaged(c, p.d.DB(), articleID, sort, limit, offset); err != nil {
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

		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, v.CreatedBy); err != nil {
			return
		}
		item.Creator = &model.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		if item.ImageUrls, err = p.GetReviseImageUrls(c, item.ID); err != nil {
			return
		}

		item.CreatedAt = v.CreatedAt
		item.UpdatedAt = v.UpdatedAt

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/article/list/revises", url.Values{
		"article_id": []string{strconv.FormatInt(articleID, 10)},
		"sort":       []string{sort},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/article/list/revises", url.Values{
		"article_id": []string{strconv.FormatInt(articleID, 10)},
		"sort":       []string{sort},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset + limit)},
	}); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	var stat *model.ArticleStat
	if stat, err = p.d.GetArticleStatByID(c, p.d.DB(), articleID); err != nil {
		return
	}

	resp.ReviseCount = int32(stat.ReviseCount)

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

	// check article
	if _, err = p.getArticle(c, tx, arg.ArticleID); err != nil {
		return
	}

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

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		if url, exist := s.Attr("src"); exist {
			u := &model.ImageURL{
				ID:         gid.NewID(),
				TargetType: model.TargetTypeRevise,
				TargetID:   item.ID,
				URL:        url,
				CreatedAt:  time.Now().Unix(),
				UpdatedAt:  time.Now().Unix(),
			}
			if err = p.d.AddImageURL(c, tx, u); err != nil {
				return
			}
		}
	})

	if err = p.d.AddRevise(c, tx, item); err != nil {
		return
	}

	if err = p.bulkCreateReviseFiles(c, tx, item.ID, arg.Files); err != nil {
		return
	}

	if err = p.d.AddReviseStat(c, tx, &model.ReviseStat{
		ReviseID:  item.ID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = p.d.IncrArticleStat(c, tx, &model.ArticleStat{ArticleID: item.ArticleID, ReviseCount: 1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	id = item.ID
	p.addCache(func() {
		p.onReviseAdded(context.Background(), item.ID, aid, time.Now().Unix())
	})

	return
}

func (p *Service) UpdateRevise(c context.Context, arg *model.ArgUpdateRevise) (err error) {
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

	var item *model.Revise
	if item, err = p.d.GetReviseByID(c, tx, arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.ReviseNotExist
		return
	}

	item.Content = arg.Content

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		err = ecode.ParseHTMLFailed
		return
	}
	item.ContentText = doc.Text()

	if err = p.d.DelImageURLByCond(c, tx, model.TargetTypeRevise, item.ID); err != nil {
		return
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		if url, exist := s.Attr("src"); exist {
			u := &model.ImageURL{
				ID:         gid.NewID(),
				TargetType: model.TargetTypeArticle,
				TargetID:   item.ID,
				URL:        url,
				CreatedAt:  time.Now().Unix(),
				UpdatedAt:  time.Now().Unix(),
			}
			if err = p.d.AddImageURL(c, tx, u); err != nil {
				return
			}
		}
	})

	if err = p.d.UpdateRevise(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelReviseCache(context.TODO(), arg.ID)
		p.onReviseUpdated(context.Background(), arg.ID, aid, time.Now().Unix())
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
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var data *model.Revise
	if data, err = p.getRevise(c, p.d.DB(), reviseID); err != nil {
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, p.d.DB(), data.ArticleID); err != nil {
		return
	}

	var acc *account.BaseInfoReply
	if acc, err = p.d.GetAccountBaseInfo(c, data.CreatedBy); err != nil {
		return
	}
	resp = &model.ReviseDetailResp{
		ID:        data.ID,
		Title:     article.Title,
		ArticleID: data.ArticleID,
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	resp.Creator = &model.Creator{
		ID:           acc.ID,
		UserName:     acc.UserName,
		Avatar:       acc.Avatar,
		Introduction: acc.Introduction,
	}

	if resp.Files, err = p.GetReviseFiles(c, reviseID); err != nil {
		return
	}

	var stat *model.ReviseStat
	if stat, err = p.d.GetReviseStatByID(c, p.d.DB(), reviseID); err != nil {
		return
	}

	resp.DislikeCount = stat.DislikeCount
	resp.LikeCount = stat.LikeCount
	resp.CommentCount = stat.CommentCount

	if resp.Fav, err = p.d.IsFav(c, aid, reviseID, model.TargetTypeRevise); err != nil {
		return
	}

	if resp.Like, err = p.d.IsLike(c, aid, reviseID, model.TargetTypeRevise); err != nil {
		return
	}

	if resp.Dislike, err = p.d.IsDislike(c, aid, reviseID, model.TargetTypeRevise); err != nil {
		return
	}

	if aid == data.CreatedBy {
		resp.CanEdit = true
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
