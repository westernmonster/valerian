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

type Comment struct {
	ID         int64         `db:"id" json:"id,string"`                   // ID ID
	Content    string        `db:"content" json:"content"`                // Content 内容
	TargetType string        `db:"target_type" json:"target_type"`        // TargetType 目标类型
	OwnerID    int64         `db:"owner_id" json:"owner_id,string"`       // OwnerID 资源ID (discussion, article, revise)
	ResourceID int64         `db:"resource_id" json:"resource_id,string"` // ResourceID 所属对象ID (discussion, article, revise, comment)
	Featured   types.BitBool `db:"featured" json:"featured"`              // Featured 是否精选
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	ReplyTo    int64         `db:"reply_to" json:"reply_to,string"`       // ReplyTo
	CreatedBy  int64         `db:"created_by" json:"created_by,string"`   // CreatedBy 创建人
	CreatedAt  int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
	OwnerType  string        `db:"owner_type" json:"owner_type"`          // OwnerType 所有者类型
}

type CommentStat struct {
	CommentID     int64 `db:"comment_id" json:"comment_id,string"`  // CommentID 账户ID
	LikeCount     int32 `db:"like_count" json:"like_count"`         // LikeCount 喜欢数
	DislikeCount  int32 `db:"dislike_count" json:"dislike_count"`   // DislikeCount 反对数
	ChildrenCount int32 `db:"children_count" json:"children_count"` // ChildrenCount 子评论数
	CreatedAt     int64 `db:"created_at" json:"created_at"`         // CreatedAt 创建时间
	UpdatedAt     int64 `db:"updated_at" json:"updated_at"`         // UpdatedAt 更新时间
}

type ArticleStat struct {
	ArticleID    int64 `db:"article_id" json:"article_id,string"` // ArticleID 讨论ID
	LikeCount    int32 `db:"like_count" json:"like_count"`        // LikeCount 喜欢数
	DislikeCount int32 `db:"dislike_count" json:"dislike_count"`  // DislikeCount 反对数
	ReviseCount  int32 `db:"revise_count" json:"revise_count"`    // ReviseCount 补充数
	CommentCount int32 `db:"comment_count" json:"comment_count"`  // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type ReviseStat struct {
	ReviseID     int64 `db:"revise_id" json:"revise_id,string"`  // ReviseID 补充ID
	LikeCount    int32 `db:"like_count" json:"like_count"`       // LikeCount 喜欢数
	DislikeCount int32 `db:"dislike_count" json:"dislike_count"` // DislikeCount 反对数
	CommentCount int32 `db:"comment_count" json:"comment_count"` // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`       // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`       // UpdatedAt 更新时间
}

type DiscussionStat struct {
	DiscussionID int64 `db:"discussion_id" json:"discussion_id,string"` // DiscussionID 讨论ID
	LikeCount    int32 `db:"like_count" json:"like_count"`              // LikeCount 喜欢数
	DislikeCount int32 `db:"dislike_count" json:"dislike_count"`        // DislikeCount 反对数
	CommentCount int32 `db:"comment_count" json:"comment_count"`        // CommentCount 评论数
	CreatedAt    int64 `db:"created_at" json:"created_at"`              // CreatedAt 创建时间
	UpdatedAt    int64 `db:"updated_at" json:"updated_at"`              // UpdatedAt 更新时间
}

type Like struct {
	ID         int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	TargetID   int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`      // TargetType 目标类型
	Deleted    types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type Dislike struct {
	ID         int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	TargetID   int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`      // TargetType 目标类型
	Deleted    types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}
