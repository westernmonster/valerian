package service

import (
	"context"

	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) getDiscussion(c context.Context, node sqalx.Node, articleID int64) (item *model.Discussion, err error) {
	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), articleID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	return
}
