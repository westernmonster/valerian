package service

import (
	"context"
	"valerian/app/interface/topic/model"
	account "valerian/app/service/account/api"
	discuss "valerian/app/service/discuss/api"
	feed "valerian/app/service/feed/api"
	"valerian/library/ecode"
	"valerian/library/xstr"
)

func (p *Service) GetTopicFeedPaged(c context.Context, topicID int64, limit, offset int) (resp *model.FeedResp, err error) {
	var data *feed.TopicFeedResp
	if data, err = p.d.GetTopicFeedPaged(c, topicID, limit, offset); err != nil {
		return
	}

	resp = &model.FeedResp{
		Items:  make([]*model.FeedItem, len(data.Items)),
		Paging: &model.FeedPaging{},
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

		var account *account.BaseInfoReply
		if account, err = p.d.GetAccountBaseInfo(c, v.ActorID); err != nil {
			return
		} else if account == nil {
			return nil, ecode.UserNotExist
		}

		item.Source.Actor = &model.Actor{
			ID:     account.ID,
			Avatar: account.Avatar,
			Name:   account.UserName,
		}

		intro := account.GetIntroductionValue()
		item.Source.Actor.Introduction = &intro

		if v.TargetType == model.TargetTypeDiscussion {
			var discuss *discuss.DiscussionInfo
			if discuss, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
				return
			}

			item.Target.Discussion = &model.TargetDiscussion{
				ID:           discuss.ID,
				Images:       make([]string, 0),
				Excerpt:      xstr.Excerpt(discuss.ContentText),
				LikeCount:    int(discuss.Stat.LikeCount),
				CommentCount: int(discuss.Stat.CommentCount),
			}

			title := discuss.GetTitleValue()
			item.Target.Discussion.Title = &title
		}

		resp.Items[i] = item
	}

	return
}
