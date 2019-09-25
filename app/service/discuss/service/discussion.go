package service

import (
	"context"
	"valerian/app/service/discuss/model"
	"valerian/library/ecode"
)

func (p *Service) GetDiscussion(c context.Context, discussionID int64) (item *model.Discussion, err error) {
	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), discussionID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	return
}
