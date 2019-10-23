package service

import (
	"context"
	"net/url"
	"strconv"
	"valerian/app/interface/recent/model"
	article "valerian/app/service/article/api"
	recent "valerian/app/service/recent/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) FromArticle(v *article.ArticleInfo) (item *model.ItemArticle) {
	item = &model.ItemArticle{
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

func (p *Service) FromTopic(v *topic.TopicInfo) (item *model.ItemTopic) {
	item = &model.ItemTopic{
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

func (p *Service) GetMemberRecentViewsPaged(c context.Context, atype string, limit, offset int) (resp *model.RecentListResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var data *recent.RecentViewsResp
	if data, err = p.d.GetRecentViewsPaged(c, aid, atype, limit, offset); err != nil {
		return
	}

	resp = &model.RecentListResp{
		Items:  make([]*model.RecentItem, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		item := &model.RecentItem{
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
		case model.TargetTypeTopic:
			var topic *topic.TopicInfo
			if topic, err = p.d.GetTopic(c, v.TargetID); err != nil {
				return
			}

			item.Topic = p.FromTopic(topic)
			break
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/list/recent", url.Values{
		"type":   []string{atype},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/list/recent", url.Values{
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
