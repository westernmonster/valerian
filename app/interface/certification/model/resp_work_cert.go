package model

type WorkCertResp struct {
	AccountID int64 `json:"account_id"`
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

	//  状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败
	Status int `json:"status"`

	// 提交时间
	CreatedAt int64 `created_at`

	// 审核时间
	AuditAt int64 `audit_at`

	// 审核结果
	Result string `result`
}
