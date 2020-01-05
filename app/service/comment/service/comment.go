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
