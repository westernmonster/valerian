package service

import (
	"context"

	"valerian/app/interface/article/model"
	article "valerian/app/service/article/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetArticleFile(c context.Context, fileID int64) (item *model.ArticleFileResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var canView bool
	if canView, err = p.d.CanView(c, &article.IDReq{Aid: aid, ID: fileID}); err != nil {
		return
	} else if !canView {
		err = ecode.NoArticleViewPermission
		return
	}

	return p.getArticleFile(c, aid, fileID)
}

func (p *Service) getArticleFile(c context.Context, aid int64, fileID int64) (item *model.ArticleFileResp, err error) {
	var data *article.ArticleFileResp
	if data, err = p.d.GetArticleFile(c, &article.IDReq{Aid: aid, ID: fileID}); err != nil {
		return
	}

	item = &model.ArticleFileResp{
		ID:        data.ID,
		FileName:  data.FileName,
		FileURL:   data.FileURL,
		FileType:  data.FileType,
		PdfURL:    data.PdfURL,
		Seq:       int(data.Seq),
		CreatedAt: data.CreatedAt,
	}

	return
}

func (p *Service) GetArticleFiles(c context.Context, articleID int64) (items []*model.ArticleFileResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var canView bool
	if canView, err = p.d.CanView(c, &article.IDReq{Aid: aid, ID: articleID}); err != nil {
		return
	} else if !canView {
		err = ecode.NoArticleViewPermission
		return
	}

	return p.getArticleFiles(c, aid, articleID)
}

func (p *Service) getArticleFiles(c context.Context, aid int64, articleID int64) (items []*model.ArticleFileResp, err error) {
	var data *article.ArticleFilesResp
	if data, err = p.d.GetArticleFiles(c, &article.IDReq{Aid: aid, ID: articleID}); err != nil {
		return
	}

	items = make([]*model.ArticleFileResp, 0)

	if data.Items != nil {
		for _, v := range data.Items {
			items = append(items, &model.ArticleFileResp{
				ID:       v.ID,
				FileName: v.FileName,
				FileURL:  v.FileURL,
				FileType: v.FileType,
				PdfURL:   v.PdfURL,
				Seq:      int(v.Seq),
			})
		}
	}

	return
}

func (p *Service) SaveArticleFiles(c context.Context, arg *model.ArgSaveArticleFiles) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &article.ArgSaveArticleFiles{
		Items:     make([]*article.ArgUpdateArticleFile, 0),
		ArticleID: arg.ArticleID,
		Aid:       aid,
	}

	if arg.Items != nil {
		for _, v := range arg.Items {
			c := &article.ArgUpdateArticleFile{
				FileName: v.FileName,
				FileURL:  v.FileURL,
				FileType: v.FileType,
				Seq:      int32(v.Seq),
			}
			if v.ID != nil {
				c.ID = &article.ArgUpdateArticleFile_IDValue{*v.ID}
			}
			item.Items = append(item.Items, c)
		}
	}

	if err = p.d.SaveArticleFiles(c, item); err != nil {
		return
	}

	return
}
