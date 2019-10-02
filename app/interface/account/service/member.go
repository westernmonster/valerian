package service

import (
	"context"

	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	relation "valerian/app/service/relation/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetMemberInfo(c context.Context, targetID int64) (resp *model.MemberInfo, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var f *model.Profile
	if f, err = p.getProfile(c, targetID); err != nil {
		return
	}

	resp = &model.MemberInfo{
		ID:             f.ID,
		UserName:       f.UserName,
		Gender:         f.Gender,
		Location:       f.Location,
		LocationString: f.LocationString,
		Introduction:   f.Introduction,
		Avatar:         f.Avatar,
		IDCert:         f.IDCert,
		WorkCert:       f.WorkCert,
		IsOrg:          f.IsOrg,
		IsVIP:          f.IsVIP,
	}

	var stat *relation.StatInfo
	if stat, err = p.d.Stat(c, aid, targetID); err != nil {
		return
	}
	resp.Stat = &model.MemberInfoStat{
		FansCount:      int(stat.Fans),
		FollowingCount: int(stat.Following),
		IsFollow:       stat.IsFollowing,
	}

	var accountStat *account.AccountStatInfo
	if accountStat, err = p.d.GetAccountStat(c, targetID); err != nil {
		return
	}

	resp.Stat.TopicCount = int(accountStat.TopicCount)
	resp.Stat.ArticleCount = int(accountStat.ArticleCount)
	resp.Stat.DiscussionCount = int(accountStat.DiscussionCount)

	return
}
