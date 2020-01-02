package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

// GetRecommendTopicsIDs 获取冷启动推荐话题
func (p *Service) GetRecommendTopicsIDs(c context.Context) (ids []int64, err error) {
	data := make([]*model.RecommendTopic, 0)
	if data, err = p.d.GetRecommendTopics(c, p.d.DB()); err != nil {
		return
	}

	ids = make([]int64, len(data))
	for i, v := range data {
		ids[i] = v.TopicID
	}

	return
}

// GetRecommendAuthTopicsIDs 获取冷启动推荐话题所包含的授权话题
func (p *Service) GetRecommendAuthTopicsIDs(c context.Context, topicIDs []int64) (ids []int64, err error) {
	if ids, err = p.d.GetAuthTopicIDs(c, p.d.DB(), topicIDs); err != nil {
		return
	}

	return
}

// GetRecommendMemberIDs 获取冷启动话题下属成员
func (p *Service) GetRecommendMemberIDs(c context.Context, topicIDs []int64) (ids []int64, err error) {
	if ids, err = p.d.GetSpecificTopicsMemberIDs(c, p.d.DB(), topicIDs); err != nil {
		return
	}
	return
}

// AddRecommendTopic 添加冷启动推荐话题
func (p *Service) AddRecommendTopic(c context.Context, arg *api.TopicReq) (err error) {
	if err = p.checkIsTopicAdmin(c, arg.Aid, arg.ID); err != nil {
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

	var t *model.Topic
	if t, err = p.getTopic(c, tx, arg.ID); err != nil {
		return
	}

	var item *model.RecommendTopic
	if item, err = p.d.GetRecommendTopicByCond(c, tx, map[string]interface{}{
		"topic_id": t.ID,
	}); err != nil {
		return
	} else if item != nil {
		err = ecode.RecommendTopicExist
		return
	}

	item = &model.RecommendTopic{
		ID:        gid.NewID(),
		TopicID:   t.ID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddRecommendTopic(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

// DelRecommendTopic 删除冷启动推荐话题
func (p *Service) DelRecommendTopic(c context.Context, arg *api.TopicReq) (err error) {
	if err = p.checkIsTopicAdmin(c, arg.Aid, arg.ID); err != nil {
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

	var t *model.Topic
	if t, err = p.getTopic(c, tx, arg.ID); err != nil {
		return
	}

	var item *model.RecommendTopic
	if item, err = p.d.GetRecommendTopicByCond(c, tx, map[string]interface{}{
		"topic_id": t.ID,
	}); err != nil {
		return
	} else if item == nil {
		return
	}

	if err = p.d.DelRecommendTopic(c, tx, item.ID); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}
