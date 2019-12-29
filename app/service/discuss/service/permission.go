package service

import (
	"context"

	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// checkTopicExist 获取话题信息
func (p *Service) checkTopicExist(c context.Context, node sqalx.Node, topicID int64) (err error) {
	var exist bool
	if exist, err = p.d.IsTopicExist(c, node, topicID); err != nil {
		return
	} else if !exist {
		err = ecode.TopicNotExist
		return
	}
	return
}

// checkIsTopicMember 检测是否话题成员
func (p *Service) checkIsTopicMember(c context.Context, node sqalx.Node, aid, topicID int64) (err error) {
	var isMember bool
	if isMember, _, err = p.d.IsTopicMember(c, node, aid, topicID); err != nil {
		return
	} else if !isMember {
		err = ecode.NotTopicMember
		return
	}

	return
}

// checkIsTopicMember 检测是否话题管理员
func (p *Service) checkIsTopicManager(c context.Context, node sqalx.Node, aid, topicID int64) (err error) {
	var isMember bool
	var role string
	if isMember, role, err = p.d.IsTopicMember(c, node, aid, topicID); err != nil {
		return
	} else if !isMember {
		err = ecode.NotTopicMember
		return
	}

	if role == model.MemberRoleUser {
		err = ecode.NotTopicAdmin
		return
	}

	return
}
