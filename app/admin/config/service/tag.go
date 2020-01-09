package service

import (
	"context"
	"strconv"
	"time"

	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
)

func (p *Service) getTag(c context.Context, node sqalx.Node, id int64) (item *model.Tag, err error) {
	if item, err = p.d.GetTagByID(c, node, id); err != nil {
		return
	} else if item == nil {
		err = ecode.TagNotExist
		return
	}

	return
}

func (p *Service) CreateTag(c context.Context, arg *model.ArgCreateTag) (tagID string, err error) {
	return p.createTag(c, p.d.ConfigDB(), arg)
}

func (p *Service) createTag(c context.Context, node sqalx.Node, arg *model.ArgCreateTag) (tagID string, err error) {
	var app *model.App
	if app, err = p.appByTree(c, node, arg.TreeID, arg.Env, arg.Zone); err != nil {
		return
	}

	item := &model.Tag{
		ID:        gid.NewID(),
		AppID:     app.ID,
		ConfigIds: arg.ConfigIDs,
		Mark:      arg.Mark,
		Force:     0,
		Operator:  "admin",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddTag(c, node, item); err != nil {
		return
	}

	tagID = strconv.FormatInt(item.ID, 10)
	return
}
