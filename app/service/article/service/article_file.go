package service

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"valerian/app/service/article/api"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/imm"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/copier"
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

// getArticleFiles 获取文章附件列表
func (p *Service) getArticleFiles(c context.Context, node sqalx.Node, articleID int64) (items []*api.ArticleFileResp, err error) {
	var addCache = true

	if items, err = p.d.ArticleFileCache(c, articleID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	var data []*model.ArticleFile
	if data, err = p.d.GetArticleFilesByCond(c, node, map[string]interface{}{"article_id": articleID}); err != nil {
		return
	}

	items = make([]*api.ArticleFileResp, 0)
	if err = copier.Copy(&items, &data); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleFileCache(context.TODO(), articleID, items)
		})
	}

	return
}

// bulkCreateFiles 批量保存附件信息
func (p *Service) bulkCreateFiles(c context.Context, node sqalx.Node, articleID int64, files []*api.ArgArticleFile) (err error) {
	var tx sqalx.Node
	if tx, err = node.Beginx(c); err != nil {
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

	for _, v := range files {
		item := &model.ArticleFile{
			ID:        gid.NewID(),
			FileName:  v.FileName,
			FileURL:   v.FileURL,
			FileType:  v.FileType,
			Seq:       v.Seq,
			ArticleID: articleID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddArticleFile(c, tx, item); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.convertOfficeFiles(context.TODO(), articleID)
		p.d.DelArticleFileCache(context.TODO(), articleID)
	})

	return
}

// convertOfficeFiles 转换Office文档
func (p *Service) convertOfficeFiles(c context.Context, articleID int64) (err error) {
	var files []*model.ArticleFile
	if files, err = p.d.GetArticleFilesByCond(c, p.d.DB(), map[string]interface{}{
		"article_id": articleID,
	}); err != nil {
		log.Error(fmt.Sprintf("service.convertOfficeFiles() error(%+v)", err))
		return
	}

	for _, v := range files {
		spew.Dump(v)
		switch v.FileType {
		case model.FileTypePPT:
			fallthrough
		case model.FileTypeExcel:
			fallthrough
		case model.FileTypeWord:
			if v.PdfURL == "" {
				req := imm.CreateCreateOfficeConversionTaskRequest()
				req.Project = "stonote"
				req.StartPage = requests.NewInteger(1)
				req.EndPage = requests.NewInteger(-1)
				req.MaxSheetRow = requests.NewInteger(-1)
				req.MaxSheetCol = requests.NewInteger(-1)
				req.MaxSheetCount = requests.NewInteger(-1)
				req.MaxSheetCount = requests.NewInteger(-1)

				var u *url.URL
				if u, err = url.Parse(v.FileURL); err != nil {
					log.Error(fmt.Sprintf("service.convertOfficeFiles() error(%+v)", err))
					return
				}
				req.SrcUri = "oss://" + p.c.Aliyun.BucketName + u.Path
				req.TgtType = "pdf"

				fName := strings.Split(path.Base(u.Path), ".")[0]
				fURL := strings.TrimRight(u.Path, path.Base(u.Path)) + fName
				req.TgtUri = "oss://" + p.c.Aliyun.BucketName + fURL
				req.SetScheme("https")

				var ret *imm.CreateOfficeConversionTaskResponse
				if ret, err = p.immClient.CreateOfficeConversionTask(req); err != nil {
					log.Error(fmt.Sprintf("service.convertOfficeFiles() error(%+v)", err))
					return
				}
				maxGetCount := 30
				getInternval := time.Second
				getCount := 0
				taskReq := imm.CreateGetOfficeConversionTaskRequest()
				taskReq.Project = "stonote"
				taskReq.TaskId = ret.TaskId
				taskReq.SetScheme("https")
				for {
					time.Sleep(getInternval)
					var taskResp *imm.GetOfficeConversionTaskResponse
					if taskResp, err = p.immClient.GetOfficeConversionTask(taskReq); err != nil {
						log.Error(fmt.Sprintf("service.convertOfficeFiles() error(%+v)", err))
						return
					}

					if taskResp.Status == "Finished" {
						v.PdfURL = u.Scheme + "://" + u.Host + strings.TrimRight(u.Path, path.Base(u.Path)) + fName + "/1.pdf"
						if err = p.d.UpdateArticleFile(c, p.d.DB(), v); err != nil {
							log.Error(fmt.Sprintf("service.convertOfficeFiles() error(%+v)", err))
							return
						}
						break
					}
					if taskResp.Status != "Running" {
						break
					}
					getCount++
					if getCount >= maxGetCount {
						fmt.Println("OfficeConversion Timeout for 30 seconds")
						break
					}
				}

			}
			break
		}
	}

	return
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
		p.convertOfficeFiles(context.TODO(), arg.ArticleID)
		p.d.DelArticleFileCache(context.TODO(), arg.ArticleID)
	})
	return
}
