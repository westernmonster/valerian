package service

import (
	"context"

	"valerian/app/interface/account/model"
)

// AdminAddAccount 添加账户
func (p *Service) AdminAddAccount(c context.Context, arg *model.ArgAdminAddAccount) (err error) {
	return
}

// AdminUpdateAccount 更新账户信息
func (p *Service) AdminUpdateAccount(c context.Context, arg *model.ArgAdminUpdateProfile) (err error) {
	return
}

// AdminLockAccount 锁定账户
func (p *Service) AdminLockAccount(c context.Context, arg *model.ArgAdminLockAccount) (err error) {
	return
}

// AdminUnlockAccount 解锁账户
func (p *Service) AdminUnlockAccount(c context.Context, arg *model.ArgAdminLockAccount) (err error) {
	return
}

// AdminDeactiveAccount 管理员注销账户
func (p *Service) AdminDeactiveAccount(c context.Context, arg *model.ArgAdminDeactiveAccount) (err error) {
	return
}
