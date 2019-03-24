package email

import (
	"github.com/denverdino/aliyungo/dm"
)

const ()

type EmailClient struct {
	*dm.Client
}

func (p *EmailClient) SendRegisterEmail(email string, valcode string) (err error) {
	args := &dm.SendSingleMailArgs{
		ReplyToAddress: false,
		ToAddress:      email,
		FromAlias:      "飞行百科",
		Subject:        "飞行百科注册验证码",
	}

	err = p.Client.SendSingleMail(args)
	return
}

func (p *EmailClient) SendResetPasswordValcode(email string, valcode string) (err error) {
	args := &dm.SendSingleMailArgs{
		ReplyToAddress: false,
		ToAddress:      email,
		FromAlias:      "飞行百科",
		Subject:        "飞行百科重置密码",
	}

	err = p.Client.SendSingleMail(args)
	return
}
