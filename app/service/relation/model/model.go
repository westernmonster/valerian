package model

import "valerian/library/database/sqlx/types"

type AccountStat struct {
	AccountID       int64 `db:"account_id" json:"account_id,string"`      // AccountID 用户ID
	Following       int   `db:"following" json:"following"`               // Following 关注数
	Fans            int   `db:"fans" json:"fans"`                         // Fans 粉丝数
	ArticleCount    int   `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int   `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	TopicCount      int   `db:"topic_count" json:"topic_count"`           // TopicCount 讨论数
	Black           int   `db:"black" json:"black"`                       // Black 黑名单数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type AccountFollowing struct {
	ID              int64         `db:"id" json:"id,string"`                               // ID ID
	AccountID       int64         `db:"account_id" json:"account_id,string"`               // AccountID 用户ID
	TargetAccountID int64         `db:"target_account_id" json:"target_account_id,string"` // TargetAccountID 被关注者ID
	Attribute       uint32        `db:"attribute" json:"attribute"`                        // Attribute 关系
	Deleted         types.BitBool `db:"deleted" json:"deleted"`                            // Deleted 是否删除
	CreatedAt       int64         `db:"created_at" json:"created_at"`                      // CreatedAt 创建时间
	UpdatedAt       int64         `db:"updated_at" json:"updated_at"`                      // UpdatedAt 更新时间
}

type AccountFans struct {
	ID              int64         `db:"id" json:"id,string"`                               // ID ID
	AccountID       int64         `db:"account_id" json:"account_id,string"`               // AccountID 用户ID
	TargetAccountID int64         `db:"target_account_id" json:"target_account_id,string"` // TargetAccountID 粉丝ID
	Attribute       uint32        `db:"attribute" json:"attribute"`                        // Attribute 关系
	Deleted         types.BitBool `db:"deleted" json:"deleted"`                            // Deleted 是否删除
	CreatedAt       int64         `db:"created_at" json:"created_at"`                      // CreatedAt 创建时间
	UpdatedAt       int64         `db:"updated_at" json:"updated_at"`                      // UpdatedAt 更新时间
}

type Account struct {
	ID           int64         `db:"id" json:"id,string"`              // ID ID
	Mobile       string        `db:"mobile" json:"mobile"`             // Mobile 手机
	UserName     string        `db:"user_name" json:"user_name"`       // UserName 用户名
	Email        string        `db:"email" json:"email"`               // Email 邮件地址
	Password     string        `db:"password" json:"password"`         // Password 密码hash
	Role         string        `db:"role" json:"role"`                 // Role 角色
	Salt         string        `db:"salt" json:"salt"`                 // Salt 盐
	Gender       int32         `db:"gender" json:"gender"`             // Gender 性别
	BirthYear    int32         `db:"birth_year" json:"birth_year"`     // BirthYear 出生年
	BirthMonth   int32         `db:"birth_month" json:"birth_month"`   // BirthMonth 出生月
	BirthDay     int32         `db:"birth_day" json:"birth_day"`       // BirthDay 出生日
	Location     int64         `db:"location" json:"location,string"`  // Location 地区
	Introduction string        `db:"introduction" json:"introduction"` // Introduction 自我介绍
	Avatar       string        `db:"avatar" json:"avatar"`             // Avatar 头像
	Source       int32         `db:"source" json:"source"`             // Source 注册来源
	IP           int64         `db:"ip" json:"ip,string"`              // IP 注册IP
	IDCert       types.BitBool `db:"id_cert" json:"id_cert"`           // IDCert 是否身份认证
	WorkCert     types.BitBool `db:"work_cert" json:"work_cert"`       // WorkCert 是否工作认证
	IsOrg        types.BitBool `db:"is_org" json:"is_org"`             // IsOrg 是否机构用户
	IsVip        types.BitBool `db:"is_vip" json:"is_vip"`             // IsVip 是否VIP用户
	Deleted      types.BitBool `db:"deleted" json:"deleted"`           // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`     // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`     // UpdatedAt 更新时间
	Prefix       string        `db:"prefix" json:"prefix"`             // Prefix 手机前缀
}

func (st *AccountStat) Count() int {
	return int(st.Following)
}

// BlackCount get count of black, including attr black.
func (st *AccountStat) BlackCount() int {
	return int(st.Black)
}

// Empty get if the stat is empty.
func (st *AccountStat) Empty() bool {
	return st.Following == 0 && st.Black == 0 && st.Fans == 0
}
