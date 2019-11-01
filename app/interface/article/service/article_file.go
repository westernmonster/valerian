package service

import (
	"context"

	"valerian/app/interface/article/model"
	article "valerian/app/service/article/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetArticleFiles(c context.Context, articleID int64) (items []*model.ArticleFileResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data *article.ArticleFilesResp
	if data, err = p.d.GetArticleFiles(c, &article.IDReq{Aid: aid, ID: articleID}); err != nil {
		return
	}

	items = make([]*model.ArticleFileResp, 0)

	if data.Items != nil {
		for _, v := range items {
			items = append(items, &model.ArticleFileResp{
				ID:       v.ID,
				FileName: v.FileName,
				FileURL:  v.FileURL,
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
