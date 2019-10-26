package service

import (
	"context"
	"crypto/rand"
	"io"
	"time"
	"valerian/app/interface/valcode/model"
	"valerian/library/ecode"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func generateValcode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (p *Service) EmailValcode(c context.Context, req *model.ArgEmailValcode) (createdTime int64, err error) {

	var exist bool
	if exist, err = p.d.IsEmailExist(c, req.Email); err != nil {
		return
	} else if exist {
		err = ecode.AccountExist
		return
	}

	code := generateValcode(6)

	switch req.CodeType {
	case model.ValcodeRegister:
		if err = p.email.SendRegisterEmail(c, req.Email, code); err != nil {
			return
		}
		break
	case model.ValcodeForgetPassword:
		if err = p.email.SendResetPasswordValcode(c, req.Email, code); err != nil {
			return
		}
		break
	}

	if err = p.d.SetEmailValcodeCache(c, req.CodeType, req.Email, code); err != nil {
		return
	}

	createdTime = time.Now().Unix()

	return
}

func (p *Service) MobileValcode(c context.Context, req *model.ArgMobileValcode) (createdTime int64, err error) {

	var exist bool
	if exist, err = p.d.IsMobileExist(c, req.Prefix, req.Mobile); err != nil {
		return
	} else if exist {
		err = ecode.AccountExist
		return
	}

	var code string
	if code, err = p.d.MobileValcodeCache(c, req.CodeType, req.Mobile); err != nil {
		return
	}

	if code != "" {
		err = ecode.MobileValcodeLimitExceed
		return
	}

	code = generateValcode(6)

	switch req.CodeType {
	case model.ValcodeLogin:
		if err = p.sms.SendLoginValcode(c, req.Prefix, req.Mobile, code); err != nil {
			return
		}
		break
	case model.ValcodeRegister:
		if err = p.sms.SendRegisterValcode(c, req.Prefix, req.Mobile, code); err != nil {
			return
		}
		break
	case model.ValcodeForgetPassword:
		if err = p.sms.SendResetPasswordValcode(c, req.Prefix, req.Mobile, code); err != nil {
			return
		}
		break
	}

	if err = p.d.SetMobileValcodeCache(c, req.CodeType, req.Prefix+req.Mobile, code); err != nil {
		return
	}

	createdTime = time.Now().Unix()

	return
}
