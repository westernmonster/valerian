package service

import (
	"context"
	"time"

	"valerian/app/service/relation/model"
)

func (p *Service) Stat(c context.Context, aid int64) (stat *model.AccountRelationStat, err error) {
	if stat, err = p.d.GetStatByID(c, p.d.DB(), aid); err != nil {
		return
	} else if stat == nil {
		stat = &model.AccountRelationStat{
			AccountID: aid,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
	}

	return
}
