package service

import (
	"context"
	"fmt"
	"time"

	api "valerian/app/service/identify/api/grpc"
	"valerian/app/service/identify/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/log"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// getAccountByID 通过ID获取用户
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

// AddAccount 添加账户
func (p *Service) addAccount(c context.Context, node sqalx.Node, item *model.Account) (err error) {
	var tx sqalx.Node
	if tx, err = node.Beginx(c); err != nil {
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

	var account *model.Account
	var e error
	if item.Mobile != "" {
		if account, e = p.d.GetAccountByMobile(c, tx, item.Mobile); e != nil {
			return e
		} else if account != nil {
			err = ecode.AccountExist
			return
		}
	} else {
		if account, e = p.d.GetAccountByEmail(c, tx, item.Email); e != nil {
			return e
		} else if account != nil {
			err = ecode.AccountExist
			return
		}
	}
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	if err = p.d.AddAccount(c, tx, item); err != nil {
		return
	}

	if err = p.d.AddAccountStat(c, tx, &model.AccountStat{
		AccountID: item.ID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = p.d.AddMessageStat(c, tx, &model.MessageStat{
		AccountID: item.ID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

// ForgetPassword 忘记密码
// 匹配验证码并生成一个有效期为5分钟的 SESSIONID
func (p *Service) ForgetPassword(c context.Context, arg *api.ForgetPasswordReq) (resp *api.ForgetPasswordResp, err error) {
	var account *model.Account
	if govalidator.IsEmail(arg.Identity) {
		if account, err = p.d.GetAccountByEmail(c, p.d.DB(), arg.Identity); err != nil {
			return
		}
		var code string
		if code, err = p.d.EmailValcodeCache(c, model.ValcodeForgetPassword, arg.Identity); err != nil {
			return
		} else if code == "" {
			err = ecode.ValcodeExpires
			return
		} else if code != arg.Valcode {
			err = ecode.ValcodeWrong
			return
		}
	} else {
		mobile := arg.Prefix + arg.Identity
		if account, err = p.d.GetAccountByMobile(c, p.d.DB(), mobile); err != nil {
			return
		}

		var code string
		if code, err = p.d.MobileValcodeCache(c, model.ValcodeForgetPassword, mobile); err != nil {
			return
		} else if code == "" {
			err = ecode.ValcodeExpires
			return
		} else if code != arg.Valcode {
			err = ecode.ValcodeWrong
			return
		}
	}

	sessionID := uuid.NewV4().String()
	if err = p.d.SetSessionResetPasswordCache(c, sessionID, account.ID); err != nil {
		return
	}

	resp = &api.ForgetPasswordResp{
		SessionID: sessionID,
	}
	return
}

// ResetPassword 重设密码
func (p *Service) ResetPassword(c context.Context, arg *api.ResetPasswordReq) (err error) {
	var aid int64
	if aid, err = p.d.SessionResetPasswordCache(c, arg.SessionID); err != nil {
		return
	} else if aid == 0 {
		return ecode.SessionExpires
	}

	var acc *model.Account
	if acc, err = p.getAccountByID(c, p.d.DB(), aid); err != nil {
		return
	}

	passwordHash, err := hashPassword(arg.Password, acc.Salt)
	if err != nil {
		return
	}

	if err = p.d.SetPassword(c, p.d.DB(), passwordHash, acc.Salt, aid); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelAccountCache(context.TODO(), aid)
		p.d.DelResetPasswordCache(context.TODO(), arg.SessionID)
		p.deleteAllToken(context.TODO(), aid)
	})

	return
}

// SetPassword 变更密码
func (p *Service) SetPassword(c context.Context, arg *api.SetPasswordReq) (err error) {
	var acc *model.Account
	if acc, err = p.getAccountByID(c, p.d.DB(), arg.Aid); err != nil {
		return
	}

	passwordHash, err := hashPassword(arg.Password, acc.Salt)
	if err != nil {
		return
	}

	if err = p.d.SetPassword(c, p.d.DB(), passwordHash, acc.Salt, arg.Aid); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelAccountCache(context.TODO(), arg.Aid)
		p.deleteAllToken(context.TODO(), arg.Aid)
	})

	return
}

// AccountLock 锁定用户
func (p *Service) AccountLock(c context.Context, req *api.LockReq) (err error) {
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

	if err = p.checkSystemAdmin(c, tx, req.Aid); err != nil {
		return
	}

	if err = p.d.AccountSetLock(c, tx, req.TargetAccountID, true); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	return
}

// AccountLock 解锁用户
func (p *Service) AccountUnlock(c context.Context, req *api.LockReq) (err error) {
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

	if err = p.checkSystemAdmin(c, tx, req.Aid); err != nil {
		return
	}

	if err = p.d.AccountSetLock(c, tx, req.TargetAccountID, false); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	return
}

// Deactive 注销用户
func (p *Service) Deactive(c context.Context, arg *api.DeactiveReq) (err error) {
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

	var account *model.Account
	if govalidator.IsEmail(arg.Identity) {
		if account, err = p.d.GetAccountByEmail(c, tx, arg.Identity); err != nil {
			return
		}
		var code string
		if code, err = p.d.EmailValcodeCache(c, model.ValcodeDeactive, arg.Identity); err != nil {
			return
		} else if code == "" {
			err = ecode.ValcodeExpires
			return
		} else if code != arg.Valcode {
			err = ecode.ValcodeWrong
			return
		}
	} else {
		mobile := arg.Prefix + arg.Identity
		if account, err = p.d.GetAccountByMobile(c, tx, mobile); err != nil {
			return
		}

		var code string
		if code, err = p.d.MobileValcodeCache(c, model.ValcodeDeactive, mobile); err != nil {
			return
		} else if code == "" {
			err = ecode.ValcodeExpires
			return
		} else if code != arg.Valcode {
			err = ecode.ValcodeWrong
			return
		}
	}

	if err = p.d.DeactiveAccount(c, tx, account.ID); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelAccountCache(context.TODO(), account.ID)
		p.deleteAllToken(context.TODO(), account.ID)
	})

	return
}
