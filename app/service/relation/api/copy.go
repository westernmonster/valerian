package api

import "valerian/app/service/relation/model"

func FromFollowingResp(items []*model.FollowingResp) *FollowingResp {
	resp := &FollowingResp{
		Items: make([]*RelationInfo, len(items)),
	}

	for i, v := range items {
		reply := &RelationInfo{
			AccountID: v.AccountID,
			Attribute: v.Attribute,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resp.Items[i] = reply
	}

	return resp
}

func FromFansResp(items []*model.FansResp) *FansResp {
	resp := &FansResp{
		Items: make([]*RelationInfo, len(items)),
	}

	for i, v := range items {
		reply := &RelationInfo{
			AccountID: v.AccountID,
			Attribute: v.Attribute,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resp.Items[i] = reply
	}

	return resp
}
