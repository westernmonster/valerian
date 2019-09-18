package model

import "valerian/library/database/sqlx/types"

// App app local cache.
type App struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Name      string        `db:"name" json:"name"`             // Name 名称
	Token     string        `db:"token" json:"token"`           // Token Token
	Env       string        `db:"env" json:"env"`               // Env ENV
	Zone      string        `db:"zone" json:"zone"`             // Zone Zone
	TreeID    int           `db:"tree_id" json:"tree_id"`       // TreeID 顺序
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type Config struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	AppID     int64         `db:"app_id" json:"app_id,string"`  // AppID APP ID
	Name      string        `db:"name" json:"name"`             // Name 名称
	Comment   string        `db:"comment" json:"comment"`       // Comment Comment
	From      int64         `db:"from" json:"from,string"`      // From From
	State     int           `db:"state" json:"state"`           // State State
	Mark      string        `db:"mark" json:"mark"`             // Mark Mark
	Operator  string        `db:"operator" json:"operator"`     // Operator Operator
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type Build struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	AppID     int64         `db:"app_id" json:"app_id,string"`  // AppID APP ID
	Name      string        `db:"name" json:"name"`             // Name 名称
	TagID     int64         `db:"tag_id" json:"tag_id,string"`  // TagID Tag ID
	Operator  string        `db:"operator" json:"operator"`     // Operator Operator
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type Force struct {
	ID        int64         `db:"id" json:"id,string"`           // ID ID
	AppID     int64         `db:"app_id" json:"app_id,string"`   // AppID APP ID
	HostName  string        `db:"host_name" json:"host_name"`    // HostName 名称
	IP        string        `db:"ip" json:"ip"`                  // IP IP
	Version   int64         `db:"version" json:"version,string"` // Version Version
	Operator  string        `db:"operator" json:"operator"`      // Operator Operator
	Deleted   types.BitBool `db:"deleted" json:"deleted"`        // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`  // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`  // UpdatedAt 更新时间
}

// Tag tag table in mysql.
type Tag struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	AppID     int64         `db:"app_id" json:"app_id,string"`  // AppID APP ID
	ConfigIds string        `db:"config_ids" json:"config_ids"` // ConfigIds Config IDs
	Mark      string        `db:"mark" json:"mark"`             // Mark Mark
	Force     int           `db:"force" json:"force"`           // Force Force
	Operator  string        `db:"operator" json:"operator"`     // Operator Operator
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type CommonConfig struct {
	ID        int64         `db:"id" json:"id,string"`           // ID ID
	TeamID    int64         `db:"team_id" json:"team_id,string"` // TeamID Team ID
	Name      string        `db:"name" json:"name"`              // Name 名称
	Comment   string        `db:"comment" json:"comment"`        // Comment Comment
	State     int           `db:"state" json:"state"`            // State State
	Mark      string        `db:"mark" json:"mark"`              // Mark Mark
	Operator  string        `db:"operator" json:"operator"`      // Operator Operator
	Deleted   types.BitBool `db:"deleted" json:"deleted"`        // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`  // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`  // UpdatedAt 更新时间
}
