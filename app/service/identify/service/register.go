package service

import (
	"context"
	"fmt"
	"time"

	api "valerian/app/service/identify/api/grpc"
	"valerian/app/service/identify/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

// AdminCreateAccount 管理员创建账户
func (p *Service) AdminCreateAccount(c context.Context, arg *api.AdminCreateAccountReq) (err error) {
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

	// 检测是否系统管理员操作
	if err = p.checkSystemAdmin(c, tx, arg.Aid); err != nil {
		return
	}

	mobile := arg.Prefix + arg.Mobile
	userName := ""
	if arg.Email != "" {
		var acc *model.Account
		if acc, err = p.d.GetAccountByEmail(c, tx, arg.Email); err != nil {
			return
		} else if acc != nil {
			err = ecode.AccountExist
			return
		}
		userName = asteriskEmailName(arg.Email)
	}

	if arg.Mobile != "" {
		var acc *model.Account
		if acc, err = p.d.GetAccountByMobile(c, tx, mobile); err != nil {
			return
		} else if acc != nil {
			err = ecode.AccountExist
			return
		}
		userName = asteriskMobile(arg.Mobile)
	}

	ip := metadata.String(c, metadata.RemoteIP)
	ipAddr := InetAtoN(ip)
	salt, err := generateSalt(16)
	if err != nil {
		return
	}
	passwordHash, err := hashPassword(arg.Password, salt)
	if err != nil {
		return
	}

	item := &model.Account{
		ID:       gid.NewID(),
		Source:   model.SourceWeb,
		IP:       ipAddr,
		Mobile:   mobile,
		Password: passwordHash,
		Prefix:   arg.Prefix,
		Email:    arg.Email,
		Salt:     salt,
		Role:     model.AccountRoleUser,
		Avatar:   "https://flywiki.oss-cn-hangzhou.aliyuncs.com/765-default-avatar.png",
		UserName: userName,
	}

	if err = p.addAccount(c, p.d.DB(), item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.onAccountAdded(context.TODO(), item.ID, time.Now().Unix())
	})

	return
}

// MobileRegister 手机注册
func (p *Service) MobileRegister(c context.Context, arg *api.MobileRegisterReq) (resp *api.LoginResp, err error) {
	var (
		code string
	)

	mobile := arg.Prefix + arg.Mobile
	if code, err = p.d.MobileValcodeCache(c, model.ValcodeRegister, mobile); err != nil {
		return
	}
	if code == "" {
		return nil, ecode.ValcodeExpires
	}
	if code != arg.Valcode {
		return nil, ecode.ValcodeWrong
	}

	if err = p.checkClient(c, arg.ClientID); err != nil {
		return
	} // Check Client

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

	var acc *model.Account
	if acc, err = p.d.GetAccountByMobile(c, tx, mobile); err != nil {
		return
	} else if acc != nil {
		err = ecode.AccountExist
		return
	}

	ip := metadata.String(c, metadata.RemoteIP)
	ipAddr := InetAtoN(ip)
	salt, err := generateSalt(16)
	if err != nil {
		return
	}
	passwordHash, err := hashPassword(arg.Password, salt)
	if err != nil {
		return
	}

	item := &model.Account{
		ID:       gid.NewID(),
		Source:   arg.Source,
		IP:       ipAddr,
		Mobile:   mobile,
		Password: passwordHash,
		Prefix:   arg.Prefix,
		Salt:     salt,
		Role:     model.AccountRoleUser,
		Avatar:   "https://flywiki.oss-cn-hangzhou.aliyuncs.com/765-default-avatar.png",
		UserName: asteriskMobile(arg.Mobile),
	}

	if err = p.addAccount(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelMobileValcodeCache(context.TODO(), model.ValcodeRegister, mobile)
		p.onAccountAdded(context.TODO(), item.ID, time.Now().Unix())
	})

	// 登录

	return p.loginAccount(c, item, arg.ClientID)
}

// EmailRegister 邮件注册
func (p *Service) EmailRegister(c context.Context, arg *api.EmailRegisterReq) (resp *api.LoginResp, err error) {
	var (
		code string
	)

	if code, err = p.d.EmailValcodeCache(c, model.ValcodeRegister, arg.Email); err != nil {
		return
	}
	if code == "" {
		return nil, ecode.ValcodeExpires
	}
	if code != arg.Valcode {
		return nil, ecode.ValcodeWrong
	}

	if err = p.checkClient(c, arg.ClientID); err != nil {
		return
	} // Check Client

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

	var acc *model.Account
	if acc, err = p.d.GetAccountByEmail(c, tx, arg.Email); err != nil {
		return
	} else if acc != nil {
		err = ecode.AccountExist
		return
	}

	ip := metadata.String(c, metadata.RemoteIP)
	ipAddr := InetAtoN(ip)
	salt, err := generateSalt(16)
	if err != nil {
		return
	}
	passwordHash, err := hashPassword(arg.Password, salt)
	if err != nil {
		return
	}

	item := &model.Account{
		ID:       gid.NewID(),
		Source:   arg.Source,
		IP:       ipAddr,
		Email:    arg.Email,
		Password: passwordHash,
		Salt:     salt,
		Role:     model.AccountRoleUser,
		Avatar:   "https://flywiki.oss-cn-hangzhou.aliyuncs.com/765-default-avatar.png",
		UserName: asteriskEmailName(arg.Email),
	}

	if err = p.addAccount(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelEmailValcodeCache(context.TODO(), model.ValcodeRegister, arg.Email)
		p.onAccountAdded(context.TODO(), item.ID, time.Now().Unix())
	})

	return p.loginAccount(c, item, arg.ClientID)
}
