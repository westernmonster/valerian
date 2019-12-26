package model

import "valerian/library/database/sqlx/types"

type Account struct {
	ID           int64         `db:"id" json:"id,string"`                // ID ID
	Mobile       string        `db:"mobile" json:"mobile"`               // Mobile 手机
	UserName     string        `db:"user_name" json:"user_name"`         // UserName 用户名
	Email        string        `db:"email" json:"email"`                 // Email 邮件地址
	Password     string        `db:"password" json:"password"`           // Password 密码hash
	Role         string        `db:"role" json:"role"`                   // Role 角色
	Salt         string        `db:"salt" json:"salt"`                   // Salt 盐
	Gender       int32         `db:"gender" json:"gender"`               // Gender 性别
	BirthYear    int32         `db:"birth_year" json:"birth_year"`       // BirthYear 出生年
	BirthMonth   int32         `db:"birth_month" json:"birth_month"`     // BirthMonth 出生月
	BirthDay     int32         `db:"birth_day" json:"birth_day"`         // BirthDay 出生日
	Location     int64         `db:"location" json:"location,string"`    // Location 地区
	Introduction string        `db:"introduction" json:"introduction"`   // Introduction 自我介绍
	Avatar       string        `db:"avatar" json:"avatar"`               // Avatar 头像
	Source       int32         `db:"source" json:"source"`               // Source 注册来源
	IP           int64         `db:"ip" json:"ip,string"`                // IP 注册IP
	IDCert       types.BitBool `db:"id_cert" json:"id_cert"`             // IDCert 是否身份认证
	WorkCert     types.BitBool `db:"work_cert" json:"work_cert"`         // WorkCert 是否工作认证
	IsOrg        types.BitBool `db:"is_org" json:"is_org"`               // IsOrg 是否机构用户
	IsVip        types.BitBool `db:"is_vip" json:"is_vip"`               // IsVip 是否VIP用户
	Deleted      types.BitBool `db:"deleted" json:"deleted"`             // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`       // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`       // UpdatedAt 更新时间
	Prefix       string        `db:"prefix" json:"prefix"`               // Prefix 手机前缀
	IsLock       types.BitBool `db:"is_lock" json:"is_lock"`             // IsLock 是否被禁用
	Deactive     types.BitBool `db:"deactive" json:"deactive,omitempty"` // Deactive
}

type AccountStat struct {
	AccountID       int64 `db:"account_id" json:"account_id,string"`      // AccountID 用户ID
	Following       int32 `db:"following" json:"following"`               // Following 关注数
	Fans            int32 `db:"fans" json:"fans"`                         // Fans 粉丝数
	ArticleCount    int32 `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int32 `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	TopicCount      int32 `db:"topic_count" json:"topic_count"`           // TopicCount 讨论数
	Black           int32 `db:"black" json:"black"`                       // Black 黑名单数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type AccountSetting struct {
	AccountID            int64         `db:"account_id" json:"account_id,string"`                  // AccountID 账户ID
	ActivityLike         types.BitBool `db:"activity_like" json:"activity_like"`                   // ActivityLike 动态-赞
	ActivityComment      types.BitBool `db:"activity_comment" json:"activity_comment"`             // ActivityComment 动态-评论
	ActivityFollowTopic  types.BitBool `db:"activity_follow_topic" json:"activity_follow_topic"`   // ActivityFollowTopic 动态-关注话题
	ActivityFollowMember types.BitBool `db:"activity_follow_member" json:"activity_follow_member"` // ActivityFollowMember 动态-关注成员
	NotifyLike           types.BitBool `db:"notify_like" json:"notify_like"`                       // NotifyLike 通知-赞
	NotifyComment        types.BitBool `db:"notify_comment" json:"notify_comment"`                 // NotifyComment 通知-评论
	NotifyNewFans        types.BitBool `db:"notify_new_fans" json:"notify_new_fans"`               // NotifyNewFans 通知-新粉丝
	NotifyNewMember      types.BitBool `db:"notify_new_member" json:"notify_new_member"`           // NotifyNewMember 通知-新成员
	Language             string        `db:"language" json:"language"`                             // Language 语言
	CreatedAt            int64         `db:"created_at" json:"created_at"`                         // CreatedAt 创建时间
	UpdatedAt            int64         `db:"updated_at" json:"updated_at"`                         // UpdatedAt 更新时间
}

type Area struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Name      string        `db:"name" json:"name"`             // Name 名称
	Code      string        `db:"code" json:"code"`             // Code 编码
	Type      string        `db:"type" json:"type"`             // Type 编码
	Parent    int64         `db:"parent" json:"parent,string"`  // Parent 父级ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type MessageStat struct {
	AccountID   int64 `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	UnreadCount int   `db:"unread_count" json:"unread_count"`    // UnreadCount 未读数
	CreatedAt   int64 `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64 `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}
