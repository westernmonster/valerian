package model

type FollowingResp struct {
	ID        int64  `json:"id,string"`         // ID ID
	AccountID int64  `json:"account_id,string"` // AccountID 用户ID
	Attribute uint32 `json:"attribute"`         // Attribute 关系
	CreatedAt int64  `json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64  `json:"updated_at"`        // UpdatedAt 更新时间
}

type FansResp struct {
	ID        int64  `json:"id,string"`         // ID ID
	AccountID int64  `json:"account_id,string"` // AccountID 用户ID
	Attribute uint32 `json:"attribute"`         // Attribute 关系
	CreatedAt int64  `json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64  `json:"updated_at"`        // UpdatedAt 更新时间
}
