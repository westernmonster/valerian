package api

import "valerian/app/service/account/model"

func FromBaseInfo(model *model.BaseInfo) *BaseInfoReply {
	reply := &BaseInfoReply{
		ID:       model.ID,
		UserName: model.UserName,
		Avatar:   model.Avatar,
		IDCert:   model.IDCert,
		WorkCert: model.WorkCert,
		IsOrg:    model.IsOrg,
		IsVIP:    model.IsVIP,
	}

	if model.Gender != nil {
		reply.Gender = &BaseInfoReply_GenderValue{int32(*model.Gender)}
	}

	if model.Introduction != nil {
		reply.Introduction = &BaseInfoReply_IntroductionValue{*model.Introduction}
	}

	return reply
}

func FromStat(model *model.AccountStat) *AccountStatInfo {
	reply := &AccountStatInfo{
		Fans:            int32(model.Fans),
		Following:       int32(model.Following),
		Black:           int32(model.Black),
		TopicCount:      int32(model.TopicCount),
		ArticleCount:    int32(model.ArticleCount),
		DiscussionCount: int32(model.DiscussionCount),
	}

	return reply
}

func FromProfileInfo(model *model.BaseInfo) *ProfileReply {
	reply := &ProfileReply{
		ID:       model.ID,
		UserName: model.UserName,
		Avatar:   model.Avatar,
		IDCert:   model.IDCert,
		WorkCert: model.WorkCert,
		IsOrg:    model.IsOrg,
		IsVIP:    model.IsVIP,
	}

	if model.Gender != nil {
		reply.Gender = &ProfileReply_GenderValue{int32(*model.Gender)}
	}

	if model.Introduction != nil {
		reply.Introduction = &ProfileReply_IntroductionValue{*model.Introduction}
	}

	return reply
}
