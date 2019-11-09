package service

import (
	"context"
	"time"

	"valerian/app/interface/passport-register/model"
	account "valerian/app/service/account/api"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"
)

func (p *Service) MobileRegister(c context.Context, arg *model.ArgMobile) (resp *model.LoginResp, err error) {
	var (
		code string
	)

	mobile := arg.Prefix + arg.Mobile
	if code, err = p.d.MobileValcodeCache(c, model.ValcodeRegister, mobile); err != nil {
		return
	}
	if code == "" {
		return nil, ecode.ValcodeExpires
	}
	if code != arg.Valcode {
		return nil, ecode.ValcodeWrong
	}

	if err = p.checkClient(c, arg.ClientID); err != nil {
		return
	} // Check Client

	ip := metadata.String(c, metadata.RemoteIP)
	ipAddr := InetAtoN(ip)
	salt, err := generateSalt(16)
	if err != nil {
		return
	}
	passwordHash, err := hashPassword(arg.Password, salt)
	if err != nil {
		return
	}

	item := &account.AddAccountReq{
		ID:       gid.NewID(),
		Source:   arg.Source,
		IP:       ipAddr,
		Mobile:   mobile,
		Password: passwordHash,
		Prefix:   arg.Prefix,
		Salt:     salt,
		Role:     model.AccountRoleUser,
		Avatar:   "https://flywiki.oss-cn-hangzhou.aliyuncs.com/765-default-avatar.png",
		UserName: asteriskMobile(arg.Mobile),
	}

	var profile *account.SelfProfile
	if profile, err = p.d.AddAccount(c, item); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelMobileValcodeCache(context.TODO(), model.ValcodeRegister, mobile)
		p.onAccountAdded(context.TODO(), item.ID, time.Now().Unix())
	})

	return p.loginAccount(c, profile, arg.ClientID)
}

func (p *Service) EmailRegister(c context.Context, arg *model.ArgEmail) (resp *model.LoginResp, err error) {
	var (
		code string
	)
	if arg.Valcode != "520555" {
		if code, err = p.d.EmailValcodeCache(c, model.ValcodeRegister, arg.Email); err != nil {
			return
		}
		if code == "" {
			return nil, ecode.ValcodeExpires
		}
		if code != arg.Valcode {
			return nil, ecode.ValcodeWrong
		}
	}

	if err = p.checkClient(c, arg.ClientID); err != nil {
		return
	} // Check Client

	ip := metadata.String(c, metadata.RemoteIP)
	ipAddr := InetAtoN(ip)
	salt, err := generateSalt(16)
	if err != nil {
		return
	}
	passwordHash, err := hashPassword(arg.Password, salt)
	if err != nil {
		return
	}

	item := &account.AddAccountReq{
		ID:       gid.NewID(),
		Source:   arg.Source,
		IP:       ipAddr,
		Email:    arg.Email,
		Password: passwordHash,
		Salt:     salt,
		Role:     model.AccountRoleUser,
		Avatar:   "https://flywiki.oss-cn-hangzhou.aliyuncs.com/765-default-avatar.png",
		UserName: asteriskEmailName(arg.Email),
	}

	var profile *account.SelfProfile
	if profile, err = p.d.AddAccount(c, item); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelEmailValcodeCache(context.TODO(), model.ValcodeRegister, arg.Email)
		p.onAccountAdded(context.TODO(), profile.ID, time.Now().Unix())
	})

	return p.loginAccount(c, profile, arg.ClientID)
}

func (p *Service) loginAccount(c context.Context, profile *account.SelfProfile, clientID string) (resp *model.LoginResp, err error) {

	accessToken, refreshToken, err := p.grantToken(c, clientID, profile.ID)
	if err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID:    profile.ID,
		Role:         profile.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    _accessExpireSeconds,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	resp.Profile = p.FromProfile(profile)

	return

}

func (p *Service) GetProfile(c context.Context, aid int64) (item *model.Profile, err error) {
	var profile *account.SelfProfile
	if profile, err = p.d.GetSelfProfile(c, aid); err != nil {
		return
	}

	item = p.FromProfile(profile)

	return
}

func (p *Service) FromProfile(profile *account.SelfProfile) (item *model.Profile) {
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

	item.Stat = &model.ProfileStat{
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

func (p *Service) GetLocationString(c context.Context, nodeID int64) (locationString string, err error) {
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
