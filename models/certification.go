package models

type WorkCertification struct {
	// 工作证照片
	WorkCertUrl string `json:"work_cert_url"`
	// 其他证件照片
	OtherCertUrl string `json:"other_cert_url"`
	// 所在单位
	Company string `json:"company"`
	// 部门
	Department string `json:"department"`
	// 职位
	Position string `json:"position"`
}
