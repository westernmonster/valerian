package service

import (
	"context"

	"valerian/app/interface/discuss/model"
)

func (p *Service) GetDiscussionFiles(c context.Context, aid, discussionID int64) (items []*model.DiscussionFileResp, err error) {
	return
}

func (p *Service) SaveDiscussionFiles(c context.Context, arg *model.ArgSaveDiscussionFiles) (err error) {
	return
}
