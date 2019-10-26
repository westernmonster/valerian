package service

import (
	"context"
	"fmt"
	"strings"

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
		IsVIP:        bool(account.IsVip),
	}
	return
}

func (p *Service) IsEmailExist(c context.Context, email string) (exist bool, err error) {
	// mobile := arg.Prefix + arg.Mobile
	if account, e := p.d.GetAccountByEmail(c, p.d.DB(), email); e != nil {
		return false, e
	} else if account != nil {
		exist = true
		return
	} else {
		exist = false
	}

	return
}

func (p *Service) IsMobileExist(c context.Context, prefix, mobile string) (exist bool, err error) {
	fullMobile := prefix + mobile
	if account, e := p.d.GetAccountByMobile(c, p.d.DB(), fullMobile); e != nil {
		return false, e
	} else if account != nil {
		exist = true
		return
	} else {
		exist = false
	}

	return
}

func (p *Service) BaseInfo(c context.Context, aid int64) (info *model.BaseInfo, err error) {
	var account *model.Account
	if account, err = p.getAccountByID(c, aid); err != nil {
		return
	}

	info = baseInfoFromAccount(account)
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

func (p *Service) GetSelfProfile(c context.Context, accountID int64) (profile *model.SelfProfile, err error) {
	var item *model.Account
	if item, err = p.getAccountByID(c, accountID); err != nil {
		return
	}

	profile = &model.SelfProfile{
		ID:           item.ID,
		Prefix:       item.Prefix,
		Mobile:       item.Mobile,
		Email:        item.Email,
		UserName:     item.UserName,
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
		IsVIP:        bool(item.IsVip),
		Role:         item.Role,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	if item.Mobile != "" {
		profile.Mobile = strings.TrimLeft(item.Mobile, item.Prefix)
	}

	if item.Location != 0 {
		if v, e := p.getLocationString(c, item.Location); e != nil {
			return nil, e
		} else {
			profile.LocationString = v
		}
	}

	return
}

func (p *Service) GetMemberProfile(c context.Context, accountID int64) (profile *model.ProfileInfo, err error) {
	var item *model.Account
	if item, err = p.getAccountByID(c, accountID); err != nil {
		return
	}

	profile = &model.ProfileInfo{
		ID:           item.ID,
		UserName:     item.UserName,
		Gender:       item.Gender,
		Location:     item.Location,
		Introduction: item.Introduction,
		Avatar:       item.Avatar,
		IDCert:       bool(item.IDCert),
		WorkCert:     bool(item.WorkCert),
		IsOrg:        bool(item.IsOrg),
		IsVIP:        bool(item.IsVip),
		Role:         item.Role,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	if item.Location != 0 {
		if v, e := p.getLocationString(c, item.Location); e != nil {
			return nil, e
		} else {
			profile.LocationString = v
		}
	}

	return
}

func (p *Service) getLocationString(c context.Context, nodeID int64) (locationString string, err error) {
	arr := []string{}

	id := nodeID
	var item *model.Area
	for {
		if item, err = p.d.GetArea(c, p.d.DB(), id); err != nil {
			return
		} else if item == nil {
			err = ecode.AreaNotExist
			return
		}

		arr = append(arr, item.Name)

		if item.Parent == 0 {
			break
		}

		id = item.Parent
	}

	locationString = ""

	for i := len(arr) - 1; i >= 0; i-- {
		locationString += arr[i] + " "
	}

	return
}

func (p *Service) GetAccountByID(c context.Context, aid int64) (account *model.Account, err error) {
	return p.getAccountByID(c, aid)
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

func (p *Service) GetAllAccounts(c context.Context) (items []*model.Account, err error) {
	return p.d.GetAccounts(c, p.d.DB())
}
