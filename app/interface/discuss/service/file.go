package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/jinzhu/copier"
)

func (p *Service) GetDiscussionFiles(c context.Context, discussionID int64) (items []*model.DiscussionFileResp, err error) {
	return p.getDiscussionFiles(c, p.d.DB(), discussionID)
}

func (p *Service) getDiscussionFiles(c context.Context, node sqalx.Node, discussionID int64) (items []*model.DiscussionFileResp, err error) {
	var addCache = true

	if items, err = p.d.DiscussionFilesCache(c, discussionID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	var data []*model.DiscussionFile
	if data, err = p.d.GetDiscussionFilesByCond(c, node, map[string]interface{}{"article_id": discussionID}); err != nil {
		return
	}

	items = make([]*model.DiscussionFileResp, 0)
	if err = copier.Copy(&items, &data); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetDiscussionFilesCache(context.TODO(), discussionID, items)
		})
	}

	return
}

func (p *Service) bulkCreateFiles(c context.Context, node sqalx.Node, discussionID int64, files []*model.AddDiscussionFile) (err error) {
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
		item := &model.DiscussionFile{
			ID:           gid.NewID(),
			FileName:     v.FileName,
			FileURL:      v.FileURL,
			Seq:          v.Seq,
			DiscussionID: discussionID,
			CreatedAt:    time.Now().Unix(),
			UpdatedAt:    time.Now().Unix(),
		}

		if err = p.d.AddDiscussionFile(c, tx, item); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelDiscussionFilesCache(context.TODO(), discussionID)
	})

	return
}

func (p *Service) SaveDiscussionFiles(c context.Context, arg *model.ArgSaveDiscussionFiles) (err error) {
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

	var article *model.Discussion
	if article, err = p.d.GetDiscussionByID(c, tx, arg.DiscussionID); err != nil {
		return
	} else if article == nil {
		return ecode.DiscussionNotExist
	}

	dbItems, err := p.d.GetDiscussionFilesByCond(c, tx, map[string]interface{}{"article_id": arg.DiscussionID})
	if err != nil {
		return
	}

	dic := make(map[int64]bool)
	for _, v := range arg.Items {
		if v.ID == nil {
			// Add
			item := &model.DiscussionFile{
				ID:           gid.NewID(),
				FileName:     v.FileName,
				FileURL:      v.FileURL,
				Seq:          v.Seq,
				DiscussionID: arg.DiscussionID,
				CreatedAt:    time.Now().Unix(),
				UpdatedAt:    time.Now().Unix(),
			}

			if err = p.d.AddDiscussionFile(c, tx, item); err != nil {
				return
			}
			continue
		}

		// Update
		dic[*v.ID] = true
		var file *model.DiscussionFile
		if file, err = p.d.GetDiscussionFileByID(c, tx, *v.ID); err != nil {
			return
		} else if file == nil {
			err = ecode.DiscussionFileNotExist
			return
		}

		file.FileName = v.FileName
		file.FileURL = v.FileURL
		file.Seq = v.Seq

		if err = p.d.UpdateDiscussionFile(c, tx, file); err != nil {
			return
		}
	}

	// Delete
	for _, v := range dbItems {
		if dic[v.ID] {
			continue
		}

		if err = p.d.DelDiscussionFile(c, tx, v.ID); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelDiscussionFilesCache(context.TODO(), arg.DiscussionID)
	})
	return
}
