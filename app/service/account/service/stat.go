package service

import (
	"context"
	"time"
	"valerian/app/service/account/model"
	"valerian/library/database/sqalx"
)

// GetAccountStat 获取用户状态
func (p *Service) GetAccountStat(c context.Context, accountID int64) (item *model.AccountStat, err error) {
	return p.getAccountStat(c, p.d.DB(), accountID)
}

// getAccountStat 获取用户状态
func (p *Service) getAccountStat(c context.Context, node sqalx.Node, accountID int64) (item *model.AccountStat, err error) {
	if item, err = p.d.GetAccountStatByID(c, node, accountID); err != nil {
		return
	} else if item == nil {
		item = &model.AccountStat{
			AccountID: accountID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
	}

	return
}
