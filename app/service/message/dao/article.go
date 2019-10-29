package dao

import (
	"context"
	"fmt"
	article "valerian/app/service/article/api"
	"valerian/library/log"
)

func (p *Dao) GetArticle(c context.Context, id int64, useMaster bool) (info *article.ArticleInfo, err error) {
	if info, err = p.articleRPC.GetArticleInfo(c, &article.IDReq{ID: id, UseMaster: useMaster}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticle, error(%+v) id(%d)", err, id))
	}
	return
}

func (p *Dao) GetRevise(c context.Context, id int64, useMaster bool) (info *article.ReviseInfo, err error) {
	if info, err = p.articleRPC.GetReviseInfo(c, &article.IDReq{ID: id, UseMaster: useMaster}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRevise, error(%+v) id(%d)", err, id))
	}
	return
}
