package service

import (
	"context"
	"time"

	"valerian/app/admin/config/model"
	"valerian/library/gid"
)

func (p *Service) CreateTag(c context.Context, arg *model.ArgCreateTag) (err error) {
	var app *model.App
	if app, err = p.appByTree(c, p.d.ConfigDB(), arg.TreeID, arg.Env, arg.Zone); err != nil {
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

	if err = p.d.AddTag(c, p.d.ConfigDB(), item); err != nil {
		return
	}
	return
}
