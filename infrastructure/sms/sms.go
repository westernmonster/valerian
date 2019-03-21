package sms

import (
	"fmt"

	"github.com/denverdino/aliyungo/sms"
)

const (
	SignName                  = "飞行百科"          // 短信签名
	RegisterTemplateCode      = "SMS_161380530" // 注册验证码模板
	ResetPasswordTemplateCode = "SMS_161380531" // 重置密码模板
)

type SMSClient struct {
	sms.Client
}

func (p *SMSClient) SendRegisterValcode(mobile string, valcode string) (err error) {
	args := &sms.SingleSendSmsArgs{
		SignName:     SignName,
		TemplateCode: RegisterTemplateCode,
		RecNum:       mobile,
		ParamString:  fmt.Sprintf(`{"code":"%s"}`, valcode),
	}

	err = p.Client.SingleSendSms(args)
	return
}

func (p *SMSClient) SendResetPasswordValcode(mobile string, valcode string) (err error) {
	args := &sms.SingleSendSmsArgs{
		SignName:     SignName,
		TemplateCode: ResetPasswordTemplateCode,
		RecNum:       mobile,
		ParamString:  fmt.Sprintf(`{"code":"%s"}`, valcode),
	}

	err = p.Client.SingleSendSms(args)
	return
}
