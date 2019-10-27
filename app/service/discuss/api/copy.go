package api

import (
	"valerian/app/service/discuss/model"
	"valerian/library/xstr"
)

func FromDiscussion(v *model.Discussion, x *model.DiscussionStat, imgs []string) *DiscussionInfo {
	reply := &DiscussionInfo{
		ID:         v.ID,
		TopicID:    v.TopicID,
		CategoryID: v.CategoryID,
		// CreatedBy:   v.CreatedBy,
		Excerpt:   xstr.Excerpt(v.ContentText),
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		ImageUrls: imgs,
		Title:     v.Title,
	}

	reply.Stat = &DiscussionStat{
		DislikeCount: int32(x.DislikeCount),
		LikeCount:    int32(x.LikeCount),
		CommentCount: int32(x.CommentCount),
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
