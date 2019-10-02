package api

import "valerian/app/service/feed/model"

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

func FromAccountFeed(items []*model.AccountFeed) *AccountFeedResp {
	resp := &AccountFeedResp{
		Items: make([]*AccountFeedInfo, len(items)),
	}

	for i, v := range items {
		reply := &AccountFeedInfo{
			ID:         v.ID,
			AccountID:  v.AccountID,
			ActionType: v.ActionType,
			ActionTime: v.ActionTime,
			ActionText: v.ActionText,
			TargetID:   v.TargetID,
			TargetType: v.TargetType,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}
		resp.Items[i] = reply
	}

	return resp
}
