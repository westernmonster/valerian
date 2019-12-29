package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/discuss/model"
	discussion "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetUserDiscussionsPaged(c context.Context, accountID int64, limit, offset int) (resp *model.DiscussListResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data *discussion.DiscussionsResp
	if data, err = p.d.GetUserDiscussionsPaged(c, accountID, aid, limit, offset); err != nil {
		return
	}

	resp = &model.DiscussListResp{
		Items:  make([]*model.DiscussItem, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		item := &model.DiscussItem{
			ID:        v.ID,
			TopicID:   v.TopicID,
			Title:     v.Title,
			Excerpt:   v.Excerpt,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			ImageUrls: make([]string, 0),
		}

		item.LikeCount = v.Stat.LikeCount
		item.DislikeCount = v.Stat.DislikeCount
		item.CommentCount = v.Stat.CommentCount

		item.Creator = &model.Creator{
			ID:           v.Creator.ID,
			UserName:     v.Creator.UserName,
			Avatar:       v.Creator.Avatar,
			Introduction: v.Creator.Introduction,
		}

		item.Category = &model.DiscussItemCategory{
			ID:   v.CategoryID,
			Name: v.CategoryInfo.Name,
		}

		if v.ImageUrls != nil {
			item.ImageUrls = v.ImageUrls
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/discussion/list/by_account", url.Values{
		"account_id": []string{strconv.FormatInt(accountID, 10)},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/discussion/list/by_account", url.Values{
		"account_id": []string{strconv.FormatInt(accountID, 10)},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset + limit)},
	}); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	return
}

func (p *Service) GetTopicDiscussionsPaged(c context.Context, topicID, categoryID int64, limit, offset int) (resp *model.DiscussListResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data *discussion.DiscussionsResp
	if data, err = p.d.GetTopicDiscussionsPaged(c, topicID, categoryID, aid, limit, offset); err != nil {
		return
	}

	resp = &model.DiscussListResp{
		Items:  make([]*model.DiscussItem, len(data.Items)),
		Paging: &model.Paging{},
	}

	resp = &model.DiscussListResp{
		Items:  make([]*model.DiscussItem, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		item := &model.DiscussItem{
			ID:        v.ID,
			TopicID:   v.TopicID,
			Title:     v.Title,
			Excerpt:   v.Excerpt,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			ImageUrls: make([]string, 0),
		}

		item.LikeCount = v.Stat.LikeCount
		item.DislikeCount = v.Stat.DislikeCount
		item.CommentCount = v.Stat.CommentCount

		item.Creator = &model.Creator{
			ID:           v.Creator.ID,
			UserName:     v.Creator.UserName,
			Avatar:       v.Creator.Avatar,
			Introduction: v.Creator.Introduction,
		}

		item.Category = &model.DiscussItemCategory{
			ID:   v.CategoryID,
			Name: v.CategoryInfo.Name,
		}

		if v.ImageUrls != nil {
			item.ImageUrls = v.ImageUrls
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/discussion/list/by_topic", url.Values{
		"topic_id":    []string{strconv.FormatInt(topicID, 10)},
		"category_id": []string{strconv.FormatInt(categoryID, 10)},
		"limit":       []string{strconv.Itoa(limit)},
		"offset":      []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/discussion/list/by_topic", url.Values{
		"topic_id":    []string{strconv.FormatInt(topicID, 10)},
		"category_id": []string{strconv.FormatInt(categoryID, 10)},
		"limit":       []string{strconv.Itoa(limit)},
		"offset":      []string{strconv.Itoa(offset + limit)},
	}); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	return
}

func (p *Service) AddDiscussion(c context.Context, arg *model.ArgAddDiscuss) (id int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	req := &discussion.ArgAddDiscussion{
		TopicID:    arg.TopicID,
		CategoryID: arg.CategoryID,
		Content:    arg.Content,
		Files:      make([]*discussion.ArgDiscussionFile, 0),
		Aid:        aid,
	}

	if arg.Title != nil {
		req.Title = &discussion.ArgAddDiscussion_TitleValue{*arg.Title}
	}

	for _, v := range arg.Files {
		req.Files = append(req.Files, &discussion.ArgDiscussionFile{
			FileName: v.FileName,
			FileType: v.FileType,
			FileURL:  v.FileURL,
			Seq:      v.Seq,
		})
	}

	if id, err = p.d.AddDiscussion(c, req); err != nil {
		return
	}

	return
}

func (p *Service) UpdateDiscussion(c context.Context, arg *model.ArgUpdateDiscuss) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	req := &discussion.ArgUpdateDiscussion{
		ID:      arg.ID,
		Content: arg.Content,
		Aid:     aid,
	}

	if arg.Title != nil {
		req.Title = &discussion.ArgUpdateDiscussion_TitleValue{*arg.Title}
	}

	if err = p.d.UpdateDiscussion(c, req); err != nil {
		return
	}

	return
}

func (p *Service) DelDiscussion(c context.Context, id int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.DelDiscussion(c, aid, id); err != nil {
		return
	}
	return
}

func (p *Service) GetDiscussion(c context.Context, discussionID int64) (resp *model.DiscussDetailResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	req := &discussion.IDReq{
		ID:      discussionID,
		Aid:     aid,
		Include: "content,content_text",
	}

	var data *discussion.DiscussionInfo
	if data, err = p.d.GetDiscussion(c, req); err != nil {
		return
	}

	resp = &model.DiscussDetailResp{
		ID:        data.ID,
		TopicID:   data.TopicID,
		Title:     data.Title,
		Content:   data.Content,
		Files:     make([]*model.DiscussFileResp, 0),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	for _, v := range data.Files {
		resp.Files = append(resp.Files, &model.DiscussFileResp{
			ID:        v.ID,
			FileName:  v.FileName,
			FileURL:   v.FileURL,
			PdfURL:    v.PdfURL,
			FileType:  v.FileType,
			Seq:       v.Seq,
			CreatedAt: v.CreatedAt,
		})
	}

	resp.Category = &model.DiscussItemCategory{
		ID:   data.CategoryID,
		Name: data.CategoryInfo.Name,
	}

	resp.Creator = &model.Creator{
		ID:           data.Creator.ID,
		UserName:     data.Creator.UserName,
		Avatar:       data.Creator.Avatar,
		Introduction: data.Creator.Introduction,
	}

	resp.LikeCount = data.Stat.LikeCount
	resp.DislikeCount = data.Stat.DislikeCount
	resp.CommentCount = data.Stat.CommentCount

	if resp.Fav, err = p.d.IsFav(c, aid, discussionID, model.TargetTypeDiscussion); err != nil {
		return
	}

	if resp.Like, err = p.d.IsLike(c, aid, discussionID, model.TargetTypeDiscussion); err != nil {
		return
	}

	if resp.Dislike, err = p.d.IsDislike(c, aid, discussionID, model.TargetTypeDiscussion); err != nil {
		return
	}

	var tp *topic.TopicInfo
	if tp, err = p.d.GetTopic(c, data.TopicID); err != nil {
		return
	}

	resp.TopicName = tp.Name

	if aid == data.Creator.ID {
		resp.CanEdit = true
		return
	}

	// TODO: 讨论权限
	return
}
