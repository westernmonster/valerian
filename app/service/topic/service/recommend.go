package service

import (
	"context"
	"valerian/app/service/topic/model"
)

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

func (p *Service) GetRecommendAuthTopicsIDs(c context.Context, topicIDs []int64) (ids []int64, err error) {
	dic := make(map[int64]bool)

	for _, v := range topicIDs {
		authIDs := make([]int64, 0)
		if authIDs, err = p.d.GetAuthTopicIDs(c, p.d.DB(), v); err != nil {
			return
		}

		for _, t := range authIDs {
			dic[t] = true
		}
	}

	ids = make([]int64, 0)
	for k, _ := range dic {
		ids = append(ids, k)
	}

	return
}

func (p *Service) GetRecommendMemberIDs(c context.Context, topicIDs []int64) (ids []int64, err error) {
	if ids, err = p.d.GetSpecificTopicsMemberIDs(c, p.d.DB(), topicIDs); err != nil {
		return
	}
	return
}
