package service

import (
	"context"
	account "valerian/app/service/account/api"
	"valerian/app/service/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error) {
	return p.d.GetAccountBaseInfo(c, aid)
}

func (p *Service) GetComment(c context.Context, commentID int64) (item *model.Comment, err error) {
	return p.getComment(c, p.d.DB(), commentID)
}

func (p *Service) getComment(c context.Context, node sqalx.Node, commentID int64) (item *model.Comment, err error) {
	// var addCache = true
	// if item, err = p.d.CommentCache(c, commentID); err != nil {
	// 	addCache = false
	// } else if item != nil {
	// 	return
	// }

	if item, err = p.d.GetCommentByID(c, p.d.DB(), commentID); err != nil {
		return
	} else if item == nil {
		err = ecode.CommentNotExist
		return
	}

	// if addCache {
	// 	p.addCache(func() {
	// 		p.d.SetCommentCache(context.TODO(), item)
	// 	})
	// }
	return
}

func (p *Service) GetCommentStat(c context.Context, commentID int64) (stat *model.CommentStat, err error) {
	return p.d.GetCommentStatByID(c, p.d.DB(), commentID)
}
