package model

type ArgAddWorkCert struct {
	AccountID  int64  `json:"account_id,string"` // AccountID 账户ID
	WorkPic    string `json:"work_pic"`          // WorkPic 工作证
	OtherPic   string `json:"other_pic"`         // OtherPic 其他证明
	Company    string `json:"company"`           // Company 公司
	Department string `json:"department"`        // Department 部门
	Position   string `json:"position"`          // Position 职位
	ExpiresAt  int64  `json:"expires_at,string"` // ExpiresAt 工作证有效期
}

type ArgAuditWorkCert struct {
	// AccountID 账户ID
	AccountID int64 `json:"account_id,string"`
	// 审核人ID
	ManagerID int64 `json:"manager_id,string"`
	// 审核原因
	AuditResult string `json:"audit_result"`
	// 通过 or 拒绝
	Approve bool `json:"approve"`
}
