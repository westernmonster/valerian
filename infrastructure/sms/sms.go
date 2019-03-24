package sms

import (
	"fmt"

	"github.com/denverdino/aliyungo/sms"
	"github.com/ztrue/tracerr"
)

const (
	SignName                  = "飞行百科"          // 短信签名
	RegisterTemplateCode      = "SMS_161380530" // 注册验证码模板
	ResetPasswordTemplateCode = "SMS_161380531" // 重置密码模板
)

type SMSClient struct {
	Client *sms.DYSmsClient
}

func (p *SMSClient) SendRegisterValcode(mobile string, valcode string) (err error) {
	args := &sms.SendSmsArgs{
		SignName:      SignName,
		TemplateCode:  RegisterTemplateCode,
		PhoneNumbers:  mobile,
		TemplateParam: fmt.Sprintf(`{"code":"%s"}`, valcode),
	}

	resp, err := p.Client.SendSms(args)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if resp.Code != "OK" {
		err = tracerr.Errorf("下发短信失败")
		return
	}
	return
}

func (p *SMSClient) SendResetPasswordValcode(mobile string, valcode string) (err error) {
	args := &sms.SendSmsArgs{
		SignName:      SignName,
		TemplateCode:  ResetPasswordTemplateCode,
		PhoneNumbers:  mobile,
		TemplateParam: fmt.Sprintf(`{"code":"%s"}`, valcode),
	}

	resp, err := p.Client.SendSms(args)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if resp.Code != "OK" {
		err = tracerr.Errorf("下发短信失败")
		return
	}
	return
}
