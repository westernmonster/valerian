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

	return reply
}
