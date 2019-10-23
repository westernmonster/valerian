package service

import (
	"context"
	"net/url"
	"strconv"
	"valerian/app/interface/topic/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	topicFeed "valerian/app/service/topic-feed/api"
	"valerian/library/ecode"
)

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
	}

	return
}

func (p *Service) GetTopicFeedPaged(c context.Context, topicID int64, limit, offset int) (resp *model.FeedResp, err error) {
	var data *topicFeed.TopicFeedResp
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
			ID:           account.ID,
			Avatar:       account.Avatar,
			Name:         account.UserName,
			Introduction: account.Introduction,
		}

		if v.TargetType == model.TargetTypeArticle {
			var article *article.ArticleInfo
			if article, err = p.d.GetArticle(c, v.TargetID); err != nil {
				return
			}

			item.Target.Article = p.FromArticle(article)
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/topic/list/activities", url.Values{
		"topic_id": []string{strconv.FormatInt(topicID, 10)},
		"limit":    []string{strconv.Itoa(limit)},
		"offset":   []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/topic/list/activities", url.Values{
		"topic_id": []string{strconv.FormatInt(topicID, 10)},
		"limit":    []string{strconv.Itoa(limit)},
		"offset":   []string{strconv.Itoa(offset + limit)},
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
