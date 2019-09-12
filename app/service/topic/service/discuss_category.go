package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) getDiscussCategories(c context.Context, node sqalx.Node, topicID int64) (items []*model.DiscussCategoryResp, err error) {
	var dbItems []*model.DiscussCategory
	if dbItems, err = p.d.GetDiscussCategoriesByCond(c, p.d.DB(), map[string]interface{}{"topic_id": topicID}); err != nil {
		return
	}

	items = make([]*model.DiscussCategoryResp, 0)
	for _, v := range dbItems {
		items = append(items, &model.DiscussCategoryResp{
			ID:      v.ID,
			TopicID: v.TopicID,
			Name:    v.Name,
			Seq:     v.Seq,
		})
	}

	return
}

func (p *Service) GetDiscussCategories(c context.Context, topicID int64) (items []*model.DiscussCategoryResp, err error) {
	return p.getDiscussCategories(c, p.d.DB(), topicID)
}

func (p *Service) loadDiscussCategoriesMap(c context.Context, node sqalx.Node, topicID int64) (dic map[int64]bool, err error) {
	dic = make(map[int64]bool)
	var dbItems []*model.DiscussCategory
	if dbItems, err = p.d.GetDiscussCategoriesByCond(c, node, map[string]interface{}{"topic_id": topicID}); err != nil {
		return
	}

	for _, v := range dbItems {
		dic[v.ID] = false
	}

	return
}

func (p *Service) SaveDiscussCategories(c context.Context, arg *model.ArgSaveDiscussCategories) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
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

	// check topic
	if err = p.checkTopic(c, tx, arg.TopicID); err != nil {
		return
	}

	// check admin role
	if err = p.checkTopicMemberAdmin(c, tx, arg.TopicID, aid); err != nil {
		return
	}

	var dic map[int64]bool
	if dic, err = p.loadDiscussCategoriesMap(c, tx, arg.TopicID); err != nil {
		return
	}

	for _, v := range arg.Items {
		if v.ID == nil {
			item := &model.DiscussCategory{
				ID:        gid.NewID(),
				TopicID:   arg.TopicID,
				Seq:       v.Seq,
				Name:      v.Name,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}
			if err = p.d.AddDiscussCategory(c, tx, item); err != nil {
				return
			}

			continue
		}

		var dItem *model.DiscussCategory
		if dItem, err = p.d.GetDiscussCategoryByID(c, tx, *v.ID); err != nil {
			return
		} else if dItem == nil {
			err = ecode.DiscussCategoryNotExist
			return
		}

		dItem.Name = v.Name
		dItem.Seq = v.Seq

		if err = p.d.UpdateDiscussCategory(c, tx, dItem); err != nil {
			return
		}

		dic[*v.ID] = true
	}

	for k, used := range dic {
		if used {
			continue
		}
		if err = p.d.DelDiscussCategory(c, tx, k); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}
