package api

import "valerian/app/service/feed/model"

func FromFeed(items []*model.Feed) *FeedResp {
	resp := &FeedResp{
		Items: make([]*FeedInfo, len(items)),
	}

	for i, v := range items {
		reply := &FeedInfo{
			ID:         v.ID,
			AccountID:  v.AccountID,
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
