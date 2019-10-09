package dao

import (
	"context"
	"fmt"
	relation "valerian/app/service/relation/api"
	"valerian/library/log"
)

func (p *Dao) GetFansIDs(c context.Context, aid int64) (resp *relation.IDsResp, err error) {
	if resp, err = p.relationRPC.GetFansIDs(c, &relation.AidReq{AccountID: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFansIDs error(%+v), aid(%d) ", err, aid))
	}

	return
}
