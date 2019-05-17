package usecase

import (
	"context"
	"time"

	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/gid"

	"github.com/ztrue/tracerr"

	"valerian/infrastructure/berr"
	"valerian/infrastructure/biz"
	"valerian/infrastructure/helper"
	"valerian/models"
	"valerian/modules/repo"
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
		// HasSentRecordsInDuration determine current identity has sent records in specified duration
		HasSentRecordsInDuration(ctx context.Context, node sqalx.Node, identity string, codeType int, duration time.Duration) (has bool, err error)
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.Valcode) (err error)
		// Update update a exist record
		Update(ctx context.Context, node sqalx.Node, item *repo.Valcode) (err error)
	}
}

func (p *ValcodeUsecase) RequestEmailValcode(c context.Context, ctx *biz.BizContext, req *models.RequestEmailValcodeReq) (createdTime int64, err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	has, err := p.ValcodeRepository.HasSentRecordsInDuration(c, tx, req.Email, req.CodeType, models.ValcodeSpan)
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

	err = p.ValcodeRepository.Insert(c, tx, item)
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

func (p *ValcodeUsecase) RequestMobileValcode(c context.Context, ctx *biz.BizContext, req *models.RequestMobileValcodeReq) (createdTime int64, err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	mobile := req.Prefix + req.Mobile
	has, err := p.ValcodeRepository.HasSentRecordsInDuration(c, tx, mobile, req.CodeType, models.ValcodeSpan)
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

	err = p.ValcodeRepository.Insert(c, tx, item)
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
