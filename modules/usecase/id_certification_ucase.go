package usecase

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure/berr"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/infrastructure/cloudauth"
	"git.flywk.com/flywiki/api/infrastructure/gid"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
)

type IDCertificationUsecase struct {
	sqalx.Node
	*sqlx.DB
	CloudAuthClient           *cloudauth.CloudAuthClient
	IDCertificationRepository interface {
		// QueryListPaged get paged records by condition
		QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*repo.IDCertification, err error)
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.IDCertification, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.IDCertification, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.IDCertification, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.IDCertification, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.IDCertification) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.IDCertification) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}
}

// Request 发起认证请求，获取 Token
func (p *IDCertificationUsecase) Request(ctx *biz.BizContext) (token cloudauth.VerifyTokenData, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.IDCertificationRepository.GetByCondition(tx, map[string]string{
		"account_id": strconv.FormatInt(*ctx.AccountID, 10),
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !exist {
		id, errInner := gid.NextID()
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
		item = &repo.IDCertification{
			ID:        id,
			AccountID: *ctx.AccountID,
			Status:    models.IDCertificationUncommitted,
		}

		errInner = p.IDCertificationRepository.Insert(tx, item)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
	}

	tickerID := strconv.FormatInt(item.ID, 10)

	resp, err := p.CloudAuthClient.GetVerifyToken(tickerID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	token = resp.Data

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// 查询认证状态
func (p *IDCertificationUsecase) GetStatus(ctx *biz.BizContext) (status int, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.IDCertificationRepository.GetByCondition(tx, map[string]string{
		"account_id": strconv.FormatInt(*ctx.AccountID, 10),
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !exist {
		err = berr.Errorf("你还未发起认证")
		return
	}

	ticketID := strconv.FormatInt(item.ID, 10)

	resp, err := p.CloudAuthClient.GetStatus(ticketID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	item.Status = resp.Data.StatusCode
	item.AuditConclusions = &resp.Data.AuditConclusions

	err = p.IDCertificationRepository.Update(tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if item.Status == models.IDCertificationSuccess {
		material, errInner := p.CloudAuthClient.GetMaterials(ticketID)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
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

		errInner = p.IDCertificationRepository.Update(tx, item)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
	}

	status = item.Status

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}
