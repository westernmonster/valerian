package api

import "valerian/app/service/account/model"

func FromBaseInfo(model *model.BaseInfo) *BaseInfoReply {
	reply := &BaseInfoReply{
		ID:           model.ID,
		UserName:     model.UserName,
		Avatar:       model.Avatar,
		IDCert:       model.IDCert,
		WorkCert:     model.WorkCert,
		IsOrg:        model.IsOrg,
		IsVIP:        model.IsVIP,
		Gender:       int32(model.Gender),
		Introduction: model.Introduction,
	}

	return reply
}

func FromStat(model *model.AccountStat) *AccountStatInfo {
	reply := &AccountStatInfo{
		FansCount:       int32(model.Fans),
		FollowingCount:  int32(model.Following),
		BlackCount:      int32(model.Black),
		TopicCount:      int32(model.TopicCount),
		ArticleCount:    int32(model.ArticleCount),
		DiscussionCount: int32(model.DiscussionCount),
	}

	return reply
}

func FromProfileInfo(model *model.ProfileInfo) *MemberInfoReply {
	reply := &MemberInfoReply{
		ID:             model.ID,
		UserName:       model.UserName,
		Avatar:         model.Avatar,
		IDCert:         model.IDCert,
		WorkCert:       model.WorkCert,
		IsOrg:          model.IsOrg,
		IsVIP:          model.IsVIP,
		CreatedAt:      model.CreatedAt,
		Gender:         int32(model.Gender),
		Introduction:   model.Introduction,
		Location:       model.Location,
		LocationString: model.LocationString,
	}

	return reply
}

func FromSelfProfile(model *model.SelfProfile) *SelfProfile {
	reply := &SelfProfile{
		ID:             model.ID,
		Email:          model.Email,
		Mobile:         model.Mobile,
		UserName:       model.UserName,
		Gender:         int32(model.Gender),
		BirthYear:      int32(model.BirthYear),
		BirthMonth:     int32(model.BirthMonth),
		BirthDay:       int32(model.BirthDay),
		Introduction:   model.Introduction,
		Avatar:         model.Avatar,
		Source:         int32(model.Source),
		Location:       model.Location,
		LocationString: model.LocationString,
		IDCert:         model.IDCert,
		IDCertStatus:   int32(model.IDCertStatus),
		WorkCert:       model.WorkCert,
		WorkCertStatus: int32(model.WorkCertStatus),
		IP:             model.IP,
		IsOrg:          model.IsOrg,
		IsVIP:          model.IsVIP,
		Role:           model.Role,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
	}

	return reply
}
