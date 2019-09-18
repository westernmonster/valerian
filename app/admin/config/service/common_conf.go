package service

import (
	"context"
	"time"
	"valerian/app/admin/config/model"
	"valerian/library/gid"
)

func (p *Service) CreateCommonConf(c context.Context, arg *model.ArgCreateCommonConfig) (err error) {
	// lint config
	if !arg.SkipLint {
		if err := lintConfig(arg.Name, arg.Comment); err != nil {
			return err
		}
	}

	item := &model.CommonConfig{
		ID:        gid.NewID(),
		TeamID:    1,
		Name:      arg.Name,
		Comment:   arg.Comment,
		State:     arg.State,
		Mark:      arg.Mark,
		Operator:  "admin",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddCommonConfig(c, p.d.ConfigDB(), item); err != nil {
		return
	}

	return
}
