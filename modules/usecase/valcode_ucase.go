package usecase

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure/berr"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/infrastructure/gid"
	"git.flywk.com/flywiki/api/infrastructure/helper"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
)

type ValcodeUsecase struct {
	sqalx.Node
	*sqlx.DB
	SMSClient interface {
		SendRegisterValcode(mobile string, valcode string) (err error)
		SendResetPasswordValcode(mobile string, valcode string) (err error)
	}

	EmailClient interface {
		SendRegisterEmail(email string, valcode string) (err error)
		SendResetPasswordValcode(email string, valcode string) (err error)
	}
	ValcodeRepository interface {
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.Valcode, err error)
		// GetByID get record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.Valcode, exist bool, err error)

		// GetByCondition get record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.Valcode, exist bool, err error)

		// HasSentRecordsInDuration determine current identity has sent records in specified duration
		HasSentRecordsInDuration(node sqalx.Node, identity string, codeType int, duration time.Duration) (has bool, err error)

		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.Valcode) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.Valcode) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}
}

func (p *ValcodeUsecase) RequestEmailValcode(ctx *biz.BizContext, req *models.RequestEmailValcodeReq) (createdTime int64, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	has, err := p.ValcodeRepository.HasSentRecordsInDuration(tx, req.Email, req.CodeType, models.ValcodeSpan)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if has {
		err = berr.Errorf("60秒下发一次验证码，请不要重复请求")
		return
	}

	valcode := helper.GenerateValcode(6)

	switch req.CodeType {
	case models.ValcodeRegister:
		if e := p.EmailClient.SendRegisterEmail(req.Email, valcode); e != nil {
			err = tracerr.Wrap(e)
			return
		}
		break
	case models.ValcodeForgetPassword:
		if e := p.EmailClient.SendResetPasswordValcode(req.Email, valcode); e != nil {
			err = tracerr.Wrap(e)
			return
		}
		break
	default:
		err = berr.Errorf("未知的验证码类型")
		return
	}

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	item := &repo.Valcode{
		ID:       id,
		CodeType: req.CodeType,
		Used:     0,
		Code:     valcode,
		Identity: req.Email,
	}

	err = p.ValcodeRepository.Insert(tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	createdTime = time.Now().Unix()
	return
}

func (p *ValcodeUsecase) RequestMobileValcode(ctx *biz.BizContext, req *models.RequestMobileValcodeReq) (createdTime int64, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	mobile := req.Prefix + req.Mobile
	has, err := p.ValcodeRepository.HasSentRecordsInDuration(tx, mobile, req.CodeType, models.ValcodeSpan)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if has {
		err = berr.Errorf("60秒下发一次验证码，请不要重复请求")
		return
	}

	valcode := helper.GenerateValcode(6)

	phone := mobile
	if req.Prefix == "86" { // if is china's number, use it without prefix
		phone = req.Mobile
	}

	switch req.CodeType {
	case models.ValcodeRegister:
		if e := p.SMSClient.SendRegisterValcode(phone, valcode); e != nil {
			err = tracerr.Wrap(e)
			return
		}
		break
	case models.ValcodeForgetPassword:
		if e := p.SMSClient.SendResetPasswordValcode(phone, valcode); e != nil {
			err = tracerr.Wrap(e)
			return
		}
		break
	case models.ValcodeLogin:
		if e := p.SMSClient.SendResetPasswordValcode(phone, valcode); e != nil {
			err = tracerr.Wrap(e)
			return
		}
		break
	default:
		err = berr.Errorf("未知的验证码类型")
		return
	}

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	item := &repo.Valcode{
		ID:       id,
		CodeType: req.CodeType,
		Used:     0,
		Code:     valcode,
		Identity: mobile,
	}

	err = p.ValcodeRepository.Insert(tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	createdTime = time.Now().Unix()
	return
}
