package model

import "valerian/library/database/sqlx/types"

type TopicCatalog struct {
	ID         int64         `db:"id" json:"id,string"`                    // ID ID
	Name       string        `db:"name" json:"name"`                       // Name 名称
	Seq        int32         `db:"seq" json:"seq"`                         // Seq 顺序
	Type       string        `db:"type" json:"type"`                       // Type 类型
	ParentID   int64         `db:"parent_id" json:"parent_id,string"`      // ParentID 父ID
	RefID      int64         `db:"ref_id" json:"ref_id,omitempty,string"`  // RefID 引用ID
	TopicID    int64         `db:"topic_id" json:"topic_id,string"`        // TopicID 话题ID
	IsPrimary  types.BitBool `db:"is_primary" json:"is_primary"`           // IsPrimary
	Permission string        `db:"permission" json:"permission,omitempty"` // Permission
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                 // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`           // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`           // UpdatedAt 更新时间
}

type TopicMember struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"`     // TopicID 分类ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 成员ID
	Role      string        `db:"role" json:"role"`                    // Role 成员角色
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type Topic struct {
	ID              int64         `db:"id" json:"id,string"`                        // ID ID
	Name            string        `db:"name" json:"name"`                           // Name 话题名
	Avatar          string        `db:"avatar" json:"avatar,omitempty"`             // Avatar 话题头像
	Bg              string        `db:"bg" json:"bg,omitempty"`                     // Bg 背景图
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

type TopicFollowRequest struct {
	ID            int64         `db:"id" json:"id,string"`                    // ID ID
	AccountID     int64         `db:"account_id" json:"account_id,string"`    // AccountID 用户ID
	TopicID       int64         `db:"topic_id" json:"topic_id,string"`        // TopicID 话题ID
	Status        int32         `db:"status" json:"status"`                   // Status 状态
	Deleted       types.BitBool `db:"deleted" json:"deleted"`                 // Deleted 是否删除
	CreatedAt     int64         `db:"created_at" json:"created_at"`           // CreatedAt 创建时间
	UpdatedAt     int64         `db:"updated_at" json:"updated_at"`           // UpdatedAt 更新时间
	Reason        string        `db:"reason" json:"reason"`                   // Reason 原因
	AllowViewCert types.BitBool `db:"allow_view_cert" json:"allow_view_cert"` // AllowViewCert 允许查看认证
	RejectReason  string        `db:"reject_reason" json:"reject_reason"`     // RejectReason 原因
}

type AccountTopicSetting struct {
	ID               int64         `db:"id" json:"id,string"`                        // ID ID
	AccountID        int64         `db:"account_id" json:"account_id,string"`        // AccountID 账户ID
	TopicID          int64         `db:"topic_id" json:"topic_id,string"`            // TopicID 话题ID
	Important        types.BitBool `db:"important" json:"important"`                 // Important 重要话题
	MuteNotification types.BitBool `db:"mute_notification" json:"mute_notification"` // MuteNotification 消息免打扰
	Fav              types.BitBool `db:"fav" json:"fav"`                             // Fav 是否收藏
	Deleted          types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt        int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt        int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type AuthTopic struct {
	ID         int64         `db:"id" json:"id,string"`                   // ID ID
	TopicID    int64         `db:"topic_id" json:"topic_id,string"`       // TopicID Topic ID
	ToTopicID  int64         `db:"to_topic_id" json:"to_topic_id,string"` // ToTopicID To Topic ID
	Permission string        `db:"permission" json:"permission"`          // Permission 权限
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}

type TopicInviteRequest struct {
	ID            int64         `db:"id" json:"id,string"`                           // ID ID
	AccountID     int64         `db:"account_id" json:"account_id,string"`           // AccountID 用户ID
	TopicID       int64         `db:"topic_id" json:"topic_id,string"`               // TopicID 话题ID
	Status        int32         `db:"status" json:"status"`                          // Status 状态
	Deleted       types.BitBool `db:"deleted" json:"deleted"`                        // Deleted 是否删除
	CreatedAt     int64         `db:"created_at" json:"created_at"`                  // CreatedAt 创建时间
	UpdatedAt     int64         `db:"updated_at" json:"updated_at"`                  // UpdatedAt 更新时间
	FromAccountID int64         `db:"from_account_id" json:"from_account_id,string"` // FromAccountID 邀请人
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

type TopicStat struct {
	TopicID         int64 `db:"topic_id" json:"topic_id,string"`          // TopicID 话题ID
	MemberCount     int32 `db:"member_count" json:"member_count"`         // MemberCount 成员数
	ArticleCount    int32 `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int32 `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type DiscussCategory struct {
	ID        int64         `db:"id" json:"id,string"`             // ID ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"` // TopicID 话题ID
	Seq       int32         `db:"seq" json:"seq"`                  // Seq 顺序
	Name      string        `db:"name" json:"name"`                // Name 话题名
	Deleted   types.BitBool `db:"deleted" json:"deleted"`          // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`    // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`    // UpdatedAt 更新时间
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

type Fav struct {
	ID         int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	TargetID   int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`      // TargetType 目标类型
	Deleted    types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}
