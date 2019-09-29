package service

import (
	"context"
	"time"

	"valerian/app/interface/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetAccountSetting(c context.Context) (resp *model.SettingResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	return p.getAccountSetting(c, p.d.DB(), aid)
}

func (p *Service) getAccountSetting(c context.Context, node sqalx.Node, accountID int64) (resp *model.SettingResp, err error) {
	var setting *model.AccountSetting
	if setting, err = p.d.GetAccountSettingByID(c, node, accountID); err != nil {
		return
	} else if setting == nil {
		setting = &model.AccountSetting{
			AccountID:            accountID,
			ActivityLike:         true,
			ActivityComment:      true,
			ActivityFollowTopic:  true,
			ActivityFollowMember: true,
			NotifyLike:           true,
			NotifyComment:        true,
			NotifyNewFans:        true,
			NotifyNewMember:      true,
			Language:             "zh-CN",
		}
	}

	resp = &model.SettingResp{
		Activity: model.ActivitySettingResp{
			Like:         bool(setting.ActivityLike),
			Comment:      bool(setting.ActivityComment),
			FollowTopic:  bool(setting.ActivityFollowTopic),
			FollowMember: bool(setting.ActivityFollowMember),
		},
		Notify: model.NotifySettingResp{
			Like:      bool(setting.NotifyLike),
			Comment:   bool(setting.NotifyComment),
			NewFans:   bool(setting.NotifyNewFans),
			NewMember: bool(setting.NotifyNewMember),
		},

		Language: model.LanguageSettingResp{
			Language: setting.Language,
		},
	}

	return
}

func (p *Service) UpdateAccountSetting(c context.Context, arg *model.ArgSetting) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	setting := &model.AccountSetting{
		AccountID:            aid,
		ActivityLike:         types.BitBool(arg.Activity.Like),
		ActivityComment:      types.BitBool(arg.Activity.Comment),
		ActivityFollowTopic:  types.BitBool(arg.Activity.FollowTopic),
		ActivityFollowMember: types.BitBool(arg.Activity.FollowMember),
		NotifyLike:           types.BitBool(arg.Notify.Like),
		NotifyComment:        types.BitBool(arg.Notify.Comment),
		NotifyNewFans:        types.BitBool(arg.Notify.NewFans),
		NotifyNewMember:      types.BitBool(arg.Notify.NewMember),
		Language:             arg.Language,
		CreatedAt:            time.Now().Unix(),
		UpdatedAt:            time.Now().Unix(),
	}

	if err = p.d.UpdateAccountSetting(c, p.d.DB(), setting); err != nil {
		return
	}

	return
}

func (p *Service) UpdateActivitySetting(c context.Context, arg *model.ArgActivitySetting) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var setting *model.AccountSetting
	if setting, err = p.d.GetAccountSettingByID(c, p.d.DB(), aid); err != nil {
		return
	} else if setting == nil {
		setting = &model.AccountSetting{
			AccountID:            aid,
			ActivityLike:         true,
			ActivityComment:      true,
			ActivityFollowTopic:  true,
			ActivityFollowMember: true,
			NotifyLike:           true,
			NotifyComment:        true,
			NotifyNewFans:        true,
			NotifyNewMember:      true,
			Language:             "zh-CN",
		}

		if err = p.d.AddAccountSetting(c, p.d.DB(), setting); err != nil {
			return
		}
	}

	setting.ActivityLike = types.BitBool(arg.Like)
	setting.ActivityComment = types.BitBool(arg.Comment)
	setting.ActivityFollowTopic = types.BitBool(arg.FollowTopic)
	setting.ActivityFollowMember = types.BitBool(arg.FollowMember)

	if err = p.d.UpdateAccountSetting(c, p.d.DB(), setting); err != nil {
		return
	}

	return
}

func (p *Service) UpdateNotifySetting(c context.Context, arg *model.ArgNotifySetting) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var setting *model.AccountSetting
	if setting, err = p.d.GetAccountSettingByID(c, p.d.DB(), aid); err != nil {
		return
	} else if setting == nil {
		setting = &model.AccountSetting{
			AccountID:            aid,
			ActivityLike:         true,
			ActivityComment:      true,
			ActivityFollowTopic:  true,
			ActivityFollowMember: true,
			NotifyLike:           true,
			NotifyComment:        true,
			NotifyNewFans:        true,
			NotifyNewMember:      true,
			Language:             "zh-CN",
		}

		if err = p.d.AddAccountSetting(c, p.d.DB(), setting); err != nil {
			return
		}
	}

	setting.NotifyLike = types.BitBool(arg.Like)
	setting.NotifyComment = types.BitBool(arg.Comment)
	setting.NotifyNewFans = types.BitBool(arg.NewFans)
	setting.NotifyNewMember = types.BitBool(arg.NewMember)

	if err = p.d.UpdateAccountSetting(c, p.d.DB(), setting); err != nil {
		return
	}

	return
}

func (p *Service) UpdateLanguageSetting(c context.Context, arg *model.ArgLanguageSetting) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var setting *model.AccountSetting
	if setting, err = p.d.GetAccountSettingByID(c, p.d.DB(), aid); err != nil {
		return
	} else if setting == nil {
		setting = &model.AccountSetting{
			AccountID:            aid,
			ActivityLike:         true,
			ActivityComment:      true,
			ActivityFollowTopic:  true,
			ActivityFollowMember: true,
			NotifyLike:           true,
			NotifyComment:        true,
			NotifyNewFans:        true,
			NotifyNewMember:      true,
			Language:             "zh-CN",
		}

		if err = p.d.AddAccountSetting(c, p.d.DB(), setting); err != nil {
			return
		}
	}

	setting.Language = arg.Language

	if err = p.d.UpdateAccountSetting(c, p.d.DB(), setting); err != nil {
		return
	}

	return
}
