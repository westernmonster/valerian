package service

import (
	"context"
	"strings"

	"valerian/app/service/account/api"
	"valerian/app/service/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

// GetSelfProfile 获取当前用户详情
func (p *Service) GetSelfProfile(c context.Context, accountID int64) (profile *api.SelfProfile, err error) {
	return p.getSelfProfile(c, p.d.DB(), accountID)
}

// getSelfProfile 获取当前用户详情
func (p *Service) getSelfProfile(c context.Context, node sqalx.Node, accountID int64) (profile *api.SelfProfile, err error) {
	var item *model.Account
	if item, err = p.getAccountByID(c, node, accountID); err != nil {
		return
	}

	profile = &api.SelfProfile{
		ID:           item.ID,
		Prefix:       item.Prefix,
		Mobile:       item.Mobile,
		Email:        item.Email,
		UserName:     item.UserName,
		Gender:       item.Gender,
		BirthYear:    item.BirthYear,
		BirthMonth:   item.BirthMonth,
		BirthDay:     item.BirthDay,
		Location:     item.Location,
		Introduction: item.Introduction,
		Avatar:       item.Avatar,
		Source:       item.Source,
		IDCert:       bool(item.IDCert),
		WorkCert:     bool(item.WorkCert),
		IsOrg:        bool(item.IsOrg),
		IsVIP:        bool(item.IsVip),
		Role:         item.Role,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	if item.Mobile != "" {
		profile.Mobile = strings.TrimLeft(item.Mobile, item.Prefix)
	}

	if item.Location != 0 {
		if v, e := p.getLocationString(c, item.Location); e != nil {
			return nil, e
		} else {
			profile.LocationString = v
		}
	}

	if profile.WorkCertStatus, err = p.getWorkCertStatus(c, node, accountID); err != nil {
		return
	}

	if profile.WorkCertStatus == 1 {
		var workCert *model.WorkCertification
		if workCert, err = p.getWorkCertByID(c, node, accountID); err != nil {
			return
		}

		profile.Company = workCert.Company
		profile.Position = workCert.Position
	}

	if profile.IDCertStatus, err = p.getIDCertStatus(c, node, accountID); err != nil {
		return
	}

	var stat *model.AccountStat
	if stat, err = p.getAccountStat(c, node, accountID); err != nil {
		return
	}

	profile.Stat = &api.AccountStatInfo{
		FansCount:       int32(stat.Fans),
		FollowingCount:  int32(stat.Following),
		BlackCount:      int32(stat.Black),
		TopicCount:      int32(stat.TopicCount),
		ArticleCount:    int32(stat.ArticleCount),
		DiscussionCount: int32(stat.DiscussionCount),
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, node, accountID); err != nil {
		return
	}

	profile.Setting = &api.Setting{
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

// GetMemberProfile 获取指定用户详情
func (p *Service) GetMemberProfile(c context.Context, accountID int64) (profile *model.ProfileInfo, err error) {
	var item *model.Account
	if item, err = p.getAccountByID(c, p.d.DB(), accountID); err != nil {
		return
	}

	profile = &model.ProfileInfo{
		ID:           item.ID,
		UserName:     item.UserName,
		Gender:       item.Gender,
		Location:     item.Location,
		Introduction: item.Introduction,
		Avatar:       item.Avatar,
		IDCert:       bool(item.IDCert),
		WorkCert:     bool(item.WorkCert),
		IsOrg:        bool(item.IsOrg),
		IsVIP:        bool(item.IsVip),
		Role:         item.Role,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
		IsLock:       bool(item.IsLock),
	}

	var workCertStatus int32
	if workCertStatus, err = p.getWorkCertStatus(c, p.d.DB(), accountID); err != nil {
		return
	}

	if workCertStatus == int32(1) {
		var workCert *model.WorkCertification
		if workCert, err = p.getWorkCertByID(c, p.d.DB(), accountID); err != nil {
			return
		}

		profile.Company = workCert.Company
		profile.Position = workCert.Position
	}

	if item.Location != 0 {
		if v, e := p.getLocationString(c, item.Location); e != nil {
			return nil, e
		} else {
			profile.LocationString = v
		}
	}

	return
}

// getLocationString 用户所在地区信息
func (p *Service) getLocationString(c context.Context, nodeID int64) (locationString string, err error) {
	arr := []string{}

	id := nodeID
	var item *model.Area
	for {
		if item, err = p.d.GetArea(c, p.d.DB(), id); err != nil {
			return
		} else if item == nil {
			err = ecode.AreaNotExist
			return
		}

		arr = append(arr, item.Name)

		if item.Parent == 0 {
			break
		}

		id = item.Parent
	}

	locationString = ""

	for i := len(arr) - 1; i >= 0; i-- {
		locationString += arr[i] + " "
	}

	return
}
