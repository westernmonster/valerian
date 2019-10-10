package service

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"valerian/app/interface/fav/model"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/gid"
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
			ID:       v.Creator.ID,
			Avatar:   v.Creator.Avatar,
			UserName: v.Creator.UserName,
		},
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
		ImageUrls:    v.ImageUrls,
		CommentCount: int(v.Stat.CommentCount),
		LikeCount:    int(v.Stat.LikeCount),
		DislikeCount: int(v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:       v.Creator.ID,
			Avatar:   v.Creator.Avatar,
			UserName: v.Creator.UserName,
		},
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
		ImageUrls:    v.ImageUrls,
		ReviseCount:  int(v.Stat.ReviseCount),
		CommentCount: int(v.Stat.CommentCount),
		LikeCount:    int(v.Stat.LikeCount),
		DislikeCount: int(v.Stat.DislikeCount),
		Creator: &model.Creator{
			ID:       v.Creator.ID,
			Avatar:   v.Creator.Avatar,
			UserName: v.Creator.UserName,
		},
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
	}

	intro := v.Creator.GetIntroductionValue()
	item.Creator.Introduction = &intro

	avatar := v.GetAvatarValue()
	item.Avatar = &avatar
	return
}

func (p *Service) Fav(c context.Context, arg *model.ArgAddFav) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": arg.TargetType,
		"target_id":   arg.TargetID,
	}); err != nil {
		return
	} else if fav != nil {
		return
	}

	switch arg.TargetType {

	case model.TargetTypeArticle:
		if _, err = p.d.GetArticle(c, arg.TargetID); err != nil {
			return
		}
		break
	case model.TargetTypeRevise:
		if _, err = p.d.GetRevise(c, arg.TargetID); err != nil {
			return
		}

		break
	case model.TargetTypeDiscussion:
		if _, err = p.d.GetDiscussion(c, arg.TargetID); err != nil {
			return
		}

		break

	case model.TargetTypeTopic:
		if _, err = p.d.GetTopic(c, arg.TargetID); err != nil {
			return
		}

		break

	}

	fav = &model.Fav{
		ID:         gid.NewID(),
		AccountID:  aid,
		TargetID:   arg.TargetID,
		TargetType: arg.TargetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddFav(c, p.d.DB(), fav); err != nil {
		return
	}

	return
}

func (p *Service) Unfav(c context.Context, arg *model.ArgDelFav) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": arg.TargetType,
		"target_id":   arg.TargetID,
	}); err != nil {
		return
	} else if fav == nil {
		return
	} else {
		if err = p.d.DelFav(c, p.d.DB(), fav.ID); err != nil {
			return
		}
	}

	return
}

func (p *Service) GetFavsPaged(c context.Context, targetType string, limit, offset int) (resp *model.FavListResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data []*model.Fav
	if data, err = p.d.GetFavsPaged(c, p.d.DB(), aid, targetType, limit, offset); err != nil {
		return
	}

	resp = &model.FavListResp{
		Items:  make([]*model.FavItem, len(data)),
		Paging: &model.Paging{},
	}

	for i, v := range data {
		item := &model.FavItem{
			ID:         v.ID,
			TargetType: v.TargetType,
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

	param := url.Values{}
	param.Set("target_type", targetType)
	param.Set("limit", strconv.Itoa(limit))
	param.Set("offset", strconv.Itoa(offset-limit))

	if resp.Paging.Prev, err = genURL("/api/v1/fav/list/all", param); err != nil {
		return
	}
	param.Set("offset", strconv.Itoa(offset+limit))
	if resp.Paging.Next, err = genURL("/api/v1/fav/list/all", param); err != nil {
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
