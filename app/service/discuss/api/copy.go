package api

import "valerian/app/service/discuss/model"

func FromDiscussion(v *model.Discussion) *DiscussionInfo {
	reply := &DiscussionInfo{
		ID:          v.ID,
		TopicID:     v.TopicID,
		CategoryID:  v.CategoryID,
		CreatedBy:   v.CreatedBy,
		Content:     v.Content,
		ContentText: v.ContentText,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}

	if v.Title != nil {
		reply.Title = &DiscussionInfo_TitleValue{*v.Title}
	}

	return reply
}
