package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/account/api"
	"valerian/app/service/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/log"

	"github.com/asaskevich/govalidator"
)

func validateBirthDay(arg *api.UpdateProfileReq) (err error) {
	if arg.BirthYear != nil {
		year := int(arg.GetBirthYearValue())
		if year < 1920 || year > time.Now().Year() {
			return ecode.InvalidBirthYear
		}
	}

	if arg.BirthMonth != nil {
		month := arg.GetBirthMonthValue()
		if month < 1 || month > 12 {
			return ecode.InvalidBirthMonth
		}
	}

	if arg.BirthDay != nil {
		day := arg.GetBirthDayValue()
		if day < 1 || day > 30 {
			return ecode.InvalidBirthDay
		}
	}

	return
}

func baseInfoFromAccount(acc *model.Account, stat *model.AccountStat) (info *api.BaseInfoReply) {
	info = &api.BaseInfoReply{
		ID:           acc.ID,
		UserName:     acc.UserName,
		Gender:       acc.Gender,
		Avatar:       acc.Avatar,
		Introduction: acc.Introduction,
		IDCert:       bool(acc.IDCert),
		WorkCert:     bool(acc.WorkCert),
		IsOrg:        bool(acc.IsOrg),
		IsVIP:        bool(acc.IsVip),
		Role:         acc.Role,
		Stat: &api.AccountStatInfo{
			FollowingCount: stat.Following,
			FansCount:      stat.Fans,
			BlackCount:     stat.Black,
			TopicCount:     stat.TopicCount,
			ArticleCount:   stat.ArticleCount,
		},
	}
	return
}

// GetAccountByEmail 通过 Email 获取账户信息
// 会排除已经注销的用户
func (p *Service) GetAccountByEmail(c context.Context, email string) (item *model.Account, err error) {
	if item, err = p.d.GetAccountByEmail(c, p.d.DB(), email); err != nil {
		return
	} else if item == nil {
		err = ecode.UserNotExist
		return
	}
	return
}

// GetAccountByMobile 通过手机号获取账户
// 会排除已经注销的账户
func (p *Service) GetAccountByMobile(c context.Context, prefix, mobile string) (item *model.Account, err error) {
	fullMobile := prefix + mobile
	if item, err = p.d.GetAccountByMobile(c, p.d.DB(), fullMobile); err != nil {
		return
	} else if item == nil {
		err = ecode.UserNotExist
		return
	}
	return
}

// IsEmailExist 邮件是否已经注册
// 会排除已经注销的账户
func (p *Service) IsEmailExist(c context.Context, email string) (exist bool, err error) {
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

// IsMobileExist 手机号是否已经注册
// 会排除已经注销的账户
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

// BaseInfo 获取账户基本信息
func (p *Service) BaseInfo(c context.Context, aid int64) (info *api.BaseInfoReply, err error) {
	var account *model.Account
	if account, err = p.getAccountByID(c, p.d.DB(), aid); err != nil {
		return
	}

	var stat *model.AccountStat
	if stat, err = p.getAccountStat(c, p.d.DB(), aid); err != nil {
		return
	}

	info = baseInfoFromAccount(account, stat)
	return
}

// BatchBaseInfo 批量获取账户基本信息
func (p *Service) BatchBaseInfo(c context.Context, aids []int64) (data map[int64]*api.BaseInfoReply, err error) {
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

	data = make(map[int64]*api.BaseInfoReply)
	for k, v := range res {
		var stat *model.AccountStat
		if stat, err = p.getAccountStat(c, p.d.DB(), v.ID); err != nil {
			return
		}
		data[k] = baseInfoFromAccount(v, stat)
	}

	if len(missA) == 0 {
		return
	}

	p.addCache(func() {
		p.d.SetBatchAccountCache(context.TODO(), missA)
	})

	return
}

// UpdateAccount 更新用户资料
func (p *Service) UpdateAccount(c context.Context, arg *api.UpdateProfileReq) (err error) {
	var account *model.Account
	if account, err = p.getAccountByID(c, p.d.DB(), arg.Aid); err != nil {
		return
	}

	if arg.Gender != nil {
		if arg.GetGenderValue() != model.GenderMale && arg.GetGenderValue() != model.GenderFemale {
			return ecode.InvalidGender
		}
		account.Gender = arg.GetGenderValue()
	}

	if arg.Avatar != nil {
		if !govalidator.IsURL(arg.GetAvatarValue()) {
			return ecode.InvalidAvatar
		}
		account.Avatar = arg.GetAvatarValue()
	}

	if arg.Introduction != nil {
		account.Introduction = arg.GetIntroductionValue()
	}

	if !account.IDCert {
		if arg.UserName != nil {
			account.UserName = arg.GetUserNameValue()
		}
	}

	if arg.BirthYear != nil {
		account.BirthYear = arg.GetBirthYearValue()
	}

	if arg.BirthMonth != nil {
		account.BirthMonth = arg.GetBirthMonthValue()
	}

	if arg.BirthDay != nil {
		account.BirthDay = arg.GetBirthDayValue()
	}

	if err = validateBirthDay(arg); err != nil {
		return
	}

	if arg.Location != nil {
		if item, e := p.d.GetArea(c, p.d.DB(), arg.GetLocationValue()); e != nil {
			return e
		} else if item == nil {
			return ecode.AreaNotExist
		}

		account.Location = arg.GetLocationValue()
	}

	if err = p.d.UpdateAccount(c, p.d.DB(), account); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelAccountCache(context.TODO(), arg.Aid)
		p.onAccountUpdated(context.TODO(), arg.Aid, time.Now().Unix())
	})

	return
}

// GetAccountByID 通过ID获取用户
func (p *Service) GetAccountByID(c context.Context, aid int64) (account *model.Account, err error) {
	return p.getAccountByID(c, p.d.DB(), aid)
}

// getAccountByID 通过ID获取用户
func (p *Service) getAccountByID(c context.Context, node sqalx.Node, aid int64) (account *model.Account, err error) {
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
			p.d.SetAccountCache(context.TODO(), account)
		})
	}
	return
}
