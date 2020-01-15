package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// SaveAuthTopics  保存授权话题
func (p *Service) SaveAuthTopics(c context.Context, arg *api.ArgSaveAuthTopics) (err error) {

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

	// 检测是否系统管理员或者话题管理员
	if err = p.checkTopicManagePermission(c, tx, arg.Aid, arg.TopicID); err != nil {
		return
	}
	if err = p.bulkSaveAuthTopics(c, tx, arg.TopicID, arg.AuthTopics); err != nil {
		return
	}

	if err = p.d.SetTopicUpdatedAt(c, tx, arg.TopicID, time.Now().Unix()); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelTopicCache(context.Background(), arg.TopicID)
		p.d.DelAuthTopicsCache(context.Background(), arg.TopicID)
	})

	return
}

// GetAuthTopics 获取授权话题
func (p *Service) GetAuthTopics(c context.Context, topicID int64) (items []*api.AuthTopicInfo, err error) {
	return p.getAuthTopicsResp(c, p.d.DB(), topicID)
}

func (p *Service) GetUserCanEditTopicIDs(c context.Context, aid int64) (ids []int64, err error) {
	if ids, err = p.d.GetUserCanEditTopicIDs(c, p.d.DB(), aid); err != nil {
		return
	}

	return
}
