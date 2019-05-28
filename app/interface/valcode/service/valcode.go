package service

import (
	"context"
	"crypto/rand"
	"io"
	"time"
	"valerian/app/interface/valcode/model"
	"valerian/library/ecode"
	"valerian/models"
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
	code := generateValcode(6)

	switch req.CodeType {
	case models.ValcodeRegister:
		if err = p.email.SendRegisterEmail(c, req.Email, code); err != nil {
			return
		}
		break
	case models.ValcodeForgetPassword:
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
	case models.ValcodeLogin:
		if err = p.sms.SendLoginValcode(c, req.Mobile, code); err != nil {
			return
		}
		break
	case models.ValcodeRegister:
		if err = p.sms.SendRegisterValcode(c, req.Mobile, code); err != nil {
			return
		}
		break
	case models.ValcodeForgetPassword:
		if err = p.sms.SendResetPasswordValcode(c, req.Mobile, code); err != nil {
			return
		}
		break
	}

	if err = p.d.SetMobileValcodeCache(c, req.CodeType, req.Mobile, code); err != nil {
		return
	}

	createdTime = time.Now().Unix()

	return
}
