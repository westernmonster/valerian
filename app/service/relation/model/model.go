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

type AccountRelation struct {
	ID          int64         `db:"id" json:"id,string"`                     // ID ID
	AccountID   int64         `db:"account_id" json:"account_id,string"`     // AccountID 用户ID
	FollowingID int64         `db:"following_id" json:"following_id,string"` // FollowingID 被关注者ID
	Attribute   uint32        `db:"Attribute" json:"attribute"`              // Attribute 关系
	Deleted     types.BitBool `db:"deleted" json:"deleted"`                  // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`            // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`            // UpdatedAt 更新时间
}

type AccountRelationStat struct {
	AccountID int64 `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	Following int   `db:"following" json:"following"`          // Following 关注数
	Fans      int   `db:"fans" json:"fans"`                    // Fans 粉丝数
	Black     int   `db:"black" json:"black"`                  // Black 黑名单数
	CreatedAt int64 `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64 `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}
