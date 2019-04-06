package email

import (
	"git.flywk.com/flywiki/api/infrastructure/email/tmpl"
	"git.flywk.com/flywiki/api/infrastructure/email/tmpl/layouts"
	"github.com/denverdino/aliyungo/dm"
)

const ()

type EmailClient struct {
	*dm.Client
}

func (p *EmailClient) SendRegisterEmail(email string, valcode string) (err error) {
	args := &dm.SendSingleMailArgs{
		SendEmailArgs: dm.SendEmailArgs{
			AccountName: "noreply@flywki.com",
			AddressType: "1",
		},
		ReplyToAddress: false,
		ToAddress:      email,
		FromAlias:      "飞行百科",
		Subject:        "飞行百科注册验证码",
	}

	body := &tmpl.RegisterValcodeBody{
		Head: &layouts.EmailPageHead{
			Title:     "飞行百科注册验证码",
			BodyClass: "",
		},
		Valcode: valcode,
	}

	args.HtmlBody = body.EmailHTML()

	err = p.Client.SendSingleMail(args)
	return
}

func (p *EmailClient) SendResetPasswordValcode(email string, valcode string) (err error) {
	args := &dm.SendSingleMailArgs{
		SendEmailArgs: dm.SendEmailArgs{
			AccountName: "noreply@flywki.com",
			AddressType: "1",
		},
		ReplyToAddress: false,
		ToAddress:      email,
		FromAlias:      "飞行百科",
		Subject:        "飞行百科重置密码",
	}

	body := &tmpl.RegisterValcodeBody{
		Head: &layouts.EmailPageHead{
			Title:     "飞行百科安全验证码",
			BodyClass: "",
		},
		Valcode: valcode,
	}

	args.HtmlBody = body.EmailHTML()

	err = p.Client.SendSingleMail(args)
	return
}
