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
	uuid "github.com/satori/go.uuid"
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
		Role:         account.Role,
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
func (p *Service) BaseInfo(c context.Context, aid int64) (info *model.BaseInfo, err error) {
	var account *model.Account
	if account, err = p.getAccountByID(c, p.d.DB(), aid); err != nil {
		return
	}

	info = baseInfoFromAccount(account)
	return
}

// BatchBaseInfo 批量获取账户基本信息
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

// ForgetPassword 忘记密码
// 匹配验证码并生成一个有效期为5分钟的 SESSIONID
func (p *Service) ForgetPassword(c context.Context, arg *api.ForgetPasswordReq) (resp *api.ForgetPasswordResp, err error) {
	var account *model.Account
	if govalidator.IsEmail(arg.Identity) {
		if account, err = p.d.GetAccountByEmail(c, p.d.DB(), arg.Identity); err != nil {
			return
		}
		var code string
		if code, err = p.d.EmailValcodeCache(c, model.ValcodeForgetPassword, arg.Identity); err != nil {
			return
		} else if code == "" {
			err = ecode.ValcodeExpires
			return
		} else if code != arg.Valcode {
			err = ecode.ValcodeWrong
			return
		}
	} else {
		mobile := arg.Prefix + arg.Identity
		if account, err = p.d.GetAccountByMobile(c, p.d.DB(), mobile); err != nil {
			return
		}

		var code string
		if code, err = p.d.MobileValcodeCache(c, model.ValcodeForgetPassword, mobile); err != nil {
			return
		} else if code == "" {
			err = ecode.ValcodeExpires
			return
		} else if code != arg.Valcode {
			err = ecode.ValcodeWrong
			return
		}
	}

	sessionID := uuid.NewV4().String()
	if err = p.d.SetSessionResetPasswordCache(c, sessionID, account.ID); err != nil {
		return
	}

	resp = &api.ForgetPasswordResp{
		SessionID: sessionID,
	}
	return
}

// ResetPassword 重设密码
func (p *Service) ResetPassword(c context.Context, arg *api.ResetPasswordReq) (err error) {
	var aid int64
	if aid, err = p.d.SessionResetPasswordCache(c, arg.SessionID); err != nil {
		return
	} else if aid == 0 {
		return ecode.SessionExpires
	}

	var acc *model.Account
	if acc, err = p.getAccountByID(c, p.d.DB(), aid); err != nil {
		return
	}

	passwordHash, err := hashPassword(arg.Password, acc.Salt)
	if err != nil {
		return
	}

	if err = p.d.SetPassword(c, p.d.DB(), passwordHash, acc.Salt, aid); err != nil {
		return
	}

	p.addCache(func() {
		// TODO: Clear this users's AccessToken Cached && Refresh Token Cache
		p.d.DelAccountCache(context.TODO(), aid)
		p.d.DelResetPasswordCache(context.TODO(), arg.SessionID)
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

	if arg.UserName != nil {
		account.UserName = arg.GetUserNameValue()
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

	if arg.Password != nil {
		passwordHash, e := hashPassword(arg.GetPasswordValue(), account.Salt)
		if e != nil {
			return e
		}

		account.Password = passwordHash

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
