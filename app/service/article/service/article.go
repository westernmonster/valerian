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
	"valerian/library/xstr"

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

func (p *Service) GetArticleInfo(c context.Context, req *api.IDReq) (resp *api.ArticleInfo, err error) {
	article, err := p.GetArticle(c, req.ID)
	if err != nil {
		return nil, err
	}

	changeDesc, err := p.GetArticleLastChangeDesc(c, req.ID)
	if err != nil {
		return nil, err
	}

	stat, err := p.GetArticleStat(c, req.ID)
	if err != nil {
		return nil, err
	}

	urls, err := p.GetArticleImageUrls(c, req.ID)
	if err != nil {
		return nil, err
	}

	m, err := p.getAccount(c, p.d.DB(), article.CreatedBy)
	if err != nil {
		return nil, err
	}

	resp = &api.ArticleInfo{
		ID:             article.ID,
		Title:          article.Title,
		Excerpt:        xstr.Excerpt(article.ContentText),
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
		ImageUrls:      urls,
		DisableRevise:  bool(article.DisableRevise),
		DisableComment: bool(article.DisableComment),
		Stat: &api.ArticleStat{
			ReviseCount:  int32(stat.ReviseCount),
			CommentCount: int32(stat.CommentCount),
			LikeCount:    int32(stat.LikeCount),
			DislikeCount: int32(stat.DislikeCount),
		},
		Creator: &api.Creator{
			ID:           m.ID,
			UserName:     m.UserName,
			Avatar:       m.Avatar,
			Introduction: m.Introduction,
		},
		ChangeDesc: changeDesc,
	}

	inc := includeParam(req.Include)

	if inc["content"] {
		resp.Content = article.Content
	}

	if inc["content_text"] {
		resp.ContentText = article.ContentText
	}

	return

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

func (p *Service) GetUserArticlesPaged(c context.Context, req *api.UserArticlesReq) (resp *api.UserArticlesResp, err error) {
	var data []*model.Article
	if data, err = p.d.GetUserArticlesPaged(c, p.d.DB(), req.AccountID, int(req.Limit), int(req.Offset)); err != nil {
		return
	}

	resp = &api.UserArticlesResp{
		Items: make([]*api.ArticleInfo, len(data)),
	}

	for i, v := range data {
		stat, err := p.GetArticleStat(c, v.ID)
		if err != nil {
			return nil, err
		}

		urls, err := p.GetArticleImageUrls(c, v.ID)
		if err != nil {
			return nil, err
		}

		m, err := p.getAccount(c, p.d.DB(), v.CreatedBy)
		if err != nil {
			return nil, err
		}

		info := &api.ArticleInfo{
			ID:        v.ID,
			Title:     v.Title,
			Excerpt:   xstr.Excerpt(v.ContentText),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			ImageUrls: urls,
			Stat: &api.ArticleStat{
				ReviseCount:  int32(stat.ReviseCount),
				CommentCount: int32(stat.CommentCount),
				LikeCount:    int32(stat.LikeCount),
				DislikeCount: int32(stat.DislikeCount),
			},
			Creator: &api.Creator{
				ID:           m.ID,
				UserName:     m.UserName,
				Avatar:       m.Avatar,
				Introduction: m.Introduction,
			},
		}

		resp.Items[i] = info

	}

	return resp, nil
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

func (p *Service) DelArticle(c context.Context, arg *api.IDReq) (err error) {
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

	if canEdit, e := p.checkEditPermission(c, tx, arg.ID, arg.Aid); e != nil {
		return e
	} else if !canEdit {
		err = ecode.NeedArticleEditPermission
		return
	}

	if err = p.d.DelFavByCond(c, tx, arg.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if err = p.d.DelAccountFeedByCond(c, tx, arg.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if err = p.d.DelFeedByCond(c, tx, arg.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if err = p.d.DelTopicFeedByCond(c, tx, arg.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if err = p.d.DelRecentPubByCond(c, tx, arg.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if err = p.d.DelFeedbacksByCond(c, tx, arg.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if err = p.d.DelMessageByCond(c, tx, arg.ID, model.TargetTypeArticle); err != nil {
		return
	}

	if err = p.d.DelArticle(c, tx, arg.ID); err != nil {
		return
	}

	if err = p.d.IncrAccountStat(c, tx, &model.AccountStat{AccountID: item.CreatedBy, ArticleCount: -1}); err != nil {
		return
	}

	var catalogs []*model.TopicCatalog
	if catalogs, err = p.d.GetTopicCatalogsByCond(c, tx, map[string]interface{}{
		"type":   model.TopicCatalogArticle,
		"ref_id": arg.ID,
	}); err != nil {
		return
	}

	for _, v := range catalogs {
		if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{TopicID: v.TopicID, ArticleCount: -1}); err != nil {
			return
		}

		if err = p.d.DelTopicCatalog(c, tx, v.ID); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelReviseCache(context.TODO(), arg.ID)
		p.onReviseDeleted(context.Background(), arg.ID, arg.Aid, time.Now().Unix())
	})
	return
}

func (p *Service) UpdateArticle(c context.Context, arg *api.ArgUpdateArticle) (err error) {
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

	if canEdit, e := p.checkEditPermission(c, tx, arg.ID, arg.Aid); e != nil {
		err = e
		return
	} else if !canEdit {
		err = ecode.NeedArticleEditPermission
		return
	}

	oldContentText := item.ContentText

	var seq int
	if seq, err = p.d.GetArticleHistoriesMaxSeq(c, tx, arg.ID); err != nil {
		return
	}

	if arg.Title != nil {
		item.Title = arg.GetTitleValue()
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
		item.DisableRevise = types.BitBool(arg.GetDisableReviseValue())
	}

	if arg.DisableComment != nil {
		item.DisableComment = types.BitBool(arg.GetDisableCommentValue())
	}

	h := &model.ArticleHistory{
		ID:          gid.NewID(),
		ArticleID:   item.ID,
		Seq:         int32(seq + 1),
		Content:     item.Content,
		ContentText: item.ContentText,
		UpdatedBy:   arg.Aid,
		ChangeDesc:  arg.ChangeDesc,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldContentText, h.ContentText, false)
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
		p.onArticleUpdated(context.Background(), arg.ID, arg.Aid, time.Now().Unix())
	})

	return
}
