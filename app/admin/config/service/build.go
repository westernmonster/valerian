package service

import (
	"context"
	"time"
	"valerian/app/admin/config/model"
	"valerian/library/gid"
)

func (p *Service) CreateBuild(c context.Context, arg *model.ArgCreateBuild) (err error) {
	var app *model.App
	if app, err = p.appByTree(c, p.d.ConfigDB(), arg.TreeID, arg.Env, arg.Zone); err != nil {
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

	if err = p.d.AddBuild(c, p.d.ConfigDB(), item); err != nil {
		return
	}
	return
}
