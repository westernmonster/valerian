package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/xstr"

	"github.com/PuerkitoBio/goquery"
)

// initDiscussionStat 初始化讨论状态
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

// getDiscussion 获取讨论信息
func (p *Service) getDiscussion(c context.Context, node sqalx.Node, articleID int64) (item *model.Discussion, err error) {
	var addCache = true
	if item, err = p.d.DiscussionCache(c, articleID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), articleID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetDiscussionCache(context.TODO(), item)
		})
	}
	return
}

// GetUserDiscussionIDsPaged 获取用户所有创建的讨论ID列表
func (p *Service) GetUserDiscussionIDsPaged(c context.Context, req *api.UserDiscussionsReq) (resp []int64, err error) {
	return p.d.GetUserDiscussionIDsPaged(c, p.d.DB(), req.AccountID, int(req.Limit), int(req.Offset))
}

// GetTopicDiscussionsPaged 分页获取话题讨论列表
func (p *Service) GetTopicDiscussionsPaged(c context.Context, req *api.TopicDiscussionsReq) (items []*api.DiscussionInfo, err error) {
	var data []*model.Discussion
	if data, err = p.d.GetTopicDiscussionsPaged(c, p.d.DB(), req.TopicID, req.CategoryID, int(req.Limit), int(req.Offset)); err != nil {
		return
	}

	items = make([]*api.DiscussionInfo, len(data))

	for i, v := range data {

		item := &api.DiscussionInfo{
			ID:         v.ID,
			TopicID:    v.TopicID,
			CategoryID: v.CategoryID,
			// CreatedBy:   v.CreatedBy,
			Excerpt:   xstr.Excerpt(v.ContentText),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			ImageUrls: make([]string, 0),
			Title:     v.Title,
		}

		var imgs []*model.ImageURL
		if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
			"target_type": model.TargetTypeDiscussion,
			"target_id":   v.ID,
		}); err != nil {
			return
		}

		for _, v := range imgs {
			item.ImageUrls = append(item.ImageUrls, v.URL)
		}

		var stat *model.DiscussionStat
		if stat, err = p.d.GetDiscussionStatByID(c, p.d.DB(), v.ID); err != nil {
			return
		}
		item.Stat = &api.DiscussionStat{
			DislikeCount: int32(stat.DislikeCount),
			LikeCount:    int32(stat.LikeCount),
			CommentCount: int32(stat.CommentCount),
		}

		if v.CategoryID != int64(-1) {
			var cate *model.DiscussCategory
			if cate, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
				return
			} else if cate == nil {
				err = ecode.DiscussCategoryNotExist
				return
			}

			item.CategoryInfo = &api.CategoryInfo{
				ID:      cate.ID,
				TopicID: cate.TopicID,
				Name:    cate.Name,
				Seq:     cate.Seq,
			}
		}

		var acc *model.Account
		if acc, err = p.getAccount(c, p.d.DB(), v.CreatedBy); err != nil {
			return
		}

		item.Creator = &api.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		items[i] = item
	}

	return
}

// GetUserDiscussionsPaged 分页获取用户创建的讨论列表
func (p *Service) GetUserDiscussionsPaged(c context.Context, req *api.UserDiscussionsReq) (items []*api.DiscussionInfo, err error) {
	var data []*model.Discussion
	if data, err = p.d.GetUserDiscussionsPaged(c, p.d.DB(), req.AccountID, int(req.Limit), int(req.Offset)); err != nil {
		return
	}

	items = make([]*api.DiscussionInfo, len(data))

	for i, v := range data {

		item := &api.DiscussionInfo{
			ID:         v.ID,
			TopicID:    v.TopicID,
			CategoryID: v.CategoryID,
			// CreatedBy:   v.CreatedBy,
			Excerpt:   xstr.Excerpt(v.ContentText),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			ImageUrls: make([]string, 0),
			Title:     v.Title,
		}

		var imgs []*model.ImageURL
		if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
			"target_type": model.TargetTypeDiscussion,
			"target_id":   v.ID,
		}); err != nil {
			return
		}

		for _, v := range imgs {
			item.ImageUrls = append(item.ImageUrls, v.URL)
		}

		var stat *model.DiscussionStat
		if stat, err = p.d.GetDiscussionStatByID(c, p.d.DB(), v.ID); err != nil {
			return
		}
		item.Stat = &api.DiscussionStat{
			DislikeCount: int32(stat.DislikeCount),
			LikeCount:    int32(stat.LikeCount),
			CommentCount: int32(stat.CommentCount),
		}

		if v.CategoryID != int64(-1) {
			var cate *model.DiscussCategory
			if cate, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
				return
			} else if cate == nil {
				err = ecode.DiscussCategoryNotExist
				return
			}

			item.CategoryInfo = &api.CategoryInfo{
				ID:      cate.ID,
				TopicID: cate.TopicID,
				Name:    cate.Name,
				Seq:     cate.Seq,
			}
		}

		var acc *model.Account
		if acc, err = p.getAccount(c, p.d.DB(), v.CreatedBy); err != nil {
			return
		}

		item.Creator = &api.Creator{
			ID:           acc.ID,
			UserName:     acc.UserName,
			Avatar:       acc.Avatar,
			Introduction: acc.Introduction,
		}

		items[i] = item
	}

	return
}

// GetDiscussion 获取指定讨论信息
func (p *Service) GetDiscussionInfo(c context.Context, req *api.IDReq) (resp *api.DiscussionInfo, err error) {
	var item *model.Discussion
	if item, err = p.getDiscussion(c, p.d.DB(), req.ID); err != nil {
		return
	}

	resp = &api.DiscussionInfo{
		ID:         item.ID,
		TopicID:    item.TopicID,
		CategoryID: item.CategoryID,
		Excerpt:    xstr.Excerpt(item.ContentText),
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
		ImageUrls:  make([]string, 0),
		Files:      make([]*api.DiscussionFile, 0),
		Title:      item.Title,
	}

	inc := includeParam(req.Include)
	if inc["content"] {
		resp.Content = item.Content
	}

	if inc["content_text"] {
		resp.ContentText = item.ContentText
	}

	// 图片
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeDiscussion,
		"target_id":   req.ID,
	}); err != nil {
		return
	}
	for _, v := range imgs {
		resp.ImageUrls = append(resp.ImageUrls, v.URL)
	}

	// 状态
	var stat *model.DiscussionStat
	if stat, err = p.getDiscussionStat(c, p.d.DB(), req.ID); err != nil {
		return
	}
	resp.Stat = &api.DiscussionStat{
		DislikeCount: int32(stat.DislikeCount),
		LikeCount:    int32(stat.LikeCount),
		CommentCount: int32(stat.CommentCount),
	}

	// 分类
	if item.CategoryID != -1 {
		var cate *model.DiscussCategory
		if cate, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), item.CategoryID); err != nil {
			return
		} else if cate == nil {
			err = ecode.DiscussCategoryNotExist
			return
		}

		resp.CategoryInfo = &api.CategoryInfo{
			ID:      cate.ID,
			TopicID: cate.TopicID,
			Name:    cate.Name,
			Seq:     cate.Seq,
		}

	} else {
		resp.CategoryInfo = &api.CategoryInfo{
			ID:      -1,
			TopicID: item.TopicID,
			Name:    "问答",
			Seq:     1,
		}
	}

	var files []*model.DiscussionFile
	if files, err = p.getDiscussionFiles(c, p.d.DB(), req.Aid, item.ID); err != nil {
		return
	}
	for _, v := range files {
		resp.Files = append(resp.Files, &api.DiscussionFile{
			ID:        v.ID,
			FileName:  v.FileName,
			FileURL:   v.FileURL,
			PdfURL:    v.PdfURL,
			FileType:  v.FileType,
			Seq:       v.Seq,
			CreatedAt: v.CreatedAt,
		})
	}

	// 创建者
	var acc *model.Account
	if acc, err = p.getAccount(c, p.d.DB(), item.CreatedBy); err != nil {
		return
	}
	resp.Creator = &api.Creator{
		ID:           acc.ID,
		UserName:     acc.UserName,
		Avatar:       acc.Avatar,
		Introduction: acc.Introduction,
	}

	return
}

// GetDiscussionStat 获取指定讨论状态
func (p *Service) getDiscussionStat(c context.Context, node sqalx.Node, discussionID int64) (item *model.DiscussionStat, err error) {
	if item, err = p.d.GetDiscussionStatByID(c, node, discussionID); err != nil {
		return
	} else {
		item = &model.DiscussionStat{
			DiscussionID: discussionID,
		}
	}
	return
}

// GetDiscussionStat 获取指定讨论状态
func (p *Service) GetDiscussionStat(c context.Context, discussionID int64) (item *model.DiscussionStat, err error) {
	return p.getDiscussionStat(c, p.d.DB(), discussionID)
}

// DelDiscussion 删除讨论
func (p *Service) DelDiscussion(c context.Context, arg *api.IDReq) (err error) {
	// TODO:  删除讨论逻辑
	if err = p.checkEditPermission(c, p.d.DB(), arg.Aid, arg.ID); err != nil {
		return
	}
	return
}

// AddDiscussion 添加讨论
func (p *Service) AddDiscussion(c context.Context, arg *api.ArgAddDiscussion) (id int64, err error) {
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

	// 检测分类
	if err = p.checkCategory(c, tx, arg.CategoryID); err != nil {
		return
	}

	// 检测话题
	if err = p.checkTopicExist(c, tx, arg.TopicID); err != nil {
		return
	}

	// 检测是否成员
	if err = p.checkIsTopicMember(c, tx, arg.Aid, arg.TopicID); err != nil {
		return
	}

	item := &model.Discussion{
		ID:         gid.NewID(),
		TopicID:    arg.TopicID,
		CategoryID: arg.CategoryID,
		CreatedBy:  arg.Aid,
		Content:    arg.Content,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	item.Title = arg.GetTitleValue()

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
	if err = p.d.IncrAccountStat(c, tx, &model.AccountStat{AccountID: arg.Aid, DiscussionCount: 1}); err != nil {
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
		p.onDiscussionAdded(context.Background(), item.ID, arg.Aid, time.Now().Unix())
	})

	return
}

// UpdateDiscussion 更新讨论
func (p *Service) UpdateDiscussion(c context.Context, arg *api.ArgUpdateDiscussion) (err error) {

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

	if err = p.checkEditPermission(c, tx, arg.Aid, arg.ID); err != nil {
		return
	}

	var item *model.Discussion
	if item, err = p.d.GetDiscussionByID(c, tx, arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	if item.CreatedBy != arg.Aid {
		if err = p.checkIsTopicManager(c, tx, item.TopicID, arg.Aid); err != nil {
			return
		}
	}

	item.Title = arg.GetTitleValue()
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
		p.onDiscussionUpdated(context.Background(), arg.ID, arg.Aid, time.Now().Unix())
	})

	return
}
