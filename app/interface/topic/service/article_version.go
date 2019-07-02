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
	if a, err = p.d.GetArticleByID(c, tx, arg.ArticleID); err != nil {
		return
	} else if a == nil {
		err = ecode.ArticleNotExist
		return
	}

	var ver *model.ArticleVersionResp
	if ver, err = p.d.GetArticleVersionByName(c, tx, arg.ArticleID, arg.Name); err != nil {
		return
	} else if ver != nil {
		err = ecode.ArticleVersionNameExist
		return
	}

	var maxSeq int
	if maxSeq, err = p.d.GetTopicVersionMaxSeq(c, tx, arg.ArticleID); err != nil {
		return
	}

	item := &model.ArticleVersion{
		ID:        gid.NewID(),
		ArticleID: arg.ArticleID,
		Name:      arg.Name,
		Content:   "",
		Seq:       maxSeq + 1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddArticleVersion(c, tx, item); err != nil {
		return
	}

	h := &model.ArticleHistory{
		ID:               gid.NewID(),
		ArticleVersionID: item.ID,
		UpdatedBy:        aid,
		Diff:             "",
		ChangeID:         "",
		Description:      "",
		Seq:              1,
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        time.Now().Unix(),
	}

	var doc *goquery.Document
	if doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content)); err != nil {
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

	id = item.ID

	p.addCache(func() {
		p.d.DelArticleVersionsCache(context.TODO(), arg.ArticleID)
	})

	return
}

func (p *Service) getArticleVersions(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleVersionResp, err error) {
	var addCache = true

	if items, err = p.d.ArticleVersionsCache(c, articleID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetArticleVersions(c, node, articleID); err != nil {
		return
	}
	fmt.Println(articleID)

	if addCache {
		p.addCache(func() {
			p.d.SetArticleVersionsCache(context.TODO(), articleID, items)
		})
	}

	return
}

func (p *Service) GetArticleVersions(c context.Context, articleID int64) (items []*model.ArticleVersionResp, err error) {
	return p.getArticleVersions(c, p.d.DB(), articleID)
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

	var t *model.Article
	if t, err = p.d.GetArticleByID(c, tx, arg.ArticleID); err != nil {
		return
	} else if t == nil {
		return ecode.ArticleNotExist
	}

	var primaryTopicID int64
	if primaryTopicID, _, err = p.getPrimaryTopicInfo(c, tx, arg.ArticleID); err != nil {
		return
	}
	if err = p.checkEditPermission(c, tx, primaryTopicID); err != nil {
		return
	}

	for _, v := range arg.Items {
		var ver *model.ArticleVersion
		if ver, err = p.d.GetArticleVersion(c, tx, v.ID); err != nil {
			return
		} else if ver == nil {
			return ecode.ArticleVersionNotExist
		}

		if m, e := p.d.GetArticleVersionByName(c, tx, arg.ArticleID, v.Name); e != nil {
			return e
		} else if m != nil && m.ID != ver.ID {
			err = ecode.ArticleVersionNameExist
			return
		}

		ver.Name = v.Name
		ver.Seq = v.Seq

		if err = p.d.UpdateArticleVersion(c, tx, ver); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelArticleVersionsCache(context.TODO(), arg.ArticleID)
	})

	return
}
