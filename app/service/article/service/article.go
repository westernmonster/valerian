package service

import (
	"context"

	account "valerian/app/service/account/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error) {
	return p.d.GetAccountBaseInfo(c, aid)
}

func (p *Service) GetArticle(c context.Context, articleID int64) (item *model.Article, err error) {
	return p.getArticle(c, p.d.DB(), articleID)
}

func (p *Service) getArticle(c context.Context, node sqalx.Node, articleID int64) (item *model.Article, err error) {
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
