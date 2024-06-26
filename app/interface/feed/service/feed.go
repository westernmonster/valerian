package service

import (
	"context"
	"net/url"
	"strconv"
	"valerian/app/interface/feed/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	feed "valerian/app/service/feed/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) FromDiscussion(v *discuss.DiscussionInfo) (item *model.TargetDiscuss) {
	item = &model.TargetDiscuss{
		ID:           v.ID,
		Excerpt:      v.Excerpt,
		ImageUrls:    v.ImageUrls,
		CommentCount: int(v.Stat.CommentCount),
		LikeCount:    int(v.Stat.LikeCount),
		DislikeCount: int(v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		Title: v.Title,
	}

	return
}

func (p *Service) FromRevise(v *article.ReviseInfo) (item *model.TargetRevise) {
	item = &model.TargetRevise{
		ID:           v.ID,
		Title:        v.Title,
		Excerpt:      v.Excerpt,
		ImageUrls:    v.ImageUrls,
		CommentCount: int(v.Stat.CommentCount),
		LikeCount:    int(v.Stat.LikeCount),
		DislikeCount: int(v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
	}

	return
}

func (p *Service) FromArticle(v *article.ArticleInfo) (item *model.TargetArticle) {
	item = &model.TargetArticle{
		ID:           v.ID,
		Title:        v.Title,
		Excerpt:      v.Excerpt,
		ImageUrls:    v.ImageUrls,
		ReviseCount:  int(v.Stat.ReviseCount),
		CommentCount: int(v.Stat.CommentCount),
		LikeCount:    int(v.Stat.LikeCount),
		DislikeCount: int(v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		ChangeDesc: v.ChangeDesc,
	}

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
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		Avatar: v.Avatar,
	}

	return
}

func (p *Service) FromComment(v *comment.CommentInfo) (item *model.TargetComment) {
	item = &model.TargetComment{
		ID:            v.ID,
		Type:          v.TargetType,
		Excerpt:       v.Content,
		CreatedAt:     v.CreatedAt,
		ResourceID:    v.ResourceID,
		ChildrenCount: v.Stat.ChildrenCount,
		LikeCount:     v.Stat.LikeCount,
	}

	return
}

func (p *Service) GetFeedPaged(c context.Context, limit, offset int) (resp *model.FeedResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var data *feed.FeedResp
	if data, err = p.d.GetFeedPaged(c, aid, limit, offset); err != nil {
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
		if acc, err = p.d.GetAccountBaseInfo(c, v.ActorID); err != nil {
			return
		} else if acc == nil {
			return nil, ecode.UserNotExist
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

			item.Target.Article = p.FromArticle(article)
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

			item.Target.Article = p.FromArticle(article)
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

			item.Target.Revise = p.FromRevise(revise)
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

			item.Target.Discussion = p.FromDiscussion(discuss)
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

			item.Target.Topic = p.FromTopic(topic)
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

			item.Target.Comment = p.FromComment(info)

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

				item.Target.Article = p.FromArticle(article)
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

				item.Target.Revise = p.FromRevise(revise)
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

				item.Target.Discussion = p.FromDiscussion(discuss)
				break

			}
			break

		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/list/activities", url.Values{
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/list/activities", url.Values{
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

	var f *account.MemberInfoReply
	if f, err = p.d.GetMemberInfo(c, targetID); err != nil {
		return
	}

	resp = &model.MemberInfo{
		ID:             f.ID,
		UserName:       f.UserName,
		Avatar:         f.Avatar,
		IDCert:         f.IDCert,
		WorkCert:       f.WorkCert,
		IsOrg:          f.IsOrg,
		IsVIP:          f.IsVIP,
		Gender:         int(f.Gender),
		Introduction:   f.Introduction,
		Location:       f.Location,
		LocationString: f.LocationString,
	}

	var isFollowing bool
	if isFollowing, err = p.d.IsFollowing(c, aid, targetID); err != nil {
		return
	}

	resp.Stat = &model.MemberInfoStat{
		FansCount:       int(f.Stat.FansCount),
		FollowingCount:  int(f.Stat.FollowingCount),
		TopicCount:      int(f.Stat.TopicCount),
		ArticleCount:    int(f.Stat.ArticleCount),
		DiscussionCount: int(f.Stat.DiscussionCount),
		IsFollow:        isFollowing,
	}
	return
}
