package models

import validation "github.com/go-ozzo/ozzo-validation"

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire,string" swaggertype:"string"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}

type CallbackParam struct {
	CallbackUrl      string `json:"callback_url"`
	CallbackBody     string `json:"callback_body"`
	CallbackBodyType string `json:"callback_body_type"`
}

type RequestOSSTokenReq struct {
	// 文件类型
	// file: 文章附件
	// image: 图片
	// certificatin: 工作证件照片
	// other: 其他
	//
	// required: true
	FileType string `json:"file_type"`
}

func (p *RequestOSSTokenReq) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileType,
			validation.Required.Error(`请输入文件类型`),
			validation.In("file", "image", "other", "certificatin").Error("文件类型不在允许范围内"),
		),
	)
}
