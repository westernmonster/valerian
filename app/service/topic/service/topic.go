package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Service) GetBelongsTopicIDs(c context.Context, aid int64) (ids []int64, err error) {

	return p.d.GetMemberBelongsTopicIDs(c, p.d.DB(), aid)
}

func (p *Service) GetTopicMemberIDs(c context.Context, aid int64) (ids []int64, err error) {
	return p.d.GetTopicMemberIDs(c, p.d.DB(), aid)
}

func (p *Service) GetUserTopicsPaged(c context.Context, aid int64, limit, offset int) (resp *api.UserTopicsResp, err error) {
	var items []*model.Topic
	if items, err = p.d.GetUserTopicsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	resp = &api.UserTopicsResp{
		Items: make([]*api.TopicInfo, len(items)),
	}

	for i, v := range items {
		var stat *model.TopicStat
		if stat, err = p.d.GetTopicStatByID(c, p.d.DB(), v.ID); err != nil {
			return
		}

		var acc *model.Account
		if acc, err = p.getAccount(c, p.d.DB(), v.CreatedBy); err != nil {
			return
		}

		item := &api.TopicInfo{
			ID:              v.ID,
			Name:            v.Name,
			Introduction:    v.Introduction,
			AllowDiscuss:    bool(v.AllowDiscuss),
			AllowChat:       bool(v.AllowChat),
			IsPrivate:       bool(v.IsPrivate),
			ViewPermission:  v.ViewPermission,
			EditPermission:  v.EditPermission,
			JoinPermission:  v.JoinPermission,
			CatalogViewType: v.CatalogViewType,
			Avatar:          v.Avatar,
			Bg:              v.Bg,
			TopicHome:       v.TopicHome,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
		}

		item.Stat = &api.TopicStat{
			MemberCount:     int32(stat.MemberCount),
			DiscussionCount: int32(stat.DiscussionCount),
			ArticleCount:    int32(stat.ArticleCount),
		}

		item.Creator = &api.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		resp.Items[i] = item
	}

	return
}

func (p *Service) GetAllTopics(c context.Context) (items []*model.Topic, err error) {
	return p.d.GetTopics(c, p.d.DB())
}

func (p *Service) GetTopic(c context.Context, topicID int64) (item *model.Topic, err error) {
	return p.getTopic(c, p.d.DB(), topicID)
}

func (p *Service) GetTopicInfo(c context.Context, topicID int64) (item *api.TopicInfo, err error) {
	return p.getTopicInfo(c, p.d.DB(), topicID)
}

func (p *Service) getTopicInfo(c context.Context, node sqalx.Node, topicID int64) (item *api.TopicInfo, err error) {
	var v *model.Topic
	if v, err = p.getTopic(c, node, topicID); err != nil {
		return
	}

	var stat *model.TopicStat
	if stat, err = p.d.GetTopicStatByID(c, node, topicID); err != nil {
		return
	}

	var acc *model.Account
	if acc, err = p.getAccount(c, node, v.CreatedBy); err != nil {
		return
	}

	item = &api.TopicInfo{
		ID:              v.ID,
		Name:            v.Name,
		Introduction:    v.Introduction,
		AllowDiscuss:    bool(v.AllowDiscuss),
		AllowChat:       bool(v.AllowChat),
		IsPrivate:       bool(v.IsPrivate),
		ViewPermission:  v.ViewPermission,
		EditPermission:  v.EditPermission,
		JoinPermission:  v.JoinPermission,
		CatalogViewType: v.CatalogViewType,
		Avatar:          v.Avatar,
		Bg:              v.Bg,
		TopicHome:       v.TopicHome,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}

	item.Stat = &api.TopicStat{
		MemberCount:     int32(stat.MemberCount),
		DiscussionCount: int32(stat.DiscussionCount),
		ArticleCount:    int32(stat.ArticleCount),
	}

	item.Creator = &api.Creator{
		ID:           acc.ID,
		UserName:     acc.UserName,
		Avatar:       acc.Avatar,
		Introduction: acc.Introduction,
	}

	return

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

// 创建
func (p *Service) CreateTopic(c context.Context, arg *api.ArgCreateTopic) (topicID int64, err error) {
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

	if topicID, err = p.createTopic(c, tx, arg.Aid, arg); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.onTopicAdded(context.Background(), topicID, arg.Aid, time.Now().Unix())
	})
	return
}

func (p *Service) UpdateTopic(c context.Context, arg *api.ArgUpdateTopic) (err error) {
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

	if err = p.checkTopic(c, tx, arg.ID); err != nil {
		return
	}

	if err = p.updateTopic(c, tx, arg.Aid, arg); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelAccountTopicSettingCache(context.TODO(), arg.Aid, arg.ID)
		p.d.DelTopicCache(context.TODO(), arg.ID)
		p.onTopicUpdated(context.Background(), arg.ID, arg.Aid, time.Now().Unix())
	})
	return
}

func (p *Service) GetTopicResp(c context.Context, aid int64, topicID int64, include string) (item *api.TopicResp, err error) {
	var t *model.Topic
	if t, err = p.getTopic(c, p.d.DB(), topicID); err != nil {
		return
	}
	item = &api.TopicResp{
		ID:              t.ID,
		Avatar:          t.Avatar,
		Bg:              t.Bg,
		Name:            t.Name,
		Introduction:    t.Introduction,
		CatalogViewType: t.CatalogViewType,
		TopicHome:       t.TopicHome,
		IsPrivate:       bool(t.IsPrivate),
		AllowChat:       bool(t.AllowChat),
		AllowDiscuss:    bool(t.AllowDiscuss),
		ViewPermission:  t.ViewPermission,
		EditPermission:  t.EditPermission,
		JoinPermission:  t.JoinPermission,
		CreatedAt:       t.CreatedAt,
	}

	var s *model.AccountTopicSetting
	if s, err = p.getAccountTopicSetting(c, p.d.DB(), aid, topicID); err != nil {
		return
	}

	item.Important = bool(s.Important)
	item.MuteNotification = bool(s.MuteNotification)

	var stat *model.TopicStat
	if stat, err = p.d.GetTopicStatByID(c, p.d.DB(), topicID); err != nil {
		return
	}
	item.Stat = &api.TopicStat{
		MemberCount:     stat.MemberCount,
		ArticleCount:    stat.ArticleCount,
		DiscussionCount: stat.DiscussionCount,
	}

	var acc *model.Account
	if acc, err = p.getAccount(c, p.d.DB(), t.CreatedBy); err != nil {
		return
	}

	item.Creator = &api.Creator{
		ID:           acc.ID,
		UserName:     acc.UserName,
		Avatar:       acc.Avatar,
		Introduction: acc.Introduction,
	}

	if item.HasCatalogTaxonomy, err = p.d.HasTaxonomy(c, p.d.DB(), topicID); err != nil {
		return
	}

	p.addCache(func() {
		p.onTopicViewed(context.Background(), topicID, aid, time.Now().Unix())
	})

	return
}

func (p *Service) DelTopic(c context.Context, topicID int64) (err error) {
	return
}
