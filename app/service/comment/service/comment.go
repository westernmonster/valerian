package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/comment/api"
	"valerian/app/service/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) GetAllChildrenComments(c context.Context, req *api.IDReq) (resp *api.ChildrenCommentListResp, err error) {
	var data []*model.Comment
	if data, err = p.d.GetAllChildrenComments(c, p.d.DB(), req.ID); err != nil {
		return
	}

	resp = &api.ChildrenCommentListResp{
		Items: make([]*api.ChildCommentItem, len(data)),
	}

	for i, v := range data {
		item := &api.ChildCommentItem{
			ID:         v.ID,
			Content:    v.Content,
			Featured:   bool(v.Featured),
			IsDelete:   bool(v.Deleted),
			OwnerID:    v.OwnerID,
			ResourceID: v.ResourceID,
			TargetType: v.TargetType,
			CreatedAt:  v.CreatedAt,
			OwnerType:  v.OwnerType,
		}

		var stat *model.CommentStat
		if stat, err = p.d.GetCommentStatByID(c, p.d.DB(), v.ID); err != nil {
			return
		}

		item.Stat = &api.CommentStat{}
		item.Stat.LikeCount = stat.LikeCount
		item.Stat.DislikeCount = stat.DislikeCount
		item.Stat.ChildrenCount = stat.ChildrenCount

		var acc *model.Account
		if acc, err = p.getAccount(c, p.d.DB(), v.CreatedBy); err != nil {
			return
		}
		item.Creator = &api.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		if v.ReplyTo != 0 {
			var acc *model.Account
			if acc, err = p.getAccount(c, p.d.DB(), v.ReplyTo); err != nil {
				return
			}
			item.ReplyTo = &api.Creator{
				ID:           acc.ID,
				UserName:     acc.UserName,
				Avatar:       acc.Avatar,
				Introduction: acc.Introduction,
			}
		}

		if item.Like, err = p.isLike(c, p.d.DB(), req.Aid, v.ID, model.TargetTypeComment); err != nil {
			return
		}

		if item.Dislike, err = p.isDislike(c, p.d.DB(), req.Aid, v.ID, model.TargetTypeComment); err != nil {
			return
		}

		resp.Items[i] = item
	}

	return
}

// GetComment 获取评论信息
func (p *Service) GetComment(c context.Context, commentID int64) (item *model.Comment, err error) {
	return p.getComment(c, p.d.DB(), commentID)
}

func (p *Service) GetCommentsPaged(c context.Context, req *api.CommentListReq) (resp *api.CommentListResp, err error) {
	var data []*model.Comment
	if data, err = p.d.GetCommentsPaged(c, p.d.DB(), req.ResourceID, req.TargetType, int(req.Limit), int(req.Offset)); err != nil {
		return
	}

	resp = &api.CommentListResp{
		Items: make([]*api.CommentInfo, len(data)),
	}

	switch req.TargetType {
	case model.TargetTypeArticle:
		var stat *model.ArticleStat
		if stat, err = p.d.GetArticleStatByID(c, p.d.DB(), req.ResourceID); err != nil {
			return
		}

		resp.CommentsCount = int32(stat.CommentCount)
		break

	case model.TargetTypeRevise:
		var stat *model.ReviseStat
		if stat, err = p.d.GetReviseStatByID(c, p.d.DB(), req.ResourceID); err != nil {
			return
		}

		resp.CommentsCount = int32(stat.CommentCount)
		break

	case model.TargetTypeDiscussion:
		var stat *model.DiscussionStat
		if stat, err = p.d.GetDiscussionStatByID(c, p.d.DB(), req.ResourceID); err != nil {
			return
		}

		resp.CommentsCount = int32(stat.CommentCount)
		break
	}

	for i, v := range data {
		item := &api.CommentInfo{
			ID:         v.ID,
			Content:    v.Content,
			Featured:   bool(v.Featured),
			IsDelete:   bool(v.Deleted),
			OwnerID:    v.OwnerID,
			ResourceID: v.ResourceID,
			TargetType: v.TargetType,
			CreatedAt:  v.CreatedAt,
			OwnerType:  v.OwnerType,
		}

		var stat *model.CommentStat
		if stat, err = p.d.GetCommentStatByID(c, p.d.DB(), v.ID); err != nil {
			return
		}

		item.Stat = &api.CommentStat{}
		item.Stat.LikeCount = stat.LikeCount
		item.Stat.DislikeCount = stat.DislikeCount
		item.Stat.ChildrenCount = stat.ChildrenCount

		item.ChildCommentsCount = stat.ChildrenCount

		if item.ChildCommentsCount > 0 {
			children := make([]*model.Comment, 0)
			if children, err = p.d.GetChildrenComments(c, p.d.DB(), v.ID, 3); err != nil {
				return
			}

			item.ChildComments = make([]*api.ChildCommentItem, len(children))

			for j, x := range children {
				child := &api.ChildCommentItem{
					ID:         x.ID,
					Content:    x.Content,
					Featured:   bool(x.Featured),
					IsDelete:   bool(x.Deleted),
					OwnerID:    x.OwnerID,
					ResourceID: x.ResourceID,
					TargetType: x.TargetType,
					CreatedAt:  x.CreatedAt,
					OwnerType:  x.OwnerType,
				}

				if child.Like, err = p.isLike(c, p.d.DB(), req.Aid, x.ID, model.TargetTypeComment); err != nil {
					return
				}

				if child.Dislike, err = p.isDislike(c, p.d.DB(), req.Aid, x.ID, model.TargetTypeComment); err != nil {
					return
				}

				var stat *model.CommentStat
				if stat, err = p.d.GetCommentStatByID(c, p.d.DB(), x.ID); err != nil {
					return
				}

				child.Stat = &api.CommentStat{}
				child.Stat.LikeCount = stat.LikeCount
				child.Stat.DislikeCount = stat.DislikeCount
				child.Stat.ChildrenCount = stat.ChildrenCount

				var acc *model.Account
				if acc, err = p.getAccount(c, p.d.DB(), x.CreatedBy); err != nil {
					return
				}
				child.Creator = &api.Creator{
					ID:           acc.ID,
					UserName:     acc.UserName,
					Avatar:       acc.Avatar,
					Introduction: acc.Introduction,
				}

				if x.ReplyTo != 0 {
					var acc *model.Account
					if acc, err = p.getAccount(c, p.d.DB(), x.ReplyTo); err != nil {
						return
					}
					child.ReplyTo = &api.Creator{
						ID:           acc.ID,
						UserName:     acc.UserName,
						Avatar:       acc.Avatar,
						Introduction: acc.Introduction,
					}
				}

				item.ChildComments[j] = child
			}
		}

		var acc *model.Account
		if acc, err = p.getAccount(c, p.d.DB(), v.CreatedBy); err != nil {
			return
		}
		item.Creator = &api.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		if v.ReplyTo != 0 {
			var acc *model.Account
			if acc, err = p.getAccount(c, p.d.DB(), v.ReplyTo); err != nil {
				return
			}
			item.ReplyTo = &api.Creator{
				ID:           acc.ID,
				UserName:     acc.UserName,
				Avatar:       acc.Avatar,
				Introduction: acc.Introduction,
			}
		}

		if item.Like, err = p.isLike(c, p.d.DB(), req.Aid, v.ID, model.TargetTypeComment); err != nil {
			return
		}

		if item.Dislike, err = p.isDislike(c, p.d.DB(), req.Aid, v.ID, model.TargetTypeComment); err != nil {
			return
		}

		resp.Items[i] = item
	}

	return
}

func (p *Service) GetCommentInfo(c context.Context, aid, commentID int64) (item *api.CommentInfo, err error) {
	var data *model.Comment
	if data, err = p.getComment(c, p.d.DB(), commentID); err != nil {
		return
	}

	var stat *model.CommentStat
	if stat, err = p.d.GetCommentStatByID(c, p.d.DB(), commentID); err != nil {
		return
	}

	var acc *model.Account
	if acc, err = p.getAccount(c, p.d.DB(), data.CreatedBy); err != nil {
		return
	}

	resp := &api.CommentInfo{
		ID:         data.ID,
		Content:    data.Content,
		TargetType: data.TargetType,
		Deleted:    bool(data.Deleted),
		Featured:   bool(data.Featured),
		OwnerID:    data.OwnerID,
		OwnerType:  data.OwnerType,
		ResourceID: data.ResourceID,
		CreatedAt:  data.CreatedAt,
		Stat: &api.CommentStat{
			ChildrenCount: int32(stat.ChildrenCount),
			LikeCount:     int32(stat.LikeCount),
			DislikeCount:  int32(stat.DislikeCount),
		},
		Creator: &api.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		},
	}

	if data.ReplyTo != int64(0) {
		var r *model.Account
		if r, err = p.getAccount(c, p.d.DB(), data.ReplyTo); err != nil {
			return
		}

		resp.ReplyTo = &api.Creator{
			ID:           r.ID,
			UserName:     r.UserName,
			Avatar:       r.Avatar,
			Introduction: r.Introduction,
		}
	}

	if resp.Like, err = p.isLike(c, p.d.DB(), aid, data.ID, model.TargetTypeComment); err != nil {
		return
	}

	if resp.Dislike, err = p.isDislike(c, p.d.DB(), aid, data.ID, model.TargetTypeComment); err != nil {
		return
	}

	return
}

// getComment 获取评论信息
func (p *Service) getComment(c context.Context, node sqalx.Node, commentID int64) (item *model.Comment, err error) {
	if item, err = p.d.GetCommentByID(c, p.d.DB(), commentID); err != nil {
		return
	} else if item == nil {
		err = ecode.CommentNotExist
		return
	}
	return
}

// GetCommentStat 获取评论状态信息
func (p *Service) GetCommentStat(c context.Context, commentID int64) (stat *model.CommentStat, err error) {
	return p.d.GetCommentStatByID(c, p.d.DB(), commentID)
}

// DelComment 删除评论
func (p *Service) DelComment(c context.Context, aid, commentID int64) (err error) {
	// 因为被删除评论也会显示，所以不做其他处理
	return p.d.DelComment(c, p.d.DB(), commentID)
}

func (p *Service) AddComment(c context.Context, arg *api.AddCommentReq) (id int64, err error) {
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

	item := &model.Comment{
		ID:         gid.NewID(),
		Content:    arg.Content,
		TargetType: arg.TargetType,
		ResourceID: arg.TargetID,
		OwnerID:    arg.TargetID,
		OwnerType:  arg.TargetType,
		CreatedBy:  arg.Aid,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if arg.TargetType == model.TargetTypeComment {
		var comment *model.Comment
		if comment, err = p.getComment(c, tx, arg.TargetID); err != nil {
			return
		}

		item.OwnerID = comment.OwnerID
		item.OwnerType = comment.OwnerType

		// 如果对象是子回复，则添加被回复人
		if comment.TargetType == model.TargetTypeComment {
			if arg.Aid != comment.CreatedBy {
				item.ReplyTo = comment.CreatedBy
			}
			item.ResourceID = comment.ResourceID

			if err = p.d.IncrCommentStat(c, tx, &model.CommentStat{
				CommentID:     comment.ResourceID,
				ChildrenCount: 1,
			}); err != nil {
				return
			}

		} else {
			// 如果被回复对象是回复  则直接设置当前的资源ID为被回复的ID
			item.ResourceID = comment.ID

			if err = p.d.IncrCommentStat(c, tx, &model.CommentStat{
				CommentID:     comment.ID,
				ChildrenCount: 1,
			}); err != nil {
				return
			}
		}

	}

	if err = p.d.AddComment(c, tx, item); err != nil {
		return
	}

	if err = p.d.AddCommentStat(c, tx, &model.CommentStat{
		CommentID: item.ID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = p.incrStat(c, tx, item.ResourceID, arg); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	id = item.ID

	p.addCache(func() {
		p.sendNotify(item.ResourceID, item.ID, arg.Aid, arg)
	})

	return
}

func (p *Service) sendNotify(resourceID, commentID int64, aid int64, arg *api.AddCommentReq) (err error) {
	switch arg.TargetType {
	case model.TargetTypeArticle:
		p.addCache(func() {
			p.onArticleCommented(context.Background(), commentID, aid, time.Now().Unix())
		})

		break
	case model.TargetTypeRevise:
		p.addCache(func() {
			p.onReviseCommented(context.Background(), commentID, aid, time.Now().Unix())
		})
		break
	case model.TargetTypeDiscussion:
		p.addCache(func() {
			p.onDiscussionCommented(context.Background(), commentID, aid, time.Now().Unix())
		})
		break
	case model.TargetTypeComment:
		p.addCache(func() {
			p.onCommentReplied(context.Background(), commentID, arg.TargetID, aid, time.Now().Unix())
		})
		break
	}

	return
}

func (p *Service) incrStat(c context.Context, node sqalx.Node, resourceID int64, arg *api.AddCommentReq) (err error) {
	targetType := arg.TargetType
	if targetType == model.TargetTypeComment {
		var comment *model.Comment
		if comment, err = p.getComment(c, node, resourceID); err != nil {
			return
		}

		// if err = p.d.IncrCommentStat(c, node, &model.CommentStat{CommentID: resourceID, ChildrenCount: 1}); err != nil {
		// 	return
		// }

		targetType = comment.TargetType
	}

	switch targetType {
	case model.TargetTypeArticle:
		if err = p.d.IncrArticleStat(c, node, &model.ArticleStat{ArticleID: arg.TargetID, CommentCount: 1}); err != nil {
			return
		}

		break
	case model.TargetTypeRevise:
		if err = p.d.IncrReviseStat(c, node, &model.ReviseStat{ReviseID: arg.TargetID, CommentCount: 1}); err != nil {
			return
		}
		break
	case model.TargetTypeDiscussion:
		if err = p.d.IncrDiscussionStat(c, node, &model.DiscussionStat{DiscussionID: arg.TargetID, CommentCount: 1}); err != nil {
			return
		}
		break
	}

	return

}
