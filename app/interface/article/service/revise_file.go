package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/jinzhu/copier"
)

func (p *Service) GetReviseFiles(c context.Context, articleID int64) (items []*model.ReviseFileResp, err error) {
	return p.getReviseFiles(c, p.d.DB(), articleID)
}

func (p *Service) getReviseFiles(c context.Context, node sqalx.Node, articleID int64) (items []*model.ReviseFileResp, err error) {
	var addCache = true

	if items, err = p.d.ReviseFileCache(c, articleID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	var data []*model.ReviseFile
	if data, err = p.d.GetReviseFilesByCond(c, node, map[string]interface{}{"article_id": articleID}); err != nil {
		return
	}

	items = make([]*model.ReviseFileResp, 0)
	if err = copier.Copy(&items, &data); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetReviseFileCache(context.TODO(), articleID, items)
		})
	}

	return
}

func (p *Service) bulkCreateReviseFiles(c context.Context, node sqalx.Node, articleID int64, files []*model.AddReviseFile) (err error) {
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
			ReviseID:  articleID,
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
		p.d.DelReviseFileCache(context.TODO(), articleID)
	})

	return
}

func (p *Service) SaveReviseFiles(c context.Context, arg *model.ArgSaveReviseFiles) (err error) {
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

	var article *model.Revise
	if article, err = p.d.GetReviseByID(c, tx, arg.ReviseID); err != nil {
		return
	} else if article == nil {
		return ecode.ReviseNotExist
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
		dic[*v.ID] = true
		var file *model.ReviseFile
		if file, err = p.d.GetReviseFileByID(c, tx, *v.ID); err != nil {
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
