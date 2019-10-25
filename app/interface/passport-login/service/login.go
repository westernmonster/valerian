package service

import (
	"context"
	"strings"

	"valerian/app/interface/passport-login/model"
	account "valerian/app/service/account/api"
	"valerian/library/ecode"
)

func (p *Service) EmailLogin(ctx context.Context, req *model.ArgEmailLogin) (resp *model.LoginResp, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	account, err := p.d.GetAccountByEmail(ctx, p.d.DB(), req.Email)
	if err != nil {
		return
	}
	if account == nil {
		err = ecode.UserNotExist
		return
	}

	if err = p.checkPassword(req.Password, account.Password, account.Salt); err != nil {
		return
	}

	accessToken, refreshToken, err := p.grantToken(ctx, req.ClientID, account.ID)
	if err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID:    account.ID,
		Role:         account.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    _accessExpireSeconds,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	if resp.Profile, err = p.GetProfile(ctx, account.ID); err != nil {
		return
	}

	p.addCache(func() {
		p.d.SetProfileCache(context.TODO(), resp.Profile)
	})

	return
}

func (p *Service) MobileLogin(ctx context.Context, req *model.ArgMobileLogin) (resp *model.LoginResp, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	mobile := req.Prefix + req.Mobile
	account, err := p.d.GetAccountByMobile(ctx, p.d.DB(), mobile)
	if err != nil {
		return
	}
	if account == nil {
		err = ecode.UserNotExist
		return
	}

	if err = p.checkPassword(req.Password, account.Password, account.Salt); err != nil {
		return
	}

	accessToken, refreshToken, err := p.grantToken(ctx, req.ClientID, account.ID)
	if err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID:    account.ID,
		Role:         account.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    _accessExpireSeconds,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	if resp.Profile, err = p.GetProfile(ctx, account.ID); err != nil {
		return
	}

	p.addCache(func() {
		p.d.SetProfileCache(context.TODO(), resp.Profile)
	})

	return
}

func (p *Service) DigitLogin(ctx context.Context, req *model.ArgDigitLogin) (resp *model.LoginResp, err error) {
	mobile := req.Prefix + req.Mobile

	var code string
	if code, err = p.d.MobileValcodeCache(ctx, model.ValcodeLogin, mobile); err != nil {
		return
	} else if code == "" {
		err = ecode.ValcodeExpires
		return
	} else if code != req.Valcode {
		err = ecode.ValcodeWrong
		return
	}

	var account *model.Account
	if account, err = p.d.GetAccountByMobile(ctx, p.d.DB(), mobile); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	accessToken, refreshToken, err := p.grantToken(ctx, req.ClientID, account.ID)
	if err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID:    account.ID,
		Role:         account.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    _accessExpireSeconds,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	if resp.Profile, err = p.GetProfile(ctx, account.ID); err != nil {
		return
	}

	p.addCache(func() {
		p.d.SetProfileCache(context.TODO(), resp.Profile)
		p.d.DelMobileValcodeCache(context.TODO(), model.ValcodeLogin, mobile)
	})

	return
}

func (p *Service) checkPassword(password, dbPassword, dbSalt string) (err error) {
	passwordHash, err := hashPassword(password, dbSalt)
	if err != nil {
		return
	}

	if !strings.EqualFold(dbPassword, passwordHash) {
		err = ecode.PasswordErr
		return
	}
	return
}

func (p *Service) GetProfile(c context.Context, aid int64) (item *model.Profile, err error) {
	var profile *account.SelfProfile
	if profile, err = p.d.GetSelfProfile(c, aid); err != nil {
		return
	}

	item = &model.Profile{
		ID:             profile.ID,
		Mobile:         profile.Mobile,
		Email:          profile.Email,
		UserName:       profile.UserName,
		Gender:         int(profile.Gender),
		BirthYear:      int(profile.BirthYear),
		BirthMonth:     int(profile.BirthMonth),
		BirthDay:       int(profile.BirthDay),
		Location:       profile.Location,
		LocationString: profile.LocationString,
		Introduction:   profile.Introduction,
		Avatar:         profile.Avatar,
		Source:         int(profile.Source),
		IDCert:         profile.IDCert,
		IDCertStatus:   int(profile.IDCertStatus),
		WorkCert:       profile.WorkCert,
		WorkCertStatus: int(profile.WorkCertStatus),
		IsOrg:          profile.IsOrg,
		IsVIP:          profile.IsVIP,
		Role:           profile.Role,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}

	item.Stat = &model.ProfileStat{
		FansCount:       int(profile.Stat.FansCount),
		FollowingCount:  int(profile.Stat.FollowingCount),
		TopicCount:      int(profile.Stat.TopicCount),
		ArticleCount:    int(profile.Stat.ArticleCount),
		DiscussionCount: int(profile.Stat.DiscussionCount),
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
