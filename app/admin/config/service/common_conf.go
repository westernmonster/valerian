package service

import (
	"context"
	"time"
	"valerian/app/admin/config/model"
	"valerian/library/gid"
)

// 创建通用配置文件
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

// 获取通用配置文件
// 系统
func (p *Service) GetCommonConf(c context.Context, id int64) (resp *model.CommonConfigResp, err error) {
	var item *model.CommonConfig
	if item, err = p.d.GetCommonConfigByID(c, p.d.ConfigDB(), id); err != nil {
		return
	}

	resp = &model.CommonConfigResp{
		ID:        item.ID,
		TeamID:    item.TeamID,
		Name:      item.Name,
		Comment:   item.Comment,
		State:     item.State,
		Mark:      item.Mark,
		Operator:  item.Operator,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	return
}
