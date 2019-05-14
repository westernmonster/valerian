package service

import (
	"context"

	"valerian/app/interface/passport-login/model"
)

func (p *Service) EmailLogin(ctx context.Context, req *model.ArgEmailLogin) (loginResult *model.LoginResp, err error) {
	return
}

func (p *Service) MobileLogin(ctx context.Context, req *model.ArgMobileLogin) (loginResult *model.LoginResp, err error) {
	return
}

func (p *Service) DigitLogin(ctx context.Context, req *model.ArgMobileLogin) (loginResult *model.LoginResp, err error) {
	return
}
