package model

import (
	"container/list"
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

// RPC rpc node value.
type RPC struct {
	Proto  string `json:"Proto"`
	Addr   string `json:"Addr"`
	Group  string `json:"Group"`
	Weight int    `json:"Weight"`
}

// Version list and map.
type Version struct {
	List *list.List
	Map  map[int64]*list.Element
}

// Databus databus rule.
type Databus struct {
	Topic       string `json:"topic"`
	Group       string `json:"group"`
	Cluster     string `json:"cluster"`
	Business    string `json:"business"`
	Operation   int8   `json:"operation"`
	Leader      string `json:"leader"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	AlarmSwitch int8   `json:"alarmSwitch"`
	Users       string `json:"users"`
	AlarmRule   string `json:"alarmRule"`
}

// Databuss databuss rules.
type Databuss struct {
	Rules []*Databus `json:"rules"`
	MD5   string     `json:"md5"`
}

// Limit limit.
type Limit struct {
	Burst int     `json:"burst"`
	Rate  float64 `json:"rate"`
}

// Limits limits.
type Limits struct {
	Apps map[string]*Limit `json:"apps"`
	MD5  string            `json:"md5"`
}

// Host host.
type Host struct {
	Name  string `json:"hostname"`
	State int    `json:"state"`
}

//CodesLangs ...
type CodesLangs struct {
	Ver  int64
	MD5  string
	Code map[int]map[string]string
}

type CodeLangs struct {
	Ver  int64
	Code int
	Msg  map[string]string
}

type Code struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Code      int           `db:"code" json:"code"`             // Code 编码
	Locale    string        `db:"locale" json:"locale"`         // Locale 语言
	Message   string        `db:"message" json:"message"`       // Message 消息
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}
