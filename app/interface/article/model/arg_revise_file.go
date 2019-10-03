package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UpdateReviseFile struct {
	ID       *int64 `json:"id,string,omitempty" swaggertype:"string"`
	FileName string `json:"file_name"` // FileName 文件名
	FileURL  string `json:"file_url"`  // FileURL 文件地址
	Seq      int    `json:"seq"`       // Seq 文件顺序
}

func (p *UpdateReviseFile) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FileName, validation.Required),
		validation.Field(&p.FileURL, validation.Required, is.URL),
	)
}

type ArgSaveReviseFiles struct {
	ReviseID int64               `json:"revise_id,string" swaggertype:"string"`
	Items    []*UpdateReviseFile `json:"items"`
}

func (p *ArgSaveReviseFiles) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.ReviseID, validation.Required),
		validation.Field(&p.Items),
	)
}
