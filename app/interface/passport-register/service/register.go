package service

import (
	"context"

	"valerian/app/interface/passport-register/model"
	account "valerian/app/service/account/api"
	identify "valerian/app/service/identify/api/grpc"
)

// MobileRegister 手机注册
func (p *Service) MobileRegister(c context.Context, req *model.ArgMobile) (resp *model.LoginResp, err error) {
	arg := &identify.MobileRegisterReq{
		Source:   req.Source,
		Mobile:   req.Mobile,
		Prefix:   req.Prefix,
		Password: req.Password,
		ClientID: req.ClientID,
		Valcode:  req.Valcode,
	}

	var data *identify.LoginResp
	if data, err = p.d.MobileRegister(c, arg); err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID:    data.Aid,
		Role:         data.Role,
		AccessToken:  data.AccessToken,
		ExpiresIn:    data.ExpiresIn,
		TokenType:    data.TokenType,
		Scope:        data.Scope,
		RefreshToken: data.RefreshToken,
	}

	if resp.Profile, err = p.GetProfile(c, data.Aid); err != nil {
		return
	}

	return
}

// EmailRegister 邮箱注册
func (p *Service) EmailRegister(c context.Context, req *model.ArgEmail) (resp *model.LoginResp, err error) {
	arg := &identify.EmailRegisterReq{
		Source:   req.Source,
		Email:    req.Email,
		Valcode:  req.Valcode,
		Password: req.Password,
		ClientID: req.ClientID,
	}

	var data *identify.LoginResp
	if data, err = p.d.EmailRegister(c, arg); err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID:    data.Aid,
		Role:         data.Role,
		AccessToken:  data.AccessToken,
		ExpiresIn:    data.ExpiresIn,
		TokenType:    data.TokenType,
		Scope:        data.Scope,
		RefreshToken: data.RefreshToken,
	}

	if resp.Profile, err = p.GetProfile(c, data.Aid); err != nil {
		return
	}

	return
}

// GetProfile 获取当前登录用户资料
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
