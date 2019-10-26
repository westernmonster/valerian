package service

import (
	"context"

	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	"valerian/library/database/sqalx"
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
	var setting *account.Setting
	if setting, err = p.d.GetAccountSetting(c, accountID); err != nil {
		return
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

	settings := make(map[string]bool)
	language := ""
	if arg.Language != nil {
		language = *arg.Language
	}

	if arg.Activity != nil {
		if arg.Activity.Like != nil {
			settings["activity_like"] = *arg.Activity.Like
		}
		if arg.Activity.Comment != nil {
			settings["activity_comment"] = *arg.Activity.Comment
		}

		if arg.Activity.FollowTopic != nil {
			settings["activity_follow_topic"] = *arg.Activity.FollowTopic
		}

		if arg.Activity.FollowMember != nil {
			settings["activity_follow_member"] = *arg.Activity.FollowMember
		}
	}

	if arg.Notify != nil {
		if arg.Notify.Like != nil {
			settings["notify_like"] = *arg.Notify.Like
		}
		if arg.Notify.Comment != nil {
			settings["notify_comment"] = *arg.Notify.Comment
		}

		if arg.Notify.NewFans != nil {
			settings["notify_new_fans"] = *arg.Notify.NewFans
		}

		if arg.Notify.NewMember != nil {
			settings["notify_new_member"] = *arg.Notify.NewMember
		}
	}

	if err = p.d.UpdateAccountSetting(c, aid, settings, language); err != nil {
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

	settings := make(map[string]bool)
	language := ""

	if arg.Like != nil {
		settings["activity_like"] = *arg.Like
	}
	if arg.Comment != nil {
		settings["activity_comment"] = *arg.Comment
	}

	if arg.FollowTopic != nil {
		settings["activity_follow_topic"] = *arg.FollowTopic
	}

	if arg.FollowMember != nil {
		settings["activity_follow_member"] = *arg.FollowMember
	}
	if err = p.d.UpdateAccountSetting(c, aid, settings, language); err != nil {
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

	settings := make(map[string]bool)
	language := ""

	if arg.Like != nil {
		settings["notify_like"] = *arg.Like
	}
	if arg.Comment != nil {
		settings["notify_comment"] = *arg.Comment
	}

	if arg.NewFans != nil {
		settings["notify_new_fans"] = *arg.NewFans
	}

	if arg.NewMember != nil {
		settings["notify_new_member"] = *arg.NewMember
	}

	if err = p.d.UpdateAccountSetting(c, aid, settings, language); err != nil {
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

	settings := make(map[string]bool)
	language := ""
	if arg.Language != nil {
		language = *arg.Language
	}

	if err = p.d.UpdateAccountSetting(c, aid, settings, language); err != nil {
		return
	}

	return
}
