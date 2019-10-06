package service

import (
	"context"
	"time"
	"valerian/app/service/account/model"
)

func (p *Service) GetAccountStat(c context.Context, accountID int64) (item *model.AccountStat, err error) {
	if item, err = p.d.GetAccountStatByID(c, p.d.DB(), accountID); err != nil {
		return
	} else if item == nil {
		item = &model.AccountStat{
			AccountID: accountID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
	}

	return
}
