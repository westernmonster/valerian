package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/account/model"
	"valerian/library/ecode"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func validateBirthDay(arg *model.ArgUpdateProfile) (err error) {
	if arg.BirthYear != nil {
		year := *arg.BirthYear
		if year < 1920 || year > time.Now().Year() {
			return ecode.InvalidBirthYear
		}
	}

	if arg.BirthMonth != nil {
		month := *arg.BirthMonth
		if month < 1 || month > 12 {
			return ecode.InvalidBirthMonth
		}
	}

	if arg.BirthDay != nil {
		day := *arg.BirthDay
		if day < 1 || day > 30 {
			return ecode.InvalidBirthDay
		}
	}

	return
}

func (p *Service) ForgetPassword(c context.Context, arg *model.ArgForgetPassword) (sessionID string, err error) {
	var account *model.Account
	if govalidator.IsEmail(arg.Identity) {
		if account, err = p.d.GetAccountByEmail(c, p.d.DB(), arg.Identity); err != nil {
			return
		} else if account == nil {
			err = ecode.UserNotExist
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
		if account, err = p.d.GetAccountByMobile(c, p.d.DB(), arg.Identity); err != nil {
			return
		} else if account == nil {
			err = ecode.UserNotExist
			return
		}

		var code string
		if code, err = p.d.MobileValcodeCache(c, model.ValcodeForgetPassword, arg.Identity); err != nil {
			return
		} else if code == "" {
			err = ecode.ValcodeExpires
			return
		} else if code != arg.Valcode {
			err = ecode.ValcodeWrong
			return
		}
	}

	sessionID = uuid.NewV4().String()
	if err = p.d.SetSessionResetPasswordCache(c, sessionID, account.ID); err != nil {
		return
	}

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

func (p *Service) ResetPassword(c context.Context, arg *model.ArgResetPassword) (err error) {
	var aid int64
	if aid, err = p.d.SessionResetPasswordCache(c, arg.SessionID); err != nil {
		return
	} else if aid == 0 {
		return ecode.SessionExpires
	}

	if _, err = p.getAccountByID(c, aid); err != nil {
		return
	}

	salt, err := generateSalt(16)
	if err != nil {
		return
	}

	passwordHash, err := hashPassword(arg.Password, salt)
	if err != nil {
		return
	}

	if err = p.d.SetPassword(c, p.d.DB(), salt, passwordHash, aid); err != nil {
		return
	}

	p.addCache(func() {
		// TODO: Clear this users's AccessToken Cached && Refresh Token Cache
		p.d.DelAccountCache(context.TODO(), aid)
		p.d.DelResetPasswordCache(context.TODO(), arg.SessionID)
	})

	return
}

func (p *Service) UpdateProfile(c context.Context, aid int64, arg *model.ArgUpdateProfile) (err error) {
	var account *model.Account
	if account, err = p.getAccountByID(c, aid); err != nil {
		return
	}

	if arg.Gender != nil {
		if *arg.Gender != model.GenderMale && *arg.Gender != model.GenderFemale {
			return ecode.InvalidGender
		}
		account.Gender = arg.Gender
	}

	if arg.Avatar != nil {
		if !govalidator.IsURL(*arg.Avatar) {
			return ecode.InvalidAvatar
		}
		account.Avatar = *arg.Avatar
	}

	if arg.Introduction != nil {
		account.Introduction = arg.Introduction
	}

	if arg.BirthYear != nil {
		account.BirthYear = arg.BirthYear
	}

	if arg.BirthMonth != nil {
		account.BirthMonth = arg.BirthMonth
	}

	if arg.BirthDay != nil {
		account.BirthDay = arg.BirthDay
	}

	if err = validateBirthDay(arg); err != nil {
		return
	}

	if arg.Password != nil {
		salt, e := generateSalt(16)
		if e != nil {
			return e
		}
		passwordHash, e := hashPassword(*arg.Password, salt)
		if e != nil {
			return e
		}

		account.Password = passwordHash
		account.Salt = salt

	}

	if arg.Location != nil {
		if item, e := p.d.GetArea(c, p.d.DB(), *arg.Location); e != nil {
			return e
		} else if item == nil {
			return ecode.AreaNotExist
		}

		account.Location = arg.Location
	}

	if err = p.d.UpdateAccount(c, p.d.DB(), account); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelAccountCache(context.TODO(), aid)
		p.onAccountUpdated(context.TODO(), aid, time.Now().Unix())
	})

	return
}

func (p *Service) ChangePassword(c context.Context, aid int64, arg *model.ArgChangePassword) (err error) {
	salt, err := generateSalt(16)
	if err != nil {
		return
	}
	passwordHash, err := hashPassword(arg.Password, salt)
	if err != nil {
		return
	}

	fmt.Println(salt)
	fmt.Println(arg.Password)
	fmt.Println(passwordHash)
	if err = p.d.SetPassword(c, p.d.DB(), salt, passwordHash, aid); err != nil {
		return
	}

	return
}

func (p *Service) GetProfile(c context.Context, aid int64) (profile *model.Profile, err error) {
	if profile, err = p.getProfile(c, aid); err != nil {
		return
	}

	var isFollowing bool
	if isFollowing, err = p.d.IsFollowing(c, aid, aid); err != nil {
		return
	}

	var stat *model.AccountStat
	if stat, err = p.d.GetAccountStatByID(c, p.d.DB(), aid); err != nil {
		return
	}

	profile.Stat = &model.MemberInfoStat{
		FansCount:       int(stat.Fans),
		FollowingCount:  int(stat.Following),
		TopicCount:      int(stat.TopicCount),
		ArticleCount:    int(stat.ArticleCount),
		DiscussionCount: int(stat.DiscussionCount),
		IsFollow:        isFollowing,
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), aid); err != nil {
		return
	}

	profile.Settings = setting

	return
}

func (p *Service) getProfile(c context.Context, accountID int64) (profile *model.Profile, err error) {
	var item *model.Account
	if item, err = p.getAccountByID(c, accountID); err != nil {
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
		Role:         item.Role,
		UserName:     item.UserName,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	ipStr := InetNtoA(item.IP)
	profile.IP = &ipStr

	if item.Location != nil {
		if v, e := p.getLocationString(c, *item.Location); e != nil {
			return nil, e
		} else {
			profile.LocationString = &v
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
