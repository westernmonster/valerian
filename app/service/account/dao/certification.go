package dao

import (
	"context"
	"fmt"
	certification "valerian/app/service/certification/api"
	"valerian/library/log"
)

func (p *Dao) GetWorkCertStatus(c context.Context, aid int64) (status int32, err error) {
	var statusResp *certification.WorkCertStatus
	if statusResp, err = p.certificationRPC.GetWorkCertStatus(c, &certification.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertStatus, aid(%d) error(%+v)", aid, err))
		return
	}

	status = statusResp.Status
	return
}

func (p *Dao) GetIDCertStatus(c context.Context, aid int64) (status int32, err error) {
	var statusResp *certification.IDCertStatus
	if statusResp, err = p.certificationRPC.GetIDCertStatus(c, &certification.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetIDCertStatus, aid(%d) error(%+v)", aid, err))
		return
	}

	status = statusResp.Status
	return
}
