package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

// SaveDiscussCategories 批量保存讨论分类
// 需要检查是否管理员全新啊
// 需要检查话题是否存在，当话题删除时候，则不能保存分类
func (p *Service) SaveDiscussCategories(c context.Context, arg *api.ArgSaveDiscussCategories) (err error) {
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

	// 检测话题
	if err = p.checkTopicExist(c, tx, arg.TopicID); err != nil {
		return
	}

	// 是否话题管理员
	if err = p.checkIsTopicManager(c, tx, arg.Aid, arg.TopicID); err != nil {
		return
	}

	var dic map[int64]bool
	if dic, err = p.loadDiscussCategoriesMap(c, tx, arg.TopicID); err != nil {
		return
	}

	for _, v := range arg.Items {
		// 获取ID为空，则新增
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

		// 如果ID不为空，则更新
		var dItem *model.DiscussCategory
		if dItem, err = p.d.GetDiscussCategoryByID(c, tx, v.GetIDValue()); err != nil {
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

		dic[v.GetIDValue()] = true
	}

	// 没有处理的视为删除
	for k, used := range dic {
		if used {
			continue
		}
		if has, e := p.d.HasDiscussionInCategory(c, tx, k); err != nil {
			err = e
			return
		} else if has {
			// 如果分类下有讨论，则不允许删除
			err = ecode.HasDiscussionInCategory
			return
		} else {
			if err = p.d.DelDiscussCategory(c, tx, k); err != nil {
				return
			}
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelDiscussionCategoriesCache(context.TODO(), arg.TopicID)
	})

	return
}

// GetDiscussCategories 获取话题所有讨论分类
func (p *Service) GetDiscussCategories(c context.Context, topicID int64) (items []*model.DiscussCategory, err error) {
	return p.getDiscussCategories(c, p.d.DB(), topicID)
}

// GetDiscussCategories 获取指定讨论分类
func (p *Service) GetDiscussCategory(c context.Context, categoryID int64) (item *model.DiscussCategory, err error) {
	if item, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), categoryID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussCategoryNotExist
		return
	}
	return
}
