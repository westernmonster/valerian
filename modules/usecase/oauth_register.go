package usecase

import (
	"context"
	"fmt"

	"github.com/ztrue/tracerr"

	"valerian/infrastructure/berr"
	"valerian/infrastructure/biz"
	"valerian/infrastructure/helper"
	"valerian/library/gid"
	"valerian/models"
	"valerian/modules/repo"
)

func (p *OauthUsecase) GetByID(c context.Context, ctx *biz.BizContext, userID int64) (item *repo.Account, err error) {
	item, exist, err := p.AccountRepository.GetByID(c, p.Node, userID)

	if !exist {
		err = tracerr.Errorf("获取用户信息失败")
		return
	}

	return
}

func (p *OauthUsecase) EmailRegister(c context.Context, ctx *biz.BizContext, req *models.EmailRegisterReq, ip string) (accountID int64, err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	ipAddr := helper.InetAtoN(ip)
	item := &repo.Account{
		ID:     id,
		Source: req.Source,
		IP:     ipAddr,
	}

	_, exist, errGet := p.AccountRepository.GetByCondition(c, tx, map[string]string{
		"email": req.Email,
	})
	if errGet != nil {
		err = tracerr.Wrap(errGet)
		return
	}
	if exist {
		err = berr.Errorf("该邮件地址已经注册")
		return
	}
	item.Email = req.Email

	salt, err := generateSalt(16)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	passwordHash, err := hashPassword(req.Password, salt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	item.Password = passwordHash
	item.Salt = salt
	item.Role = "user"

	// Valcode
	correct, valcodeItem, errValcode := p.ValcodeRepository.IsCodeCorrect(c, tx, req.Email, models.ValcodeRegister, req.Valcode)
	if errValcode != nil {
		err = tracerr.Wrap(errValcode)
		return
	}
	if !correct {
		err = berr.Errorf("验证码不正确或已经使用")
		return
	}
	valcodeItem.Used = 1

	err = p.ValcodeRepository.Update(c, tx, valcodeItem)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.AccountRepository.Insert(c, tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	accountID = id
	return

}

func (p *OauthUsecase) MobileRegister(c context.Context, ctx *biz.BizContext, req *models.MobileRegisterReq, ip string) (accountID int64, err error) {
	tx, err := p.Node.Beginx(c)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	ipAddr := helper.InetAtoN(ip)
	item := &repo.Account{
		ID:     id,
		Source: req.Source,
		IP:     ipAddr,
	}

	mobile := req.Prefix + req.Mobile

	_, exist, errGet := p.AccountRepository.GetByCondition(c, tx, map[string]string{
		"mobile": mobile,
	})
	if errGet != nil {
		err = tracerr.Wrap(errGet)
		return
	}
	if exist {
		err = berr.Errorf("该手机号已经注册")
		return
	}
	item.Mobile = mobile

	salt, err := generateSalt(16)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	fmt.Printf("salt: %s\n", salt)
	passwordHash, err := hashPassword(req.Password, salt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	fmt.Printf("hash: %s\n", passwordHash)
	item.Password = passwordHash
	item.Salt = salt
	item.Role = "user"

	// Valcode
	correct, valcodeItem, errValcode := p.ValcodeRepository.IsCodeCorrect(c, tx, mobile, models.ValcodeRegister, req.Valcode)
	if errValcode != nil {
		err = tracerr.Wrap(errValcode)
		return
	}
	if !correct {
		err = berr.Errorf("验证码不正确或已经使用")
		return
	}
	valcodeItem.Used = 1

	err = p.ValcodeRepository.Update(c, tx, valcodeItem)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.AccountRepository.Insert(c, tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	accountID = id
	return
}
