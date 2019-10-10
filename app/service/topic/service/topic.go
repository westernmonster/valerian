package service

import (
	"context"
	"time"
	account "valerian/app/service/account/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
)

func (p *Service) GetBelongsTopicIDs(c context.Context, aid int64) (ids []int64, err error) {
	return p.d.GetMemberBelongsTopicIDs(c, p.d.DB(), aid)
}

func (p *Service) GetTopicMemberIDs(c context.Context, aid int64) (ids []int64, err error) {
	return p.d.GetTopicMemberIDs(c, p.d.DB(), aid)
}

func (p *Service) GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error) {
	return p.d.GetAccountBaseInfo(c, aid)
}

func (p *Service) GetUserTopicsPaged(c context.Context, aid int64, limit, offset int) (items []*model.Topic, err error) {
	return p.d.GetUserTopicsPaged(c, p.d.DB(), aid, limit, offset)
}

func (p *Service) GetTopic(c context.Context, topicID int64) (item *model.Topic, err error) {
	return p.getTopic(c, p.d.DB(), topicID)
}

func (p *Service) GetTopicStat(c context.Context, topicID int64) (stat *model.TopicStat, err error) {
	if stat, err = p.d.GetTopicStatByID(c, p.d.DB(), topicID); err != nil {
		return
	} else if stat == nil {
		stat = &model.TopicStat{
			TopicID: topicID,
		}
	}
	return
}

func (p *Service) getTopic(c context.Context, node sqalx.Node, topicID int64) (item *model.Topic, err error) {
	var addCache = true
	if item, err = p.d.TopicCache(c, topicID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if item == nil {
		return nil, ecode.TopicNotExist
	}

	if addCache {
		p.addCache(func() {
			p.d.SetTopicCache(context.TODO(), item)
		})
	}

	return
}

func (p *Service) GetTopicManagerRole(c context.Context, topicID, aid int64) (isMember bool, role string, err error) {
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, p.d.DB(), map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if member == nil {
		isMember = false
	} else {
		isMember = true
		role = member.Role
	}

	return
}

func (p *Service) getAccountTopicSetting(c context.Context, node sqalx.Node, aid, topicID int64) (item *model.AccountTopicSetting, err error) {
	var addCache = true
	if item, err = p.d.AccountTopicSettingCache(c, aid, topicID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetAccountTopicSettingByCond(c, node, map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if item == nil {
		item = &model.AccountTopicSetting{
			ID:               gid.NewID(),
			AccountID:        aid,
			TopicID:          topicID,
			Important:        false,
			MuteNotification: false,
			Fav:              false,
			CreatedAt:        time.Now().Unix(),
			UpdatedAt:        time.Now().Unix(),
		}
	}

	if addCache {
		p.addCache(func() {
			p.d.SetAccountTopicSettingCache(context.TODO(), item)
		})
	}

	return
}
