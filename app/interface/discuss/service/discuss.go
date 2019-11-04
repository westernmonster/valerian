package service

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"valerian/app/interface/discuss/model"
	account "valerian/app/service/account/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
	"valerian/library/xstr"

	"github.com/PuerkitoBio/goquery"
)

func (p *Service) initDiscussionStat(c context.Context, discussionID int64) (err error) {
	if err = p.d.AddDiscussionStat(c, p.d.DB(), &model.DiscussionStat{
		DiscussionID: discussionID,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}); err != nil {
		return
	}

	return
}

func (p *Service) GetUserDiscussionsPaged(c context.Context, aid int64, limit, offset int) (resp *model.DiscussListResp, err error) {
	var data []*model.Discussion
	if data, err = p.d.GetUserDiscussionsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	resp = &model.DiscussListResp{
		Items:  make([]*model.DiscussItem, len(data)),
		Paging: &model.Paging{},
	}

	for i, v := range data {
		item := &model.DiscussItem{
			ID:        v.ID,
			TopicID:   v.TopicID,
			Title:     v.Title,
			Excerpt:   xstr.Excerpt(v.ContentText),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		var stat *model.DiscussionStat
		if stat, err = p.d.GetDiscussionStatByID(c, p.d.DB(), v.ID); err != nil {
			return
		}

		item.LikeCount = stat.LikeCount
		item.DislikeCount = stat.DislikeCount
		item.CommentCount = stat.CommentCount

		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, v.CreatedBy); err != nil {
			return
		}
		item.Creator = &model.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		if v.CategoryID == -1 {
			item.Category = &model.DiscussItemCategory{
				ID:   v.CategoryID,
				Name: "问答",
			}
		} else {
			var category *model.DiscussCategory
			if category, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
				return
			} else if category == nil {
				err = ecode.DiscussCategoryNotExist
				return
			}

			item.Category = &model.DiscussItemCategory{
				ID:   v.CategoryID,
				Name: category.Name,
			}

		}

		if item.ImageUrls, err = p.GetDiscussionImageUrls(c, v.ID); err != nil {
			return
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/discussion/list/by_account", url.Values{
		"account_id": []string{strconv.FormatInt(aid, 10)},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/discussion/list/by_account", url.Values{
		"account_id": []string{strconv.FormatInt(aid, 10)},
		"limit":      []string{strconv.Itoa(limit)},
		"offset":     []string{strconv.Itoa(offset + limit)},
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

func (p *Service) GetTopicDiscussionsPaged(c context.Context, topicID, categoryID int64, limit, offset int) (resp *model.DiscussListResp, err error) {
	var data []*model.Discussion
	if data, err = p.d.GetTopicDiscussionsPaged(c, p.d.DB(), topicID, categoryID, limit, offset); err != nil {
		return
	}

	resp = &model.DiscussListResp{
		Items:  make([]*model.DiscussItem, len(data)),
		Paging: &model.Paging{},
	}

	for i, v := range data {
		item := &model.DiscussItem{
			ID:        v.ID,
			TopicID:   v.TopicID,
			Title:     v.Title,
			Excerpt:   xstr.Excerpt(v.ContentText),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		var stat *model.DiscussionStat
		if stat, err = p.d.GetDiscussionStatByID(c, p.d.DB(), v.ID); err != nil {
			return
		}

		item.LikeCount = stat.LikeCount
		item.DislikeCount = stat.DislikeCount
		item.CommentCount = stat.CommentCount

		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, v.CreatedBy); err != nil {
			return
		}
		item.Creator = &model.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		if v.CategoryID == -1 {
			item.Category = &model.DiscussItemCategory{
				ID:   v.CategoryID,
				Name: "问答",
			}
		} else {
			var category *model.DiscussCategory
			if category, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
				return
			} else if category == nil {
				err = ecode.DiscussCategoryNotExist
				return
			}

			item.Category = &model.DiscussItemCategory{
				ID:   v.CategoryID,
				Name: category.Name,
			}

		}

		if item.ImageUrls, err = p.GetDiscussionImageUrls(c, v.ID); err != nil {
			return
		}

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/discussion/list/by_topic", url.Values{
		"topic_id":    []string{strconv.FormatInt(topicID, 10)},
		"category_id": []string{strconv.FormatInt(categoryID, 10)},
		"limit":       []string{strconv.Itoa(limit)},
		"offset":      []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/discussion/list/by_topic", url.Values{
		"topic_id":    []string{strconv.FormatInt(topicID, 10)},
		"category_id": []string{strconv.FormatInt(categoryID, 10)},
		"limit":       []string{strconv.Itoa(limit)},
		"offset":      []string{strconv.Itoa(offset + limit)},
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

func (p *Service) GetDiscussionImageUrls(c context.Context, discussionID int64) (urls []string, err error) {
	urls = make([]string, 0)
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeDiscussion,
		"target_id":   discussionID,
	}); err != nil {
		return
	}

	for _, v := range imgs {
		urls = append(urls, v.URL)
	}

	return
}

func (p *Service) AddDiscussion(c context.Context, arg *model.ArgAddDiscuss) (id int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.checkCategory(c, p.d.DB(), arg.CategoryID); err != nil {
		return
	}

	// 检测话题
	if _, err = p.d.GetTopic(c, arg.TopicID); err != nil {
		return
	}

	// 检测是否成员
	if err = p.checkTopicMember(c, arg.TopicID, aid); err != nil {
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	item := &model.Discussion{
		ID:         gid.NewID(),
		TopicID:    arg.TopicID,
		CategoryID: arg.CategoryID,
		CreatedBy:  aid,
		Content:    arg.Content,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if arg.Title != nil {
		item.Title = *arg.Title
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		err = ecode.ParseHTMLFailed
		return
	}

	item.ContentText = doc.Text()

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		if url, exist := s.Attr("src"); exist {
			u := &model.ImageURL{
				ID:         gid.NewID(),
				TargetType: model.TargetTypeDiscussion,
				TargetID:   item.ID,
				URL:        url,
				CreatedAt:  time.Now().Unix(),
				UpdatedAt:  time.Now().Unix(),
			}
			if err = p.d.AddImageURL(c, tx, u); err != nil {
				return
			}
		}
	})

	if err = p.d.AddDiscussion(c, tx, item); err != nil {
		return
	}

	if err = p.bulkCreateFiles(c, tx, item.ID, arg.Files); err != nil {
		return
	}

	if err = p.d.AddDiscussionStat(c, tx, &model.DiscussionStat{
		DiscussionID: item.ID,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}); err != nil {
		return
	}

	// Update Stat
	if err = p.d.IncrAccountStat(c, tx, &model.AccountStat{AccountID: aid, DiscussionCount: 1}); err != nil {
		return
	}
	if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{TopicID: arg.TopicID, DiscussionCount: 1}); err != nil {
		return
	}

	if err = p.d.AddDiscussionStat(c, tx, &model.DiscussionStat{
		DiscussionID: item.ID,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	id = item.ID

	p.addCache(func() {
		p.onDiscussionAdded(context.Background(), item.ID, aid, time.Now().Unix())
	})

	return
}

func (p *Service) checkCategory(c context.Context, node sqalx.Node, categoryID int64) (err error) {
	if categoryID == -1 {
		return
	}

	var t *model.DiscussCategory
	if t, err = p.d.GetDiscussCategoryByID(c, node, categoryID); err != nil {
		return
	} else if t == nil {
		return ecode.DiscussCategoryNotExist
	}

	return
}

func (p *Service) UpdateDiscussion(c context.Context, arg *model.ArgUpdateDiscuss) (err error) {

	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var item *model.Discussion
	if item, err = p.d.GetDiscussionByID(c, tx, arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	if item.CreatedBy != aid {
		if err = p.checkTopicManager(c, item.TopicID, aid); err != nil {
			return
		}
	}

	if arg.Title != nil {
		item.Title = *arg.Title
	}
	item.Content = arg.Content
	item.UpdatedAt = time.Now().Unix()

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		err = ecode.ParseHTMLFailed
		return
	}

	item.ContentText = doc.Text()

	if err = p.d.DelImageURLByCond(c, tx, model.TargetTypeDiscussion, item.ID); err != nil {
		return
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		if url, exist := s.Attr("src"); exist {
			u := &model.ImageURL{
				ID:         gid.NewID(),
				TargetType: model.TargetTypeDiscussion,
				TargetID:   item.ID,
				URL:        url,
				CreatedAt:  time.Now().Unix(),
				UpdatedAt:  time.Now().Unix(),
			}
			if err = p.d.AddImageURL(c, tx, u); err != nil {
				return
			}
		}
	})

	if err = p.d.UpdateDiscussion(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.onDiscussionUpdated(context.Background(), arg.ID, aid, time.Now().Unix())
	})

	return
}

func (p *Service) DelDiscussion(c context.Context, id int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var item *model.Discussion
	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), id); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	if item.CreatedBy != aid {
		if err = p.checkTopicManager(c, item.TopicID, aid); err != nil {
			return
		}
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	if err = p.d.DelDiscussion(c, tx, id); err != nil {
		return
	}

	if err = p.d.DelDiscussionFiles(c, tx, id); err != nil {
		return

	}

	if err = p.d.IncrAccountStat(c, tx, &model.AccountStat{AccountID: aid, DiscussionCount: -1}); err != nil {
		return
	}

	if err = p.d.IncrTopicStat(c, tx, &model.TopicStat{TopicID: item.TopicID, DiscussionCount: -1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.onDiscussionDeleted(context.Background(), id, aid, time.Now().Unix())
		p.d.DelDiscussionFilesCache(context.TODO(), id)
	})

	return
}

func (p *Service) GetDiscussion(c context.Context, discussionID int64) (resp *model.DiscussDetailResp, err error) {

	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data *model.Discussion
	if data, err = p.getDiscussion(c, p.d.DB(), discussionID); err != nil {
		return
	}

	resp = &model.DiscussDetailResp{
		ID:        data.ID,
		TopicID:   data.TopicID,
		Title:     data.Title,
		Content:   data.Content,
		Files:     make([]*model.DiscussFileResp, 0),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	if data.CategoryID == -1 {
		resp.Category = &model.DiscussItemCategory{
			ID:   data.CategoryID,
			Name: "问答",
		}
	} else {
		var category *model.DiscussCategory
		if category, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), data.CategoryID); err != nil {
			return
		} else if category == nil {
			err = ecode.DiscussCategoryNotExist
			return
		}

		resp.Category = &model.DiscussItemCategory{
			ID:   data.CategoryID,
			Name: category.Name,
		}

	}
	if acc, e := p.d.GetAccountBaseInfo(c, data.CreatedBy); e != nil {
		err = e
		return
	} else {
		resp.Creator = &model.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}
	}

	var stat *model.DiscussionStat
	if stat, err = p.d.GetDiscussionStatByID(c, p.d.DB(), discussionID); err != nil {
		return
	}

	resp.LikeCount = stat.LikeCount
	resp.DislikeCount = stat.DislikeCount
	resp.CommentCount = stat.CommentCount

	if resp.Fav, err = p.d.IsFav(c, aid, discussionID, model.TargetTypeDiscussion); err != nil {
		return
	}

	if resp.Like, err = p.d.IsLike(c, aid, discussionID, model.TargetTypeDiscussion); err != nil {
		return
	}

	if resp.Dislike, err = p.d.IsDislike(c, aid, discussionID, model.TargetTypeDiscussion); err != nil {
		return
	}

	var tp *topic.TopicInfo
	if tp, err = p.d.GetTopic(c, data.TopicID); err != nil {
		return
	}

	resp.TopicName = tp.Name

	if aid == data.CreatedBy {
		resp.CanEdit = true
		return
	}

	var info *topic.MemberRoleReply
	if info, err = p.d.GetTopicMemberRole(c, data.TopicID, aid); err != nil {
		return
	}

	if info.IsMember && info.Role != model.MemberRoleUser {
		resp.CanEdit = true
	}

	return
}

func (p *Service) getDiscussion(c context.Context, node sqalx.Node, discussionID int64) (item *model.Discussion, err error) {
	if item, err = p.d.GetDiscussionByID(c, node, discussionID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	return
}
