package service

import (
	"context"
	"valerian/app/interface/dm/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
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

func (p *Service) GetUserMessagesPaged(c context.Context, atype string, limit, offset int) (resp *model.NotificationResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data []*model.Message
	if data, err = p.d.GetUserMessagesPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	var stat *model.MessageStat
	if stat, err = p.d.GetMessageStatByID(c, p.d.DB(), aid); err != nil {
		return
	}

	resp = &model.NotificationResp{
		Items:       make([]*model.MessageItem, len(data)),
		Paging:      &model.Paging{},
		UnreadCount: stat.UnreadCount,
	}

	for i, v := range data {
		item := &model.MessageItem{
			ID:         v.ID,
			Verb:       v.ActionText,
			IsRead:     bool(v.IsRead),
			MergeCount: v.MergeCount,
			Type:       v.ActionType,
			CreatedAt:  v.CreatedAt,
			TargetType: v.TargetType,
		}
		item.Content.Actors = make([]*model.Actor, 0)
		item.Content.Target.ID = v.TargetID
		item.Content.Target.Type = v.TargetType

		switch v.TargetType {
		case model.TargetTypeArticle:
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, v.TargetID); err != nil {
				return
			}

			item.Target = p.FromArticle(article)
			break
		case model.TargetTypeRevise:
			var revise *article.ReviseInfo
			if revise, err = p.d.GetRevise(c, v.TargetID); err != nil {
				return
			}

			item.Target = p.FromRevise(revise)
			break
		case model.TargetTypeDiscussion:
			var discuss *discuss.DiscussionInfo
			if discuss, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
				return
			}

			item.Target = p.FromDiscussion(discuss)
			break

		case model.TargetTypeTopic:
			var topic *topic.TopicInfo
			if topic, err = p.d.GetTopic(c, v.TargetID); err != nil {
				return
			}

			item.Target = p.FromTopic(topic)
			break

		case model.TargetTypeMember:
			var info *model.MemberInfo
			if info, err = p.GetMemberInfo(c, v.TargetID); err != nil {
				return
			}

			item.Target = info
			break
		}

		resp.Items[i] = item
	}

	return
}

func (p *Service) GetMemberInfo(c context.Context, targetID int64) (resp *model.MemberInfo, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var f *account.ProfileReply
	if f, err = p.d.GetAccountProfile(c, targetID); err != nil {
		return
	}

	resp = &model.MemberInfo{
		ID:       f.ID,
		UserName: f.UserName,
		Avatar:   f.Avatar,
		IDCert:   f.IDCert,
		WorkCert: f.WorkCert,
		IsOrg:    f.IsOrg,
		IsVIP:    f.IsVIP,
	}

	if f.Gender != nil {
		g := int(f.GetGenderValue())
		resp.Gender = &g
	}

	if f.Location != nil {
		loc := f.GetLocationValue()
		resp.Location = &loc

		if f.LocationString != nil {
			locString := f.GetLocationStringValue()
			resp.LocationString = &locString
		}
	}

	if f.Introduction != nil {
		intro := f.GetIntroductionValue()
		resp.Introduction = &intro
	}

	var isFollowing bool
	if isFollowing, err = p.d.IsFollowing(c, aid, targetID); err != nil {
		return
	}

	resp.Stat = &model.MemberInfoStat{
		FansCount:       int(f.Stat.Fans),
		FollowingCount:  int(f.Stat.Following),
		TopicCount:      int(f.Stat.TopicCount),
		ArticleCount:    int(f.Stat.ArticleCount),
		DiscussionCount: int(f.Stat.DiscussionCount),
		IsFollow:        isFollowing,
	}
	return
}
