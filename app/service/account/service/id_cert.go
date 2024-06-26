package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"valerian/app/service/account/model"
	"valerian/library/cloudauth"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

// RequestIDCert 申请身份认证
// 使用阿里云实人认证
func (p *Service) RequestIDCert(c context.Context, aid int64) (token cloudauth.VerifyTokenData, err error) {
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

	var item *model.IDCertification
	if item, err = p.d.GetIDCertificationByCond(c, tx, map[string]interface{}{"account_id": aid}); err != nil {
		return
	} else if item == nil {
		item = &model.IDCertification{
			ID:        gid.NewID(),
			AccountID: aid,
			Status:    model.IDCertificationUncommitted,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddIDCertification(c, tx, item); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	tickerID := strconv.FormatInt(item.ID, 10)
	resp, err := p.cloudauth.GetVerifyToken(c, tickerID)
	if err != nil {
		return
	}

	token = resp.Data

	return

}

// GetIDCertStatus 获取身份认证状态
func (p *Service) GetIDCertStatus(c context.Context, aid int64) (status int32, err error) {
	return p.getIDCertStatus(c, p.d.DB(), aid)
}

func (p *Service) getIDCertStatus(c context.Context, node sqalx.Node, aid int64) (status int32, err error) {
	var item *model.IDCertification
	if item, err = p.getIDCertByID(c, node, aid); err != nil {
		if ecode.Cause(err) == ecode.IDCertificationNotExist {
			return model.IDCertificationUncommitted, nil
		}
		return
	} else {
		return item.Status, nil
	}
}

// RefreshIDCertStatus 刷新认证状态
func (p *Service) RefreshIDCertStatus(c context.Context, aid int64) (status int32, err error) {
	var item *model.IDCertification
	if item, err = p.getIDCertByID(c, p.d.DB(), aid); err != nil {
		return
	}

	ticketID := strconv.FormatInt(item.ID, 10)

	var resp *cloudauth.GetStatusResponse
	if resp, err = p.cloudauth.GetStatus(c, ticketID); err != nil {
		return
	}

	item.Status = int32(resp.Data.StatusCode)
	item.AuditConclusions = resp.Data.AuditConclusions

	if item.Status == model.IDCertificationSuccess {
		var material *cloudauth.GetMaterialsResponse
		if material, err = p.cloudauth.GetMaterials(c, ticketID); err != nil {
			return
		}

		item.Name = material.Data.Name
		item.IdentificationNumber = material.Data.IdentificationNumber
		item.IDCardType = material.Data.IdCardType
		item.IDCardStartDate = material.Data.IdCardStartDate
		item.IDCardExpiry = material.Data.IdCardExpiry
		item.Address = material.Data.Address
		item.Sex = material.Data.Sex
		item.IDCardFrontPic = material.Data.IdCardFrontPic
		item.IDCardBackPic = material.Data.IdCardBackPic
		item.FacePic = material.Data.FacePic
		item.EthnicGroup = material.Data.EthnicGroup
		item.UpdatedAt = time.Now().Unix()

	}

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

	if err = p.d.UpdateIDCertification(c, tx, item); err != nil {
		return
	}

	if item.Status == model.IDCertificationSuccess {
		if err = p.d.UpdateAccountIDCert(c, tx, aid, true); err != nil {
			return
		}

		if item.Name != "" {
			if err = p.d.UpdateAccountRealName(c, tx, aid, item.Name); err != nil {
				return
			}
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	status = item.Status

	p.addCache(func() {
		p.d.DelAccountCache(context.Background(), item.AccountID)
		p.d.DelIDCertCache(context.Background(), item.AccountID)
	})
	return
}

// GetIDCert 通过ID获取身份认证
func (p *Service) GetIDCert(c context.Context, aid int64) (item *model.IDCertification, err error) {
	return p.getIDCertByID(c, p.d.DB(), aid)
}

// getIDCertByID 通过ID获取身份认证
func (p *Service) getIDCertByID(c context.Context, node sqalx.Node, aid int64) (item *model.IDCertification, err error) {
	var needCache = true

	if item, err = p.d.IDCertCache(c, aid); err != nil {
		needCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetIDCertificationByCond(c, node, map[string]interface{}{"account_id": aid}); err != nil {
		return
	} else if item == nil {
		item = &model.IDCertification{
			ID:        gid.NewID(),
			AccountID: aid,
			Status:    model.IDCertificationUncommitted,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		return
	}

	if needCache {
		p.addCache(func() {
			p.d.SetIDCertCache(context.Background(), item)
		})
	}
	return
}
