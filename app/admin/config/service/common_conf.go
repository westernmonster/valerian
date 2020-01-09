package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) GetCommonConfsPaged(c context.Context, cond map[string]interface{}, page, pageSize int32) (resp *model.CommonConfigListResp, err error) {
	var count int32
	var data []*model.CommonConfig
	if count, data, err = p.d.GetCommonConfigsByCondPaged(c, p.d.ConfigDB(), cond, page, pageSize); err != nil {
		return
	}

	resp = &model.CommonConfigListResp{
		Total:    count,
		Page:     page,
		PageSize: pageSize,
		Items:    make([]*model.CommonConfigItem, len(data)),
	}

	for i, v := range data {
		item := &model.CommonConfigItem{
			ID:        v.ID,
			Name:      v.Name,
			Comment:   v.Comment,
			Mark:      v.Mark,
			State:     v.State,
			Operator:  v.Operator,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		resp.Items[i] = item
	}

	return
}

// 创建通用配置文件
func (p *Service) CreateCommonConf(c context.Context, arg *model.ArgCreateCommonConfig) (err error) {
	// lint config
	if !arg.SkipLint {
		if err := lintConfig(arg.Name, arg.Comment); err != nil {
			return err
		}
	}

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

	var data *model.CommonConfig
	if data, err = p.d.GetCommonConfigByCond(c, tx, map[string]interface{}{
		"name": arg.Name,
	}); err != nil {
		return
	} else if data != nil {
		err = ecode.CommonConfExist
		return
	}

	item := &model.CommonConfig{
		ID:        gid.NewID(),
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

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

func (p *Service) getCommonConf(c context.Context, node sqalx.Node, id int64) (item *model.CommonConfig, err error) {
	if item, err = p.d.GetCommonConfigByID(c, node, id); err != nil {
		return
	} else if item == nil {
		err = ecode.CommonConfNotExist
		return
	}

	return
}

// 获取通用配置文件
func (p *Service) GetCommonConf(c context.Context, id int64) (resp *model.CommonConfigItem, err error) {
	var item *model.CommonConfig
	if item, err = p.getCommonConf(c, p.d.ConfigDB(), id); err != nil {
		return
	}

	resp = &model.CommonConfigItem{
		ID:        item.ID,
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

func (p *Service) UpdatCommonConf(c context.Context, arg *model.ArgUpdateCommonConfig) (err error) {
	if !arg.SkipLint {
		if err = lintConfig(arg.Name, arg.Comment); err != nil {
			return
		}
	}

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
	var confDB *model.CommonConfig
	if confDB, err = p.getCommonConf(c, tx, arg.ID); err != nil {
		return
	}

	if confDB.State == model.ConfigStateInProgress {
		// 如果有状态为正在配置中的
		// 判断其时间戳是否跟传入来的相等，不相等则报错
		if arg.UpdatedAt != confDB.UpdatedAt {
			err = ecode.TargetBlocked
			return
		}
	}

	confDB.UpdatedAt = time.Now().Unix()
	confDB.Comment = arg.Comment
	confDB.State = arg.State
	confDB.Mark = arg.Mark
	confDB.Name = arg.Name

	if err = p.d.UpdateCommonConfig(c, tx, confDB); err != nil {
		return
	}
	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	return
}

func (p *Service) getConfiguringCommonConf(c context.Context, node sqalx.Node, name string) (conf *model.CommonConfig, err error) {
	if conf, err = p.d.GetCommonConfigByCond(c, p.d.ConfigDB(), map[string]interface{}{
		"name":  name,
		"state": model.ConfigStateInProgress,
	}); err != nil {
		return
	} else if conf == nil {
		err = ecode.CommonConfNotExist
		return
	}

	return
}
