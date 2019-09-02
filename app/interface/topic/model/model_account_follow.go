package model

import "valerian/library/database/sqlx/types"

type AccountFollower struct {
	ID         int64         `db:"id" json:"id,string"`                   // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"`   // AccountID 用户ID
	FollowerID int64         `db:"follower_id" json:"follower_id,string"` // FollowerID 关注者ID
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}