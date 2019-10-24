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

func (p *Service) CreateApp(c context.Context, arg *model.ArgCreateApp) (err error) {
	creates := []string{"dev", "uat", "prod"}

	var tx sqalx.Node
	if tx, err = p.d.ConfigDB().Beginx(c); err != nil {
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
