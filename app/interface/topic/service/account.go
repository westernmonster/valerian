package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) getAccountByID(c context.Context, node sqalx.Node, aid int64) (account *model.Account, err error) {
	var (
		addCache = true
	)

	if account, err = p.d.AccountCache(c, aid); err != nil {
		addCache = false
	} else if account != nil {
		return
	}

	if account, err = p.d.GetAccountByID(c, node, aid); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetAccountCache(context.TODO(), account)
		})
	}
	return
}
