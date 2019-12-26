package service

import (
	"context"
)

// AccountLock 锁定用户
func (p *Service) AccountLock(ctx context.Context, accountID int64) (err error) {
	err = p.d.AccountSetLock(ctx, p.d.DB(), accountID, true)
	return
}

// AccountLock 解锁用户
func (p *Service) AccountUnlock(ctx context.Context, accountID int64) (err error) {
	err = p.d.AccountSetLock(ctx, p.d.DB(), accountID, false)
	return
}
