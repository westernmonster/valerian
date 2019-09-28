package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/passport-login/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) MobileRegister(c context.Context, arg *model.ArgMobile) (resp *model.LoginResp, err error) {
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
		ID:        gid.NewID(),
		Source:    arg.Source,
		IP:        ipAddr,
		Mobile:    mobile,
		Password:  passwordHash,
		Salt:      salt,
		Role:      model.AccountRoleUser,
		Avatar:    "https://flywiki.oss-cn-hangzhou.aliyuncs.com/765-default-avatar.png",
		UserName:  asteriskMobile(arg.Mobile),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
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

	if account, e := p.d.GetAccountByMobile(c, tx, mobile); e != nil {
		return nil, e
	} else if account != nil {
		err = ecode.AccountExist
		return
	}

	if err = p.d.AddAccount(c, tx, item); err != nil {
		return
	}
	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelMobileValcodeCache(context.TODO(), model.ValcodeRegister, mobile)
	})

	return p.loginAccount(c, item.ID, arg.ClientID)
}
