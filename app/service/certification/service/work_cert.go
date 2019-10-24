package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/certification/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) RequestWorkCert(c context.Context, arg *model.ArgAddWorkCert) (err error) {
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var cert *model.IDCertification
	if cert, err = p.getIDCertByID(c, tx, arg.AccountID); err != nil {
		if ecode.Cause(err) == ecode.IDCertificationNotExist {
			err = ecode.IDCertFirst
			return
		}

		return
	} else if cert.Status != model.IDCertificationSuccess {
		err = ecode.IDCertFirst
		return
	}

	var item *model.WorkCertification
	if item, err = p.d.GetWorkCertificationByCond(c, tx, map[string]interface{}{"account_id": arg.AccountID}); err != nil {
		return
	} else if item != nil {
		err = ecode.WorkCertExist
		return
	}

	item = &model.WorkCertification{
		ID:         gid.NewID(),
		AccountID:  item.AccountID,
		Status:     model.WorkCertificationInProgress,
		WorkPic:    item.WorkPic,
		OtherPic:   item.OtherPic,
		Company:    item.Company,
		Department: item.Department,
		Position:   item.Position,
		ExpiresAt:  item.ExpiresAt,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddWorkCertification(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

func (p *Service) AuditWorkCert(c context.Context, arg *model.ArgAuditWorkCert) (err error) {
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var cert *model.IDCertification
	if cert, err = p.getIDCertByID(c, tx, arg.AccountID); err != nil {
		if ecode.Cause(err) == ecode.IDCertificationNotExist {
			err = ecode.IDCertFirst
			return
		}

		return
	} else if cert.Status != model.IDCertificationSuccess {
		err = ecode.IDCertFirst
		return
	}

	var item *model.WorkCertification
	if item, err = p.getWorkCertByID(c, tx, arg.AccountID); err != nil {
		return
	}

	if arg.Approve {
		item.Status = model.WorkCertificationSuccess
		var acc *model.Account
		if acc, err = p.getAccountByID(c, tx, arg.AccountID); err != nil {
			return
		}

		acc.WorkCert = true

		if err = p.d.UpdateAccount(c, tx, acc); err != nil {
			return
		}
	} else {
		item.Status = model.WorkCertificationFailed
	}

	item.AuditResult = arg.AuditResult

	if err = p.d.UpdateWorkCertification(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

func (p *Service) GetWorkCertStatus(c context.Context, aid int64) (status int, err error) {
	var item *model.WorkCertification
	if item, err = p.getWorkCertByID(c, p.d.DB(), aid); err != nil {
		if ecode.Cause(err) == ecode.WorkCertificationNotExist {
			return model.WorkCertificationUncommitted, nil
		}
		return
	} else {
		return item.Status, nil
	}
}

func (p *Service) GetWorkCert(c context.Context, aid int64) (item *model.WorkCertification, err error) {
	return p.getWorkCertByID(c, p.d.DB(), aid)
}

func (p *Service) getWorkCertByID(c context.Context, node sqalx.Node, aid int64) (item *model.WorkCertification, err error) {
	var needCache = true

	if item, err = p.d.WorkCertCache(c, aid); err != nil {
		needCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetWorkCertificationByCond(c, node, map[string]interface{}{"account_id": aid}); err != nil {
		return
	} else if item == nil {
		err = ecode.WorkCertificationNotExist
		return
	}

	if needCache {
		p.addCache(func() {
			p.d.SetWorkCertCache(context.TODO(), item)
		})
	}
	return
}