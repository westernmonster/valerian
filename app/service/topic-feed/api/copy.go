package api

import "valerian/app/service/topic-feed/model"

func FromTopicFeed(items []*model.TopicFeed) *TopicFeedResp {
	resp := &TopicFeedResp{
		Items: make([]*TopicFeedInfo, len(items)),
	}

	for i, v := range items {
		reply := &TopicFeedInfo{
			ID:         v.ID,
			TopicID:    v.TopicID,
			ActionType: v.ActionType,
			ActionTime: v.ActionTime,
			ActionText: v.ActionText,
			ActorID:    v.ActorID,
			ActorType:  v.ActorType,
			TargetID:   v.TargetID,
			TargetType: v.TargetType,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}
		resp.Items[i] = reply
	}

	return resp
}
