package service

import (
	"context"
	"valerian/app/service/discuss/model"
	"valerian/library/ecode"
)

func (p *Service) GetDiscussion(c context.Context, discussionID int64) (item *model.Discussion, imageUrls []string, err error) {
	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), discussionID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	imageUrls = make([]string, 0)
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeDiscussion,
		"target_id":   discussionID,
	}); err != nil {
		return
	}

	for _, v := range imgs {
		imageUrls = append(imageUrls, v.URL)
	}

	return
}

func (p *Service) GetDiscussionStat(c context.Context, discussionID int64) (item *model.DiscussionStat, err error) {
	if item, err = p.d.GetDiscussionStatByID(c, p.d.DB(), discussionID); err != nil {
		return
	}
	return
}
