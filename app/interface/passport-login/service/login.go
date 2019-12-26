package service

import (
	"context"

	"valerian/app/interface/passport-login/model"
	account "valerian/app/service/account/api"
	identify "valerian/app/service/identify/api/grpc"
)

// EmailLogin 邮件登录
func (p *Service) EmailLogin(ctx context.Context, req *model.ArgEmailLogin) (resp *model.LoginResp, err error) {
	arg := &identify.EmailLoginReq{
		Source:   req.Source,
		Email:    req.Email,
		Password: req.Password,
		ClientID: req.ClientID,
	}

	var data *identify.LoginResp
	if data, err = p.d.EmailLogin(ctx, arg); err != nil {
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

	if resp.Profile, err = p.GetProfile(ctx, data.Aid); err != nil {
		return
	}

	return
}

// MobileLogin 手机登录
func (p *Service) MobileLogin(ctx context.Context, req *model.ArgMobileLogin) (resp *model.LoginResp, err error) {
	arg := &identify.MobileLoginReq{
		Source:   req.Source,
		Mobile:   req.Mobile,
		Prefix:   req.Prefix,
		Password: req.Password,
		ClientID: req.ClientID,
	}

	var data *identify.LoginResp
	if data, err = p.d.MobileLogin(ctx, arg); err != nil {
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

	if resp.Profile, err = p.GetProfile(ctx, data.Aid); err != nil {
		return
	}

	return
}

// DigitLogin 验证码登录
func (p *Service) DigitLogin(ctx context.Context, req *model.ArgDigitLogin) (resp *model.LoginResp, err error) {
	arg := &identify.DigitLoginReq{
		Source:   req.Source,
		Prefix:   req.Prefix,
		Valcode:  req.Valcode,
		Mobile:   req.Mobile,
		ClientID: req.ClientID,
	}

	var data *identify.LoginResp
	if data, err = p.d.DigitLogin(ctx, arg); err != nil {
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

	if resp.Profile, err = p.GetProfile(ctx, data.Aid); err != nil {
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
