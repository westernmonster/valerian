package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	fav "valerian/app/service/fav/api"
	identify "valerian/app/service/identify/api/grpc"
	message "valerian/app/service/message/api"
	recent "valerian/app/service/recent/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

// ForgetPassword 忘记密码
func (p *Service) ForgetPassword(c context.Context, arg *model.ArgForgetPassword) (sessionID string, err error) {
	resp, err := p.d.ForgetPassword(c, arg.Identity, arg.Valcode, arg.Prefix, arg.IdentityType)
	if err != nil {
		return
	}

	sessionID = resp.SessionID
	return
}

// ResetPassword 重设密码
func (p *Service) ResetPassword(c context.Context, arg *model.ArgResetPassword) (err error) {
	return p.d.ResetPassword(c, arg.Password, arg.SessionID)
}

// Deactive 注销
func (p *Service) Deactive(c context.Context, arg *model.ArgDeactiveAccount) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	req := &identify.DeactiveReq{
		Aid:          aid,
		IdentityType: arg.IdentityType,
		Identity:     arg.Identity,
		Valcode:      arg.Valcode,
		Prefix:       arg.Prefix,
	}
	return p.d.Deactive(c, req)
}

// UpdateProfile 更新用户资料
func (p *Service) UpdateProfile(c context.Context, arg *model.ArgUpdateProfile) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	req := &account.UpdateProfileReq{Aid: aid, AccountID: aid}

	if arg.Gender != nil {
		req.Gender = &account.UpdateProfileReq_GenderValue{*arg.Gender}
	}

	if arg.Avatar != nil {
		req.Avatar = &account.UpdateProfileReq_AvatarValue{*arg.Avatar}
	}

	if arg.Introduction != nil {
		req.Avatar = &account.UpdateProfileReq_AvatarValue{*arg.Avatar}
	}

	if arg.UserName != nil {
		req.UserName = &account.UpdateProfileReq_UserNameValue{*arg.UserName}
	}

	if arg.BirthYear != nil {
		req.BirthYear = &account.UpdateProfileReq_BirthYearValue{*arg.BirthYear}
	}

	if arg.BirthMonth != nil {
		req.BirthMonth = &account.UpdateProfileReq_BirthMonthValue{*arg.BirthMonth}
	}

	if arg.BirthDay != nil {
		req.BirthDay = &account.UpdateProfileReq_BirthDayValue{*arg.BirthDay}
	}

	if arg.Password != nil {
		req.Password = &account.UpdateProfileReq_PasswordValue{*arg.Password}

	}

	if arg.Location != nil {
		req.Location = &account.UpdateProfileReq_LocationValue{*arg.Location}
	}

	return p.d.UpdateProfile(c, req)
}

// ChangePassword 更改密码
func (p *Service) ChangePassword(c context.Context, aid int64, arg *model.ArgChangePassword) (err error) {
	req := &account.UpdateProfileReq{Aid: aid}

	req.Password = &account.UpdateProfileReq_PasswordValue{arg.Password}

	return p.d.UpdateProfile(c, req)
}

// GetProfile 获取当前登录用户个人资料
func (p *Service) GetProfile(c context.Context, aid int64) (item *model.Profile, err error) {
	var profile *account.SelfProfile
	if profile, err = p.d.GetSelfProfile(c, aid); err != nil {
		return
	}

	item = &model.Profile{
		ID:             profile.ID,
		Prefix:         profile.Prefix,
		Mobile:         profile.Mobile,
		Email:          profile.Email,
		UserName:       profile.UserName,
		Gender:         (profile.Gender),
		BirthYear:      (profile.BirthYear),
		BirthMonth:     (profile.BirthMonth),
		BirthDay:       (profile.BirthDay),
		Location:       profile.Location,
		LocationString: profile.LocationString,
		Introduction:   profile.Introduction,
		Avatar:         profile.Avatar,
		Source:         (profile.Source),
		IDCert:         profile.IDCert,
		IDCertStatus:   (profile.IDCertStatus),
		WorkCert:       profile.WorkCert,
		WorkCertStatus: (profile.WorkCertStatus),
		IsOrg:          profile.IsOrg,
		IsVIP:          profile.IsVIP,
		Role:           profile.Role,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
		Company:        profile.Company,
		Position:       profile.Position,
	}

	item.Stat = &model.MemberInfoStat{
		FansCount:       (profile.Stat.FansCount),
		FollowingCount:  (profile.Stat.FollowingCount),
		TopicCount:      (profile.Stat.TopicCount),
		ArticleCount:    (profile.Stat.ArticleCount),
		DiscussionCount: (profile.Stat.DiscussionCount),
	}

	item.Settings = &model.SettingResp{
		Activity: model.ActivitySettingResp{
			Like:         profile.Setting.ActivityLike,
			Comment:      profile.Setting.ActivityComment,
			FollowTopic:  profile.Setting.ActivityFollowTopic,
			FollowMember: profile.Setting.ActivityFollowMember,
		},
		Notify: model.NotifySettingResp{
			Like:      profile.Setting.NotifyLike,
			Comment:   profile.Setting.NotifyComment,
			NewFans:   profile.Setting.NotifyNewFans,
			NewMember: profile.Setting.NotifyNewMember,
		},
		Language: model.LanguageSettingResp{
			Language: profile.Setting.Language,
		},
	}

	var msgStat *message.MessageStat
	if msgStat, err = p.d.GetMessageStat(c, aid); err != nil {
		return
	}
	item.Stat.MsgCount = msgStat.UnreadCount

	return
}

// GetUserTopicsPaged 获取当前登录用户话题
func (p *Service) GetUserTopicsPaged(c context.Context, cate string, limit, offset int) (resp *model.AccountTopicsResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	resp = &model.AccountTopicsResp{
		Items:  make([]*model.TopicItem, 0),
		Paging: &model.Paging{},
	}

	switch cate {
	case model.CateManaged:
		var data *topic.IDsResp
		if data, err = p.d.GetManageTopicIDsPaged(c, aid, int32(limit), int32(offset)); err != nil {
			return
		}
		if data.IDs == nil || len(data.IDs) == 0 {
			return
		}

		for _, v := range data.IDs {
			item := &model.TopicItem{}

			var t *topic.TopicInfo
			if t, err = p.d.GetTopic(c, v); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				item.Target = p.fromTopic(t)
			}

			resp.Items = append(resp.Items, item)
		}
		break
	case model.CateFollowed:
		var data *topic.IDsResp
		if data, err = p.d.GetFollowedTopicIDsPaged(c, aid, int32(limit), int32(offset)); err != nil {
			return
		}
		if data.IDs == nil || len(data.IDs) == 0 {
			return
		}

		for _, v := range data.IDs {
			item := &model.TopicItem{}
			var t *topic.TopicInfo
			if t, err = p.d.GetTopic(c, v); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				item.Target = p.fromTopic(t)
			}

			resp.Items = append(resp.Items, item)
		}
		break
	case model.CateFaved:
		var data *fav.FavsResp
		if data, err = p.d.GetUserFavsPaged(c, aid, model.TargetTypeTopic, int32(limit), int32(offset)); err != nil {
			return
		}
		if data.Items == nil || len(data.Items) == 0 {
			return
		}

		for _, v := range data.Items {
			item := &model.TopicItem{}
			var t *topic.TopicInfo
			if t, err = p.d.GetTopic(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				item.Target = p.fromTopic(t)
			}

			resp.Items = append(resp.Items, item)
		}
		break
	case model.CateRecentViewed:
		var data *recent.RecentViewsResp
		if data, err = p.d.GetRecentViewsPaged(c, aid, model.TargetTypeTopic, int32(limit), int32(offset)); err != nil {
			return
		}
		if data.Items == nil || len(data.Items) == 0 {
			return
		}

		for _, v := range data.Items {
			item := &model.TopicItem{}
			var t *topic.TopicInfo
			if t, err = p.d.GetTopic(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				item.Target = p.fromTopic(t)
			}

			resp.Items = append(resp.Items, item)
		}
		break
	}

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/my_topics", url.Values{
		"cate":   []string{cate},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/my_topics", url.Values{
		"cate":   []string{cate},
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

// GetUserArticlesPaged 获取当前登录用户文章
func (p *Service) GetUserArticlesPaged(c context.Context, cate string, limit, offset int) (resp *model.AccountArticlesResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	resp = &model.AccountArticlesResp{
		Items:  make([]*model.ArticleItem, 0),
		Paging: &model.Paging{},
	}

	switch cate {
	case model.CateCreated:
		var data *article.IDsResp
		if data, err = p.d.GetUserArticleIDsPaged(c, aid, limit, offset); err != nil {
			return
		}
		if data.IDs == nil || len(data.IDs) == 0 {
			return
		}

		for _, v := range data.IDs {
			item := &model.ArticleItem{}
			var t *article.ArticleInfo
			if t, err = p.d.GetArticle(c, v); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				tg := p.fromArticle(t)
				tg.ChangeDesc = ""
				item.Target = tg
			}

			resp.Items = append(resp.Items, item)
		}
		break

	case model.CateFaved:
		var data *fav.FavsResp
		if data, err = p.d.GetUserFavsPaged(c, aid, model.TargetTypeArticle, int32(limit), int32(offset)); err != nil {
			return
		}
		if data.Items == nil || len(data.Items) == 0 {
			return
		}

		for _, v := range data.Items {
			item := &model.ArticleItem{}
			var t *article.ArticleInfo
			if t, err = p.d.GetArticle(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				tg := p.fromArticle(t)
				tg.ChangeDesc = ""
				item.Target = tg
			}

			resp.Items = append(resp.Items, item)
		}
		break
	case model.CateRecentViewed:
		var data *recent.RecentViewsResp
		if data, err = p.d.GetRecentViewsPaged(c, aid, model.TargetTypeArticle, int32(limit), int32(offset)); err != nil {
			return
		}
		if data.Items == nil || len(data.Items) == 0 {
			return
		}

		for _, v := range data.Items {
			item := &model.ArticleItem{}
			var t *article.ArticleInfo
			if t, err = p.d.GetArticle(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				tg := p.fromArticle(t)
				tg.ChangeDesc = ""
				item.Target = tg
			}

			resp.Items = append(resp.Items, item)
		}
		break
	}

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/my_articles", url.Values{
		"cate":   []string{cate},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/my_articles", url.Values{
		"cate":   []string{cate},
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

// GetUserArticlesPaged 获取当前登录用户补充
func (p *Service) GetUserRevisesPaged(c context.Context, cate string, limit, offset int) (resp *model.AccountRevisesResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	resp = &model.AccountRevisesResp{
		Items:  make([]*model.ReviseItem, 0),
		Paging: &model.Paging{},
	}

	switch cate {
	case model.CateCreated:
		var data *article.IDsResp
		if data, err = p.d.GetUserReviseIDsPaged(c, aid, limit, offset); err != nil {
			return
		}
		if data.IDs == nil || len(data.IDs) == 0 {
			return
		}

		for _, v := range data.IDs {
			item := &model.ReviseItem{}
			var t *article.ReviseInfo
			if t, err = p.d.GetRevise(c, v); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				item.Target = p.fromRevise(t)
			}

			resp.Items = append(resp.Items, item)
		}
		break

	case model.CateFaved:
		var data *fav.FavsResp
		if data, err = p.d.GetUserFavsPaged(c, aid, model.TargetTypeRevise, int32(limit), int32(offset)); err != nil {
			return
		}
		if data.Items == nil || len(data.Items) == 0 {
			return
		}

		for _, v := range data.Items {
			item := &model.ReviseItem{}
			var t *article.ReviseInfo
			if t, err = p.d.GetRevise(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				item.Target = p.fromRevise(t)
			}

			resp.Items = append(resp.Items, item)
		}
		break
	}

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/my_revises", url.Values{
		"cate":   []string{cate},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/my_revises", url.Values{
		"cate":   []string{cate},
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

// GetUserDiscussionsPaged 获取当前登录用户讨论信息
func (p *Service) GetUserDiscussionsPaged(c context.Context, cate string, limit, offset int) (resp *model.AccountDiscussionsResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	resp = &model.AccountDiscussionsResp{
		Items:  make([]*model.DiscussItem, 0),
		Paging: &model.Paging{},
	}

	switch cate {
	case model.CateCreated:
		var data *discuss.IDsResp
		if data, err = p.d.GetUserDiscussionIDsPaged(c, aid, limit, offset); err != nil {
			return
		}
		if data.IDs == nil || len(data.IDs) == 0 {
			return
		}

		for _, v := range data.IDs {
			item := &model.DiscussItem{}
			var t *discuss.DiscussionInfo
			if t, err = p.d.GetDiscussion(c, v); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				item.Target = p.fromDiscussion(t)
			}

			resp.Items = append(resp.Items, item)
		}
		break

	case model.CateFaved:
		var data *fav.FavsResp
		if data, err = p.d.GetUserFavsPaged(c, aid, model.TargetTypeDiscussion, int32(limit), int32(offset)); err != nil {
			return
		}
		if data.Items == nil || len(data.Items) == 0 {
			return
		}

		for _, v := range data.Items {
			item := &model.DiscussItem{}
			var t *discuss.DiscussionInfo
			if t, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
				if ecode.IsNotExistEcode(err) {
					item.Deleted = true
				} else {
					return
				}
			} else {
				item.Target = p.fromDiscussion(t)
			}

			resp.Items = append(resp.Items, item)
		}
		break
	}

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/my_discussions", url.Values{
		"cate":   []string{cate},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/account/list/my_discussions", url.Values{
		"cate":   []string{cate},
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
