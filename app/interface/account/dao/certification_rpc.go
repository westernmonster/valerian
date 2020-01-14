package dao

import (
	"context"
	"fmt"

	account "valerian/app/service/account/api"
	"valerian/library/log"
)

func (p *Dao) RequestIDCert(c context.Context, aid int64) (info *account.RequestIDCertResp, err error) {
	if info, err = p.accountRPC.RequestIDCert(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.RequestIDCert err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) RefreshIDCertStatus(c context.Context, aid int64) (info *account.IDCertStatus, err error) {
	if info, err = p.accountRPC.RefreshIDCertStatus(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.RefreshIDCertStatus err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetIDCert(c context.Context, aid int64) (info *account.IDCertInfo, err error) {
	if info, err = p.accountRPC.GetIDCert(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetIDCert err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetIDCertStatus(c context.Context, aid int64) (info *account.IDCertStatus, err error) {
	if info, err = p.accountRPC.GetIDCertStatus(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetIDCertStatus err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) RequestWorkCert(c context.Context, req *account.WorkCertReq) (err error) {
	if _, err = p.accountRPC.RequestWorkCert(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.RequestWorkCert err(%+v) req(%+v)", err, req))
	}
	return
}

func (p *Dao) AuditWorkCert(c context.Context, req *account.AuditWorkCertReq) (err error) {
	if _, err = p.accountRPC.AuditWorkCert(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AuditWorkCert err(%+v) req(%+v)", err, req))
	}
	return
}

func (p *Dao) GetWorkCert(c context.Context, aid int64) (info *account.WorkCertInfo, err error) {
	if info, err = p.accountRPC.GetWorkCert(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCert err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetWorkCertStatus(c context.Context, aid int64) (info *account.WorkCertStatus, err error) {
	if info, err = p.accountRPC.GetWorkCertStatus(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertStatus err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetWorkCertsPaged(c context.Context, req *account.WorkCertPagedReq) (info *account.WorkCertPagedResp, err error) {
	if info, err = p.accountRPC.GetWorkCertsPaged(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertsPaged err(%+v) req(%+v)", err, req))
	}
	return
}

func (p *Dao) GetWorkCertHistoriesPaged(c context.Context, req *account.WorkCertHistoriesPagedReq) (info *account.WorkCertHistoriesPagedResp, err error) {
	if info, err = p.accountRPC.GetWorkCertHistoriesPaged(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetWorkCertHistoriesPaged err(%+v) req(%+v)", err, req))
	}
	return
}
