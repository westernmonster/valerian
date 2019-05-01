package usecase

import (
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ztrue/tracerr"

	"valerian/infrastructure/berr"
	"valerian/infrastructure/biz"
	"valerian/infrastructure/gid"
	"valerian/infrastructure/helper"
	"valerian/models"
	"valerian/modules/repo"
)

func (p *OauthUsecase) ForgetPassword(ctx *biz.BizContext, req *models.ForgetPasswordReq) (sessionID int64, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	var accountID int64

	if govalidator.IsEmail(req.Identity) {
		item, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
			"email": req.Identity,
		})
		if errGet != nil {
			err = tracerr.Wrap(errGet)
			return
		}
		if !exist {
			err = berr.Errorf("该邮件未注册")
			return
		}

		accountID = item.ID

	} else {
		item, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
			"mobile": req.Identity,
		})
		if errGet != nil {
			err = tracerr.Wrap(errGet)
			return
		}
		if !exist {
			err = berr.Errorf("该手机未注册")
			return
		}
		accountID = item.ID
	}

	// Valcode
	correct, valcodeItem, errValcode := p.ValcodeRepository.IsCodeCorrect(tx, req.Identity, models.ValcodeForgetPassword, req.Valcode)
	if errValcode != nil {
		err = tracerr.Wrap(errValcode)
		return
	}
	if !correct {
		err = berr.Errorf("验证码不正确或已经使用")
		return
	}
	valcodeItem.Used = 1

	err = p.ValcodeRepository.Update(tx, valcodeItem)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	session := &repo.Session{
		ID:          id,
		SessionType: models.SessionTypeResetPassword,
		Used:        0,
		AccountID:   accountID,
	}

	err = p.SessionRepository.Insert(tx, session)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	sessionID = id

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

}

func (p *OauthUsecase) ResetPassword(ctx *biz.BizContext, req *models.ResetPasswordReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	idStr, err := helper.Base64Decode(req.SessionID)
	if err != nil {
		err = berr.Errorf("错误的 Session ID")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		err = berr.Errorf("错误的 Session ID")
		return
	}

	session, exist, err := p.SessionRepository.GetByID(tx, id)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !exist {
		err = berr.Errorf("未获取到当前Session")
		return
	}

	if session.Used == 1 {
		err = berr.Errorf("当前Session已失效")
		return
	}

	if time.Now().Sub(time.Unix(session.CreatedAt, 0)) > 2*time.Minute {
		err = berr.Errorf("当前Session已失效")
		return
	}

	account, exist, err := p.AccountRepository.GetByID(tx, session.AccountID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("未找到当前用户")
		return
	}

	account.Password = req.Password

	err = p.AccountRepository.Update(tx, account)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}
