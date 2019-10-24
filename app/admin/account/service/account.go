package service

import (
	"context"

	"valerian/app/admin/account/model"
	account "valerian/app/service/account/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetSelfProfile(c context.Context) (profile *model.SelfProfile, err error) {
	uid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var item *account.SelfProfile
	if item, err = p.d.GetSelfProfileInfo(c, uid); err != nil {
		return
	}

	profile = &model.SelfProfile{
		ID:           item.ID,
		Mobile:       item.Mobile,
		Email:        item.Email,
		Gender:       int(item.Gender),
		BirthYear:    int(item.BirthYear),
		BirthMonth:   int(item.BirthMonth),
		BirthDay:     int(item.BirthDay),
		Location:     item.Location,
		Introduction: item.Introduction,
		Avatar:       item.Avatar,
		Source:       int(item.Source),
		IDCert:       bool(item.IDCert),
		WorkCert:     bool(item.WorkCert),
		IsOrg:        bool(item.IsOrg),
		IsVIP:        bool(item.IsVIP),
		IP:           item.IP,
		Role:         item.Role,
		UserName:     item.UserName,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
		Stat: &model.AccountStatInfo{
			FollowingCount:  int(item.Stat.FollowingCount),
			FansCount:       int(item.Stat.FansCount),
			TopicCount:      int(item.Stat.TopicCount),
			ArticleCount:    int(item.Stat.ArticleCount),
			DiscussionCount: int(item.Stat.DiscussionCount),
		},
		Settings: &model.SettingInfo{
			ActivityLike:         item.Setting.ActivityLike,
			ActivityComment:      item.Setting.ActivityComment,
			ActivityFollowTopic:  item.Setting.ActivityFollowTopic,
			ActivityFollowMember: item.Setting.ActivityFollowMember,
			NotifyLike:           item.Setting.NotifyLike,
			NotifyComment:        item.Setting.NotifyComment,
			NotifyNewFans:        item.Setting.NotifyNewFans,
			NotifyNewMember:      item.Setting.NotifyNewMember,
			Language:             item.Setting.Language,
		},
	}

	return
}
