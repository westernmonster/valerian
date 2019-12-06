package model

import "valerian/library/database/sqlx/types"

type DiscussCategory struct {
	ID        int64         `db:"id" json:"id,string"`             // ID ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"` // TopicID 话题ID
	Seq       int32         `db:"seq" json:"seq"`                  // Seq 顺序
	Name      string        `db:"name" json:"name"`                // Name 话题名
	Deleted   types.BitBool `db:"deleted" json:"deleted"`          // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`    // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`    // UpdatedAt 更新时间
}

type Topic struct {
	ID              int64         `db:"id" json:"id,string"`                        // ID ID
	Name            string        `db:"name" json:"name"`                           // Name 话题名
	Avatar          string        `db:"avatar" json:"avatar"`                       // Avatar 话题头像
	Bg              string        `db:"bg" json:"bg"`                               // Bg 背景图
	Introduction    string        `db:"introduction" json:"introduction"`           // Introduction 话题简介
	AllowDiscuss    types.BitBool `db:"allow_discuss" json:"allow_discuss"`         // AllowDiscuss 允许讨论
	AllowChat       types.BitBool `db:"allow_chat" json:"allow_chat"`               // AllowChat 开启群聊
	IsPrivate       types.BitBool `db:"is_private" json:"is_private"`               // IsPrivate 是否私密
	ViewPermission  string        `db:"view_permission" json:"view_permission"`     // ViewPermission 查看权限
	EditPermission  string        `db:"edit_permission" json:"edit_permission"`     // EditPermission 编辑权限
	JoinPermission  string        `db:"join_permission" json:"join_permission"`     // JoinPermission 加入权限
	CatalogViewType string        `db:"catalog_view_type" json:"catalog_view_type"` // CatalogViewType 分类视图
	TopicHome       string        `db:"topic_home" json:"topic_home"`               // TopicHome 话题首页
	CreatedBy       int64         `db:"created_by" json:"created_by,string"`        // CreatedBy 创建人
	Deleted         types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt       int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt       int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type Article struct {
	ID             int64         `db:"id" json:"id,string"`                    // ID ID
	Title          string        `db:"title" json:"title"`                     // Title 标题
	Content        string        `db:"content" json:"content"`                 // Content 内容
	ContentText    string        `db:"content_text" json:"content_text"`       // ContentText
	DisableRevise  types.BitBool `db:"disable_revise" json:"disable_revise"`   // DisableRevise 禁止补充
	DisableComment types.BitBool `db:"disable_comment" json:"disable_comment"` // DisableComment 禁止评论
	CreatedBy      int64         `db:"created_by" json:"created_by,string"`    // CreatedBy 创建人
	Deleted        types.BitBool `db:"deleted" json:"deleted"`                 // Deleted 是否删除
	CreatedAt      int64         `db:"created_at" json:"created_at"`           // CreatedAt 创建时间
	UpdatedAt      int64         `db:"updated_at" json:"updated_at"`           // UpdatedAt 更新时间
}

type Revise struct {
	ID          int64         `db:"id" json:"id,string"`                 // ID ID
	ArticleID   int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Title       string        `db:"title" json:"title"`                  // Title 标题
	Content     string        `db:"content" json:"content"`              // Content 内容
	ContentText string        `db:"content_text" json:"content_text"`    // ContentText 内容纯文本
	CreatedBy   int64         `db:"created_by" json:"created_by,string"` // CreatedBy 创建人
	Deleted     types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
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
	IsLock       types.BitBool `db:"is_lock" json:"is_lock"`           // IsLock 是否被禁用
}

type ArticleHistory struct {
	ID          int64         `db:"id" json:"id,string"`                 // ID ID
	ArticleID   int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Content     string        `db:"content" json:"content"`              // Content 内容
	ContentText string        `db:"content_text" json:"content_text"`    // ContentText 内容
	Seq         int           `db:"seq" json:"seq"`                      // Seq 顺序
	Diff        string        `db:"diff" json:"diff"`                    // Diff 更改内容
	UpdatedBy   int64         `db:"updated_by" json:"updated_by,string"` // UpdatedBy 更新人
	ChangeDesc  string        `db:"change_desc" json:"change_desc"`      // ChangeDesc 修订说明
	Deleted     types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type Discussion struct {
	ID          int64         `db:"id" json:"id,string"`                   // ID ID
	TopicID     int64         `db:"topic_id" json:"topic_id,string"`       // TopicID 话题ID
	CategoryID  int64         `db:"category_id" json:"category_id,string"` // CategoryID 分类ID
	CreatedBy   int64         `db:"created_by" json:"created_by,string"`   // CreatedBy 创建用户ID
	Title       string        `db:"title" json:"title"`                    // Title 标题
	Content     string        `db:"content" json:"content"`                // Content 内容HTML
	ContentText string        `db:"content_text" json:"content_text"`      // ContentText 内容纯文本
	Deleted     types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}
