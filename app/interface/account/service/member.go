package service

import (
	"context"

	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	relation "valerian/app/service/relation/api"
)

func (p *Service) GetMemberInfo(c context.Context, aid int64) (resp *model.MemberInfo, err error) {
	var f *model.Profile
	if f, err = p.getProfile(c, aid); err != nil {
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
	if stat, err = p.d.Stat(c, aid); err != nil {
		return
	}
	resp.Stat = &model.MemberInfoStat{
		FansCount:      int(stat.Fans),
		FollowingCount: int(stat.Following),
	}

	var accountStat *account.AccountStatInfo
	if accountStat, err = p.d.GetAccountStat(c, aid); err != nil {
		return
	}

	resp.Stat.TopicCount = int(accountStat.TopicCount)
	resp.Stat.ArticleCount = int(accountStat.ArticleCount)
	resp.Stat.DiscussionCount = int(accountStat.DiscussionCount)

	return
}
