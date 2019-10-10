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

func (p *Dao) GetUserArticlesPaged(c context.Context, aid int64, limit, offset int) (resp *article.UserArticlesResp, err error) {
	if resp, err = p.articleRPC.GetUserArticlesPaged(c, &article.UserArticlesReq{AccountID: aid, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserArticlesPaged, error(%+v), aid(%d), limit(%d), offset(%d)`", err, aid, limit, offset))
	}
	return
}