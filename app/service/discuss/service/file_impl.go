package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"
)

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
			p.d.SetDiscussionFilesCache(context.Background(), discussionID, items)
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
		p.d.DelDiscussionFilesCache(context.Background(), discussionID)
	})

	return
}
