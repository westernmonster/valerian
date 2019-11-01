package dao

import (
	"context"
	"fmt"
	article "valerian/app/service/article/api"
	"valerian/library/log"
)

func (p *Dao) GetArticleInfo(c context.Context, req *article.IDReq) (resp *article.ArticleInfo, err error) {
	if resp, err = p.articleRPC.GetArticleInfo(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleInfo() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) GetReviseStat(c context.Context, req *article.IDReq) (resp *article.ReviseStat, err error) {
	if resp, err = p.articleRPC.GetReviseStat(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetReviseStat() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) AddArticle(c context.Context, req *article.ArgAddArticle) (resp *article.IDResp, err error) {
	if resp, err = p.articleRPC.AddArticle(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticle() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) UpdateArticle(c context.Context, req *article.ArgUpdateArticle) (err error) {
	if _, err = p.articleRPC.UpdateArticle(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticle() error(%+v), req(%+v)", err, req))
		return
	}
	return
}

func (p *Dao) DelArticle(c context.Context, req *article.IDReq) (err error) {
	if _, err = p.articleRPC.DelArticle(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticle() error(%+v), req(%+v)", err, req))
		return
	}
	return
}

func (p *Dao) GetArticleFiles(c context.Context, req *article.IDReq) (resp *article.ArticleFilesResp, err error) {
	if resp, err = p.articleRPC.GetArticleFiles(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleFiles() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) SaveArticleFiles(c context.Context, req *article.ArgSaveArticleFiles) (err error) {
	if _, err = p.articleRPC.SaveArticleFiles(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SaveArticleFiles() error(%+v), req(%+v)", err, req))
		return
	}
	return
}

func (p *Dao) GetArticleHistoriesPaged(c context.Context, req *article.ArgArticleHistoriesPaged) (resp *article.ArticleHistoryListResp, err error) {
	if resp, err = p.articleRPC.GetArticleHistoriesPaged(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoriesPaged() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) GetArticleHistory(c context.Context, req *article.IDReq) (resp *article.ArticleHistoryResp, err error) {
	if resp, err = p.articleRPC.GetArticleHistory(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistory() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) GetArticleRelations(c context.Context, req *article.IDReq) (resp *article.ArticleRelationsResp, err error) {
	if resp, err = p.articleRPC.GetArticleRelations(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleRelations() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) UpdateArticleRelation(c context.Context, req *article.ArgUpdateArticleRelation) (err error) {
	if _, err = p.articleRPC.UpdateArticleRelation(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticleRelation() error(%+v), req(%+v)", err, req))
		return
	}
	return
}

func (p *Dao) SetPrimary(c context.Context, req *article.ArgSetPrimaryArticleRelation) (err error) {
	if _, err = p.articleRPC.SetPrimary(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SetPrimary() error(%+v), req(%+v)", err, req))
		return
	}
	return
}

func (p *Dao) AddArticleRelation(c context.Context, req *article.ArgAddArticleRelation) (err error) {
	if _, err = p.articleRPC.AddArticleRelation(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleRelation() error(%+v), req(%+v)", err, req))
		return
	}
	return
}

func (p *Dao) DelArticleRelation(c context.Context, req *article.ArgDelArticleRelation) (err error) {
	if _, err = p.articleRPC.DelArticleRelation(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticleRelation() error(%+v), req(%+v)", err, req))
		return
	}
	return
}

func (p *Dao) GetArticleStat(c context.Context, req *article.IDReq) (resp *article.ArticleStat, err error) {
	if resp, err = p.articleRPC.GetArticleStat(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleStat() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) GetReviseInfo(c context.Context, req *article.IDReq) (resp *article.ReviseInfo, err error) {
	if resp, err = p.articleRPC.GetReviseInfo(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetReviseInfo() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) GetUserArticlesPaged(c context.Context, req *article.UserArticlesReq) (resp *article.UserArticlesResp, err error) {
	if resp, err = p.articleRPC.GetUserArticlesPaged(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserArticlesPaged() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}
