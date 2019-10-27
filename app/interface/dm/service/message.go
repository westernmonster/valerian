package service

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"valerian/app/interface/dm/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
	"valerian/library/xstr"
)

func (p *Service) FromDiscussion(v *discuss.DiscussionInfo) (item *model.TargetDiscuss) {
	item = &model.TargetDiscuss{
		ID:           v.ID,
		Excerpt:      v.Excerpt,
		CommentCount: int(v.Stat.CommentCount),
		LikeCount:    int(v.Stat.LikeCount),
		DislikeCount: int(v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Title:     v.Title,
	}

	if v.ImageUrls == nil {
		item.ImageUrls = make([]string, 0)
	} else {
		item.ImageUrls = v.ImageUrls
	}

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
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
	if v.ImageUrls == nil {
		item.ImageUrls = make([]string, 0)
	} else {
		item.ImageUrls = v.ImageUrls
	}

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
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
	if v.ImageUrls == nil {
		item.ImageUrls = make([]string, 0)
	} else {
		item.ImageUrls = v.ImageUrls
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
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Avatar:    v.Avatar,
	}

	return
}

func (p *Service) FromAccount(v *account.BaseInfoReply) (item *model.Actor) {
	item = &model.Actor{
		ID:     v.ID,
		Avatar: v.Avatar,
		Name:   v.UserName,
	}

	return
}

func (p *Service) GetActors(c context.Context, actorsFields string) (items []*model.Actor, err error) {
	actorsIds := strings.Split(actorsFields, ",")
	items = make([]*model.Actor, 0)
	for _, v := range actorsIds {
		var id int64
		if id, err = strconv.ParseInt(v, 10, 64); err != nil {
			return
		}

		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, id); err != nil {
			return
		}

		items = append(items, p.FromAccount(acc))
	}

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

		var actors []*model.Actor
		if actors, err = p.GetActors(c, v.Actors); err != nil {
			return
		}

		item.Content.Actors = actors

		switch v.TargetType {
		case model.TargetTypeArticle:
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, v.TargetID); err != nil {
				return
			}

			item.Content.Target.Link = fmt.Sprintf("pronote://article/%d", article.ID)
			item.Content.Target.Text = article.Title

			item.Target = p.FromArticle(article)
			break
		case model.TargetTypeRevise:
			var revise *article.ReviseInfo
			if revise, err = p.d.GetRevise(c, v.TargetID); err != nil {
				return
			}

			item.Content.Target.Link = fmt.Sprintf("pronote://revise/%d", revise.ID)
			item.Content.Target.Text = revise.Title
			item.Target = p.FromRevise(revise)
			break
		case model.TargetTypeDiscussion:
			var discuss *discuss.DiscussionInfo
			if discuss, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
				return
			}

			item.Content.Target.Link = fmt.Sprintf("pronote://discuss/%d", discuss.ID)
			item.Content.Target.Text = discuss.Title
			item.Target = p.FromDiscussion(discuss)
			break

		case model.TargetTypeTopic:
			var topic *topic.TopicInfo
			if topic, err = p.d.GetTopic(c, v.TargetID); err != nil {
				return
			}

			item.Content.Target.Link = fmt.Sprintf("pronote://topic/%d", topic.ID)
			item.Content.Target.Text = topic.Name
			item.Target = p.FromTopic(topic)
			break

		case model.TargetTypeMember:
			var info *model.MemberInfo
			if info, err = p.GetMemberInfo(c, v.TargetID); err != nil {
				return
			}

			item.Content.Target.Link = fmt.Sprintf("pronote://user/%d", info.ID)
			item.Content.Target.Text = info.UserName
			item.Target = info
			break

		case model.TargetTypeComment:
			var ct *comment.CommentInfo
			if ct, err = p.d.GetComment(c, v.TargetID); err != nil {
				return
			}

			switch ct.TargetType {
			case model.TargetTypeArticle:
				if item.Target, err = p.GetCommentTarget(c, ct); err != nil {
					return
				}

				item.Content.Target.Link = fmt.Sprintf("pronote://%s/%d/comment/%d", model.TargetTypeArticle, ct.OwnerID, ct.ID)
				item.Content.Target.Text = xstr.Excerpt(ct.Content)
				break
			case model.TargetTypeRevise:
				if item.Target, err = p.GetCommentTarget(c, ct); err != nil {
					return
				}
				item.Content.Target.Link = fmt.Sprintf("pronote://%s/%d/comment/%d", model.TargetTypeRevise, ct.OwnerID, ct.ID)
				item.Content.Target.Text = xstr.Excerpt(ct.Content)
				break
			case model.TargetTypeDiscussion:
				if item.Target, err = p.GetCommentTarget(c, ct); err != nil {
					return
				}
				item.Content.Target.Link = fmt.Sprintf("pronote://%s/%d/comment/%d", model.TargetTypeDiscussion, ct.OwnerID, ct.ID)
				item.Content.Target.Text = xstr.Excerpt(ct.Content)
				break
			case model.TargetTypeComment:
				var tg *model.TargetComment
				if tg, err = p.GetCommentTarget(c, ct); err != nil {
					return
				}
				item.Target = tg

				switch tg.OwnerType {
				case model.TargetTypeArticle:
					owner := tg.Owner.(*model.TargetArticle)
					item.Content.Target.Link = fmt.Sprintf("pronote://{%s}/{%d}/comment/%d/sub/%d",
						tg.OwnerType,
						owner.ID,
						tg.ParentComment.ID,
						tg.ID)
					item.Content.Target.Text = xstr.Excerpt(ct.Content)
					break
				case model.TargetTypeRevise:
					owner := tg.Owner.(*model.TargetRevise)
					item.Content.Target.Link = fmt.Sprintf("pronote://{%s}/{%d}/comment/%d/sub/%d",
						tg.OwnerType,
						owner.ID,
						tg.ParentComment.ID,
						tg.ID)
					break
				case model.TargetTypeDiscussion:
					owner := tg.Owner.(*model.TargetDiscuss)
					item.Content.Target.Link = fmt.Sprintf("pronote://{%s}/{%d}/comment/%d/sub/%d",
						tg.OwnerType,
						owner.ID,
						tg.ParentComment.ID,
						tg.ID)
					break

				}
				break
			}
			break

		case model.TargetTypeTopicFollowRequest:
			var req *model.TopicFollowRequest
			if req, err = p.d.GetTopicFollowRequestByID(c, p.d.DB(), v.TargetID); err != nil {
				return
			} else if req == nil {
				err = ecode.TopicFollowRequestNotExist
				return
			}

			tg := &model.TargetTopicFollowRequest{
				ID:        req.ID,
				Status:    req.Status,
				Reason:    req.Reason,
				CreatedAt: req.CreatedAt,
			}

			var acc *account.BaseInfoReply
			if acc, err = p.d.GetAccountBaseInfo(c, req.AccountID); err != nil {
				return
			}

			tg.Member = &model.Creator{
				ID:           acc.ID,
				Avatar:       acc.Avatar,
				UserName:     acc.UserName,
				Introduction: acc.Introduction,
			}

			var topic *topic.TopicInfo
			if topic, err = p.d.GetTopic(c, req.TopicID); err != nil {
				return
			}

			tg.Topic = p.FromTopic(topic)

			item.Target = tg
			break
		case model.TargetTypeTopicInviteRequest:
			var req *model.TopicInviteRequest
			if req, err = p.d.GetTopicInviteRequestByID(c, p.d.DB(), v.TargetID); err != nil {
				return
			} else if req == nil {
				err = ecode.TopicInviteRequestNotExist
				return
			}

			tg := &model.TargetTopicInviteRequest{
				ID:        req.ID,
				Status:    req.Status,
				CreatedAt: req.CreatedAt,
			}

			var acc *account.BaseInfoReply
			if acc, err = p.d.GetAccountBaseInfo(c, req.FromAccountID); err != nil {
				return
			}

			tg.Creator = &model.Creator{
				ID:           acc.ID,
				Avatar:       acc.Avatar,
				UserName:     acc.UserName,
				Introduction: acc.Introduction,
			}

			var topic *topic.TopicInfo
			if topic, err = p.d.GetTopic(c, req.TopicID); err != nil {
				return
			}

			tg.Topic = p.FromTopic(topic)

			item.Target = tg
			break

		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/dm/list", url.Values{
		"type":   []string{atype},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/dm/list", url.Values{
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

func (p *Service) GetCommentTarget(c context.Context, v *comment.CommentInfo) (resp *model.TargetComment, err error) {
	resp = &model.TargetComment{
		ID:         v.ID,
		Excerpt:    xstr.Excerpt(v.Content),
		TargetType: v.TargetType,
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		CreatedAt: v.CreatedAt,
	}

	switch v.TargetType {
	case model.TargetTypeArticle:
		var article *article.ArticleInfo
		if article, err = p.d.GetArticle(c, v.OwnerID); err != nil {
			return
		}
		resp.Owner = p.FromArticle(article)
		resp.OwnerType = model.TargetTypeArticle
		break
	case model.TargetTypeRevise:
		var item *article.ReviseInfo
		if item, err = p.d.GetRevise(c, v.OwnerID); err != nil {
			return
		}
		resp.Owner = p.FromRevise(item)
		resp.OwnerType = model.TargetTypeRevise
		break
	case model.TargetTypeDiscussion:
		var item *discuss.DiscussionInfo
		if item, err = p.d.GetDiscussion(c, v.OwnerID); err != nil {
			return
		}
		resp.Owner = p.FromDiscussion(item)
		resp.OwnerType = model.TargetTypeDiscussion
		break
	case model.TargetTypeComment:
		var replyComment *comment.CommentInfo
		if replyComment, err = p.d.GetComment(c, v.ResourceID); err != nil {
			return
		}

		resp.ParentComment = &model.TargetReplyComment{
			ID:      replyComment.ID,
			Excerpt: xstr.Excerpt(replyComment.Content),
			Creator: &model.Creator{
				ID:           replyComment.Creator.ID,
				Avatar:       replyComment.Creator.Avatar,
				UserName:     replyComment.Creator.UserName,
				Introduction: replyComment.Creator.Introduction,
			},
			CreatedAt:  replyComment.CreatedAt,
			TargetType: replyComment.TargetType,
		}

		if v.ReplyTo != 0 {
			var acc *account.BaseInfoReply
			if acc, err = p.d.GetAccountBaseInfo(c, v.ReplyTo); err != nil {
				return
			}

			resp.ReplyTo = &model.Creator{
				ID:           acc.ID,
				Avatar:       acc.Avatar,
				UserName:     acc.UserName,
				Introduction: acc.Introduction,
			}

		}
		switch replyComment.TargetType {
		case model.TargetTypeArticle:
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, replyComment.OwnerID); err != nil {
				return
			}
			resp.Owner = p.FromArticle(article)
			resp.OwnerType = model.TargetTypeArticle
			break
		case model.TargetTypeRevise:
			var item *article.ReviseInfo
			if item, err = p.d.GetRevise(c, replyComment.OwnerID); err != nil {
				return
			}
			resp.Owner = p.FromRevise(item)
			resp.OwnerType = model.TargetTypeRevise
			break
		case model.TargetTypeDiscussion:
			var item *discuss.DiscussionInfo
			if item, err = p.d.GetDiscussion(c, replyComment.OwnerID); err != nil {
				return
			}
			resp.Owner = p.FromDiscussion(item)
			resp.OwnerType = model.TargetTypeDiscussion
			break
		}

		break
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
		ID:           f.ID,
		UserName:     f.UserName,
		Avatar:       f.Avatar,
		IDCert:       f.IDCert,
		WorkCert:     f.WorkCert,
		IsOrg:        f.IsOrg,
		IsVIP:        f.IsVIP,
		Introduction: f.Introduction,
		Gender:       int(f.Gender),
	}

	if f.Location != 0 {
		resp.Location = f.Location

		if f.LocationString != "" {
			locString := f.LocationString
			resp.LocationString = locString
		}
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
