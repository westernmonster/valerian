package service

import (
	"context"
	"strings"

	"valerian/app/interface/passport-login/model"
	"valerian/library/ecode"
	"valerian/models"
)

func (p *Service) EmailLogin(ctx context.Context, req *model.ArgEmailLogin) (resp *model.LoginResp, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	account, err := p.d.GetAccountByEmail(ctx, p.d.DB(), req.Email)
	if err != nil {
		return
	}
	if account == nil {
		err = ecode.UserNotExist
		return
	}

	if err = p.checkPassword(req.Password, account.Password, account.Salt); err != nil {
		return
	}

	accessToken, refreshToken, err := p.grantToken(ctx, req.ClientID, account.ID)
	if err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID:    account.ID,
		Role:         account.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    models.ExpiresIn,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	if resp.Profile, err = p.GetProfile(ctx, account.ID); err != nil {
		return
	}

	p.addCache(func() {
		p.d.SetAccessTokenCache(context.TODO(), accessToken)
	})

	return
}

func (p *Service) MobileLogin(ctx context.Context, req *model.ArgMobileLogin) (resp *model.LoginResp, err error) {
	if err = p.checkClient(ctx, req.ClientID); err != nil {
		return
	} // Check Client

	account, err := p.d.GetAccountByMobile(ctx, p.d.DB(), req.Mobile)
	if err != nil {
		return
	}
	if account == nil {
		err = ecode.UserNotExist
		return
	}

	if err = p.checkPassword(req.Password, account.Password, account.Salt); err != nil {
		return
	}

	accessToken, refreshToken, err := p.grantToken(ctx, req.ClientID, account.ID)
	if err != nil {
		return
	}

	resp = &model.LoginResp{
		AccountID:    account.ID,
		Role:         account.Role,
		AccessToken:  accessToken.Token,
		ExpiresIn:    models.ExpiresIn,
		TokenType:    "Bearer",
		Scope:        "",
		RefreshToken: refreshToken.Token,
	}

	if resp.Profile, err = p.GetProfile(ctx, account.ID); err != nil {
		return
	}

	p.addCache(func() {
		p.d.SetAccessTokenCache(context.TODO(), accessToken)
	})

	return
}

func (p *Service) DigitLogin(ctx context.Context, req *model.ArgDigitLogin) (resp *model.LoginResp, err error) {
	return
}

func (p *Service) checkPassword(password, dbPassword, dbSalt string) (err error) {
	passwordHash, err := hashPassword(password, dbSalt)
	if err != nil {
		return
	}

	if !strings.EqualFold(dbPassword, passwordHash) {
		err = ecode.PasswordErr
		return
	}
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
