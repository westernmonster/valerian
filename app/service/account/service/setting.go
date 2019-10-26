package service

import (
	"context"
	"fmt"

	"valerian/app/service/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/log"
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

func (p *Service) UpdateAccountSetting(c context.Context, aid int64, req map[string]bool, language string) (err error) {
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()
	var setting *model.AccountSetting
	if setting, err = p.d.GetAccountSettingByID(c, tx, aid); err != nil {
		return
	}

	if setting == nil {
		setting = &model.AccountSetting{
			AccountID:            aid,
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

		if err = p.d.AddAccountSetting(c, tx, setting); err != nil {
			return
		}
	}

	if v, ok := req["activity_like"]; ok {
		setting.ActivityLike = types.BitBool(v)
	}

	if v, ok := req["activity_comment"]; ok {
		setting.ActivityComment = types.BitBool(v)
	}

	if v, ok := req["activity_follow_topic"]; ok {
		setting.ActivityFollowTopic = types.BitBool(v)
	}

	if v, ok := req["activity_follow_member"]; ok {
		setting.ActivityFollowMember = types.BitBool(v)
	}

	if v, ok := req["notify_like"]; ok {
		setting.NotifyLike = types.BitBool(v)
	}

	if v, ok := req["notify_comment"]; ok {
		setting.NotifyComment = types.BitBool(v)
	}

	if v, ok := req["notify_new_fans"]; ok {
		setting.NotifyNewFans = types.BitBool(v)
	}

	if v, ok := req["notify_new_member"]; ok {
		setting.NotifyNewMember = types.BitBool(v)
	}

	if language != "" {
		setting.Language = language
	}

	if err = p.d.UpdateAccountSetting(c, tx, setting); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}
