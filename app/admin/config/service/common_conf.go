package service

import (
	"context"
	"time"
	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
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

	var team *model.Team
	if team, err = p.getTeamByName(c, p.d.ConfigDB(), arg.Name, arg.Env, arg.Zone); err != nil {
		return
	}

	item := &model.CommonConfig{
		ID:        gid.NewID(),
		TeamID:    team.ID,
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

func (p *Service) getCommonConf(c context.Context, node sqalx.Node, id int64) (item *model.CommonConfig, err error) {
	if item, err = p.d.GetCommonConfigByID(c, p.d.ConfigDB(), id); err != nil {
		return
	} else if item == nil {
		err = ecode.CommonConfNotExist
		return
	}

	return
}

// 获取通用配置文件
func (p *Service) GetCommonConf(c context.Context, id int64) (resp *model.CommonConfigResp, err error) {
	var item *model.CommonConfig
	if item, err = p.getCommonConf(c, p.d.ConfigDB(), id); err != nil {
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

func (p *Service) GetCommonConfsByName(c context.Context, teamName, env, zone, name string) (items []*model.CommonConfigResp, err error) {
	var team *model.Team
	if team, err = p.getTeamByName(c, p.d.ConfigDB(), teamName, env, zone); err != nil {
		return
	}
	var data []*model.CommonConfig
	if data, err = p.d.GetCommonConfigsByCond(c, p.d.ConfigDB(), map[string]interface{}{
		"name":    name,
		"team_id": team.ID,
	}); err != nil {
		return
	}

	items = make([]*model.CommonConfigResp, len(data))
	for i, v := range data {
		items[i] = &model.CommonConfigResp{
			ID:        v.ID,
			TeamID:    v.TeamID,
			Name:      v.Name,
			Comment:   v.Comment,
			State:     v.State,
			Mark:      v.Mark,
			Operator:  v.Operator,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return
}

func (p *Service) UpdatCommonConf(c context.Context, arg *model.ArgUpdateCommonConfig) (err error) {
	if !arg.SkipLint {
		if err = lintConfig(arg.Name, arg.Comment); err != nil {
			return
		}
	}

	var confDB *model.CommonConfig
	if confDB, err = p.getCommonConf(c, p.d.ConfigDB(), arg.ID); err != nil {
		return
	}

	if confDB.State == model.ConfigStateInProgress { //judge config is configIng.
		if arg.UpdatedAt != confDB.UpdatedAt {
			err = ecode.TargetBlocked
			return
		}

		confDB.UpdatedAt = 0
		confDB.Comment = arg.Comment
		confDB.State = arg.State
		confDB.Mark = arg.Mark
		confDB.Name = arg.Name

		if err = p.d.UpdateCommonConfig(c, p.d.ConfigDB(), confDB); err != nil {
			return
		}
		return
	}
	if _, err = p.getConfiguringCommonConf(c, p.d.ConfigDB(), confDB.Name, confDB.TeamID); err == nil { //judge have configing.
		err = ecode.TargetBlocked
		return
	}

	return
}

func (p *Service) getConfiguringCommonConf(c context.Context, node sqalx.Node, name string, teamID int64) (conf *model.CommonConfig, err error) {
	if conf, err = p.d.GetCommonConfigByCond(c, p.d.ConfigDB(), map[string]interface{}{
		"name":    name,
		"team_id": teamID,
		"state":   model.ConfigStateInProgress,
	}); err != nil {
		return
	} else if conf == nil {
		err = ecode.CommonConfNotExist
		return
	}

	return
}
