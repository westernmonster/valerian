package dao

import (
	"context"
	"fmt"

	comment "valerian/app/service/comment/api"
	"valerian/library/log"
)

func (p *Dao) GetCommentsPaged(c context.Context, req *comment.CommentListReq) (info *comment.CommentListResp, err error) {
	if info, err = p.commentRPC.GetCommentsPaged(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCommentsPaged, error(%+v) req(%+v)", err, req))
	}
	return
}

func (p *Dao) GetCommentInfo(c context.Context, req *comment.IDReq) (info *comment.CommentInfo, err error) {
	if info, err = p.commentRPC.GetCommentInfo(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCommentInfo, error(%+v) req(%+v)", err, req))
	}
	return
}

func (p *Dao) AddComment(c context.Context, req *comment.AddCommentReq) (err error) {
	if _, err = p.commentRPC.AddComment(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddComment, error(%+v) req(%+v)", err, req))
	}
	return
}

func (p *Dao) DelComment(c context.Context, req *comment.DeleteReq) (err error) {
	if _, err = p.commentRPC.DeleteComment(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelComment, error(%+v) req(%+v)", err, req))
	}
	return
}

func (p *Dao) GetAllChildrenComment(c context.Context, req *comment.IDReq) (info *comment.ChildrenCommentListResp, err error) {
	if info, err = p.commentRPC.GetAllChildrenComment(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllChildrenComment, error(%+v) req(%+v)", err, req))
	}
	return
}
