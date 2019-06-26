package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) AddArticle(c context.Context, arg *model.ArgAddArticle) (id int64, err error) {
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

	item := &model.Article{
		ID:           gid.NewID(),
		Title:        arg.Title,
		Content:      arg.Content,
		Cover:        arg.Cover,
		Introduction: arg.Introduction,
		Private:      types.BitBool(arg.Private),
		Locale:       arg.Locale,
		VersionName:  arg.VersionName,
		CreatedBy:    aid,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}

	if arg.ArticleSetID == nil {
		set := &model.ArticleSet{
			ID:        gid.NewID(),
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		item.ArticleSetID = set.ID

		if err = p.d.AddArticleSet(c, tx, set); err != nil {
			return
		}
	} else {
		if v, e := p.d.GetArticleVersionByName(c, tx, *arg.ArticleSetID, arg.VersionName); e != nil {
			err = e
			return
		} else if v != nil {
			err = ecode.ArticleVersionNameExist
			return
		}

		item.ArticleSetID = *arg.ArticleSetID
	}

	var locale *model.Locale
	if locale, err = p.d.GetLocaleByCondition(c, tx, map[string]interface{}{"locale": arg.Locale}); err != nil {
		return
	} else if locale == nil {
		err = ecode.LocaleNotExist
		return
	}

	if err = p.d.AddArticle(c, tx, item); err != nil {
		return
	}

	h := &model.ArticleHistory{
		ID:          gid.NewID(),
		ArticleID:   item.ID,
		UpdatedBy:   aid,
		Content:     item.Content,
		Diff:        "",
		Description: "",
		Seq:         1,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	var doc *goquery.Document
	if doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content)); err != nil {
		return
	}

	h.ContentText = doc.Text()
	if err = p.d.AddArticleHistory(c, tx, h); err != nil {
		return
	}

	if err = p.bulkCreateFiles(c, tx, item.ID, arg.Files); err != nil {
		return
	}

	if err = p.bulkCreateArticleCatalogs(c, tx, item.ID, item.Title, arg.Relations); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelArticleVersionCache(context.TODO(), item.ArticleSetID)
	})

	return
}

func (p *Service) UpdateArticle(c context.Context, arg *model.ArgUpdateArticle) (err error) {
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

	var primaryTopicID int64
	if primaryTopicID, err = p.getPrimaryTopicID(c, tx, arg.ID); err != nil {
		return
	}

	if err = p.checkEditPermission(c, tx, primaryTopicID); err != nil {
		return
	}

	var item *model.Article
	if item, err = p.d.GetArticleByID(c, tx, arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.ArticleNotExist
		return
	}

	if arg.Title != nil {
		item.Title = *arg.Title
	}

	if arg.Cover != nil {
		item.Cover = arg.Cover
	}

	if arg.Locale != nil {
		item.Locale = *arg.Locale

		var locale *model.Locale
		if locale, err = p.d.GetLocaleByCondition(c, tx, map[string]interface{}{"locale": *arg.Locale}); err != nil {
			return
		} else if locale == nil {
			err = ecode.LocaleNotExist
			return
		}
	}

	if arg.Introduction != nil {
		item.Introduction = *arg.Introduction
	}

	if arg.VersionName != nil {
		if resp, e := p.d.GetArticleVersionByName(c, tx, item.ArticleSetID, *arg.VersionName); e != nil {
			err = e
			return
		} else if resp != nil && resp.ArticleID != arg.ID {
			err = ecode.ArticleVersionNameExist
			return
		}

		item.VersionName = *arg.VersionName
	}

	if arg.Private != nil && aid == item.CreatedBy {
		if *arg.Private {
			if count, e := p.d.GetOrderMemberArticleHistoriesCount(c, tx, item.ID, item.CreatedBy); e != nil {
				err = e
				return
			} else if count > 0 {
				err = ecode.ArticleEditedByOthers
				return
			}
		}

		item.Private = types.BitBool(*arg.Private)
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelArticleCache(context.TODO(), arg.ID)
		p.d.DelArticleVersionCache(context.TODO(), item.ArticleSetID)
	})

	return
}

func (p *Service) getPrimaryTopicID(c context.Context, node sqalx.Node, articleID int64) (topicID int64, err error) {
	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByCondition(c, node, map[string]interface{}{
		"type":       model.TopicCatalogTaxonomy,
		"ref_id":     articleID,
		"is_primary": 1,
	}); err != nil {
		return
	} else if catalog == nil {
		err = ecode.TopicCatalogNotExist
		return
	}

	topicID = catalog.TopicID

	return
}

func (p *Service) DelArticle(c context.Context, id int64) (err error) {
	return
}

func (p *Service) GetArticle(c context.Context, id int64) (item *model.ArticleResp, err error) {
	if item, err = p.getArticle(c, p.d.DB(), id); err != nil {
		return
	}

	if item.Files, err = p.getArticleFiles(c, p.d.DB(), id); err != nil {
		return
	}

	if item.Relations, err = p.getArticleRelations(c, p.d.DB(), id); err != nil {
		return
	}

	if item.Versions, err = p.getArticleVersionsResp(c, p.d.DB(), id); err != nil {
		return
	}

	return
}

func (p *Service) getArticle(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleResp, err error) {
	var addCache = true
	if item, err = p.d.ArticleCache(c, articleID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	var a *model.Article
	if a, err = p.d.GetArticleByID(c, p.d.DB(), articleID); err != nil {
		return
	} else if a == nil {
		err = ecode.ArticleNotExist
		return
	}

	item = &model.ArticleResp{
		ID:           a.ID,
		Title:        a.Title,
		Content:      a.Content,
		ArticleSetID: a.ArticleSetID,
		Locale:       a.Locale,
		Cover:        a.Cover,
		Introduction: a.Introduction,
		Private:      bool(a.Private),
		VersionName:  a.VersionName,
		Seq:          a.Seq,
		Files:        make([]*model.ArticleFileResp, 0),
		Relations:    make([]*model.ArticleRelationResp, 0),
		Versions:     make([]*model.ArticleVersionResp, 0),
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleCache(context.TODO(), item)
		})
	}
	return
}

func (p *Service) checkEditPermission(c context.Context, node sqalx.Node, topicID int64) (err error) {
	var t *model.TopicResp
	if t, err = p.getTopic(c, node, topicID); err != nil {
		return
	}
	var meta *model.TopicMeta
	if meta, err = p.GetTopicMeta(c, t); err != nil {
		return
	}

	if !meta.CanEdit {
		err = ecode.NeedEditPermission
		return
	}

	return
}
