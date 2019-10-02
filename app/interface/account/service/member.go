package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	feed "valerian/app/service/feed/api"
	relation "valerian/app/service/relation/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
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

	var stat *relation.StatInfo
	if stat, err = p.d.Stat(c, aid, targetID); err != nil {
		return
	}
	resp.Stat = &model.MemberInfoStat{
		FansCount:      int(stat.Fans),
		FollowingCount: int(stat.Following),
		IsFollow:       stat.IsFollowing,
	}

	var accountStat *account.AccountStatInfo
	if accountStat, err = p.d.GetAccountStat(c, targetID); err != nil {
		return
	}

	resp.Stat.TopicCount = int(accountStat.TopicCount)
	resp.Stat.ArticleCount = int(accountStat.ArticleCount)
	resp.Stat.DiscussionCount = int(accountStat.DiscussionCount)

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
