package service

import (
	"context"
	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/model"
	"valerian/library/ecode"
)

func (p *Service) GetUserDiscussionsPaged(c context.Context, aid int64, limit, offset int) (items []*api.DiscussionInfo, err error) {
	var data []*model.Discussion
	if data, err = p.d.GetUserDiscussionsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	items = make([]*api.DiscussionInfo, len(data))

	for i, v := range data {

		imageUrls := make([]string, 0)
		var imgs []*model.ImageURL
		if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
			"target_type": model.TargetTypeDiscussion,
			"target_id":   v.ID,
		}); err != nil {
			return
		}

		for _, v := range imgs {
			imageUrls = append(imageUrls, v.URL)
		}

		var stat *model.DiscussionStat
		if stat, err = p.d.GetDiscussionStatByID(c, p.d.DB(), v.ID); err != nil {
			return
		}

		items[i] = api.FromDiscussion(v, stat, imageUrls)

	}

	return
}

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
	} else {
		item = &model.DiscussionStat{
			DiscussionID: discussionID,
		}
	}
	return
}
