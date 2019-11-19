package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) InitProd(c context.Context) (err error) {
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

	var builds []*model.Build
	if builds, err = p.d.GetBuilds(c, tx); err != nil {
		return
	}

	for _, v := range builds {
		var app *model.App
		if app, err = p.d.GetAppByID(c, tx, v.AppID); err != nil {
			return
		} else if app == nil {
			err = ecode.AppNotExist
			return
		}

		var prodApp *model.App
		if prodApp, err = p.d.GetAppByCond(c, tx, map[string]interface{}{
			"name":    app.Name,
			"env":     "prod",
			"tree_id": app.TreeID,
			"zone":    app.Zone,
		}); err != nil {
			return
		} else if prodApp == nil {
			err = ecode.AppNotExist
			return
		}

		var tag *model.Tag
		if tag, err = p.d.GetTagByID(c, tx, v.TagID); err != nil {
			return
		} else if tag == nil {
			err = ecode.TagNotExist
			return
		}

		var configID int64
		if configID, err = strconv.ParseInt(tag.ConfigIds, 10, 64); err != nil {
			return
		}

		var conf *model.Config
		if conf, err = p.d.GetConfigByID(c, tx, configID); err != nil {
			return
		} else if conf == nil {
			err = ecode.ConfigsNotExist
			return
		}

		conf.ID = gid.NewID()
		conf.Mark = app.Name
		conf.AppID = prodApp.ID
		conf.CreatedAt = time.Now().Unix()
		conf.UpdatedAt = time.Now().Unix()

		if err = p.d.AddConfig(c, tx, conf); err != nil {
			return
		}

		tag.ID = gid.NewID()
		tag.ConfigIds = strconv.FormatInt(conf.ID, 10)
		tag.AppID = prodApp.ID
		tag.Mark = prodApp.Name
		tag.CreatedAt = time.Now().Unix()
		tag.UpdatedAt = time.Now().Unix()

		if err = p.d.AddTag(c, tx, tag); err != nil {
			return
		}

		build := &model.Build{
			ID:        gid.NewID(),
			AppID:     prodApp.ID,
			Name:      "0.1",
			TagID:     tag.ID,
			Operator:  "admin",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddBuild(c, tx, build); err != nil {
			return
		}

	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}
