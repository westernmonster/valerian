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
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/imm"
	"github.com/jinzhu/copier"
)

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
			p.d.SetArticleFileCache(context.Background(), articleID, items)
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
		p.convertOfficeFiles(context.Background(), articleID)
		p.d.DelArticleFileCache(context.Background(), articleID)
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
