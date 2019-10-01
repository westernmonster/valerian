package api

import "valerian/app/service/discuss/model"

func FromDiscussion(v *model.Discussion, x *model.DiscussionStat, imgs []string) *DiscussionInfo {
	reply := &DiscussionInfo{
		ID:          v.ID,
		TopicID:     v.TopicID,
		CategoryID:  v.CategoryID,
		CreatedBy:   v.CreatedBy,
		Content:     v.Content,
		ContentText: v.ContentText,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Deleted:     bool(v.Deleted),
		ImageUrls:   imgs,
	}

	reply.Stat = &DiscussionStat{
		LikeCount:    int32(x.LikeCount),
		CommentCount: int32(x.CommentCount),
	}

	if v.Title != nil {
		reply.Title = &DiscussionInfo_TitleValue{*v.Title}
	}

	return reply
}

func FromCategory(v *model.DiscussCategory) *CategoryInfo {
	reply := &CategoryInfo{
		ID:      v.ID,
		TopicID: v.TopicID,
		Name:    v.Name,
		Seq:     int32(v.Seq),
	}

	return reply
}
