package service

import (
	"context"
	"fmt"
	"time"

	account "valerian/app/service/account/api"
	discuss "valerian/app/service/discuss/api"
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

func (p *Service) GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error) {
	return p.d.GetAccountBaseInfo(c, aid)
}

func (p *Service) GetUserTopicsPaged(c context.Context, aid int64, limit, offset int) (items []*model.Topic, err error) {
	return p.d.GetUserTopicsPaged(c, p.d.DB(), aid, limit, offset)
}

func (p *Service) GetAllTopics(c context.Context) (items []*model.Topic, err error) {
	return p.d.GetTopics(c, p.d.DB())
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
	if item, err = p.getTopicResp(c, p.d.DB(), aid, topicID); err != nil {
		return
	}
	inc := includeParam(include)
	if inc["members"] {
		// if item.MemberCount, item.Members, err = p.getTopicMembers(c, p.d.DB(), topicID, 10); err != nil {
		// 	return
		// }
	}

	if inc["catalogs"] {
		if item.Catalogs, err = p.getCatalogHierarchyOfAll(c, p.d.DB(), topicID); err != nil {
			return
		}
	}

	if inc["discuss_categories"] {
		var resp *discuss.CategoriesResp
		if resp, err = p.d.GetDiscussionCategories(c, topicID); err != nil {
			return
		}

		item.DiscussCategories = make([]*api.DiscussCategoryInfo, len(resp.Items))

		for i, v := range resp.Items {
			item.DiscussCategories[i] = &api.DiscussCategoryInfo{
				ID:      v.ID,
				TopicID: v.TopicID,
				Name:    v.Name,
				Seq:     (v.Seq),
			}
		}
	}

	if inc["auth_topics"] {
		if item.AuthTopics, err = p.getAuthTopicsResp(c, p.d.DB(), topicID); err != nil {
			return
		}
	}

	if inc["meta"] {
		if item.TopicMeta, err = p.GetTopicMeta(c, aid, topicID); err != nil {
			return
		}
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
