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

func (p *Dao) GetArticleDetail(c context.Context, req *article.IDReq) (resp *article.ArticleDetail, err error) {
	if resp, err = p.articleRPC.GetArticleDetail(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleDetail() error(%+v), req(%+v)", err, req))
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

func (p *Dao) GetReviseDetail(c context.Context, req *article.IDReq) (resp *article.ReviseDetail, err error) {
	if resp, err = p.articleRPC.GetReviseDetail(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetReviseDetail() error(%+v), req(%+v)", err, req))
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

func (p *Dao) AddRevise(c context.Context, req *article.ArgAddRevise) (resp *article.IDResp, err error) {
	if resp, err = p.articleRPC.AddRevise(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddRevise() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) UpdateRevise(c context.Context, req *article.ArgUpdateRevise) (err error) {
	if _, err = p.articleRPC.UpdateRevise(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateRevise() error(%+v), req(%+v)", err, req))
		return
	}
	return
}

func (p *Dao) DelRevise(c context.Context, req *article.IDReq) (err error) {
	if _, err = p.articleRPC.DelRevise(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelRevise() error(%+v), req(%+v)", err, req))
		return
	}
	return
}

func (p *Dao) GetArticleRevisesPaged(c context.Context, req *article.ArgArticleRevisesPaged) (resp *article.ReviseListResp, err error) {
	if resp, err = p.articleRPC.GetArticleRevisesPaged(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleRevisesPaged() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) GetReviseFiles(c context.Context, req *article.IDReq) (resp *article.ReviseFilesResp, err error) {
	if resp, err = p.articleRPC.GetReviseFiles(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetReviseFiles() error(%+v), req(%+v)", err, req))
		return
	}
	return resp, err
}

func (p *Dao) SaveReviseFiles(c context.Context, req *article.ArgSaveReviseFiles) (err error) {
	if _, err = p.articleRPC.SaveReviseFiles(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetReviseFiles() error(%+v), req(%+v)", err, req))
		return
	}
	return err
}

func (p *Dao) CanEdit(c context.Context, req *article.IDReq) (canEdit bool, err error) {
	var resp *article.BoolResp
	if resp, err = p.articleRPC.CanEdit(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CanEdit() error(%+v), req(%+v)", err, req))
		return
	}
	canEdit = resp.Result
	return
}

func (p *Dao) CanView(c context.Context, req *article.IDReq) (canView bool, err error) {
	var resp *article.BoolResp
	if resp, err = p.articleRPC.CanView(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CanView() error(%+v), req(%+v)", err, req))
		return
	}
	canView = resp.Result
	return
}

func (p *Dao) GetUserCanEditArticleIDs(c context.Context, arg *article.AidReq) (resp *article.IDsResp, err error) {
	if resp, err = p.articleRPC.GetUserCanEditArticleIDs(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserCanEditArticleIDs err(%+v) arg(%+v)", err, arg))
	}
	return
}
