package service

import (
	"context"
	"valerian/app/admin/account/model"
	"valerian/library/ecode"
)

func (p *Service) GetProfile(c context.Context, accountID int64) (profile *model.Profile, err error) {
	var item *model.Account
	if item, err = p.getAccountByID(c, accountID); err != nil {
		return
	} else if item == nil {
		err = ecode.UserNotExist
		return
	}

	profile = &model.Profile{
		ID:           item.ID,
		Mobile:       item.Mobile,
		Email:        item.Email,
		Gender:       item.Gender,
		BirthYear:    item.BirthYear,
		BirthMonth:   item.BirthMonth,
		BirthDay:     item.BirthDay,
		Location:     item.Location,
		Introduction: item.Introduction,
		Avatar:       item.Avatar,
		Source:       item.Source,
		IDCert:       bool(item.IDCert),
		WorkCert:     bool(item.WorkCert),
		IsOrg:        bool(item.IsOrg),
		IsVIP:        bool(item.IsVIP),
		Role:         item.Role, UserName: item.UserName,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	profile.IP = InetNtoA(item.IP)
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
