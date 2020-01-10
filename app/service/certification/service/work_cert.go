package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/certification/api"
	"valerian/app/service/certification/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

// RequestWorkCert 申请工作认证
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
	} else if item == nil {
		item = &model.WorkCertification{
			ID:         gid.NewID(),
			AccountID:  arg.AccountID,
			Status:     model.WorkCertificationInProgress,
			WorkPic:    arg.WorkPic,
			OtherPic:   arg.OtherPic,
			Company:    arg.Company,
			Department: arg.Department,
			Position:   arg.Position,
			ExpiresAt:  arg.ExpiresAt,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddWorkCertification(c, tx, item); err != nil {
			return
		}
	} else {
		item.Status = model.WorkCertificationInProgress
		item.WorkPic = arg.WorkPic
		item.OtherPic = arg.OtherPic
		item.Company = arg.Company
		item.Department = arg.Department
		item.Position = arg.Position
		item.ExpiresAt = arg.ExpiresAt
		item.UpdatedAt = time.Now().Unix()
		if err = p.d.UpdateWorkCertification(c, tx, item); err != nil {
			return
		}

		if err = p.d.UpdateAccountWorkCert(c, tx, arg.AccountID, false); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelAccountCache(context.Background(), item.AccountID)
	})

	return
}

// AuditWorkCert 审批工作认证
func (p *Service) AuditWorkCert(c context.Context, arg *api.AuditWorkCertReq) (err error) {
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

	// 检测操作人是否系统管理员
	if err = p.checkSystemAdmin(c, tx, arg.Aid); err != nil {
		return
	}

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
		if err = p.d.UpdateAccountWorkCert(c, tx, arg.AccountID, true); err != nil {
			return
		}
	} else {
		item.Status = model.WorkCertificationFailed
	}

	item.AuditResult = arg.AuditResult

	if err = p.d.UpdateWorkCertification(c, tx, item); err != nil {
		return
	}
	// add result to history
	if err = p.d.AddWorkCertHistory(c, tx, &model.WorkCertHistory{
		ID:          gid.NewID(),
		AccountID:   item.AccountID,
		Status:      item.Status,
		WorkPic:     item.WorkPic,
		OtherPic:    item.OtherPic,
		Company:     item.Company,
		Department:  item.Department,
		Position:    item.Position,
		ExpiresAt:   item.ExpiresAt,
		AuditResult: arg.AuditResult,
		ManagerID:   arg.Aid,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelAccountCache(context.Background(), item.AccountID)
	})

	return
}

// GetWorkCertStatus 获取工作认证状态
func (p *Service) GetWorkCertStatus(c context.Context, aid int64) (status int32, err error) {
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

// GetWorkCert 获取工作认证
func (p *Service) GetWorkCert(c context.Context, aid int64) (item *model.WorkCertification, err error) {
	return p.getWorkCertByID(c, p.d.DB(), aid)
}

// getWorkCertByID 通过ID获取工作认证
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
		item = &model.WorkCertification{
			ID:        gid.NewID(),
			AccountID: aid,
			Status:    model.WorkCertificationUncommitted,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		return
	}

	if needCache {
		p.addCache(func() {
			p.d.SetWorkCertCache(context.Background(), item)
		})
	}
	return
}

// GetWorkCertificationsPaged 获取工作认证信息
func (p *Service) GetWorkCertificationsPaged(c context.Context, arg *api.WorkCertPagedReq) (resp *api.WorkCertPagedResp, err error) {
	// 检测操作人是否系统管理员
	if err = p.checkSystemAdmin(c, p.d.DB(), arg.Aid); err != nil {
		return
	}

	cond := make(map[string]interface{})
	if arg.Status != int32(0) {
		cond["status"] = arg.Status
	}

	var items []*model.WorkCertification
	if items, err = p.d.GetWorkCertificationsPaged(c, p.d.DB(), cond, arg.Limit, arg.Offset); err != nil {
		return
	}

	resp = &api.WorkCertPagedResp{
		Items: make([]*api.WorkCertItem, len(items)),
	}

	for i, v := range items {
		item := &api.WorkCertItem{
			ID:          v.ID,
			AccountID:   v.AccountID,
			Status:      v.Status,
			WorkPic:     v.WorkPic,
			OtherPic:    v.OtherPic,
			Company:     v.Company,
			Department:  v.Department,
			Position:    v.Position,
			ExpiresAt:   v.ExpiresAt,
			AuditResult: v.AuditResult,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}

		resp.Items[i] = item
	}

	return
}

func (p *Service) GetWorkCertHistoriesPaged(c context.Context, arg *api.WorkCertHistoriesPagedReq) (resp *api.WorkCertHistoriesPagedResp, err error) {
	var data []*model.WorkCertHistory
	if data, err = p.d.GetWorkCertHistoriesPaged(c, p.d.DB(), arg.AccountID, arg.Limit, arg.Offset); err != nil {
		return
	}

	resp = &api.WorkCertHistoriesPagedResp{
		Items: make([]*api.WorkCertHistoryItem, len(data)),
	}

	for i, v := range data {
		item := &api.WorkCertHistoryItem{
			ID:          v.ID,
			AccountID:   v.AccountID,
			Status:      v.Status,
			WorkPic:     v.WorkPic,
			OtherPic:    v.OtherPic,
			Company:     v.Company,
			Department:  v.Department,
			Position:    v.Position,
			ExpiresAt:   v.ExpiresAt,
			AuditResult: v.AuditResult,
			ManagerID:   v.ManagerID,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}

		resp.Items[i] = item
	}

	return
}
