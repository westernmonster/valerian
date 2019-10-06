package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	discuss "valerian/app/service/discuss/api"
	feed "valerian/app/service/feed/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
	"valerian/library/xstr"
)

func (p *Service) GetMemberInfo(c context.Context, targetID int64) (resp *model.MemberInfo, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var f *model.Profile
	if f, err = p.getProfile(c, targetID); err != nil {
		return
	}

	resp = &model.MemberInfo{
		ID:             f.ID,
		UserName:       f.UserName,
		Gender:         f.Gender,
		Location:       f.Location,
		LocationString: f.LocationString,
		Introduction:   f.Introduction,
		Avatar:         f.Avatar,
		IDCert:         f.IDCert,
		WorkCert:       f.WorkCert,
		IsOrg:          f.IsOrg,
		IsVIP:          f.IsVIP,
	}

	var isFollowing bool
	if isFollowing, err = p.d.IsFollowing(c, aid, targetID); err != nil {
		return
	}

	var stat *account.AccountStatInfo
	if stat, err = p.d.GetAccountStat(c, aid); err != nil {
		return
	}

	resp.Stat = &model.MemberInfoStat{
		FansCount:       int(stat.Fans),
		FollowingCount:  int(stat.Following),
		TopicCount:      int(stat.TopicCount),
		ArticleCount:    int(stat.ArticleCount),
		DiscussionCount: int(stat.DiscussionCount),
		IsFollow:        isFollowing,
	}
	return
}

func (p *Service) GetMemberActivitiesPaged(c context.Context, aid int64, limit, offset int) (resp *model.FeedResp, err error) {
	var data *feed.AccountFeedResp
	if data, err = p.d.GetAccountFeedPaged(c, aid, limit, offset); err != nil {
		return
	}

	resp = &model.FeedResp{
		Items:  make([]*model.FeedItem, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		item := &model.FeedItem{
			TargetType: v.TargetType,
			Source: &model.FeedSource{
				ActionTime: v.ActionTime,
				ActionText: v.ActionText,
				ActionType: v.ActionType,
			},
			Target: &model.FeedTarget{
				ID:   v.TargetID,
				Type: v.TargetType,
			},
		}

		// if v.TargetType == model.TargetTypeDiscussion {
		// var discuss *discuss.DiscussionInfo
		// if discuss, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
		// 	return
		// }

		// item.Target.Discussion = &model.TargetDiscussion{
		// 	ID:           discuss.ID,
		// 	Images:       make([]string, 0),
		// 	Excerpt:      xstr.Excerpt(discuss.ContentText),
		// 	LikeCount:    int(discuss.Stat.LikeCount),
		// 	CommentCount: int(discuss.Stat.CommentCount),
		// }

		// title := discuss.GetTitleValue()
		// item.Target.Discussion.Title = &title
		// }

		resp.Items[i] = item
	}

	param := url.Values{}
	param.Set("account_id", strconv.FormatInt(aid, 10))
	param.Set("limit", strconv.Itoa(limit))
	param.Set("offset", strconv.Itoa(offset-limit))

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/activities", param); err != nil {
		return
	}
	param.Set("offset", strconv.Itoa(offset+limit))
	if resp.Paging.Next, err = genURL("/api/v1/account/list/activities", param); err != nil {
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

func (p *Service) GetMemberDiscussionsPaged(c context.Context, aid int64, limit, offset int) (resp *model.MemberDiscussResp, err error) {
	var data *discuss.UserDiscussionsResp
	if data, err = p.d.GetUserDiscussionsPaged(c, aid, limit, offset); err != nil {
		return
	}

	resp = &model.MemberDiscussResp{
		Items:  make([]*model.MemberDiscuss, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		item := &model.MemberDiscuss{
			ID:        v.ID,
			ImageUrls: v.ImageUrls,
			Excerpt:   xstr.Excerpt(v.ContentText),
			LikeCount: int(v.Stat.LikeCount),
			// DislikeCount: int(v.Stat.DislikeCount),
			CommentCount: int(v.Stat.CommentCount),
		}

		title := v.GetTitleValue()
		item.Title = &title

		resp.Items[i] = item
	}

	param := url.Values{}
	param.Set("account_id", strconv.FormatInt(aid, 10))
	param.Set("limit", strconv.Itoa(limit))
	param.Set("offset", strconv.Itoa(offset-limit))

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/discussions", param); err != nil {
		return
	}
	param.Set("offset", strconv.Itoa(offset+limit))
	if resp.Paging.Next, err = genURL("/api/v1/account/list/discussions", param); err != nil {
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

func (p *Service) GetMemberTopicsPaged(c context.Context, aid int64, limit, offset int) (resp *model.MemberTopicResp, err error) {
	var data *topic.UserTopicsResp
	if data, err = p.d.GetUserTopicsPaged(c, aid, limit, offset); err != nil {
		return
	}

	resp = &model.MemberTopicResp{
		Items:  make([]*model.MemberTopic, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		item := &model.MemberTopic{
			ID:           v.ID,
			Name:         v.Name,
			MemberCount:  int(v.Stat.MemberCount),
			Introduction: v.Introduction,
		}

		avatar := v.GetAvatarValue()
		item.Avatar = &avatar
		resp.Items[i] = item
	}

	param := url.Values{}
	param.Set("account_id", strconv.FormatInt(aid, 10))
	param.Set("limit", strconv.Itoa(limit))
	param.Set("offset", strconv.Itoa(offset-limit))

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/topics", param); err != nil {
		return
	}
	param.Set("offset", strconv.Itoa(offset+limit))
	if resp.Paging.Next, err = genURL("/api/v1/account/list/topics", param); err != nil {
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
