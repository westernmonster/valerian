package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"valerian/app/interface/article/model"
	account "valerian/app/service/account/api"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/PuerkitoBio/goquery"
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
		ID:             gid.NewID(),
		Title:          arg.Title,
		Content:        arg.Content,
		DisableComment: types.BitBool(arg.DisableComment),
		DisableRevise:  types.BitBool(arg.DisableRevise),
		CreatedBy:      aid,
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
		UpdatedBy:   aid,
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

	var relations []*model.ArticleRelationResp
	if relations, err = p.getArticleRelations(c, p.d.DB(), item.ID); err != nil {
		return
	}

	p.addCache(func() {
		p.onArticleAdded(context.Background(), item.ID, aid, time.Now().Unix())
		for _, v := range relations {
			p.onCatalogArticleAdded(context.Background(), item.ID, v.ToTopicID, aid, time.Now().Unix())
		}
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

	var item *model.Article
	if item, err = p.d.GetArticleByID(c, tx, arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.ArticleNotExist
		return
	}

	oldContentText := item.ContentText

	var seq int
	if seq, err = p.d.GetArticleHistoriesMaxSeq(c, tx, arg.ID); err != nil {
		return
	}

	if arg.Title != nil {
		item.Title = *arg.Title
	}

	item.Content = arg.Content

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		err = ecode.ParseHTMLFailed
		return
	}
	item.ContentText = doc.Text()

	if err = p.d.DelImageURLByCond(c, tx, model.TargetTypeArticle, item.ID); err != nil {
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

	if arg.DisableRevise != nil {
		item.DisableRevise = types.BitBool(*arg.DisableRevise)
	}

	if arg.DisableComment != nil {
		item.DisableComment = types.BitBool(*arg.DisableComment)
	}

	h := &model.ArticleHistory{
		ID:          gid.NewID(),
		ArticleID:   item.ID,
		Seq:         seq + 1,
		Content:     item.Content,
		ContentText: item.ContentText,
		UpdatedBy:   aid,
		ChangeDesc:  arg.ChangeDesc,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldContentText, h.ContentText, false)
	fmt.Println(dmp.DiffPrettyText(diffs))
	h.Diff = dmp.DiffPrettyHtml(diffs)

	if err = p.d.UpdateArticle(c, tx, item); err != nil {
		return
	}

	if err = p.d.AddArticleHistory(c, tx, h); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelArticleCache(context.TODO(), arg.ID)
		p.onArticleUpdated(context.Background(), arg.ID, aid, time.Now().Unix())
	})

	return
}

// func (p *Service) DelArticle(c context.Context, id int64) (err error) {
// 	return
// }

func (p *Service) GetArticle(c context.Context, id int64, include string) (item *model.ArticleResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	inc := includeParam(include)
	var data *model.Article
	if data, err = p.getArticle(c, p.d.DB(), id); err != nil {
		return
	}

	item = &model.ArticleResp{
		ID:        data.ID,
		Title:     data.Title,
		Content:   data.Content,
		CreatedBy: data.CreatedBy,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Files:     make([]*model.ArticleFileResp, 0),
		Relations: make([]*model.ArticleRelationResp, 0),
	}

	var acc *account.BaseInfoReply
	if acc, err = p.d.GetAccountBaseInfo(c, item.CreatedBy); err != nil {
		return
	}

	item.Creator = &model.Creator{
		ID:           acc.ID,
		UserName:     acc.UserName,
		Avatar:       acc.Avatar,
		Introduction: acc.Introduction,
	}

	var history *model.ArticleHistory
	if history, err = p.d.GetLastArticleHistory(c, p.d.DB(), id); err != nil {
		return
	} else if history != nil {
		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, history.UpdatedBy); err != nil {
			return
		}

		item.Updator = &model.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		item.ChangeDesc = history.ChangeDesc
	}

	if inc["files"] {
		if item.Files, err = p.getArticleFiles(c, p.d.DB(), id); err != nil {
			return
		}
	}

	if inc["relations"] {
		if item.Relations, err = p.getArticleRelations(c, p.d.DB(), id); err != nil {
			return
		}
	}

	if inc["meta"] {
		if item.ArticleMeta, err = p.getArticleMeta(c, p.d.DB(), data); err != nil {
			return
		}
	}

	p.addCache(func() {
		p.onArticleViewed(context.Background(), id, aid, time.Now().Unix())
	})

	return
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

func (p *Service) getArticleMeta(c context.Context, node sqalx.Node, article *model.Article) (item *model.ArticleMeta, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	item = new(model.ArticleMeta)

	if aid == article.CreatedBy {
		item.CanEdit = true
	}

	if item.Like, err = p.d.IsLike(c, aid, article.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if item.Fav, err = p.d.IsFav(c, aid, article.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if item.Dislike, err = p.d.IsDislike(c, aid, article.ID, model.TargetTypeArticle); err != nil {
		return
	}

	var stat *model.ArticleStat
	if stat, err = p.d.GetArticleStatByID(c, node, article.ID); err != nil {
		return
	}

	item.LikeCount = stat.LikeCount
	item.DislikeCount = stat.DislikeCount
	item.ReviseCount = stat.ReviseCount
	item.CommentCount = stat.CommentCount

	return
}
