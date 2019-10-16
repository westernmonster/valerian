package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/account/model"
	feed "valerian/app/service/account-feed/api"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	recent "valerian/app/service/recent/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) FromDiscussion(v *discuss.DiscussionInfo) (item *model.TargetDiscuss) {
	item = &model.TargetDiscuss{
		ID:           v.ID,
		Excerpt:      v.Excerpt,
		CommentCount: int(v.Stat.CommentCount),
		LikeCount:    int(v.Stat.LikeCount),
		DislikeCount: int(v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:       v.Creator.ID,
			Avatar:   v.Creator.Avatar,
			UserName: v.Creator.UserName,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}

	if v.ImageUrls == nil {
		item.ImageUrls = make([]string, 0)
	} else {
		item.ImageUrls = v.ImageUrls
	}

	title := v.GetTitleValue()
	item.Title = &title

	intro := v.Creator.GetIntroductionValue()
	item.Creator.Introduction = &intro
	return
}

func (p *Service) FromRevise(v *article.ReviseInfo) (item *model.TargetRevise) {
	item = &model.TargetRevise{
		ID:           v.ID,
		Title:        v.Title,
		Excerpt:      v.Excerpt,
		CommentCount: int(v.Stat.CommentCount),
		LikeCount:    int(v.Stat.LikeCount),
		DislikeCount: int(v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:       v.Creator.ID,
			Avatar:   v.Creator.Avatar,
			UserName: v.Creator.UserName,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
	if v.ImageUrls == nil {
		item.ImageUrls = make([]string, 0)
	} else {
		item.ImageUrls = v.ImageUrls
	}

	intro := v.Creator.GetIntroductionValue()
	item.Creator.Introduction = &intro
	return
}

func (p *Service) FromArticle(v *article.ArticleInfo) (item *model.TargetArticle) {
	item = &model.TargetArticle{
		ID:           v.ID,
		Title:        v.Title,
		Excerpt:      v.Excerpt,
		ReviseCount:  int(v.Stat.ReviseCount),
		CommentCount: int(v.Stat.CommentCount),
		LikeCount:    int(v.Stat.LikeCount),
		DislikeCount: int(v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:       v.Creator.ID,
			Avatar:   v.Creator.Avatar,
			UserName: v.Creator.UserName,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
	if v.ImageUrls == nil {
		item.ImageUrls = make([]string, 0)
	} else {
		item.ImageUrls = v.ImageUrls
	}

	intro := v.Creator.GetIntroductionValue()
	item.Creator.Introduction = &intro
	return
}

func (p *Service) FromTopic(v *topic.TopicInfo) (item *model.TargetTopic) {
	item = &model.TargetTopic{
		ID:              v.ID,
		Name:            v.Name,
		Introduction:    v.Introduction,
		MemberCount:     int(v.Stat.MemberCount),
		DiscussionCount: int(v.Stat.DiscussionCount),
		ArticleCount:    int(v.Stat.ArticleCount),
		Creator: &model.Creator{
			ID:       v.Creator.ID,
			Avatar:   v.Creator.Avatar,
			UserName: v.Creator.UserName,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}

	intro := v.Creator.GetIntroductionValue()
	item.Creator.Introduction = &intro

	avatar := v.GetAvatarValue()
	item.Avatar = &avatar
	return
}

func (p *Service) GetMemberRecentPubsPaged(c context.Context, aid int64, atype string, limit, offset int) (resp *model.RecentPublishResp, err error) {
	var data *recent.RecentPubsResp
	if data, err = p.d.GetRecentPubsPaged(c, aid, limit, offset); err != nil {
		return
	}

	resp = &model.RecentPublishResp{
		Items:  make([]*model.PublishItem, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		item := &model.PublishItem{
			Type: v.TargetType,
		}

		switch v.TargetType {
		case model.TargetTypeArticle:
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, v.TargetID); err != nil {
				return
			}

			item.Article = p.FromArticle(article)
			break
		case model.TargetTypeRevise:
			var revise *article.ReviseInfo
			if revise, err = p.d.GetRevise(c, v.TargetID); err != nil {
				return
			}

			item.Revise = p.FromRevise(revise)
			break
		case model.TargetTypeDiscussion:
			var discuss *discuss.DiscussionInfo
			if discuss, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
				return
			}

			item.Discussion = p.FromDiscussion(discuss)
			break
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/recent", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"type":   []string{atype},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/recent", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"type":   []string{atype},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset + limit)},
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
	if stat, err = p.d.GetAccountStat(c, targetID); err != nil {
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

		var account *model.Account
		if account, err = p.getAccountByID(c, v.AccountID); err != nil {
			return
		}

		item.Source.Actor = &model.Actor{
			ID:           account.ID,
			Avatar:       account.Avatar,
			Name:         account.UserName,
			Introduction: account.Introduction,
		}

		switch v.TargetType {

		case model.TargetTypeArticle:
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, v.TargetID); err != nil {
				return
			}

			item.Target.Article = p.FromArticle(article)
			break
		case model.TargetTypeRevise:
			var revise *article.ReviseInfo
			if revise, err = p.d.GetRevise(c, v.TargetID); err != nil {
				return
			}

			item.Target.Revise = p.FromRevise(revise)
			break
		case model.TargetTypeDiscussion:
			var discuss *discuss.DiscussionInfo
			if discuss, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
				return
			}

			item.Target.Discussion = p.FromDiscussion(discuss)
			break

		case model.TargetTypeTopic:
			var topic *topic.TopicInfo
			if topic, err = p.d.GetTopic(c, v.TargetID); err != nil {
				return
			}

			item.Target.Topic = p.FromTopic(topic)
			break

		case model.TargetTypeMember:
			var info *model.MemberInfo
			if info, err = p.GetMemberInfo(c, v.TargetID); err != nil {
				return
			}

			item.Target.Member = info
			break
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/activities", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/activities", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset + limit)},
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

func (p *Service) GetMemberArticlesPaged(c context.Context, aid int64, limit, offset int) (resp *model.MemberArticleResp, err error) {
	var data *article.UserArticlesResp
	if data, err = p.d.GetUserArticlesPaged(c, aid, limit, offset); err != nil {
		return
	}

	resp = &model.MemberArticleResp{
		Items:  make([]*model.MemberArticle, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		item := &model.MemberArticle{
			ID:           v.ID,
			Title:        v.Title,
			Excerpt:      v.Excerpt,
			LikeCount:    int(v.Stat.LikeCount),
			DislikeCount: int(v.Stat.DislikeCount),
			CommentCount: int(v.Stat.CommentCount),
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}

		if v.ImageUrls == nil {
			item.ImageUrls = make([]string, 0)
		} else {
			item.ImageUrls = v.ImageUrls
		}

		resp.Items[i] = item
	}
	if resp.Paging.Prev, err = genURL("/api/v1/account/list/articles", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/articles", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset + limit)},
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
			ID:           v.ID,
			Excerpt:      v.Excerpt,
			LikeCount:    int(v.Stat.LikeCount),
			DislikeCount: int(v.Stat.DislikeCount),
			CommentCount: int(v.Stat.CommentCount),
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}

		if v.ImageUrls == nil {
			item.ImageUrls = make([]string, 0)
		} else {
			item.ImageUrls = v.ImageUrls
		}

		title := v.GetTitleValue()
		item.Title = &title

		resp.Items[i] = item
	}
	if resp.Paging.Prev, err = genURL("/api/v1/account/list/discussions", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/discussions", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset + limit)},
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
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}

		avatar := v.GetAvatarValue()
		item.Avatar = &avatar
		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/topics", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/topics", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset + limit)},
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
