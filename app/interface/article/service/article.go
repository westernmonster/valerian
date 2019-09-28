package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/article/model"
	account "valerian/app/service/account/api"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
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

	if err = p.d.AddArticle(c, tx, item); err != nil {
		return
	}

	if err = p.bulkCreateFiles(c, tx, item.ID, arg.Files); err != nil {
		return
	}

	if err = p.bulkCreateArticleRelations(c, tx, item.ID, item.Title, arg.Relations); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	id = item.ID

	return
}

func (p *Service) UpdateArticle(c context.Context, arg *model.ArgUpdateArticle) (err error) {
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

	if arg.Title != nil {
		item.Title = *arg.Title
	}

	item.Content = arg.Content

	if arg.DisableRevise != nil {
		item.DisableRevise = types.BitBool(*arg.DisableRevise)
	}

	if arg.DisableComment != nil {
		item.DisableComment = types.BitBool(*arg.DisableComment)
	}

	if err = p.d.UpdateArticle(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelArticleCache(context.TODO(), arg.ID)
	})

	return
}

// func (p *Service) DelArticle(c context.Context, id int64) (err error) {
// 	return
// }

func (p *Service) GetArticle(c context.Context, id int64, include string) (item *model.ArticleResp, err error) {
	inc := includeParam(include)
	if item, err = p.getArticle(c, p.d.DB(), id); err != nil {
		return
	}

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, item.CreatedBy); err != nil {
		return
	}

	item.Creator = &model.Creator{
		ID:       account.ID,
		UserName: account.UserName,
		Avatar:   account.Avatar,
	}
	intro := account.GetIntroductionValue()
	item.Creator.Introduction = &intro

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
		if item.ArticleMeta, err = p.getArticleMeta(c, p.d.DB(), id); err != nil {
			return
		}
	}

	return
}

func (p *Service) getArticle(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleResp, err error) {
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

	item = &model.ArticleResp{
		ID:        a.ID,
		Title:     a.Title,
		Content:   a.Content,
		CreatedBy: a.CreatedBy,
		Files:     make([]*model.ArticleFileResp, 0),
		Relations: make([]*model.ArticleRelationResp, 0),
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleCache(context.TODO(), item)
		})
	}
	return
}

func (p *Service) FavArticle(c context.Context, articleID int64) (isFav bool, err error) {
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

	var attr *model.AccountArticleAttr
	if attr, err = p.d.GetAccountArticleAttr(c, tx, aid, articleID); err != nil {
		return
	} else if attr == nil {
		attr := &model.AccountArticleAttr{
			ID:        gid.NewID(),
			AccountID: aid,
			ArticleID: articleID,
			Read:      false,
			Like:      false,
			Fav:       true,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		isFav = true
		if err = p.d.AddAccountArticleAttr(c, tx, attr); err != nil {
			return
		}

	} else {
		attr.Fav = !attr.Fav
		attr.UpdatedAt = time.Now().Unix()
		isFav = bool(attr.Fav)
		if err = p.d.UpdateAccountArticleAttr(c, tx, attr); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	return
}

func (p *Service) LikeArticle(c context.Context, articleID int64) (isLiked bool, err error) {
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

	var attr *model.AccountArticleAttr
	if attr, err = p.d.GetAccountArticleAttr(c, tx, aid, articleID); err != nil {
		return
	} else if attr == nil {
		attr := &model.AccountArticleAttr{
			ID:        gid.NewID(),
			AccountID: aid,
			ArticleID: articleID,
			Read:      false,
			Like:      true,
			Fav:       false,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		isLiked = true
		if err = p.d.AddAccountArticleAttr(c, tx, attr); err != nil {
			return
		}
	} else {
		attr.Like = !attr.Like
		attr.UpdatedAt = time.Now().Unix()
		isLiked = bool(attr.Like)
		if err = p.d.UpdateAccountArticleAttr(c, tx, attr); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	return
}

func (p *Service) getArticleMeta(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleMeta, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	item = new(model.ArticleMeta)

	var attr *model.AccountArticleAttr
	if attr, err = p.d.GetAccountArticleAttr(c, node, aid, articleID); err != nil {
		return
	} else if attr == nil {
		item.Fav = false
		item.Read = false
		item.Like = false
	} else {
		item.Fav = bool(attr.Fav)
		item.Read = bool(attr.Read)
		item.Like = bool(attr.Like)
	}

	return
}
