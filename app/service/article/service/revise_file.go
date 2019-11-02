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

func (p *Service) GetReviseFiles(c context.Context, req *api.IDReq) (items []*api.ReviseFileResp, err error) {
	return p.getReviseFiles(c, p.d.DB(), req.ID)
}

func (p *Service) getReviseFiles(c context.Context, node sqalx.Node, reviseID int64) (items []*api.ReviseFileResp, err error) {
	var addCache = true

	if items, err = p.d.ReviseFileCache(c, reviseID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	var data []*model.ReviseFile
	if data, err = p.d.GetReviseFilesByCond(c, node, map[string]interface{}{"revise_id": reviseID}); err != nil {
		return
	}

	items = make([]*api.ReviseFileResp, 0)
	if err = copier.Copy(&items, &data); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetReviseFileCache(context.TODO(), reviseID, items)
		})
	}

	return
}

func (p *Service) bulkCreateReviseFiles(c context.Context, node sqalx.Node, reviseID int64, files []*api.AddReviseFile) (err error) {
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
		item := &model.ReviseFile{
			ID:        gid.NewID(),
			FileName:  v.FileName,
			FileURL:   v.FileURL,
			Seq:       v.Seq,
			ReviseID:  reviseID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddReviseFile(c, tx, item); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelReviseFileCache(context.TODO(), reviseID)
	})

	return
}

func (p *Service) SaveReviseFiles(c context.Context, arg *api.ArgSaveReviseFiles) (err error) {
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

	var article *model.Revise
	if article, err = p.d.GetReviseByID(c, tx, arg.ReviseID); err != nil {
		return
	} else if article == nil {
		return ecode.ReviseNotExist
	}

	if canEdit, e := p.checkEditPermission(c, tx, article.ArticleID, arg.Aid); e != nil {
		err = e
		return
	} else if !canEdit {
		err = ecode.NeedArticleEditPermission
		return
	}

	dbItems, err := p.d.GetReviseFilesByCond(c, tx, map[string]interface{}{"article_id": arg.ReviseID})
	if err != nil {
		return
	}

	dic := make(map[int64]bool)
	for _, v := range arg.Items {
		if v.ID == nil {
			// Add
			item := &model.ReviseFile{
				ID:        gid.NewID(),
				FileName:  v.FileName,
				FileURL:   v.FileURL,
				Seq:       v.Seq,
				ReviseID:  arg.ReviseID,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}

			if err = p.d.AddReviseFile(c, tx, item); err != nil {
				return
			}
			continue
		}

		// Update
		dic[v.GetIDValue()] = true
		var file *model.ReviseFile
		if file, err = p.d.GetReviseFileByID(c, tx, v.GetIDValue()); err != nil {
			return
		} else if file == nil {
			err = ecode.ReviseFileNotExist
			return
		}

		file.FileName = v.FileName
		file.FileURL = v.FileURL
		file.Seq = v.Seq

		if err = p.d.UpdateReviseFile(c, tx, file); err != nil {
			return
		}
	}

	// Delete
	for _, v := range dbItems {
		if dic[v.ID] {
			continue
		}

		if err = p.d.DelReviseFile(c, tx, v.ID); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelReviseFileCache(context.TODO(), arg.ReviseID)
	})
	return
}
