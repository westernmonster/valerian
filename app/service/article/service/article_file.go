package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

// GetArticleFile 获取指定文章附件
func (p *Service) GetArticleFile(c context.Context, id int64) (item *api.ArticleFileResp, err error) {
	var data *model.ArticleFile
	if data, err = p.d.GetArticleFileByID(c, p.d.DB(), id); err != nil {
		return
	} else if data == nil {
		err = ecode.ArticleFileNotExist
		return
	}

	item = &api.ArticleFileResp{
		ID:        data.ID,
		FileName:  data.FileName,
		FileURL:   data.FileURL,
		PdfURL:    data.PdfURL,
		FileType:  data.FileType,
		Seq:       data.Seq,
		CreatedAt: data.CreatedAt,
	}

	return
}

// GetArticleFiles 获取文章附件列表
func (p *Service) GetArticleFiles(c context.Context, articleID int64) (items []*api.ArticleFileResp, err error) {
	return p.getArticleFiles(c, p.d.DB(), articleID)
}

// SaveArticleFiles 批量保存文章附件
func (p *Service) SaveArticleFiles(c context.Context, arg *api.ArgSaveArticleFiles) (err error) {
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

	if err = p.checkEditPermission(c, tx, arg.Aid, arg.ArticleID); err != nil {
		return
	}

	var article *model.Article
	if article, err = p.d.GetArticleByID(c, tx, arg.ArticleID); err != nil {
		return
	} else if article == nil {
		return ecode.ArticleNotExist
	}

	dbItems, err := p.d.GetArticleFilesByCond(c, tx, map[string]interface{}{"article_id": arg.ArticleID})
	if err != nil {
		return
	}

	dic := make(map[int64]bool)
	for _, v := range arg.Items {
		if v.ID == nil {
			// Add
			item := &model.ArticleFile{
				ID:        gid.NewID(),
				FileName:  v.FileName,
				FileURL:   v.FileURL,
				FileType:  v.FileType,
				Seq:       v.Seq,
				ArticleID: arg.ArticleID,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}

			if err = p.d.AddArticleFile(c, tx, item); err != nil {
				return
			}
			continue
		}

		// Update
		dic[v.GetIDValue()] = true
		var file *model.ArticleFile
		if file, err = p.d.GetArticleFileByID(c, tx, v.GetIDValue()); err != nil {
			return
		} else if file == nil {
			err = ecode.ArticleFileNotExist
			return
		}

		if file.FileURL != v.FileURL {
			file.PdfURL = ""
		}
		file.FileName = v.FileName
		file.FileURL = v.FileURL
		file.FileType = v.FileType
		file.Seq = v.Seq

		if err = p.d.UpdateArticleFile(c, tx, file); err != nil {
			return
		}
	}

	// Delete
	for _, v := range dbItems {
		if dic[v.ID] {
			continue
		}

		if err = p.d.DelArticleFile(c, tx, v.ID); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.convertOfficeFiles(context.Background(), arg.ArticleID)
		p.d.DelArticleFileCache(context.Background(), arg.ArticleID)
	})
	return
}
