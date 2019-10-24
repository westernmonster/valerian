package service

import (
	"context"
	"time"

	"valerian/app/service/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
)

func (p *Service) GetAccountSetting(c context.Context, aid int64) (resp *model.SettingResp, err error) {
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

func (p *Service) UpdateAccountSetting(c context.Context, arg *model.ArgSetting) (err error) {
	setting := &model.AccountSetting{
		AccountID:            arg.AccountID,
		ActivityLike:         types.BitBool(arg.ActivityLike),
		ActivityComment:      types.BitBool(arg.ActivityComment),
		ActivityFollowTopic:  types.BitBool(arg.ActivityFollowTopic),
		ActivityFollowMember: types.BitBool(arg.ActivityFollowMember),
		NotifyLike:           types.BitBool(arg.NotifyLike),
		NotifyComment:        types.BitBool(arg.NotifyComment),
		NotifyNewFans:        types.BitBool(arg.NotifyNewFans),
		NotifyNewMember:      types.BitBool(arg.NotifyNewMember),
		Language:             arg.Language,
		CreatedAt:            time.Now().Unix(),
		UpdatedAt:            time.Now().Unix(),
	}

	if err = p.d.UpdateAccountSetting(c, p.d.DB(), setting); err != nil {
		return
	}

	return
}
