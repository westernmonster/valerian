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

// GetArticle 获取文章信息
func (p *Service) GetArticle(c context.Context, articleID int64) (item *model.Article, err error) {
	return p.getArticle(c, p.d.DB(), articleID)
}

// GetArticleDetail 获取文章详情
func (p *Service) GetArticleDetail(c context.Context, req *api.IDReq) (resp *api.ArticleDetail, err error) {
	article, err := p.GetArticle(c, req.ID)
	if err != nil {
		return nil, err
	}

	lastHistory, err := p.d.GetLastArticleHistory(c, p.d.DB(), req.ID)
	if err != nil {
		return nil, err
	}

	stat, err := p.d.GetArticleStatByID(c, p.d.DB(), req.ID)
	if err != nil {
		return nil, err
	}

	urls, err := p.getArticleImageUrls(c, p.d.DB(), req.ID)
	if err != nil {
		return nil, err
	}

	m, err := p.getAccount(c, p.d.DB(), article.CreatedBy)
	if err != nil {
		return nil, err
	}

	resp = &api.ArticleDetail{
		ID:             article.ID,
		Title:          article.Title,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
		ImageUrls:      urls,
		Content:        article.Content,
		ContentText:    article.ContentText,
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
	}

	if lastHistory != nil {
		m, err := p.getAccount(c, p.d.DB(), lastHistory.UpdatedBy)
		if err != nil {
			return nil, err
		}
		resp.LastHistory = &api.ArticleHistoryResp{
			ID:         lastHistory.ID,
			ArticleID:  lastHistory.ArticleID,
			Seq:        lastHistory.Seq,
			ChangeDesc: lastHistory.ChangeDesc,
			Diff:       lastHistory.Diff,
			UpdatedAt:  lastHistory.UpdatedAt,
			CreatedAt:  lastHistory.CreatedAt,
			Updator: &api.Creator{
				ID:           m.ID,
				UserName:     m.UserName,
				Avatar:       m.Avatar,
				Introduction: m.Introduction,
			},
		}

	}

	return

}

// GetArticleInfos 批量获取文章信息
func (p *Service) GetArticleInfos(c context.Context, req *api.IDsReq) (resp *api.ArticleInfosResp, err error) {
	if dl, ok := c.Deadline(); ok {
		ctimeout := time.Until(dl)
		fmt.Println(ctimeout)
	}

	resp = &api.ArticleInfosResp{
		Items: make(map[int64]*api.ArticleInfo),
	}

	for _, v := range req.IDs {
		var item *api.ArticleInfo
		if item, err = p.getArticleInfo(c, p.d.DB(), v, ""); err != nil {
			if ecode.IsNotExistEcode(err) {
				resp.Items[v] = nil
				continue
			}
			return
		}

		resp.Items[v] = item
	}
	return
}

// GetArticleInfo 获取文章基本信息
func (p *Service) GetArticleInfo(c context.Context, req *api.IDReq) (resp *api.ArticleInfo, err error) {
	if req.UseMaster {
		c = sqalx.NewContext(c, true)
	}
	return p.getArticleInfo(c, p.d.DB(), req.ID, req.Include)
}

// getArticleInfo 获取文章基本信息
func (p *Service) getArticleInfo(c context.Context, node sqalx.Node, id int64, include string) (resp *api.ArticleInfo, err error) {
	article, err := p.getArticle(c, node, id)
	if err != nil {
		return nil, err
	}

	changeDesc, err := p.getArticleLastChangeDesc(c, node, id)
	if err != nil {
		return nil, err
	}

	stat, err := p.d.GetArticleStatByID(c, node, id)
	if err != nil {
		return nil, err
	}

	urls, err := p.getArticleImageUrls(c, node, id)
	if err != nil {
		return nil, err
	}

	m, err := p.getAccount(c, node, article.CreatedBy)
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

	inc := includeParam(include)

	if inc["content"] {
		resp.Content = article.Content
	}

	if inc["content_text"] {
		resp.ContentText = article.ContentText
	}

	return

}

// getArticleLastChangeDesc 获取文章最后一次改动备注
func (p *Service) getArticleLastChangeDesc(c context.Context, node sqalx.Node, articleID int64) (changeDesc string, err error) {
	var history *model.ArticleHistory
	if history, err = p.d.GetLastArticleHistory(c, node, articleID); err != nil {
		return
	} else if history != nil {
		changeDesc = history.ChangeDesc
	} else {
		changeDesc = ""
	}

	return
}

// GetAllArticles 获取所有文章信息
func (p *Service) GetAllArticles(c context.Context) (items []*model.Article, err error) {
	return p.d.GetArticles(c, p.d.DB())
}

// getArticle 获取文章信息
func (p *Service) getArticle(c context.Context, node sqalx.Node, articleID int64) (item *model.Article, err error) {
	var addCache = true
	if item, err = p.d.ArticleCache(c, articleID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetArticleByID(c, node, articleID); err != nil {
		return
	} else if item == nil {
		err = ecode.ArticleNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleCache(context.Background(), item)
		})
	}
	return
}

// getArticleImageUrls 获取文章图片
func (p *Service) getArticleImageUrls(c context.Context, node sqalx.Node, articleID int64) (urls []string, err error) {
	urls = make([]string, 0)
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, node, map[string]interface{}{
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

// GetArticleStat 获取文章状态值
func (p *Service) GetArticleStat(c context.Context, articleID int64) (stat *model.ArticleStat, err error) {
	return p.d.GetArticleStatByID(c, p.d.DB(), articleID)
}

// GetUserArticleIDsPaged 获取用户创建的文章ID列表
func (p *Service) GetUserArticleIDsPaged(c context.Context, req *api.UserArticlesReq) (ids []int64, err error) {
	return p.d.GetUserArticleIDsPaged(c, p.d.DB(), req.AccountID, int(req.Limit), int(req.Offset))
}

// GetUserArticlesPaged  获取用户创建的文章信息列表
func (p *Service) GetUserArticlesPaged(c context.Context, req *api.UserArticlesReq) (resp *api.UserArticlesResp, err error) {
	var data []*model.Article
	if data, err = p.d.GetUserArticlesPaged(c, p.d.DB(), req.AccountID, int(req.Limit), int(req.Offset)); err != nil {
		return
	}

	resp = &api.UserArticlesResp{
		Items: make([]*api.ArticleInfo, len(data)),
	}

	for i, v := range data {
		stat, err := p.d.GetArticleStatByID(c, p.d.DB(), v.ID)
		if err != nil {
			return nil, err
		}

		urls, err := p.getArticleImageUrls(c, p.d.DB(), v.ID)
		if err != nil {
			return nil, err
		}

		m, err := p.getAccount(c, p.d.DB(), v.CreatedBy)
		if err != nil {
			return nil, err
		}

		changeDesc, err := p.getArticleLastChangeDesc(c, p.d.DB(), v.ID)
		if err != nil {
			return nil, err
		}

		info := &api.ArticleInfo{
			ID:         v.ID,
			Title:      v.Title,
			Excerpt:    xstr.Excerpt(v.ContentText),
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			ImageUrls:  urls,
			ChangeDesc: changeDesc,
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

// AddArticle 添加文章
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

	if arg.Relations == nil || len(arg.Relations) == 0 {
		err = ecode.NeedArticleRelation
		return
	}

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

	var ids []int64
	if ids, err = p.bulkCreateArticleRelations(c, tx, arg.Aid, item.ID, item.Title, arg.Relations); err != nil {
		return
	}

	if err = p.d.IncrAccountStat(c, tx, &model.AccountStat{AccountID: item.CreatedBy, ArticleCount: 1}); err != nil {
		return
	}

	for _, v := range ids {
		if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{TopicID: v, ArticleCount: 1}); err != nil {
			return
		}
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
			p.onCatalogArticleAdded(context.Background(), item.ID, h.ID, v.ToTopicID, arg.Aid, time.Now().Unix())
		}
	})
	return
}

// DelArticle 删除文章
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

	// 检查编辑权限
	if err = p.checkEditPermission(c, tx, arg.Aid, arg.ID); err != nil {
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
		p.d.DelArticleCache(context.Background(), arg.ID)

		for _, v := range catalogs {
			p.d.DelTopicCatalogCache(context.Background(), v.TopicID)
		}

		p.onArticleDeleted(context.Background(), arg.ID, arg.Aid, time.Now().Unix())
	})
	return
}

// UpdateArticle 更新文章
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

	if err = p.checkEditPermission(c, tx, arg.Aid, arg.ID); err != nil {
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

	item.UpdatedAt = time.Now().Unix()

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
		p.d.DelArticleCache(context.Background(), arg.ID)
		p.onArticleUpdated(context.Background(), arg.ID, h.ID, arg.Aid, time.Now().Unix())
	})

	return
}
