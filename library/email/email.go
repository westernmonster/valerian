package email

import (
	"context"
	"encoding/json"
	"fmt"
	"valerian/library/email/tmpl"
	"valerian/library/email/tmpl/layouts"
	"valerian/library/tracing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

const (
	EndPoint       = "dm.aliyuncs.com"
	SingleSendMail = "SingleSendMail"
	BatchSendMail  = "BatchSendMail"
	Version        = "2015-11-23"
)

// emailResponse .
type emailResponse struct {
	RequestID string `json:"RequestId"`
}

type EmailClient struct {
	Client *sdk.Client
}

func (p *EmailClient) SendRegisterEmail(c context.Context, email string, valcode string) (err error) {
	if span := opentracing.SpanFromContext(c); span != nil {
		span := tracing.StartSpan("email", opentracing.ChildOf(span.Context()))
		span.SetTag("param.email", email)
		span.SetTag("param.type", "Register")
		ext.SpanKindRPCClient.Set(span)
		defer span.Finish()
		c = opentracing.ContextWithSpan(c, span)
	}

	body := &tmpl.RegisterValcodeBody{
		Head: &layouts.EmailPageHead{
			Title:     "他石笔记注册验证码",
			BodyClass: "",
		},
		Valcode: valcode,
	}

	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = SingleSendMail
	request.QueryParams["Action"] = SingleSendMail
	request.QueryParams["AccountName"] = "noreply@flywki.com"
	request.QueryParams["AddressType"] = "1"
	request.QueryParams["ReplyToAddress"] = "false"
	request.QueryParams["ToAddress"] = email
	request.QueryParams["FromAlias"] = "他石笔记"
	request.QueryParams["Subject"] = "他石笔记注册验证码"
	request.QueryParams["HtmlBody"] = body.EmailHTML()

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix": "email",
			"method": "SingleSendMail",
			"email":  email,
		}).Error(fmt.Sprintf("ProcessCommonRequest: %v", err))
		err = tracerr.Errorf("下发短信失败")
		return
	}

	if !response.IsSuccess() {
		logrus.WithFields(logrus.Fields{
			"prefix":       "email",
			"method":       "SingleSendMail",
			"email":        email,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("HTTP Status: %v", err))

		err = tracerr.Errorf("下发短信失败")
		return
	}
	data := response.GetHttpContentBytes()
	sr := new(emailResponse)
	err = json.Unmarshal(data, sr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":       "email",
			"method":       "SingleSendMail",
			"email":        email,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Unmarshal Response: %v", err))
		err = tracerr.Errorf("下发短信失败")
		return
	}

	return
}

func (p *EmailClient) SendResetPasswordValcode(c context.Context, email string, valcode string) (err error) {
	if span := opentracing.SpanFromContext(c); span != nil {
		span := tracing.StartSpan("email", opentracing.ChildOf(span.Context()))
		span.SetTag("param.email", email)
		span.SetTag("param.type", "ResetPassword")
		ext.SpanKindRPCClient.Set(span)
		defer span.Finish()
		c = opentracing.ContextWithSpan(c, span)
	}

	body := &tmpl.RegisterValcodeBody{
		Head: &layouts.EmailPageHead{
			Title:     "他石笔记安全验证码",
			BodyClass: "",
		},
		Valcode: valcode,
	}

	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = SingleSendMail
	request.QueryParams["Action"] = SingleSendMail
	request.QueryParams["AccountName"] = "noreply@flywki.com"
	request.QueryParams["AddressType"] = "1"
	request.QueryParams["ReplyToAddress"] = "false"
	request.QueryParams["ToAddress"] = email
	request.QueryParams["FromAlias"] = "他石笔记"
	request.QueryParams["Subject"] = "他石笔记安全验证码"
	request.QueryParams["HtmlBody"] = body.EmailHTML()

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix": "email",
			"method": "SingleSendMail",
			"email":  email,
		}).Error(fmt.Sprintf("ProcessCommonRequest: %v", err))
		err = tracerr.Errorf("下发短信失败")
		return
	}

	if !response.IsSuccess() {
		logrus.WithFields(logrus.Fields{
			"prefix":       "email",
			"method":       "SingleSendMail",
			"email":        email,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("HTTP Status: %v", err))

		err = tracerr.Errorf("下发短信失败")
		return
	}
	data := response.GetHttpContentBytes()
	sr := new(emailResponse)
	err = json.Unmarshal(data, sr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":       "email",
			"method":       "SingleSendMail",
			"email":        email,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Unmarshal Response: %v", err))
		err = tracerr.Errorf("下发短信失败")
		return
	}

	return
}
