package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

// GetDiscussionFiles 获取讨论文件
func (p *Service) GetDiscussionFiles(c context.Context, req *api.IDReq) (resp *api.DiscussionFilesResp, err error) {
	var data []*model.DiscussionFile
	if data, err = p.getDiscussionFiles(c, p.d.DB(), req.Aid, req.ID); err != nil {
		return
	}

	resp = &api.DiscussionFilesResp{
		Items: make([]*api.DiscussionFile, len(data)),
	}

	for i, v := range data {
		resp.Items[i] = &api.DiscussionFile{
			ID:        gid.NewID(),
			FileName:  v.FileName,
			FileURL:   v.FileURL,
			PdfURL:    v.PdfURL,
			FileType:  v.FileType,
			Seq:       int32(v.Seq),
			CreatedAt: time.Now().Unix(),
		}
	}

	return
}

// getDiscussionFiles 获取讨论文件
func (p *Service) getDiscussionFiles(c context.Context, node sqalx.Node, aid, discussionID int64) (items []*model.DiscussionFile, err error) {
	var addCache = true
	if items, err = p.d.DiscussionFilesCache(c, discussionID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetDiscussionFilesByCond(c, node, map[string]interface{}{"discussion_id": discussionID}); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetDiscussionFilesCache(context.TODO(), discussionID, items)
		})
	}

	return
}

// bulkCreateFiles 批量添加附件（讨论新建时）
func (p *Service) bulkCreateFiles(c context.Context, node sqalx.Node, discussionID int64, files []*api.ArgDiscussionFile) (err error) {
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
			Seq:          int32(v.Seq),
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

// SaveDiscussionFiles 批量更新讨论附件
func (p *Service) SaveDiscussionFiles(c context.Context, arg *api.ArgSaveDiscussionFiles) (err error) {
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

	var discussion *model.Discussion
	if discussion, err = p.d.GetDiscussionByID(c, tx, arg.DiscussionID); err != nil {
		return
	} else if discussion == nil {
		return ecode.DiscussionNotExist
	}

	dbItems, err := p.d.GetDiscussionFilesByCond(c, tx, map[string]interface{}{"discussion_id": arg.DiscussionID})
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
				Seq:          int32(v.Seq),
				DiscussionID: arg.DiscussionID,
				FileType:     v.FileType,
				CreatedAt:    time.Now().Unix(),
				UpdatedAt:    time.Now().Unix(),
			}

			if err = p.d.AddDiscussionFile(c, tx, item); err != nil {
				return
			}
			continue
		}

		// Update
		dic[v.GetIDValue()] = true
		var file *model.DiscussionFile
		if file, err = p.d.GetDiscussionFileByID(c, tx, v.GetIDValue()); err != nil {
			return
		} else if file == nil {
			err = ecode.DiscussionFileNotExist
			return
		}

		file.FileName = v.FileName
		file.FileURL = v.FileURL
		file.Seq = int32(v.Seq)

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
