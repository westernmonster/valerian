package service

import (
	"bytes"
	"context"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"valerian/app/admin/config/model"
	"valerian/app/admin/config/pkg/lint"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
)

func lintConfig(filename, content string) error {
	ext := strings.TrimLeft(filepath.Ext(filename), ".")
	err := lint.Lint(ext, bytes.NewBufferString(content))
	if err != nil && err != lint.ErrLintNotExists {
		return ecode.Error(ecode.RequestErr, err.Error())
	}
	return nil
}

func (p *Service) CreateConf(c context.Context, arg *model.ArgCreateConfig) (configID string, err error) {
	// lint config
	if !arg.SkipLint {
		if err = lintConfig(arg.Name, arg.Comment); err != nil {
			return
		}
	}
	app, err := p.appByTree(c, p.d.ConfigDB(), arg.TreeID, arg.Env, arg.Zone)
	if err != nil {
		return
	}

	item := &model.Config{
		ID:        gid.NewID(),
		AppID:     app.ID,
		Name:      arg.Name,
		Comment:   arg.Comment,
		From:      arg.From,
		State:     arg.State,
		Mark:      arg.Mark,
		Operator:  "admin",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if _, err = p.isConfiguring(c, p.d.ConfigDB(), arg.Name, app.ID); err == nil {
		err = ecode.TargetBlocked
		return
	}

	if err = p.d.AddConfig(c, p.d.ConfigDB(), item); err != nil {
		return
	}

	configID = strconv.FormatInt(item.ID, 10)
	return
}

func (p *Service) isConfiguring(c context.Context, node sqalx.Node, name string, appID int64) (conf *model.Config, err error) {
	if conf, err = p.d.GetConfigByCond(c, node, map[string]interface{}{
		"name":   name,
		"app_id": appID,
		"state":  model.ConfigStateInProgress,
	}); err != nil {
		return
	} else if conf == nil {
		err = ecode.NothingFound
		return
	}

	return
}
