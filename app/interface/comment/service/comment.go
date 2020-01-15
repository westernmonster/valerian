package service

import (
	"context"
	"net/url"
	"strconv"
	"valerian/app/interface/comment/model"
	comment "valerian/app/service/comment/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetCommentsPaged(c context.Context, resourceID int64, targetType string, limit, offset int) (resp *model.CommentListResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var data *comment.CommentListResp
	if data, err = p.d.GetCommentsPaged(c, &comment.CommentListReq{
		ResourceID: resourceID,
		TargetType: targetType,
		Limit:      int32(limit),
		Offset:     int32(offset),
		Aid:        aid,
	}); err != nil {
		return
	}

	resp = &model.CommentListResp{
		Paging: &model.Paging{},
		Items:  make([]*model.CommentItem, len(data.Items)),
	}

	resp.CommentsCount = data.CommentsCount
	resp.FeaturedCount = data.FeaturedCount

	for i, v := range data.Items {
		item := &model.CommentItem{
			ID:            v.ID,
			Content:       v.Content,
			Featured:      bool(v.Featured),
			IsDelete:      bool(v.Deleted),
			CreatedAt:     v.CreatedAt,
			ChildComments: make([]*model.ChildCommentItem, 0),
			Creator: &model.CommentCreator{
				ID:           v.Creator.ID,
				UserName:     v.Creator.UserName,
				Avatar:       v.Creator.Avatar,
				Introduction: v.Creator.Introduction,
			},
			LikeCount:          v.Stat.LikeCount,
			DislikeCount:       v.Stat.DislikeCount,
			ChildCommentsCount: v.Stat.ChildrenCount,
		}

		if v.ChildComments != nil {
			for _, x := range v.ChildComments {
				child := &model.ChildCommentItem{
					ID:           x.ID,
					Content:      x.Content,
					Featured:     x.Featured,
					IsDelete:     x.Deleted,
					CreatedAt:    x.CreatedAt,
					Like:         x.Like,
					Dislike:      x.Dislike,
					LikeCount:    x.Stat.LikeCount,
					DislikeCount: x.Stat.DislikeCount,
					Creator: &model.CommentCreator{
						ID:           x.Creator.ID,
						UserName:     x.Creator.UserName,
						Avatar:       x.Creator.Avatar,
						Introduction: x.Creator.Introduction,
					},
				}

				if x.ReplyTo != nil {
					child.ReplyTo = &model.CommentCreator{
						ID:           x.ReplyTo.ID,
						UserName:     x.ReplyTo.UserName,
						Avatar:       x.ReplyTo.Avatar,
						Introduction: x.ReplyTo.Introduction,
					}
				}

				item.ChildComments = append(item.ChildComments, child)
			}
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/comment/list/comments", url.Values{
		"resource_id": []string{strconv.FormatInt(resourceID, 10)},
		"target_type": []string{targetType},
		"limit":       []string{strconv.Itoa(limit)},
		"offset":      []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/comment/list/comments", url.Values{
		"resource_id": []string{strconv.FormatInt(resourceID, 10)},
		"target_type": []string{targetType},
		"limit":       []string{strconv.Itoa(limit)},
		"offset":      []string{strconv.Itoa(offset + limit)},
	}); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	return
}

func (p *Service) AddComment(c context.Context, arg *model.ArgAddComment) (id int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if id, err = p.d.AddComment(c, &comment.AddCommentReq{
		TargetType: arg.TargetType,
		TargetID:   arg.TargetID,
		Content:    arg.Content,
		Aid:        aid,
	}); err != nil {
		return
	}

	return
}

func (p *Service) GetComment(c context.Context, commentID int64) (resp *model.CommentResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var data *comment.CommentInfo
	if data, err = p.d.GetCommentInfo(c, &comment.IDReq{ID: commentID, Aid: aid}); err != nil {
		return
	}

	if data.TargetType == model.TargetTypeComment {
		if data, err = p.d.GetCommentInfo(c, &comment.IDReq{ID: data.ResourceID, Aid: aid}); err != nil {
			return
		}
	}

	resp = &model.CommentResp{
		ID:            data.ID,
		Content:       data.Content,
		Featured:      bool(data.Featured),
		IsDelete:      bool(data.Deleted),
		CreatedAt:     data.CreatedAt,
		ChildComments: make([]*model.ChildCommentItem, 0),
		Like:          data.Like,
		Dislike:       data.Dislike,
		Creator: &model.CommentCreator{
			ID:           data.Creator.ID,
			UserName:     data.Creator.UserName,
			Avatar:       data.Creator.Avatar,
			Introduction: data.Creator.Introduction,
		},
	}

	resp.LikeCount = data.Stat.LikeCount
	resp.DislikeCount = data.Stat.DislikeCount
	resp.ChildCommentsCount = data.Stat.ChildrenCount

	if data.Stat.ChildrenCount > 0 {
		var childData *comment.ChildrenCommentListResp
		if childData, err = p.d.GetAllChildrenComment(c, &comment.IDReq{Aid: aid, ID: data.ID}); err != nil {
			return
		}
		for _, x := range childData.Items {
			child := &model.ChildCommentItem{
				ID:           x.ID,
				Content:      x.Content,
				Featured:     x.Featured,
				IsDelete:     x.Deleted,
				CreatedAt:    x.CreatedAt,
				Like:         x.Like,
				Dislike:      x.Dislike,
				LikeCount:    x.Stat.LikeCount,
				DislikeCount: x.Stat.DislikeCount,
				Creator: &model.CommentCreator{
					ID:           x.Creator.ID,
					UserName:     x.Creator.UserName,
					Avatar:       x.Creator.Avatar,
					Introduction: x.Creator.Introduction,
				},
			}

			if x.ReplyTo != nil {
				child.ReplyTo = &model.CommentCreator{
					ID:           x.ReplyTo.ID,
					UserName:     x.ReplyTo.UserName,
					Avatar:       x.ReplyTo.Avatar,
					Introduction: x.ReplyTo.Introduction,
				}
			}

			resp.ChildComments = append(resp.ChildComments, child)
		}
	}

	return

}

func (p *Service) DelComment(c context.Context, commentID int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	// 因为被删除评论也会显示，所以不做其他处理
	return p.d.DelComment(c, &comment.DeleteReq{ID: commentID, Aid: aid})
}
