package service

import (
	"context"

	account "valerian/app/service/account/api"
	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/xstr"
)

func (p *Service) getDiscussion(c context.Context, node sqalx.Node, articleID int64) (item *model.Discussion, err error) {
	var addCache = true
	if item, err = p.d.DiscussionCache(c, articleID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), articleID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetDiscussionCache(context.TODO(), item)
		})
	}
	return
}

func (p *Service) GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error) {
	return p.d.GetAccountBaseInfo(c, aid)
}

func (p *Service) GetUserDiscussionIDsPaged(c context.Context, req *api.UserDiscussionsReq) (resp []int64, err error) {
	return p.d.GetUserDiscussionIDsPaged(c, p.d.DB(), req.AccountID, int(req.Limit), int(req.Offset))
}

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

		item := &api.DiscussionInfo{
			ID:         v.ID,
			TopicID:    v.TopicID,
			CategoryID: v.CategoryID,
			// CreatedBy:   v.CreatedBy,
			Excerpt:   xstr.Excerpt(v.ContentText),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			ImageUrls: imageUrls,
			Title:     v.Title,
		}

		item.Stat = &api.DiscussionStat{
			DislikeCount: int32(stat.DislikeCount),
			LikeCount:    int32(stat.LikeCount),
			CommentCount: int32(stat.CommentCount),
		}

		if v.CategoryID != int64(-1) {
			var cate *model.DiscussCategory
			if cate, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
				return
			} else if cate == nil {
				err = ecode.DiscussCategoryNotExist
				return
			}

			item.CategoryInfo = &api.CategoryInfo{
				ID:      cate.ID,
				TopicID: cate.TopicID,
				Name:    cate.Name,
				Seq:     cate.Seq,
			}
		}

		var acc *account.BaseInfoReply
		if acc, err = p.GetAccountBaseInfo(c, v.CreatedBy); err != nil {
			return
		}

		item.Creator = &api.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		items[i] = item
	}

	return
}

func (p *Service) GetAllDiscussions(c context.Context) (items []*api.DiscussionInfo, err error) {
	var data []*model.Discussion
	if data, err = p.d.GetDiscussions(c, p.d.DB()); err != nil {
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

		item := &api.DiscussionInfo{
			ID:         v.ID,
			TopicID:    v.TopicID,
			CategoryID: v.CategoryID,
			// CreatedBy:   v.CreatedBy,
			Excerpt:   xstr.Excerpt(v.ContentText),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			ImageUrls: imageUrls,
			Title:     v.Title,
		}

		item.Stat = &api.DiscussionStat{
			DislikeCount: int32(stat.DislikeCount),
			LikeCount:    int32(stat.LikeCount),
			CommentCount: int32(stat.CommentCount),
		}

		if v.CategoryID != int64(-1) {
			var cate *model.DiscussCategory
			if cate, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
				return
			} else if cate == nil {
				err = ecode.DiscussCategoryNotExist
				return
			}

			item.CategoryInfo = &api.CategoryInfo{
				ID:      cate.ID,
				TopicID: cate.TopicID,
				Name:    cate.Name,
				Seq:     cate.Seq,
			}
		}

		var acc *account.BaseInfoReply
		if acc, err = p.GetAccountBaseInfo(c, v.CreatedBy); err != nil {
			return
		}

		item.Creator = &api.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		items[i] = item

	}

	return
}

func (p *Service) GetDiscussion(c context.Context, discussionID int64) (item *model.Discussion, imageUrls []string, err error) {
	if item, err = p.getDiscussion(c, p.d.DB(), discussionID); err != nil {
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

func (p *Service) DelDiscussion(c context.Context, aid, discussionID int64) (err error) {
	return
}
