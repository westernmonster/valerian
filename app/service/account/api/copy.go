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

func FromStat(model *model.AccountResStat) *AccountStatInfo {
	reply := &AccountStatInfo{
		AccountID:       model.AccountID,
		TopicCount:      int32(model.TopicCount),
		ArticleCount:    int32(model.ArticleCount),
		DiscussionCount: int32(model.DiscussionCount),
	}

	return reply
}
