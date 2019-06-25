package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) bulkCreateFiles(c context.Context, node sqalx.Node, articleID int64, files []*model.AddArticleFile) (err error) {
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

	return
}

func (p *Service) SaveArticleFiles(c context.Context, arg *model.ArgSaveArticleFiles) (err error) {
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

	// TODO: check edit permission

	var article *model.Article
	if article, err = p.d.GetArticleByID(c, tx, arg.ArticleID); err != nil {
		return
	} else if article == nil {
		return ecode.ArticleNotExist
	}

	dbItems, err := p.d.GetArticleFiles(c, tx, arg.ArticleID)
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
		dic[*v.ID] = true
		var file *model.ArticleFile
		if file, err = p.d.GetArticleFileByID(c, tx, *v.ID); err != nil {
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
	return
}
