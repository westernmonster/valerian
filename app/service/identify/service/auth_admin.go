package service

import (
	"context"
	"strconv"

	"valerian/app/service/identify/model"
	"valerian/library/ecode"
)

func (p *Service) AuthAdmin(c context.Context, reqSID string) (sid string, aid int64, uname string, err error) {
	si := p.session(c, reqSID)
	if si.Get(_sessUIDKey) == nil {
		err = ecode.Unauthorized
		return
	}

	sid = si.Sid
	uidStr := si.Get(_sessUIDKey).(string)
	if aid, err = strconv.ParseInt(uidStr, 10, 64); err != nil {
		return
	}

	var u *model.Account
	if u, err = p.getAccountByID(c, aid); err != nil {
		return
	}

	aid = u.ID
	uname = si.Get(_sessUnameKey).(string)

	return
}

func (p *Service) getAccountByID(c context.Context, aid int64) (account *model.Account, err error) {
	var needCache = true

	if account, err = p.d.AccountCache(c, aid); err != nil {
		needCache = false
	} else if account != nil {
		return
	}

	if account, err = p.d.GetAccountByID(c, p.d.DB(), aid); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	if needCache {
		p.addCache(func() {
			p.d.SetAccountCache(context.TODO(), account)
		})
	}
	return
}
