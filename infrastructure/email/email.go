package email

import (
	"encoding/json"
	"fmt"

	"valerian/infrastructure/email/tmpl"
	"valerian/infrastructure/email/tmpl/layouts"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
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

func (p *EmailClient) SendRegisterEmail(email string, valcode string) (err error) {
	body := &tmpl.RegisterValcodeBody{
		Head: &layouts.EmailPageHead{
			Title:     "飞行百科注册验证码",
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
	request.QueryParams["FromAlias"] = "飞行百科"
	request.QueryParams["Subject"] = "飞行百科注册验证码"
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

func (p *EmailClient) SendResetPasswordValcode(email string, valcode string) (err error) {
	body := &tmpl.RegisterValcodeBody{
		Head: &layouts.EmailPageHead{
			Title:     "飞行百科安全验证码",
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
	request.QueryParams["FromAlias"] = "飞行百科"
	request.QueryParams["Subject"] = "飞行百科安全验证码"
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
