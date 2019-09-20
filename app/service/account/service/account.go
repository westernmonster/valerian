package service

import (
	"context"
	"fmt"

	"valerian/app/service/account/model"
	"valerian/library/ecode"
	"valerian/library/log"
)

func baseInfoFromAccount(account *model.Account) (info *model.BaseInfo) {
	info = &model.BaseInfo{
		ID:           account.ID,
		UserName:     account.UserName,
		Gender:       account.Gender,
		Avatar:       account.Avatar,
		Introduction: account.Introduction,
		IDCert:       bool(account.IDCert),
		WorkCert:     bool(account.WorkCert),
		IsOrg:        bool(account.IsOrg),
		IsVIP:        bool(account.IsVIP),
	}
	return
}

func (p *Service) BaseInfo(c context.Context, aid int64) (info *model.BaseInfo, err error) {
	var needCache = true

	var account *model.Account
	if account, err = p.d.AccountCache(c, aid); err != nil {
		needCache = false
	} else if account != nil {
		info = baseInfoFromAccount(account)
		return
	}

	if account, err = p.d.GetAccountByID(c, p.d.DB(), aid); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	info = baseInfoFromAccount(account)

	if needCache {
		p.addCache(func() {
			p.d.SetAccountCache(context.TODO(), account)
		})
	}
	return
}

func (p *Service) BatchBaseInfo(c context.Context, aids []int64) (data map[int64]*model.BaseInfo, err error) {
	if len(aids) > 100 {
		err = ecode.MemberOverLimit
		return
	}

	var (
		missed  []int64
		account *model.Account
		res     map[int64]*model.Account
	)
	if res, missed, err = p.d.BatchAccountCache(c, aids); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.BatchAccountCache err(%v)", err))
		missed = aids
	}

	var missA []*model.Account
	for _, aid := range missed {
		if account, err = p.d.GetAccountByID(c, p.d.DB(), aid); err != nil {
			return
		} else if account == nil {
			err = ecode.UserNotExist
			return
		} else {
			res[aid] = account
			missA = append(missA, account)
		}
	}

	data = make(map[int64]*model.BaseInfo)
	for k, v := range res {
		data[k] = baseInfoFromAccount(v)
	}

	if len(missA) == 0 {
		return
	}

	p.addCache(func() {
		p.d.SetBatchAccountCache(context.TODO(), missA)
	})

	return
}
