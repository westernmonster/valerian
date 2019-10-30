package service

import (
	"context"
	"fmt"

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
	if err = p.bulkSaveAuthTopics(c, tx, arg.TopicID, arg.AuthTopics); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelAuthTopicsCache(context.TODO(), arg.TopicID)
	})

	return
}

// GetAuthTopics 获取授权话题
func (p *Service) GetAuthTopics(c context.Context, topicID int64) (items []*api.AuthTopicInfo, err error) {
	return p.getAuthTopicsResp(c, p.d.DB(), topicID)
}
