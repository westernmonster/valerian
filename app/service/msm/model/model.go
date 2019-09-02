package model

import (
	"valerian/library/database/sqlx/types"
)

const (
	// CodePlatDefaut db common plat code.
	CodePlatDefaut = 1
	// CodePlatDefautMsg db common plat msg.
	CodePlatDefautMsg = "common"
	// CodeDelStatus db delete status.
	CodeDelStatus = 2
	// HostOffline host offline state.
	HostOffline = 0
)

//CodesLangs ...
type CodesLangs struct {
	Ver  int64
	MD5  string
	Code map[int]map[string]string
}

type Code struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Code      int           `db:"code" json:"code"`             // Code 编码
	Message   string        `db:"message" json:"message"`       // Message 消息
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type CodeLang struct {
	Code          int    `db:"code" json:"code"`                       // Code 编码
	Message       string `db:"message" json:"message"`                 // Message 消息
	CreatedAt     int64  `db:"created_at" json:"created_at"`           // CreatedAt 创建时间
	Locale        string `db:"locale" json:"locale"`                   // Locale 语言
	LangMessage   string `db:"lang_message" json:"lang_message"`       //
	LangCreatedAt int64  `db:"lang_created_at" json:"lang_created_at"` // CreatedAt 创建时间
}
