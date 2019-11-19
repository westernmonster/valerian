package service

import (
	"context"
	"strconv"
	"time"
	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
)

func (p *Service) CreateBuild(c context.Context, arg *model.ArgCreateBuild) (buildID string, err error) {
	return p.createBuild(c, p.d.ConfigDB(), arg)
}

func (p *Service) createBuild(c context.Context, node sqalx.Node, arg *model.ArgCreateBuild) (buildID string, err error) {
	var app *model.App
	if app, err = p.appByTree(c, node, arg.TreeID, arg.Env, arg.Zone); err != nil {
		return
	}

	item := &model.Build{
		ID:        gid.NewID(),
		AppID:     app.ID,
		Name:      arg.Name,
		TagID:     arg.TagID,
		Operator:  "admin",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddBuild(c, node, item); err != nil {
		return
	}

	buildID = strconv.FormatInt(item.ID, 10)
	return
}
