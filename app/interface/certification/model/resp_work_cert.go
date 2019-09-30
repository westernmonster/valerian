package model

type WorkCertResp struct {
	// 实名认证姓名
	IDName string `json:"id_name"`

	// 工作证
	WorkPic string `json:"work_pic"`
	// 其他证明
	OtherPic string `json:"other_pic"`
	// 公司
	Company string `json:"company"`
	// 部门
	Department string `json:"department"`
	// 职位
	Position string `json:""position`

	// 工作证有效期
	// 过期时间
	// Unix时间戳
	ExpiresAt int64 `json:"expires_at"`

	// 认证状态
	// 0 审核中
	// 1 通过审核
	// 2 审核失败
	Status int `json:"status"`

	// 提交时间
	CreatedAt int64 `created_at`

	// 审核时间
	AuditAt int64 `audit_at`

	// 审核结果
	Result string `result`
}
