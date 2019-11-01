package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/PuerkitoBio/goquery"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func (p *Service) getAccount(c context.Context, node sqalx.Node, aid int64) (info *model.Account, err error) {
	if info, err = p.d.GetAccountByID(c, node, aid); err != nil {
		return
	} else if info == nil {
		err = ecode.UserNotExist
		return
	}

	return
}

func (p *Service) getTopic(c context.Context, node sqalx.Node, tid int64) (info *model.Topic, err error) {
	if info, err = p.d.GetTopicByID(c, node, tid); err != nil {
		return
	} else if info == nil {
		err = ecode.TopicNotExist
		return
	}

	return
}

func (p *Service) GetArticle(c context.Context, articleID int64) (item *model.Article, err error) {
	return p.getArticle(c, p.d.DB(), articleID)
}

func (p *Service) GetArticleLastChangeDesc(c context.Context, articleID int64) (changeDesc string, err error) {
	var history *model.ArticleHistory
	if history, err = p.d.GetLastArticleHistory(c, p.d.DB(), articleID); err != nil {
		return
	} else if history != nil {
		changeDesc = history.ChangeDesc
	} else {
		changeDesc = ""
	}

	return
}

func (p *Service) GetAllArticles(c context.Context) (items []*model.Article, err error) {
	return p.d.GetArticles(c, p.d.DB())
}

func (p *Service) getArticle(c context.Context, node sqalx.Node, articleID int64) (item *model.Article, err error) {
	var addCache = true
	if item, err = p.d.ArticleCache(c, articleID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetArticleByID(c, p.d.DB(), articleID); err != nil {
		return
	} else if item == nil {
		err = ecode.ArticleNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleCache(context.TODO(), item)
		})
	}
	return
}

func (p *Service) GetArticleImageUrls(c context.Context, articleID int64) (urls []string, err error) {
	urls = make([]string, 0)
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeArticle,
		"target_id":   articleID,
	}); err != nil {
		return
	}

	for _, v := range imgs {
		urls = append(urls, v.URL)
	}

	return
}

func (p *Service) GetArticleStat(c context.Context, articleID int64) (stat *model.ArticleStat, err error) {
	return p.d.GetArticleStatByID(c, p.d.DB(), articleID)
}

func (p *Service) GetUserArticlesPaged(c context.Context, aid int64, limit, offset int) (items []*model.Article, err error) {
	if items, err = p.d.GetUserArticlesPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	return
}

func (p *Service) AddArticle(c context.Context, arg *api.ArgAddArticle) (id int64, err error) {
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
		ID:             gid.NewID(),
		Title:          arg.Title,
		Content:        arg.Content,
		DisableComment: types.BitBool(arg.DisableComment),
		DisableRevise:  types.BitBool(arg.DisableRevise),
		CreatedBy:      arg.Aid,
		CreatedAt:      time.Now().Unix(),
		UpdatedAt:      time.Now().Unix(),
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

	h := &model.ArticleHistory{
		ID:          gid.NewID(),
		ArticleID:   item.ID,
		Seq:         1,
		Content:     item.Content,
		ContentText: item.ContentText,
		UpdatedBy:   arg.Aid,
		ChangeDesc:  "创建文章",
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain("", item.ContentText, false)
	h.Diff = dmp.DiffPrettyHtml(diffs)

	if err = p.d.AddArticle(c, tx, item); err != nil {
		return
	}

	if err = p.d.AddArticleHistory(c, tx, h); err != nil {
		return
	}

	if err = p.d.AddArticleStat(c, tx, &model.ArticleStat{
		ArticleID: item.ID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = p.bulkCreateFiles(c, tx, item.ID, arg.Files); err != nil {
		return
	}

	if err = p.bulkCreateArticleRelations(c, tx, item.ID, item.Title, arg.Relations); err != nil {
		return
	}

	if err = p.d.IncrAccountStat(c, tx, &model.AccountStat{AccountID: item.CreatedBy, ArticleCount: 1}); err != nil {
		return
	}

	if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{ArticleCount: 1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	id = item.ID

	var relations []*api.ArticleRelationResp
	if relations, err = p.getArticleRelations(c, p.d.DB(), item.ID); err != nil {
		return
	}

	p.addCache(func() {
		p.onArticleAdded(context.Background(), item.ID, arg.Aid, time.Now().Unix())
		for _, v := range relations {
			p.onCatalogArticleAdded(context.Background(), item.ID, v.ToTopicID, arg.Aid, time.Now().Unix())
		}
	})
	return
}
