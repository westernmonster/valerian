package dao

import (
	"context"
	"fmt"

	article "valerian/app/service/article/api"
	"valerian/library/log"
)

func (p *Dao) GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error) {
	if info, err = p.articleRPC.GetArticleInfo(c, &article.IDReq{ID: id}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticle, error(%+v) id(%d)", err, id))
	}
	return
}

func (p *Dao) GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error) {
	if info, err = p.articleRPC.GetReviseInfo(c, &article.IDReq{ID: id}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRevise, error(%+v) id(%d)", err, id))
	}
	return
}

func (p *Dao) GetAllArticles(c context.Context) (items []*article.DBArticle, err error) {
	var resp *article.ArticlesResp
	if resp, err = p.articleRPC.GetAllArticles(c, &article.EmptyStruct{}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllArticles, error(%+v) ", err))
		return
	}

	items = make([]*article.DBArticle, 0)
	if resp.Items != nil {
		items = resp.Items
	}

	return
}
