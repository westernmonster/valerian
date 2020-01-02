package service

import (
	"context"
	"time"

	api "valerian/app/service/identify/api/grpc"
	"valerian/app/service/identify/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"
)

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

	if err = p.addAccount(c, p.d.DB(), item); err != nil {
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

	if err = p.addAccount(c, p.d.DB(), item); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelEmailValcodeCache(context.TODO(), model.ValcodeRegister, arg.Email)
		p.onAccountAdded(context.TODO(), item.ID, time.Now().Unix())
	})

	return p.loginAccount(c, item, arg.ClientID)
}
