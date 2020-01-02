package service

import (
	"context"
	"time"

	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// initDiscussionStat 初始化讨论状态
func (p *Service) initDiscussionStat(c context.Context, discussionID int64) (err error) {
	if err = p.d.AddDiscussionStat(c, p.d.DB(), &model.DiscussionStat{
		DiscussionID: discussionID,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}); err != nil {
		return
	}

	return
}

// getDiscussion 获取讨论信息
func (p *Service) getDiscussion(c context.Context, node sqalx.Node, discussionID int64) (item *model.Discussion, err error) {
	var addCache = true
	if item, err = p.d.DiscussionCache(c, discussionID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), discussionID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetDiscussionCache(context.TODO(), item)
		})
	}
	return
}

// GetDiscussionStat 获取指定讨论状态
func (p *Service) getDiscussionStat(c context.Context, node sqalx.Node, discussionID int64) (item *model.DiscussionStat, err error) {
	if item, err = p.d.GetDiscussionStatByID(c, node, discussionID); err != nil {
		return
	} else {
		item = &model.DiscussionStat{
			DiscussionID: discussionID,
		}
	}
	return
}
