package service

import (
	"context"
	"valerian/app/service/topic-feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) getAccount(c context.Context, node sqalx.Node, aid int64) (account *model.Account, err error) {
	var needCache = true

	if account, err = p.d.AccountCache(c, aid); err != nil {
		needCache = false
	} else if account != nil {
		return
	}

	if account, err = p.d.GetAccountByID(c, node, aid); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	if needCache {
		p.addCache(func() {
			p.d.SetAccountCache(context.Background(), account)
		})
	}
	return
}
