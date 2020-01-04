package service

import (
	"context"
	"valerian/app/service/identify/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// IsSystemAdmin 是否系统管理员
func (p *Service) IsSystemAdmin(c context.Context, aid int64) (ret bool, err error) {
	return p.isSystemAdmin(c, p.d.DB(), aid)
}

// isSystemAdmin 是否系统管理员
func (p *Service) isSystemAdmin(c context.Context, node sqalx.Node, aid int64) (ret bool, err error) {
	var acc *model.Account
	if acc, err = p.getAccountByID(c, node, aid); err != nil {
		return
	}

	if acc.Role == "admin" || acc.Role == "superadmin" {
		ret = true
		return
	}

	return
}

// checkSystemAdmin 检测是否系统管理员
func (p *Service) checkSystemAdmin(c context.Context, node sqalx.Node, aid int64) (err error) {
	var isSystemAdmin bool
	if isSystemAdmin, err = p.isSystemAdmin(c, node, aid); err != nil {
		return
	}

	if !isSystemAdmin {
		err = ecode.MethodNoPermission
		return
	}

	return
}
