package service

import (
	"bytes"
	"context"

	"valerian/app/interface/common/model"
	"valerian/library/ecode"

	"github.com/PuerkitoBio/goquery"
	retry "gopkg.in/h2non/gentleman-retry.v2"
	"gopkg.in/h2non/gentleman.v2"
)

func (p *Service) LinkInfo(c context.Context, arg *model.ArgLinkInfo) (resp *model.LinkInfoResp, err error) {
	cli := gentleman.New()
	cli.Use(retry.New(retry.ConstantBackoff))
	cli.URL(arg.Link)
	req := cli.Request()

	var r *gentleman.Response
	if r, err = req.Send(); err != nil {
		err = ecode.GrabLinkFailed
		return
	}

	if !r.Ok {
		err = ecode.GrabLinkFailed
		return
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(bytes.NewReader(r.Bytes()))
	if err != nil {
		err = ecode.ParseHTMLFailed
		return
	}

	title := doc.Find("title").Text()

	resp = &model.LinkInfoResp{
		Title: title,
		Link:  arg.Link,
	}

	return
}
