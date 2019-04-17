package sms

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

const (
	SignName                  = "飞行百科"          // 短信签名
	RegisterTemplateCode      = "SMS_161380530" // 注册验证码模板
	ResetPasswordTemplateCode = "SMS_161380531" // 重置密码模板
	SendSms                   = "SendSms"
	SendBatchSms              = "SendBatchSms"
	QuerySendDetails          = "QuerySendDetails"
	SmsReport                 = "SmsReport"
	SmsUp                     = "SmsUp"
	Version                   = "2017-05-25"
	EndPoint                  = "dysmsapi.aliyuncs.com"
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

func (p *SMSClient) SendRegisterValcode(mobile string, valcode string) (err error) {
	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = SendSms
	request.QueryParams["PhoneNumbers"] = mobile
	request.QueryParams["SignName"] = SignName
	request.QueryParams["TemplateCode"] = RegisterTemplateCode
	request.QueryParams["Action"] = SendSms
	request.QueryParams["TemplateParam"] = fmt.Sprintf(`{"code":"%s"}`, valcode)

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix": "sms",
			"method": "SendRegisterValcode",
			"mobile": mobile,
		}).Error(fmt.Sprintf("ProcessCommonRequest: %v", err))
		err = tracerr.Errorf("下发短信失败")
		return
	}

	if !response.IsSuccess() {
		logrus.WithFields(logrus.Fields{
			"prefix":       "sms",
			"method":       "SendRegisterValcode",
			"mobile":       mobile,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("HTTP Status: %v", err))

		err = tracerr.Errorf("下发短信失败")
		return
	}
	data := response.GetHttpContentBytes()
	sr := new(smsResponse)
	err = json.Unmarshal(data, sr)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":       "sms",
			"method":       "SendRegisterValcode",
			"mobile":       mobile,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Unmarshal Response: %v", err))
		err = tracerr.Errorf("下发短信失败")
		return
	}

	if sr.Code != "OK" {
		logrus.WithFields(logrus.Fields{
			"prefix":       "sms",
			"method":       "SendRegisterValcode",
			"mobile":       mobile,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Response Fail Message: %v", sr.Message))
		err = tracerr.Errorf("下发短信失败")
		return

	}

	return
}

func (p *SMSClient) SendResetPasswordValcode(mobile string, valcode string) (err error) {
	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = SendSms
	request.QueryParams["PhoneNumbers"] = mobile
	request.QueryParams["SignName"] = SignName
	request.QueryParams["TemplateCode"] = ResetPasswordTemplateCode
	request.QueryParams["Action"] = SendSms
	request.QueryParams["TemplateParam"] = fmt.Sprintf(`{"code":"%s"}`, valcode)

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix": "sms",
			"method": "SendRegisterValcode",
			"mobile": mobile,
		}).Error(fmt.Sprintf("ProcessCommonRequest: %v", err))
		err = tracerr.Errorf("下发短信失败")
		return
	}

	if !response.IsSuccess() {
		logrus.WithFields(logrus.Fields{
			"prefix":       "sms",
			"method":       "SendRegisterValcode",
			"mobile":       mobile,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("HTTP Status: %v", err))

		err = tracerr.Errorf("下发短信失败")
		return
	}
	data := response.GetHttpContentBytes()
	sr := new(smsResponse)
	err = json.Unmarshal(data, sr)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":       "sms",
			"method":       "SendRegisterValcode",
			"mobile":       mobile,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Unmarshal Response: %v", err))
		err = tracerr.Errorf("下发短信失败")
		return
	}

	if sr.Code != "OK" {
		logrus.WithFields(logrus.Fields{
			"prefix":       "sms",
			"method":       "SendRegisterValcode",
			"mobile":       mobile,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Response Fail Message: %v", sr.Message))
		err = tracerr.Errorf("下发短信失败")
		return

	}

	return
}
