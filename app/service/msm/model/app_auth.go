package model

import (
	xtime "valerian/library/time"
)

// apps
type App struct {
	ID        int64      `gorm:"column:id" json:"id"`
	AppTreeID int64      `gorm:"column:app_tree_id" json:"app_tree_id"`
	AppID     string     `gorm:"column:app_id" json:"app_id"`
	Limit     int64      `gorm:"column:limit" json:"limit"`
	CTime     xtime.Time `gorm:"column:ctime" json:"ctime"`
	MTime     xtime.Time `gorm:"column:mtime" json:"mtime"`
}

// app_auth
type Auth struct {
	ID            int64      `gorm:"column:id" json:"id"`
	ServiceTreeID int64      `gorm:"column:service_tree_id" json:"service_tree_id"`
	AppTreeID     int64      `gorm:"column:app_tree_id" json:"app_tree_id"`
	ServiceID     string     `gorm:"column:service_id" json:"service_id"`
	AppID         string     `gorm:"column:app_id" json:"app_id"`
	RPCMethod     string     `gorm:"column:rpc_method" json:"rpc_method"`
	HTTPMethod    string     `gorm:"column:http_method" json:"http_method"`
	Quota         int64      `gorm:"column:quota" json:"quota"`
	CTime         xtime.Time `gorm:"column:ctime" json:"ctime"`
	MTime         xtime.Time `gorm:"column:mtime" json:"mtime"`
}

type CodeMsg struct {
	ID       int64      `gorm:"column:id" json:"cid"`
	CodeID   int64      `gorm:"column:code_id" json:"code_id"`
	Locale   string     `gorm:"column:locale" json:"locale"`
	Msg      string     `gorm:"column:msg" json:"msg"`
	Operator string     `gorm:"column:operator" json:"c_operator"`
	CTime    xtime.Time `gorm:"column:ctime" json:"c_ctime"`
	MTime    xtime.Time `gorm:"column:mtime" json:"c_mtime"`
}

// Codes codes
type Codes struct {
	ID          int64      `gorm:"column:id" json:"id"`
	Code        int32      `gorm:"column:code" json:"code"`
	Message     string     `gorm:"column:message" json:"message"`
	Operator    string     `gorm:"column:operator" json:"operator"`
	HantMessage string     `gorm:"column:hant_message" json:"hant_message"`
	Level       int8       `gorm:"column:level" json:"level"`
	CTime       xtime.Time `gorm:"column:ctime" json:"ctime"`
	MTime       xtime.Time `gorm:"column:mtime" json:"mtime"`
}

// AppInfo App info.
//
type AppInfo struct {
	AppTreeID int64      `json:"app_tree_id"`
	AppID     string     `json:"app_id"`
	Limit     int32      `json:"limit"`
	MTime     xtime.Time `json:"mtime"`
}

// AppAuth AppAuth.
type AppAuth struct {
	ServiceTreeID int64      `json:"service_tree_id"`
	AppTreeID     int64      `json:"app_tree_id"`
	RPCMethod     string     `json:"rpc_method"`
	HTTPMethod    string     `json:"http_method"`
	Quota         int32      `json:"quota"`
	MTime         xtime.Time `json:"mtime"`
}

// Scope Scope.
type Scope struct {
	AppTreeID   int64    `json:"app_tree_id"`
	RPCMethods  []string `json:"rpc_methods"`
	HTTPMethods []string `json:"http_methods"`
	Quota       int32    `json:"quota"`
	Sign        string   `json:"sign"`
}

// AppToken AppToken.
type AppToken struct {
	AppTreeID int64  `json:"app_tree_id"`
	AppID     string `json:"app_id"`
	AppAuth   string `json:"app_auth"`
}
