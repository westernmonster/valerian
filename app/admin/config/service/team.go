package service

import (
	"context"
	"time"
	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
)

func (p *Service) CreateTeam(c context.Context, arg *model.ArgCreateTeam) (err error) {
	var team *model.Team
	if team, err = p.d.GetTeamByCond(c, p.d.ConfigDB(), map[string]interface{}{
		"name": arg.Name,
		"env":  arg.Env,
		"zone": arg.Zone,
	}); err != nil {
		return
	} else if team != nil {
		err = ecode.TeamExist
		return
	}

	item := &model.Team{
		ID:        gid.NewID(),
		Name:      arg.Name,
		Env:       arg.Env,
		Zone:      arg.Zone,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddTeam(c, p.d.ConfigDB(), item); err != nil {
		return
	}

	return
}

func (p *Service) GetTeamByName(c context.Context, name, env, zone string) (item *model.Team, err error) {
	return p.getTeamByName(c, p.d.ConfigDB(), name, env, zone)
}

func (p *Service) getTeamByName(c context.Context, node sqalx.Node, name, env, zone string) (item *model.Team, err error) {
	if item, err = p.d.GetTeamByCond(c, node, map[string]interface{}{
		"name": name,
		"env":  env,
		"zone": zone,
	}); err != nil {
		return
	} else if item == nil {
		err = ecode.TeamNotExist
		return
	}

	return
}
