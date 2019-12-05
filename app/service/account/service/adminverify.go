package service

import (
	"context"
)

func (p *Service) AccountLock(ctx context.Context, accountID int64) (err error) {
	err = p.d.AccountSetLock(ctx, p.d.DB(), accountID, true)
	return
}

func (p *Service) AccountUnlock(ctx context.Context, accountID int64) (err error) {
	err = p.d.AccountSetLock(ctx, p.d.DB(), accountID, false)
	return
}
