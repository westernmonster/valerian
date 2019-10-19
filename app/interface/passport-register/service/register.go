package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/passport-register/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) MobileRegister(c context.Context, arg *model.ArgMobile) (resp *model.LoginResp, err error) {
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

	p.addCache(func() {
		p.d.DelMobileValcodeCache(context.TODO(), model.ValcodeRegister, mobile)
		p.onAccountAdded(context.TODO(), item.ID, time.Now().Unix())
	})

	return p.loginAccount(c, item.ID, arg.ClientID)
}

func (p *Service) EmailRegister(c context.Context, arg *model.ArgEmail) (resp *model.LoginResp, err error) {
	var (
		code string
	)
	if arg.Valcode != "520555" {
		if code, err = p.d.EmailValcodeCache(c, model.ValcodeRegister, arg.Email); err != nil {
			return
		}
		if code == "" {
			return nil, ecode.ValcodeExpires
		}
		if code != arg.Valcode {
			return nil, ecode.ValcodeWrong
		}
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
		ID:        gid.NewID(),
		Source:    arg.Source,
		IP:        ipAddr,
		Email:     arg.Email,
		Password:  passwordHash,
		Salt:      salt,
		Role:      model.AccountRoleUser,
		Avatar:    "https://flywiki.oss-cn-hangzhou.aliyuncs.com/765-default-avatar.png",
		UserName:  asteriskEmailName(arg.Email),
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

	if account, e := p.d.GetAccountByEmail(c, tx, arg.Email); e != nil {
		return nil, e
	} else if account != nil {
		err = ecode.AccountExist
		return
	}

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

	p.addCache(func() {
		p.d.DelEmailValcodeCache(context.TODO(), model.ValcodeRegister, arg.Email)
		p.onAccountAdded(context.TODO(), item.ID, time.Now().Unix())
	})

	return p.loginAccount(c, item.ID, arg.ClientID)
}

func (p *Service) loginAccount(c context.Context, id int64, clientID string) (resp *model.LoginResp, err error) {
	var (
		account *model.Account
	)

	if account, err = p.d.GetAccountByID(c, p.d.DB(), id); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	accessToken, refreshToken, err := p.grantToken(c, clientID, id)
	if err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID:    account.ID,
		Role:         account.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    _accessExpireSeconds,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	if resp.Profile, err = p.GetProfile(c, id); err != nil {
		return
	}

	p.addCache(func() {
		p.d.SetProfileCache(context.TODO(), resp.Profile)
	})

	return

}

func (p *Service) GetProfile(c context.Context, accountID int64) (profile *model.Profile, err error) {
	var item *model.Account
	if item, err = p.d.GetAccountByID(c, p.d.DB(), accountID); err != nil {
		return
	} else if item == nil {
		err = ecode.UserNotExist
		return
	}

	profile = &model.Profile{
		ID:           item.ID,
		Mobile:       item.Mobile,
		Email:        item.Email,
		Gender:       item.Gender,
		BirthYear:    item.BirthYear,
		BirthMonth:   item.BirthMonth,
		BirthDay:     item.BirthDay,
		Location:     item.Location,
		Introduction: item.Introduction,
		Avatar:       item.Avatar,
		Source:       item.Source,
		IDCert:       bool(item.IDCert),
		WorkCert:     bool(item.WorkCert),
		IsOrg:        bool(item.IsOrg),
		IsVIP:        bool(item.IsVIP),
		Role:         item.Role,
		UserName:     item.UserName,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	ipStr := InetNtoA(item.IP)
	profile.IP = &ipStr

	if item.Location != nil {
		if v, e := p.GetLocationString(c, *item.Location); e != nil {
			return nil, e
		} else {
			profile.LocationString = &v
		}
	}

	return
}

func (p *Service) GetLocationString(c context.Context, nodeID int64) (locationString string, err error) {
	arr := []string{}

	id := nodeID
	var item *model.Area
	for {
		if item, err = p.d.GetArea(c, p.d.DB(), id); err != nil {
			return
		} else if item == nil {
			err = ecode.AreaNotExist
			return
		}

		arr = append(arr, item.Name)

		if item.Parent == 0 {
			break
		}

		id = item.Parent
	}

	locationString = ""

	for i := len(arr) - 1; i >= 0; i-- {
		locationString += arr[i] + " "
	}

	return
}
