package dao

import (
	"context"
	"fmt"
	comment "valerian/app/service/comment/api"
	"valerian/library/log"
)

func (p *Dao) GetComment(c context.Context, id int64, useMaster bool) (info *comment.CommentInfo, err error) {
	if info, err = p.commentRPC.GetCommentInfo(c, &comment.IDReq{ID: id, UseMaster: useMaster}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetComment, error(%+v) id(%d)", err, id))
	}
	return
}
