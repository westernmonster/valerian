package service

import (
	"context"

	"valerian/app/service/account-feed/model"
	"valerian/library/database/sqalx"
)

// getAccountSetting 获取用户设置
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
			NotifyLike:           false,
			NotifyComment:        true,
			NotifyNewFans:        true,
			NotifyNewMember:      true,
			Language:             "zh-CN",
		}
	}

	resp = &model.SettingResp{
		ActivityLike:         bool(setting.ActivityLike),
		ActivityComment:      bool(setting.ActivityComment),
		ActivityFollowTopic:  bool(setting.ActivityFollowTopic),
		ActivityFollowMember: bool(setting.ActivityFollowMember),
		NotifyLike:           bool(setting.NotifyLike),
		NotifyComment:        bool(setting.NotifyComment),
		NotifyNewFans:        bool(setting.NotifyNewFans),
		NotifyNewMember:      bool(setting.NotifyNewMember),
		Language:             setting.Language,
	}

	return
}
