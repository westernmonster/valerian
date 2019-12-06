package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
	"strconv"

	"valerian/app/interface/file/model"
	"valerian/library/conf/env"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/h2non/filetype"
)

func (p *Service) UploadImageURL(c context.Context, arg *model.ArgUploadURL) (resp *model.UploadURLResp, err error) {
	var data []byte
	if data, err = p.downloadImage(c, arg.FileURL); err != nil {
		err = ecode.DownloadImageFailed
		return
	}

	if !filetype.IsImage(data) {
		err = ecode.InvalidImage
		return
	}

	bucketName := "flywiki"
	if env.DeployEnv == env.DeployEnvProd {
		bucketName = "stonote"
	}

	bucket, err := p.ossClient.Bucket(bucketName)
	if err != nil {
		return
	}

	r := bytes.NewReader(data)

	id := gid.NewID()
	ext := filepath.Ext(path.Base(arg.FileURL))
	name := strconv.FormatInt(id, 10)
	if ext != "" {
		name = name + ext
	}
	if err = bucket.PutObject(p.c.OSS.ImageDir+name, r); err != nil {
		return
	}

	resp = &model.UploadURLResp{}
	if env.DeployEnv == env.DeployEnvProd {
		resp.FileURL = "https://res.stonote.cn/" + p.c.OSS.ImageDir + name
	} else {
		resp.FileURL = "https://res.flywk.com/" + p.c.OSS.ImageDir + name
	}
	return
}

func (p *Service) downloadImage(c context.Context, url string) (resp []byte, err error) {
	reply, err := http.Get(url)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("downloadImage(), error(%+v)", err))
		return
	}
	defer reply.Body.Close()

	if reply.StatusCode != 200 {
		err = errors.New("curl error http code not equal to 200")
		log.For(c).Error(fmt.Sprintf("downloadImage() status_code(%d)", reply.StatusCode))
		return
	}
	resp, err = ioutil.ReadAll(reply.Body)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("downloadImage() read_err:%v", err))
		return
	}
	return
}
