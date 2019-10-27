package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"

	"valerian/library/database/sqalx"
	"valerian/library/ecode"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func validateBirthDay(arg *model.ArgUpdateProfile) (err error) {
	if arg.BirthYear != nil {
		year := int(*arg.BirthYear)
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

func (p *Service) ResetPassword(c context.Context, arg *model.ArgResetPassword) (err error) {
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

	if err = p.d.SetPassword(c, p.d.DB(), acc.Salt, passwordHash, aid); err != nil {
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
	if account, err = p.getAccountByID(c, p.d.DB(), aid); err != nil {
		return
	}

	if arg.Gender != nil {
		if *arg.Gender != model.GenderMale && *arg.Gender != model.GenderFemale {
			return ecode.InvalidGender
		}
		account.Gender = *arg.Gender
	}

	if arg.Avatar != nil {
		if !govalidator.IsURL(*arg.Avatar) {
			return ecode.InvalidAvatar
		}
		account.Avatar = *arg.Avatar
	}

	if arg.Introduction != nil {
		account.Introduction = *arg.Introduction
	}

	if arg.UserName != nil {
		account.UserName = *arg.UserName
	}

	if arg.BirthYear != nil {
		account.BirthYear = *arg.BirthYear
	}

	if arg.BirthMonth != nil {
		account.BirthMonth = *arg.BirthMonth
	}

	if arg.BirthDay != nil {
		account.BirthDay = *arg.BirthDay
	}

	if err = validateBirthDay(arg); err != nil {
		return
	}

	if arg.Password != nil {
		passwordHash, e := hashPassword(*arg.Password, account.Salt)
		if e != nil {
			return e
		}

		account.Password = passwordHash

	}

	if arg.Location != nil {
		if item, e := p.d.GetArea(c, p.d.DB(), *arg.Location); e != nil {
			return e
		} else if item == nil {
			return ecode.AreaNotExist
		}

		account.Location = *arg.Location
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

	var acc *model.Account
	if acc, err = p.getAccountByID(c, p.d.DB(), aid); err != nil {
		return
	}

	fmt.Printf("arg Password: %s\n", arg.Password)
	fmt.Printf("old Salt: %s\n", acc.Salt)
	fmt.Printf("old Hash: %s\n", acc.Password)
	passwordHash, err := hashPassword(arg.Password, acc.Salt)
	if err != nil {
		return
	}
	fmt.Printf("new Salt: %s\n", acc.Salt)
	fmt.Printf("new Hash: %s\n", passwordHash)

	if err = p.d.SetPassword(c, p.d.DB(), acc.Salt, passwordHash, aid); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelAccountCache(context.TODO(), aid)
	})
	return
}

func (p *Service) GetProfile(c context.Context, aid int64) (item *model.Profile, err error) {
	var profile *account.SelfProfile
	if profile, err = p.d.GetSelfProfile(c, aid); err != nil {
		return
	}

	item = &model.Profile{
		ID:             profile.ID,
		Prefix:         profile.Prefix,
		Mobile:         profile.Mobile,
		Email:          profile.Email,
		UserName:       profile.UserName,
		Gender:         (profile.Gender),
		BirthYear:      (profile.BirthYear),
		BirthMonth:     (profile.BirthMonth),
		BirthDay:       (profile.BirthDay),
		Location:       profile.Location,
		LocationString: profile.LocationString,
		Introduction:   profile.Introduction,
		Avatar:         profile.Avatar,
		Source:         (profile.Source),
		IDCert:         profile.IDCert,
		IDCertStatus:   (profile.IDCertStatus),
		WorkCert:       profile.WorkCert,
		WorkCertStatus: (profile.WorkCertStatus),
		IsOrg:          profile.IsOrg,
		IsVIP:          profile.IsVIP,
		Role:           profile.Role,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}

	item.Stat = &model.MemberInfoStat{
		FansCount:       (profile.Stat.FansCount),
		FollowingCount:  (profile.Stat.FollowingCount),
		TopicCount:      (profile.Stat.TopicCount),
		ArticleCount:    (profile.Stat.ArticleCount),
		DiscussionCount: (profile.Stat.DiscussionCount),
	}

	item.Settings = &model.SettingResp{
		Activity: model.ActivitySettingResp{
			Like:         profile.Setting.ActivityLike,
			Comment:      profile.Setting.ActivityComment,
			FollowTopic:  profile.Setting.ActivityFollowTopic,
			FollowMember: profile.Setting.ActivityFollowMember,
		},
		Notify: model.NotifySettingResp{
			Like:      profile.Setting.NotifyLike,
			Comment:   profile.Setting.NotifyComment,
			NewFans:   profile.Setting.NotifyNewFans,
			NewMember: profile.Setting.NotifyNewMember,
		},
		Language: model.LanguageSettingResp{
			Language: profile.Setting.Language,
		},
	}

	return
}
