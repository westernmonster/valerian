package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/jinzhu/copier"
)

func (p *Service) GetArticleFiles(c context.Context, articleID int64) (items []*api.ArticleFileResp, err error) {
	return p.getArticleFiles(c, p.d.DB(), articleID)
}

func (p *Service) getArticleFiles(c context.Context, node sqalx.Node, articleID int64) (items []*api.ArticleFileResp, err error) {
	var addCache = true

	if items, err = p.d.ArticleFileCache(c, articleID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	var data []*model.ArticleFile
	if data, err = p.d.GetArticleFilesByCond(c, node, map[string]interface{}{"article_id": articleID}); err != nil {
		return
	}

	items = make([]*api.ArticleFileResp, 0)
	if err = copier.Copy(&items, &data); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleFileCache(context.TODO(), articleID, items)
		})
	}

	return
}

func (p *Service) bulkCreateFiles(c context.Context, node sqalx.Node, articleID int64, files []*api.ArgArticleFile) (err error) {
	var tx sqalx.Node
	if tx, err = node.Beginx(c); err != nil {
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

	for _, v := range files {
		item := &model.ArticleFile{
			ID:        gid.NewID(),
			FileName:  v.FileName,
			FileURL:   v.FileURL,
			Seq:       v.Seq,
			ArticleID: articleID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddArticleFile(c, tx, item); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelArticleFileCache(context.TODO(), articleID)
	})

	return
}

func (p *Service) SaveArticleFiles(c context.Context, arg *api.ArgSaveArticleFiles) (err error) {
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

	if canEdit, e := p.checkEditPermission(c, tx, arg.ArticleID, arg.Aid); e != nil {
		return e
	} else if !canEdit {
		err = ecode.NeedArticleEditPermission
		return
	}

	var article *model.Article
	if article, err = p.d.GetArticleByID(c, tx, arg.ArticleID); err != nil {
		return
	} else if article == nil {
		return ecode.ArticleNotExist
	}

	dbItems, err := p.d.GetArticleFilesByCond(c, tx, map[string]interface{}{"article_id": arg.ArticleID})
	if err != nil {
		return
	}

	dic := make(map[int64]bool)
	for _, v := range arg.Items {
		if v.ID == nil {
			// Add
			item := &model.ArticleFile{
				ID:        gid.NewID(),
				FileName:  v.FileName,
				FileURL:   v.FileURL,
				Seq:       v.Seq,
				ArticleID: arg.ArticleID,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}

			if err = p.d.AddArticleFile(c, tx, item); err != nil {
				return
			}
			continue
		}

		// Update
		dic[v.GetIDValue()] = true
		var file *model.ArticleFile
		if file, err = p.d.GetArticleFileByID(c, tx, v.GetIDValue()); err != nil {
			return
		} else if file == nil {
			err = ecode.ArticleFileNotExist
			return
		}

		file.FileName = v.FileName
		file.FileURL = v.FileURL
		file.Seq = v.Seq

		if err = p.d.UpdateArticleFile(c, tx, file); err != nil {
			return
		}
	}

	// Delete
	for _, v := range dbItems {
		if dic[v.ID] {
			continue
		}

		if err = p.d.DelArticleFile(c, tx, v.ID); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelArticleFileCache(context.TODO(), arg.ArticleID)
	})
	return
}
