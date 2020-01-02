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
	// 检测讨论
	if _, err = p.getDiscussion(c, p.d.DB(), req.ID); err != nil {
		return
	}

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

	// 检测讨论和所属话题
	var discuss *model.Discussion
	if discuss, err = p.getDiscussion(c, tx, arg.DiscussionID); err != nil {
		return
	}
	if err = p.checkTopicExist(c, tx, discuss.TopicID); err != nil {
		return
	}

	// 检测讨论的编辑权限
	if err = p.checkEditPermission(c, tx, arg.Aid, discuss.ID); err != nil {
		return
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
