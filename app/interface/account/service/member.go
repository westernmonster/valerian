package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/account/model"
	feed "valerian/app/service/account-feed/api"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	recent "valerian/app/service/recent/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

// GetMemberRecentPubsPaged 获取用户最近发布列表
func (p *Service) GetMemberRecentPubsPaged(c context.Context, aid int64, atype string, limit, offset int) (resp *model.RecentPublishResp, err error) {
	var data *recent.RecentPubsResp
	if data, err = p.d.GetRecentPubsPaged(c, aid, atype, limit, offset); err != nil {
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
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Article = p.fromArticle(article)
			item.Article.ChangeDesc = ""
			break
		case model.TargetTypeRevise:
			var revise *article.ReviseInfo
			if revise, err = p.d.GetRevise(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Revise = p.fromRevise(revise)
			break
		case model.TargetTypeDiscussion:
			var discuss *discuss.DiscussionInfo
			if discuss, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Discussion = p.fromDiscussion(discuss)
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

// GetMemberInfo 获取指定用户信息
func (p *Service) GetMemberInfo(c context.Context, targetID int64) (resp *model.MemberInfo, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var f *account.MemberInfoReply
	if f, err = p.d.GetMemberInfo(c, targetID); err != nil {
		return
	}
	resp = &model.MemberInfo{
		ID:             f.ID,
		UserName:       f.UserName,
		Gender:         (f.Gender),
		Location:       f.Location,
		LocationString: f.LocationString,
		Introduction:   f.Introduction,
		Avatar:         f.Avatar,
		IDCert:         f.IDCert,
		WorkCert:       f.WorkCert,
		IsOrg:          f.IsOrg,
		IsVIP:          f.IsVIP,
		Company:        f.Company,
		Position:       f.Position,
	}

	var isFollowing bool
	if isFollowing, err = p.d.IsFollowing(c, aid, targetID); err != nil {
		return
	}

	resp.Stat = &model.MemberInfoStat{
		FansCount:       (f.Stat.FansCount),
		FollowingCount:  (f.Stat.FollowingCount),
		TopicCount:      (f.Stat.TopicCount),
		ArticleCount:    (f.Stat.ArticleCount),
		DiscussionCount: (f.Stat.DiscussionCount),
		IsFollow:        isFollowing,
	}
	return
}

// GetMemberActivitiesPaged 获取指定用户动态
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

		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, v.AccountID); err != nil {
			return
		}

		item.Source.Actor = &model.Actor{
			ID:           acc.ID,
			Avatar:       acc.Avatar,
			Name:         acc.UserName,
			Introduction: acc.Introduction,
		}

		switch v.TargetType {

		case model.TargetTypeArticle:
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Target.Article = p.fromArticle(article)
			item.Target.Article.ChangeDesc = ""
			break
		case model.TargetTypeArticleHistory:
			var h *article.ArticleHistoryResp
			if h, err = p.d.GetArticleHistory(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, h.ArticleID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Target.Article = p.fromArticle(article)
			item.Target.Article.ChangeDesc = h.ChangeDesc

			break
		case model.TargetTypeRevise:
			var revise *article.ReviseInfo
			if revise, err = p.d.GetRevise(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Target.Revise = p.fromRevise(revise)
			break
		case model.TargetTypeDiscussion:
			var discuss *discuss.DiscussionInfo
			if discuss, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Target.Discussion = p.fromDiscussion(discuss)
			break

		case model.TargetTypeTopic:
			var topic *topic.TopicInfo
			if topic, err = p.d.GetTopic(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Target.Topic = p.fromTopic(topic)
			break

		case model.TargetTypeMember:
			var info *model.MemberInfo
			if info, err = p.GetMemberInfo(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Target.Member = info
			break
		case model.TargetTypeComment:
			var info *comment.CommentInfo
			if info, err = p.d.GetComment(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
					break
				}
				return
			}

			item.Target.Comment = p.fromComment(info)

			switch info.TargetType {
			case model.TargetTypeArticle:
				var article *article.ArticleInfo
				if article, err = p.d.GetArticle(c, info.ResourceID); err != nil {
					if ecode.IsNotExistEcode(err) {
						item.Deleted = true
						break
					}
					return
				}

				item.Target.Article = p.fromArticle(article)
				item.Target.Article.ChangeDesc = ""
				break
			case model.TargetTypeRevise:
				var revise *article.ReviseInfo
				if revise, err = p.d.GetRevise(c, info.ResourceID); err != nil {
					if ecode.IsNotExistEcode(err) {
						item.Deleted = true
						break
					}
					return
				}

				item.Target.Revise = p.fromRevise(revise)
				break
			case model.TargetTypeDiscussion:
				var discuss *discuss.DiscussionInfo
				if discuss, err = p.d.GetDiscussion(c, info.ResourceID); err != nil {
					if ecode.IsNotExistEcode(err) {
						item.Deleted = true
						break
					}
					return
				}

				item.Target.Discussion = p.fromDiscussion(discuss)
				break

			}
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

// GetMemberArticlesPaged 获取指定用户文章列表
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
			ChangeDesc:   v.ChangeDesc,
			LikeCount:    (v.Stat.LikeCount),
			DislikeCount: (v.Stat.DislikeCount),
			CommentCount: (v.Stat.CommentCount),
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

// GetMemberDiscussionsPaged 获取指定用户讨论列表
func (p *Service) GetMemberDiscussionsPaged(c context.Context, aid int64, limit, offset int) (resp *model.MemberDiscussResp, err error) {
	var data *discuss.DiscussionsResp
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
			LikeCount:    (v.Stat.LikeCount),
			DislikeCount: (v.Stat.DislikeCount),
			CommentCount: (v.Stat.CommentCount),
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			Title:        v.Title,
		}

		if v.ImageUrls == nil {
			item.ImageUrls = make([]string, 0)
		} else {
			item.ImageUrls = v.ImageUrls
		}

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

// GetMemberManageTopicsPaged 获取指定用户管理的话题列表
func (p *Service) GetMemberManageTopicsPaged(c context.Context, aid int64, limit, offset int) (resp *model.MemberTopicResp, err error) {
	var data *topic.IDsResp
	if data, err = p.d.GetManageTopicIDsPaged(c, aid, int32(limit), int32(offset)); err != nil {
		return
	}

	resp = &model.MemberTopicResp{
		Items:  make([]*model.TargetTopic, 0),
		Paging: &model.Paging{},
	}

	if data.IDs == nil || len(data.IDs) == 0 {
		resp.Paging.IsEnd = true
		return
	}

	for _, v := range data.IDs {
		var t *topic.TopicInfo
		if t, err = p.d.GetTopic(c, v); err != nil {
			return
		}
		resp.Items = append(resp.Items, p.fromTopic(t))
	}

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/managed_topics", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/managed_topics", url.Values{
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

// GetMemberFollowedTopicsPaged 获取指定用户加入的话题列表
func (p *Service) GetMemberFollowedTopicsPaged(c context.Context, aid int64, limit, offset int) (resp *model.MemberTopicResp, err error) {
	var data *topic.IDsResp
	if data, err = p.d.GetFollowedTopicIDsPaged(c, aid, int32(limit), int32(offset)); err != nil {
		return
	}

	resp = &model.MemberTopicResp{
		Items:  make([]*model.TargetTopic, 0),
		Paging: &model.Paging{},
	}

	if data.IDs == nil || len(data.IDs) == 0 {
		resp.Paging.IsEnd = true
		return
	}

	for _, v := range data.IDs {
		var t *topic.TopicInfo
		if t, err = p.d.GetTopic(c, v); err != nil {
			return
		}
		resp.Items = append(resp.Items, p.fromTopic(t))
	}

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/followed_topics", url.Values{
		"id":     []string{strconv.FormatInt(aid, 10)},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/followed_topics", url.Values{
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

// GetMemberCert 获取指定用户的认证信息
func (p *Service) GetMemberCert(c context.Context, targetID int64) (resp *model.MemberCertInfo, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var f *account.MemberInfoReply
	if f, err = p.d.GetMemberInfo(c, targetID); err != nil {
		return
	}

	resp = &model.MemberCertInfo{}

	resp.Member = &model.MemberInfo{
		ID:             f.ID,
		UserName:       f.UserName,
		Gender:         (f.Gender),
		Location:       f.Location,
		LocationString: f.LocationString,
		Introduction:   f.Introduction,
		Avatar:         f.Avatar,
		IDCert:         f.IDCert,
		WorkCert:       f.WorkCert,
		IsOrg:          f.IsOrg,
		IsVIP:          f.IsVIP,
		Company:        f.Company,
		Position:       f.Position,
	}

	var isFollowing bool
	if isFollowing, err = p.d.IsFollowing(c, aid, targetID); err != nil {
		return
	}

	resp.Member.Stat = &model.MemberInfoStat{
		FansCount:       (f.Stat.FansCount),
		FollowingCount:  (f.Stat.FollowingCount),
		TopicCount:      (f.Stat.TopicCount),
		ArticleCount:    (f.Stat.ArticleCount),
		DiscussionCount: (f.Stat.DiscussionCount),
		IsFollow:        isFollowing,
	}

	var v *account.IDCertInfo
	if v, err = p.d.GetIDCert(c, targetID); err != nil {
		return
	}

	var w *account.WorkCertInfo
	if w, err = p.d.GetWorkCert(c, targetID); err != nil {
		return
	}

	resp.IDCert = &model.IDCertificationInfo{
		ID:         v.AccountID,
		Name:       v.Name,
		IDCardType: v.IDCardType,
		Status:     v.Status,
		UpdatedAt:  v.UpdatedAt,
	}

	if len(v.IdentificationNumber) > 0 {
		numbers := []rune(v.IdentificationNumber)
		resp.IDCert.IdentificationNumber = string(numbers[0]) + "*********" + string(numbers[len(numbers)-1])
	}

	resp.WorkCert = &model.WorkCertificationInfo{
		ID:         w.AccountID,
		Company:    w.Company,
		Department: w.Department,
		Position:   w.Position,
		Status:     w.Status,
		UpdatedAt:  w.UpdatedAt,
	}

	return
}

// GetMemberTopicsPaged 获取用户创建的话题列表
func (p *Service) GetMemberTopicsPaged(c context.Context, aid int64, limit, offset int) (resp *model.MemberTopicResp, err error) {
	var data *topic.UserTopicsResp
	if data, err = p.d.GetUserTopicsPaged(c, aid, int32(limit), int32(offset)); err != nil {
		return
	}

	resp = &model.MemberTopicResp{
		Items:  make([]*model.TargetTopic, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		item := &model.TargetTopic{
			ID:           v.ID,
			Name:         v.Name,
			MemberCount:  (v.Stat.MemberCount),
			Introduction: v.Introduction,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			Avatar:       v.Avatar,
		}

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
