package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/article/model"
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
		ID:           gid.NewID(),
		Title:        arg.Title,
		Content:      arg.Content,
		Cover:        arg.Cover,
		Introduction: arg.Introduction,
		Private:      types.BitBool(arg.Private),
		VersionName:  arg.VersionName,
		CreatedBy:    aid,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}

	if arg.ArticleSetID == nil {
		set := &model.ArticleSet{
			ID:        gid.NewID(),
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		item.ArticleSetID = set.ID

		if err = p.d.AddArticleSet(c, tx, set); err != nil {
			return
		}
	} else {
		if v, e := p.d.GetArticleVersionByName(c, tx, *arg.ArticleSetID, arg.VersionName); e != nil {
			err = e
			return
		} else if v != nil {
			err = ecode.ArticleVersionNameExist
			return
		}

		item.ArticleSetID = *arg.ArticleSetID
	}

	if err = p.d.AddArticle(c, tx, item); err != nil {
		return
	}

	if err = p.bulkCreateFiles(c, tx, item.ID, arg.Files); err != nil {
		return
	}

	if err = p.bulkCreateCatalogs(c, tx, item.ID, item.Title, arg.Relations); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	return
}
