package repo

import types "valerian/library/database/sqlx/types"

type OauthRole struct {
	ID        string        `db:"id" json:"id"`                 // ID Role ID
	Name      string        `db:"name" json:"name"`             // Name Role Name
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}
