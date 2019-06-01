package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) getAccountByID(c context.Context, node sqalx.Node, aid int64) (item *model.Account, err error) {
	var (
		account   *model.Account
		needCache bool
	)

	if account, err = p.d.AccountCache(c, aid); err != nil {
		needCache = true
	}

	if account == nil {
		if account, err = p.d.GetAccountByID(c, node, aid); err != nil {
			return
		} else if account == nil {
			err = ecode.UserNotExist
			return
		}
	}

	if needCache {
		p.addCache(func() {
			p.d.SetAccountCache(context.TODO(), account)
		})
	}
	return
}
