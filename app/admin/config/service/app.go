package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	uuid "github.com/satori/go.uuid"
)

// AppByTree 获取App信息
func (p *Service) AppByTree(c context.Context, treeID int, env, zone string) (app *model.App, err error) {
	return p.appByTree(c, p.d.ConfigDB(), treeID, env, zone)
}

//appByTree 获取App信息
func (p *Service) appByTree(c context.Context, node sqalx.Node, treeID int, env, zone string) (app *model.App, err error) {
	if app, err = p.d.GetAppByCond(c, node, map[string]interface{}{
		"tree_id": treeID,
		"env":     env,
		"zone":    zone,
	}); err != nil {
		return
	} else if app == nil {
		err = ecode.NothingFound
		return
	}

	return
}

// CreateApp 新建App
func (p *Service) CreateApp(c context.Context, arg *model.ArgCreateApp) (err error) {
	return p.createApp(c, p.d.ConfigDB(), arg)
}

// createApp 新建App
func (p *Service) createApp(c context.Context, node sqalx.Node, arg *model.ArgCreateApp) (err error) {
	creates := []string{"dev", "uat", "prod"}

	var tx sqalx.Node
	if tx, err = node.Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	for _, v := range creates {
		bytes := [16]byte(uuid.NewV1())
		token := hex.EncodeToString(bytes[:])

		var app *model.App
		if app, err = p.d.GetAppByCond(c, tx, map[string]interface{}{
			"env":     v,
			"zone":    model.DefaultZone,
			"tree_id": arg.TreeID,
		}); err != nil {
			return
		} else if app != nil {
			err = ecode.AppExist
			return
		}

		item := &model.App{
			ID:        gid.NewID(),
			Name:      arg.AppName,
			Token:     token,
			Env:       v,
			Zone:      model.DefaultZone,
			TreeID:    arg.TreeID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddApp(c, tx, item); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

func (p *Service) UpdateToken(c context.Context, env, zone string, treeID int64) (err error) {
	bytes := [16]byte(uuid.NewV1())
	token := hex.EncodeToString(bytes[:])

	var app *model.App
	if app, err = p.d.GetAppByCond(c, p.d.ConfigDB(), map[string]interface{}{
		"tree_id": treeID,
		"env":     env,
		"zone":    zone,
	}); err != nil {
		return
	} else if app == nil {
		err = ecode.NothingFound
		return
	}

	app.Token = token
	if err = p.d.UpdateApp(c, p.d.ConfigDB(), app); err != nil {
		return
	}

	if err = p.SetToken(c, treeID, env, zone, token); err != nil {
		return
	}
	return
}

func (p *Service) AppByTreeZone(c context.Context, treeID int, zone string) (app []*model.App, err error) {
	return p.appByTreeZone(c, p.d.ConfigDB(), treeID, zone)
}

func (p *Service) appByTreeZone(c context.Context, node sqalx.Node, treeID int, zone string) (apps []*model.App, err error) {
	apps = make([]*model.App, 0)
	if apps, err = p.d.GetAppsByCond(c, node, map[string]interface{}{
		"tree_id": treeID,
		"zone":    zone,
	}); err != nil {
		return
	}

	return
}
