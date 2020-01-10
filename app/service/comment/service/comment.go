package service

import (
	"context"
	account "valerian/app/service/account/api"
	"valerian/app/service/comment/api"
	"valerian/app/service/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// GetAccountBaseInfo 获取用户信息
func (p *Service) GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error) {
	return p.d.GetAccountBaseInfo(c, aid)
}

// GetComment 获取评论信息
func (p *Service) GetComment(c context.Context, commentID int64) (item *model.Comment, err error) {
	return p.getComment(c, p.d.DB(), commentID)
}

func (p *Service) GetCommentInfo(c context.Context, commentID int64) (item *api.CommentInfo, err error) {

	var data *model.Comment
	if data, err = p.getComment(c, p.d.DB(), commentID); err != nil {
		return
	}

	var stat *model.CommentStat
	if stat, err = p.d.GetCommentStatByID(c, p.d.DB(), commentID); err != nil {
		return
	}

	var acc *account.BaseInfoReply
	if acc, err = p.d.GetAccountBaseInfo(c, data.CreatedBy); err != nil {
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
		var r *account.BaseInfoReply
		if r, err = p.d.GetAccountBaseInfo(c, data.ReplyTo); err != nil {
			return
		}

		resp.ReplyTo = &api.Creator{
			ID:           r.ID,
			UserName:     r.UserName,
			Avatar:       r.Avatar,
			Introduction: r.Introduction,
		}
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
func (p *Service) DelComment(c context.Context, commentID int64) (err error) {
	// 因为被删除评论也会显示，所以不做其他处理
	return p.d.DelComment(c, p.d.DB(), commentID)
}
