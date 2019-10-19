package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"valerian/app/interface/certification/model"
	"valerian/library/cloudauth"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

// RequestIDCertification 发起认证请求，获取 Token
func (p *Service) RequestIDCertification(c context.Context) (token cloudauth.VerifyTokenData, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
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

	var item *model.IDCertification
	if item, err = p.d.GetUserIDCertification(c, tx, aid); err != nil {
		return
	}

	if item == nil {
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

// 查询认证状态
func (p *Service) GetIDCertificationStatus(c context.Context) (status int, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var item *model.IDCertification
	if item, err = p.d.GetUserIDCertification(c, p.d.DB(), aid); err != nil {
		return
	} else if item == nil {
		err = ecode.IDCertificationNotExist
		return
	}

	ticketID := strconv.FormatInt(item.ID, 10)

	var resp *cloudauth.GetStatusResponse
	if resp, err = p.cloudauth.GetStatus(c, ticketID); err != nil {
		return
	}

	item.Status = resp.Data.StatusCode
	item.AuditConclusions = &resp.Data.AuditConclusions

	if item.Status == model.IDCertificationSuccess {
		var material *cloudauth.GetMaterialsResponse
		if material, err = p.cloudauth.GetMaterials(c, ticketID); err != nil {
			return
		}

		item.Name = &material.Data.Name
		item.IdentificationNumber = &material.Data.IdentificationNumber
		item.IDCardType = &material.Data.IdCardType
		item.IDCardStartDate = &material.Data.IdCardStartDate
		item.IDCardExpiry = &material.Data.IdCardExpiry
		item.Address = &material.Data.Address
		item.Sex = &material.Data.Sex
		// TODO: 图片下载
		item.IDCardFrontPic = &material.Data.IdCardFrontPic
		item.IDCardBackPic = &material.Data.IdCardBackPic
		item.FacePic = &material.Data.FacePic
		item.EthnicGroup = &material.Data.EthnicGroup

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
		var acc *model.Account
		if acc, err = p.getAccountByID(c, tx, item.AccountID); err != nil {
			return
		}

		acc.IDCert = true

		if err = p.d.UpdateAccount(c, tx, acc); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	status = item.Status

	p.addCache(func() {
		p.d.DelAccountCache(context.TODO(), item.AccountID)
	})
	return
}

func (p *Service) getAccountByID(c context.Context, node sqalx.Node, aid int64) (account *model.Account, err error) {
	var needCache = true

	if account, err = p.d.AccountCache(c, aid); err != nil {
		needCache = false
	} else if account != nil {
		return
	}

	if account, err = p.d.GetAccountByID(c, node, aid); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	if needCache {
		p.addCache(func() {
			p.d.SetAccountCache(context.TODO(), account)
		})
	}
	return
}
