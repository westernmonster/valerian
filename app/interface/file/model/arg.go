package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgOSSToken struct {
	// 文件类型
	// file: 文章附件
	// image: 图片
	// certification: 工作证件照片
	// other: 其他
	FileType string `json:"file_type"`

	// 文件名
	FileName string `json:"file_name"`
}

func (p *ArgOSSToken) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileType,
			validation.Required.Error(`请输入文件类型`),
			validation.In("file", "image", "other", "certificatin").Error("文件类型不在允许范围内"),
		),
		validation.Field(&p.FileName,
			validation.Required.Error(`请输入文件名`)),
	)
}
