package service

import (
	"context"
	"valerian/app/service/message/api"
	"valerian/app/service/message/model"
)

func (p *Service) GetMessageStat(c context.Context, arg *api.AidReq) (st *api.MessageStat, err error) {
	var stat *model.MessageStat
	if stat, err = p.d.GetMessageStatByID(c, p.d.DB(), arg.AccountID); err != nil {
		return
	}

	st = &api.MessageStat{
		AccountID:   stat.AccountID,
		UnreadCount: int32(stat.UnreadCount),
	}
	return
}
