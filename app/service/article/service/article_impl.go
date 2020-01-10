package service

import (
	"context"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// getAccount 获取用户信息
func (p *Service) getAccount(c context.Context, node sqalx.Node, aid int64) (info *model.Account, err error) {
	if info, err = p.d.GetAccountByID(c, node, aid); err != nil {
		return
	} else if info == nil {
		err = ecode.UserNotExist
		return
	}

	return
}

// getTopic 获取话题信息
func (p *Service) getTopic(c context.Context, node sqalx.Node, tid int64) (info *model.Topic, err error) {
	if info, err = p.d.GetTopicByID(c, node, tid); err != nil {
		return
	} else if info == nil {
		err = ecode.TopicNotExist
		return
	}

	return
}
