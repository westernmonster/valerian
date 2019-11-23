package dao

import (
	"context"
	"fmt"

	certification "valerian/app/service/certification/api"
	"valerian/library/log"
)

func (p *Dao) RequestIDCert(c context.Context, aid int64) (info *certification.RequestIDCertResp, err error) {
	if info, err = p.certificationRPC.RequestIDCert(c, &certification.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.RequestIDCert err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) RefreshIDCertStatus(c context.Context, aid int64) (info *certification.IDCertStatus, err error) {
	if info, err = p.certificationRPC.RefreshIDCertStatus(c, &certification.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.RefreshIDCertStatus err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetIDCert(c context.Context, aid int64) (info *certification.IDCertInfo, err error) {
	if info, err = p.certificationRPC.GetIDCert(c, &certification.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetIDCert err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetIDCertStatus(c context.Context, aid int64) (info *certification.IDCertStatus, err error) {
	if info, err = p.certificationRPC.GetIDCertStatus(c, &certification.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetIDCertStatus err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) RequestWorkCert(c context.Context, req *certification.WorkCertReq) (err error) {
	if _, err = p.certificationRPC.RequestWorkCert(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.RequestWorkCert err(%+v) req(%+v)", err, req))
	}
	return
}

func (p *Dao) AuditWorkCert(c context.Context, req *certification.AuditWorkCertReq) (err error) {
	if _, err = p.certificationRPC.AuditWorkCert(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AuditWorkCert err(%+v) req(%+v)", err, req))
	}
	return
}

func (p *Dao) GetWorkCert(c context.Context, aid int64) (info *certification.WorkCertInfo, err error) {
	if info, err = p.certificationRPC.GetWorkCert(c, &certification.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCert err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetWorkCertStatus(c context.Context, aid int64) (info *certification.WorkCertStatus, err error) {
	if info, err = p.certificationRPC.GetWorkCertStatus(c, &certification.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertStatus err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetWorkCertP(c context.Context, aid int64) (info *certification.WorkCertStatus, err error) {
	if info, err = p.certificationRPC.GetW(c, &certification.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertStatus err(%+v) aid(%d)", err, aid))
	}
	return
}
