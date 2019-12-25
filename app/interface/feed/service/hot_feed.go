package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/feed/model"
)

func (p *Service) GetHotFeedPaged(c context.Context, limit, offset int) (resp *model.HotFeedResp, err error) {
	// aid, ok := metadata.Value(c, metadata.Aid).(int64)
	// if !ok {
	// 	err = ecode.AcquireAccountIDFailed
	// 	return
	// }
	// var data *feed.FeedResp
	// if data, err = p.d.GetFeedPaged(c, aid, limit, offset); err != nil {
	// 	return
	// }

	resp = &model.HotFeedResp{
		Items:  make([]*model.HotFeedItem, 0),
		Paging: &model.Paging{},
	}

	// for i, v := range data.Items {
	// 	item := &model.HotFeedItem{
	// 		TargetType: v.TargetType,
	// 		Target: &model.FeedTarget{
	// 			ID:   v.TargetID,
	// 			Type: v.TargetType,
	// 		},
	// 	}

	// 	switch v.TargetType {
	// 	case model.TargetTypeArticle:
	// 		var article *article.ArticleInfo
	// 		if article, err = p.d.GetArticle(c, v.TargetID); err != nil {
	// 			if ecode.IsNotExistEcode(err) {
	// 				item.Deleted = true
	// 				break
	// 			}
	// 			return
	// 		}

	// 		item.Target.Article = p.FromArticle(article)
	// 		break
	// 	case model.TargetTypeRevise:
	// 		var revise *article.ReviseInfo
	// 		if revise, err = p.d.GetRevise(c, v.TargetID); err != nil {
	// 			if ecode.IsNotExistEcode(err) {
	// 				item.Deleted = true
	// 				break
	// 			}
	// 			return
	// 		}

	// 		item.Target.Revise = p.FromRevise(revise)
	// 		break
	// 	case model.TargetTypeDiscussion:
	// 		var discuss *discuss.DiscussionInfo
	// 		if discuss, err = p.d.GetDiscussion(c, v.TargetID); err != nil {
	// 			if ecode.IsNotExistEcode(err) {
	// 				item.Deleted = true
	// 				break
	// 			}
	// 			return
	// 		}

	// 		item.Target.Discussion = p.FromDiscussion(discuss)
	// 		break

	// 	case model.TargetTypeTopic:
	// 		var topic *topic.TopicInfo
	// 		if topic, err = p.d.GetTopic(c, v.TargetID); err != nil {
	// 			if ecode.IsNotExistEcode(err) {
	// 				item.Deleted = true
	// 				break
	// 			}
	// 			return
	// 		}

	// 		item.Target.Topic = p.FromTopic(topic)
	// 		break

	// 	case model.TargetTypeComment:
	// 		var info *comment.CommentInfo
	// 		if info, err = p.d.GetComment(c, v.TargetID); err != nil {
	// 			if ecode.IsNotExistEcode(err) {
	// 				item.Deleted = true
	// 				break
	// 			}
	// 			return
	// 		}

	// 		item.Target.Comment = p.FromComment(info)

	// 		switch info.TargetType {
	// 		case model.TargetTypeArticle:
	// 			var article *article.ArticleInfo
	// 			if article, err = p.d.GetArticle(c, info.ResourceID); err != nil {
	// 				if ecode.IsNotExistEcode(err) {
	// 					item.Deleted = true
	// 					break
	// 				}
	// 				return
	// 			}

	// 			item.Target.Article = p.FromArticle(article)
	// 			break
	// 		case model.TargetTypeRevise:
	// 			var revise *article.ReviseInfo
	// 			if revise, err = p.d.GetRevise(c, info.ResourceID); err != nil {
	// 				if ecode.IsNotExistEcode(err) {
	// 					item.Deleted = true
	// 					break
	// 				}
	// 				return
	// 			}

	// 			item.Target.Revise = p.FromRevise(revise)
	// 			break
	// 		case model.TargetTypeDiscussion:
	// 			var discuss *discuss.DiscussionInfo
	// 			if discuss, err = p.d.GetDiscussion(c, info.ResourceID); err != nil {
	// 				if ecode.IsNotExistEcode(err) {
	// 					item.Deleted = true
	// 					break
	// 				}
	// 				return
	// 			}

	// 			item.Target.Discussion = p.FromDiscussion(discuss)
	// 			break

	// 		}
	// 		break

	// 	}

	// 	resp.Items[i] = item
	// }

	if resp.Paging.Prev, err = genURL("/api/v1/list/hot", url.Values{
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/list/hot", url.Values{
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
