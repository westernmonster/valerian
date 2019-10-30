package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"valerian/library/log"
	"valerian/library/tracing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	SignName                       = "飞行百科"          // 短信签名
	RegisterTemplateCode           = "SMS_166865016" // 注册验证码模板
	ResetPasswordTemplateCode      = "SMS_166865016" // 重置密码模板
	LoginTemplateCode              = "SMS_166776690" // 注册验证码模板
	ChinaSignName                  = "飞行百科"          // 短信签名
	ChinaRegisterTemplateCode      = "SMS_161380530" // 注册验证码模板
	ChinaResetPasswordTemplateCode = "SMS_161380531" // 重置密码模板
	ChinaLoginTemplateCode         = "SMS_161380530" // 注册验证码模板
	SendSms                        = "SendSms"
	SendBatchSms                   = "SendBatchSms"
	QuerySendDetails               = "QuerySendDetails"
	SmsReport                      = "SmsReport"
	SmsUp                          = "SmsUp"
	Version                        = "2017-05-25"
	EndPoint                       = "dysmsapi.aliyuncs.com"
)

// smsResponse .
type smsResponse struct {
	RequestID string `json:"RequestId"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	BizID     string `json:"BizId"`
}

type SMSClient struct {
	Client *sdk.Client
}

func (p *SMSClient) SendLoginValcode(c context.Context, prefix string, mobile, valcode string) (err error) {
	if prefix == "86" {
		return p.sendLoginValcode(c, mobile, valcode, ChinaSignName, ChinaLoginTemplateCode)
	} else {
		return p.sendLoginValcode(c, prefix+mobile, valcode, SignName, LoginTemplateCode)
	}
}

func (p *SMSClient) sendLoginValcode(c context.Context, mobile string, valcode string, signName, template string) (err error) {
	if span := opentracing.SpanFromContext(c); span != nil {
		span := tracing.StartSpan("sms", opentracing.ChildOf(span.Context()))
		span.SetTag("param.mobile", mobile)
		span.SetTag("param.type", "Login")
		ext.SpanKindRPCClient.Set(span)
		defer span.Finish()
		c = opentracing.ContextWithSpan(c, span)
	}

	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = SendSms
	request.QueryParams["PhoneNumbers"] = mobile
	request.QueryParams["SignName"] = signName
	request.QueryParams["TemplateCode"] = template
	request.QueryParams["Action"] = SendSms
	request.QueryParams["TemplateParam"] = fmt.Sprintf(`{"code":"%s"}`, valcode)

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendLoginValcode error(%+v), mobile(%s)", err, mobile))
		return
	}

	if !response.IsSuccess() {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendLoginValcode error(%+v), mobile(%s) resp(%+v, %+v)", err, mobile, response.GetHttpStatus(), response.GetHttpContentString()))
		return
	}
	data := response.GetHttpContentBytes()
	sr := new(smsResponse)
	err = json.Unmarshal(data, sr)

	if err != nil {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendLoginValcode error(%+v), mobile(%s) resp(%+v, %+v)", err, mobile, response.GetHttpStatus(), response.GetHttpContentString()))
		return
	}

	if sr.Code != "OK" {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendLoginValcode error(%+v), mobile(%s) resp(%+v, %+v)", err, mobile, response.GetHttpStatus(), response.GetHttpContentString()))
		return

	}

	return
}

func (p *SMSClient) SendRegisterValcode(c context.Context, prefix, mobile, valcode string) (err error) {
	if prefix == "86" {
		return p.sendRegisterValcode(c, mobile, valcode, ChinaSignName, ChinaRegisterTemplateCode)
	} else {
		return p.sendRegisterValcode(c, prefix+mobile, valcode, SignName, RegisterTemplateCode)
	}
}

func (p *SMSClient) sendRegisterValcode(c context.Context, mobile string, valcode, signName, template string) (err error) {
	if span := opentracing.SpanFromContext(c); span != nil {
		span := tracing.StartSpan("sms", opentracing.ChildOf(span.Context()))
		span.SetTag("param.mobile", mobile)
		span.SetTag("param.type", "Register")
		ext.SpanKindRPCClient.Set(span)
		defer span.Finish()
		c = opentracing.ContextWithSpan(c, span)
	}

	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = SendSms
	request.QueryParams["PhoneNumbers"] = mobile
	request.QueryParams["SignName"] = signName
	request.QueryParams["TemplateCode"] = template
	request.QueryParams["Action"] = SendSms
	request.QueryParams["TemplateParam"] = fmt.Sprintf(`{"code":"%s"}`, valcode)

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendRegisterValcode error(%+v), mobile(%s)", err, mobile))
		return
	}

	if !response.IsSuccess() {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendRegisterValcode error(%+v), mobile(%s) resp(%+v, %+v)", err, mobile, response.GetHttpStatus(), response.GetHttpContentString()))
		return
	}
	data := response.GetHttpContentBytes()
	sr := new(smsResponse)
	err = json.Unmarshal(data, sr)

	if err != nil {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendRegisterValcode error(%+v), mobile(%s) resp(%+v, %+v)", err, mobile, response.GetHttpStatus(), response.GetHttpContentString()))
		return
	}

	if sr.Code != "OK" {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendRegisterValcode error(%+v), mobile(%s) resp(%+v, %+v)", err, mobile, response.GetHttpStatus(), response.GetHttpContentString()))
		return

	}

	return
}

func (p *SMSClient) SendResetPasswordValcode(c context.Context, prefix, mobile string, valcode string) (err error) {
	if prefix == "86" {
		return p.sendResetPasswordValcode(c, mobile, valcode, ChinaSignName, ChinaResetPasswordTemplateCode)
	} else {
		return p.sendResetPasswordValcode(c, prefix+mobile, valcode, SignName, ResetPasswordTemplateCode)
	}
}
func (p *SMSClient) sendResetPasswordValcode(c context.Context, mobile string, valcode, signName, template string) (err error) {
	if span := opentracing.SpanFromContext(c); span != nil {
		span := tracing.StartSpan("sms", opentracing.ChildOf(span.Context()))
		span.SetTag("param.mobile", mobile)
		span.SetTag("param.type", "ResetPassword")
		ext.SpanKindRPCClient.Set(span)
		defer span.Finish()
		c = opentracing.ContextWithSpan(c, span)
	}
	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = SendSms
	request.QueryParams["PhoneNumbers"] = mobile
	request.QueryParams["SignName"] = signName
	request.QueryParams["TemplateCode"] = ResetPasswordTemplateCode
	request.QueryParams["Action"] = SendSms
	request.QueryParams["TemplateParam"] = fmt.Sprintf(`{"code":"%s"}`, valcode)

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendResetPasswordValcode error(%+v), mobile(%s)", err, mobile))
		return
	}

	if !response.IsSuccess() {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendResetPasswordValcode error(%+v), mobile(%s) resp(%+v, %+v)", err, mobile, response.GetHttpStatus(), response.GetHttpContentString()))
		return
	}
	data := response.GetHttpContentBytes()
	sr := new(smsResponse)
	err = json.Unmarshal(data, sr)

	if err != nil {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendResetPasswordValcode error(%+v), mobile(%s) resp(%+v, %+v)", err, mobile, response.GetHttpStatus(), response.GetHttpContentString()))
		return
	}

	if sr.Code != "OK" {
		log.For(c).Error(fmt.Sprintf("SMSClient.SendResetPasswordValcode error(%+v), mobile(%s) resp(%+v, %+v)", err, mobile, response.GetHttpStatus(), response.GetHttpContentString()))
		return

	}

	return
}
