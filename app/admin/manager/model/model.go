package model

import "valerian/library/database/sqlx/types"

type User struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	UserName  string        `db:"user_name" json:"user_name"`   // UserName 用户名
	NickName  string        `db:"nick_name" json:"nick_name"`   // NickName 昵称
	Email     string        `db:"email" json:"email"`           // Email 邮件地址
	Password  string        `db:"password" json:"password"`     // Password 密码hash
	Salt      string        `db:"salt" json:"salt"`             // Salt 盐
	Role      string        `db:"role" json:"role"`             // Role 角色
	State     int           `db:"state" json:"state"`           // State 状态
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type UserRole struct {
	ID        string        `db:"id" json:"id"`                 // ID 角色ID
	Name      string        `db:"name" json:"name"`             // Name 角色名
	Perms     string        `db:"perms" json:"perms"`           // Perms 权限
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}
