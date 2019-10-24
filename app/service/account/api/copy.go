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

func FromProfileInfo(model *model.ProfileInfo) *ProfileReply {
	reply := &ProfileReply{
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
