package repo

import types "valerian/library/database/sqlx/types"

type OauthScope struct {
	ID          int64         `db:"id" json:"id,string"`            // ID ID
	Scope       string        `db:"scope" json:"scope"`             // Scope Scope
	Description string        `db:"description" json:"description"` // Description Description
	IsDefault   int           `db:"is_default" json:"is_default"`   // IsDefault 是否默认, 0 否，1 是
	Deleted     types.BitBool `db:"deleted" json:"deleted"`         // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`   // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`   // UpdatedAt 更新时间
}
