package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/xstr"

	"github.com/PuerkitoBio/goquery"
)

// GetReviseStat 获取补充状态信息
func (p *Service) GetReviseStat(c context.Context, reviseID int64) (item *model.ReviseStat, err error) {
	return p.d.GetReviseStatByID(c, p.d.DB(), reviseID)
}

// GetRevise 获取补充
func (p *Service) GetRevise(c context.Context, reviseID int64) (item *model.Revise, err error) {
	if item, err = p.getRevise(c, p.d.DB(), reviseID); err != nil {
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, p.d.DB(), item.ArticleID); err != nil {
		return
	}

	item.Title = article.Title
	return
}

// getRevise 获取补充
func (p *Service) getRevise(c context.Context, node sqalx.Node, reviseID int64) (item *model.Revise, err error) {
	var addCache = true
	if item, err = p.d.ReviseCache(c, reviseID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetReviseByID(c, p.d.DB(), reviseID); err != nil {
		return
	} else if item == nil {
		err = ecode.ReviseNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetReviseCache(context.Background(), item)
		})
	}
	return
}

// GetReviseImageUrls 获取补充图片地址
func (p *Service) GetReviseImageUrls(c context.Context, reviseID int64) (urls []string, err error) {
	urls = make([]string, 0)
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeRevise,
		"target_id":   reviseID,
	}); err != nil {
		return
	}

	for _, v := range imgs {
		urls = append(urls, v.URL)
	}

	return
}

// GetReviseInfo 获取补充基本信息
func (p *Service) GetReviseInfo(c context.Context, req *api.IDReq) (item *api.ReviseInfo, err error) {
	revise, err := p.GetRevise(c, req.ID)
	if err != nil {
		return nil, err
	}
	article, err := p.GetArticle(c, revise.ArticleID)
	if err != nil {
		return nil, err
	}

	stat, err := p.GetReviseStat(c, req.ID)
	if err != nil {
		return nil, err
	}

	urls, err := p.GetReviseImageUrls(c, req.ID)
	if err != nil {
		return nil, err
	}

	m, err := p.getAccount(c, p.d.DB(), article.CreatedBy)
	if err != nil {
		return nil, err
	}

	resp := &api.ReviseInfo{
		ID:        revise.ID,
		Title:     article.Title,
		Excerpt:   xstr.Excerpt(revise.ContentText),
		CreatedAt: revise.CreatedAt,
		UpdatedAt: revise.UpdatedAt,
		ImageUrls: urls,
		Stat: &api.ReviseStat{
			CommentCount: int32(stat.CommentCount),
			LikeCount:    int32(stat.LikeCount),
			DislikeCount: int32(stat.DislikeCount),
		},
		Creator: &api.Creator{
			ID:           m.ID,
			UserName:     m.UserName,
			Avatar:       m.Avatar,
			Introduction: m.Introduction,
		},
		ArticleID: revise.ArticleID,
	}

	inc := includeParam(req.Include)

	if inc["content"] {
		resp.Content = article.Content
	}

	if inc["content_text"] {
		resp.ContentText = article.ContentText
	}

	return resp, nil

}

// GetReviseDetail 获取补充详情
func (p *Service) GetReviseDetail(c context.Context, req *api.IDReq) (item *api.ReviseDetail, err error) {
	revise, err := p.GetRevise(c, req.ID)
	if err != nil {
		return nil, err
	}
	article, err := p.GetArticle(c, revise.ArticleID)
	if err != nil {
		return nil, err
	}

	stat, err := p.GetReviseStat(c, req.ID)
	if err != nil {
		return nil, err
	}

	urls, err := p.GetReviseImageUrls(c, req.ID)
	if err != nil {
		return nil, err
	}

	m, err := p.getAccount(c, p.d.DB(), article.CreatedBy)
	if err != nil {
		return nil, err
	}

	files, err := p.getReviseFiles(c, p.d.DB(), req.ID)
	if err != nil {
		return nil, err
	}

	resp := &api.ReviseDetail{
		ID:          revise.ID,
		Title:       article.Title,
		CreatedAt:   revise.CreatedAt,
		UpdatedAt:   revise.UpdatedAt,
		ImageUrls:   urls,
		Content:     revise.Content,
		ContentText: revise.ContentText,
		Stat: &api.ReviseStat{
			CommentCount: int32(stat.CommentCount),
			LikeCount:    int32(stat.LikeCount),
			DislikeCount: int32(stat.DislikeCount),
		},
		Creator: &api.Creator{
			ID:           m.ID,
			UserName:     m.UserName,
			Avatar:       m.Avatar,
			Introduction: m.Introduction,
		},
		ArticleID: revise.ArticleID,
		Files:     make([]*api.ReviseFileResp, 0),
	}

	for _, v := range files {
		resp.Files = append(resp.Files, &api.ReviseFileResp{
			ID:       v.ID,
			FileName: v.FileName,
			FileURL:  v.FileURL,
			Seq:      v.Seq,
		})
	}

	return resp, nil

}

// AddRevise 添加补充
func (p *Service) AddRevise(c context.Context, arg *api.ArgAddRevise) (id int64, err error) {
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

	// check article
	if _, err = p.getArticle(c, tx, arg.ArticleID); err != nil {
		return
	}

	if err = p.checkEditPermission(c, tx, arg.Aid, arg.ArticleID); err != nil {
		return
	}

	item := &model.Revise{
		ID:        gid.NewID(),
		ArticleID: arg.ArticleID,
		Content:   arg.Content,
		CreatedBy: arg.Aid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
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
				TargetType: model.TargetTypeRevise,
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

	if err = p.d.AddRevise(c, tx, item); err != nil {
		return
	}

	if err = p.bulkCreateReviseFiles(c, tx, item.ID, arg.Files); err != nil {
		return
	}

	if err = p.d.AddReviseStat(c, tx, &model.ReviseStat{
		ReviseID:  item.ID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = p.d.IncrArticleStat(c, tx, &model.ArticleStat{ArticleID: item.ArticleID, ReviseCount: 1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	id = item.ID
	p.addCache(func() {
		p.onReviseAdded(context.Background(), item.ID, arg.Aid, time.Now().Unix())
	})

	return
}

// UpdateRevise 更新补充
func (p *Service) UpdateRevise(c context.Context, arg *api.ArgUpdateRevise) (err error) {
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

	var item *model.Revise
	if item, err = p.d.GetReviseByID(c, tx, arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.ReviseNotExist
		return
	}

	if err = p.checkEditPermission(c, tx, arg.Aid, item.ArticleID); err != nil {
		return
	}

	item.Content = arg.Content

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		err = ecode.ParseHTMLFailed
		return
	}
	item.ContentText = doc.Text()

	if err = p.d.DelImageURLByCond(c, tx, model.TargetTypeRevise, item.ID); err != nil {
		return
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		if url, exist := s.Attr("src"); exist {
			u := &model.ImageURL{
				ID:         gid.NewID(),
				TargetType: model.TargetTypeArticle,
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

	item.UpdatedAt = time.Now().Unix()

	if err = p.d.UpdateRevise(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelReviseCache(context.Background(), arg.ID)
		p.onReviseUpdated(context.Background(), arg.ID, arg.Aid, time.Now().Unix())
	})

	return
}

// DelRevise 删除补充
func (p *Service) DelRevise(c context.Context, arg *api.IDReq) (err error) {
	return p.delRevise(c, arg.Aid, arg.ID)
}

// GetUserReviseIDsPaged 获取用户创建的补充ID列表
func (p *Service) GetUserReviseIDsPaged(c context.Context, req *api.UserRevisesReq) (items []int64, err error) {
	return p.d.GetUserReviseIDsPaged(c, p.d.DB(), req.AccountID, int(req.Limit), int(req.Offset))
}

// GetArticleRevisesPaged 获取文章补充列表
func (p *Service) GetArticleRevisesPaged(c context.Context, req *api.ArgArticleRevisesPaged) (resp *api.ReviseListResp, err error) {
	var data []*model.Revise
	if data, err = p.d.GetArticleRevisesPaged(c, p.d.DB(), req.ArticleID, req.Sort, int(req.Limit), int(req.Offset)); err != nil {
		return
	}

	resp = &api.ReviseListResp{
		Items: make([]*api.ReviseInfo, len(data)),
	}

	for i, v := range data {
		article, err := p.GetArticle(c, v.ArticleID)
		if err != nil {
			return nil, err
		}

		stat, err := p.GetReviseStat(c, v.ID)
		if err != nil {
			return nil, err
		}

		urls, err := p.GetReviseImageUrls(c, v.ID)
		if err != nil {
			return nil, err
		}

		m, err := p.getAccount(c, p.d.DB(), article.CreatedBy)
		if err != nil {
			return nil, err
		}

		item := &api.ReviseInfo{
			ID:        v.ID,
			Title:     article.Title,
			Excerpt:   xstr.Excerpt(v.ContentText),
			ImageUrls: urls,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Stat: &api.ReviseStat{
				CommentCount: int32(stat.CommentCount),
				LikeCount:    int32(stat.LikeCount),
				DislikeCount: int32(stat.DislikeCount),
			},
			Creator: &api.Creator{
				ID:           m.ID,
				UserName:     m.UserName,
				Avatar:       m.Avatar,
				Introduction: m.Introduction,
			},
			ArticleID: v.ArticleID,
		}

		resp.Items[i] = item
	}

	var stat *model.ArticleStat
	if stat, err = p.d.GetArticleStatByID(c, p.d.DB(), req.ArticleID); err != nil {
		return
	}

	resp.ReviseCount = int32(stat.ReviseCount)

	return
}
