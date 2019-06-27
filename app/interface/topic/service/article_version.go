package service

import (
	"context"
	"fmt"
	"strings"
	"time"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"

	"github.com/PuerkitoBio/goquery"
)

func (p *Service) AddArticleVersion(c context.Context, arg *model.ArgNewArticleVersion) (id int64, err error) {
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

	var a *model.Article
	if a, err = p.d.GetArticleByID(c, tx, arg.FromArticleID); err != nil {
		return
	} else if a == nil {
		err = ecode.ArticleNotExist
		return
	}

	if v, e := p.d.GetArticleVersionByName(c, tx, a.ArticleSetID, arg.VersionName); e != nil {
		return 0, e
	} else if v != nil {
		err = ecode.ArticleVersionNameExist
		return
	}

	a.ID = gid.NewID()
	a.VersionName = arg.VersionName
	a.CreatedBy = aid
	a.CreatedAt = time.Now().Unix()
	a.UpdatedAt = time.Now().Unix()

	if err = p.d.AddArticle(c, tx, a); err != nil {
		return
	}

	if err = p.copyArticleRelations(c, tx, aid, arg.FromArticleID, a.ID); err != nil {
		return
	}

	if err = p.copyArticleFiles(c, tx, aid, arg.FromArticleID, a.ID); err != nil {
		return
	}

	h := &model.ArticleHistory{
		ID:          gid.NewID(),
		ArticleID:   a.ID,
		UpdatedBy:   aid,
		Content:     a.Content,
		Diff:        "",
		Description: "",
		Seq:         1,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	var doc *goquery.Document
	if doc, err = goquery.NewDocumentFromReader(strings.NewReader(a.Content)); err != nil {
		return
	}

	h.ContentText = doc.Text()
	if err = p.d.AddArticleHistory(c, tx, h); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	id = a.ID

	return
}

func (p *Service) copyArticleRelations(c context.Context, node sqalx.Node, aid int64, fromArticleID, toArticleID int64) (err error) {
	var catalogs []*model.TopicCatalog
	if catalogs, err = p.d.GetTopicCatalogsByCondition(c, node, map[string]interface{}{
		"ref_id": fromArticleID,
		"type":   model.TopicCatalogArticle,
	}); err != nil {
		return
	}

	for _, v := range catalogs {
		v.ID = gid.NewID()
		v.CreatedAt = time.Now().Unix()
		v.UpdatedAt = time.Now().Unix()
		v.RefID = &toArticleID

		if err = p.d.AddTopicCatalog(c, node, v); err != nil {
			return
		}
	}

	return
}

func (p *Service) copyArticleFiles(c context.Context, node sqalx.Node, aid int64, fromArticleID, toArticleID int64) (err error) {
	var files []*model.ArticleFile
	if files, err = p.d.GetArticleFiles(c, node, fromArticleID); err != nil {
		return
	}

	for _, v := range files {
		v.ID = gid.NewID()
		v.CreatedAt = time.Now().Unix()
		v.UpdatedAt = time.Now().Unix()
		v.ArticleID = toArticleID
		if err = p.d.AddArticleFile(c, node, v); err != nil {
			return
		}
	}
	return
}

func (p *Service) getArticleVersionsResp(c context.Context, node sqalx.Node, articleSetID int64) (items []*model.ArticleVersionResp, err error) {
	items = make([]*model.ArticleVersionResp, 0)
	var data []int64
	if data, err = p.getArticleVersions(c, node, articleSetID); err != nil {
		return
	}

	for _, v := range data {
		t, e := p.getArticle(c, node, v)
		if e != nil {
			err = e
			return
		}

		version := &model.ArticleVersionResp{
			ArticleSetID: t.ArticleSetID,
			ArticleID:    t.ID,
			Seq:          t.Seq,
			VersionName:  t.VersionName,
			ArticleTitle: t.Title,
		}

		items = append(items, version)
	}

	return
}

func (p *Service) getArticleVersions(c context.Context, node sqalx.Node, articleSetID int64) (items []int64, err error) {
	var addCache = true

	if items, err = p.d.ArticleVersionCache(c, articleSetID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetArticleVersions(c, node, articleSetID); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleVersionCache(context.TODO(), articleSetID, items)
		})
	}

	return
}

func (p *Service) GetArticleVersions(c context.Context, articleSetID int64) (items []*model.ArticleVersionResp, err error) {
	return p.getArticleVersionsResp(c, p.d.DB(), articleSetID)
}

func (p *Service) SaveArticleVersions(c context.Context, arg *model.ArgSaveArticleVersions) (err error) {
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

	for _, v := range arg.Items {
		var t *model.Article
		if t, err = p.d.GetArticleByID(c, tx, v.ArticleID); err != nil {
			return
		} else if t == nil {
			return ecode.ArticleNotExist
		}

		t.VersionName = v.VersionName
		t.Seq = v.Seq

		if err = p.d.UpdateArticle(c, tx, t); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelArticleVersionCache(context.TODO(), arg.ArticleSetID)
		for _, v := range arg.Items {
			p.d.DelArticleCache(context.TODO(), v.ArticleID)
		}
	})

	return
}
