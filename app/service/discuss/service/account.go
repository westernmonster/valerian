package service

import (
	"context"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// getAccountByID 通过ID获取用户
func (p *Service) getAccount(c context.Context, node sqalx.Node, aid int64) (account *model.Account, err error) {
	if account, err = p.d.AccountCache(c, aid); err != nil {
	} else if account != nil {
		return
	}
	if account, err = p.d.GetAccountByID(c, node, aid); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	return
}
