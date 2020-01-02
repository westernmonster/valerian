package model

import "valerian/library/database/sqlx/types"

type Message struct {
	ID         int64         `db:"id" json:"id,string"`                   // ID ID
	AccountID  int64         `db:"account_id" json:"account_id,string"`   // AccountID 账户ID
	ActionType string        `db:"action_type" json:"action_type"`        // ActionType 行为类型
	ActionTime int64         `db:"action_time" json:"action_time,string"` // ActionTime 行为发起时间
	ActionText string        `db:"action_text" json:"action_text"`        // ActionText 行为文字内容
	Actors     string        `db:"actors" json:"actors"`                  // Actors 行为发起者
	Extend     string        `db:"extend" json:"extend"`                  // Extend 扩展内容
	MergeCount int           `db:"merge_count" json:"merge_count"`        // MergeCount 合并条数
	ActorType  string        `db:"actor_type" json:"actor_type"`          // ActorType 行为发起者类型
	TargetID   int64         `db:"target_id" json:"target_id,string"`     // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`        // TargetType 目标类型
	IsRead     types.BitBool `db:"is_read" json:"is_read"`                // IsRead 是否阅读
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}

type MessageStat struct {
	AccountID   int64 `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	UnreadCount int   `db:"unread_count" json:"unread_count"`    // UnreadCount 未读数
	CreatedAt   int64 `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64 `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
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

type Feedback struct {
	ID           int64         `db:"id" json:"id,string"`                 // ID ID
	TargetID     int64         `db:"target_id" json:"target_id,string"`   // TargetID 目标ID
	TargetType   int32         `db:"target_type" json:"target_type"`      // TargetType 目标类型
	TargetDesc   string        `db:"target_desc" json:"target_desc"`      // TargetDesc 目标描述（用于通知）
	FeedbackType int32         `db:"feedback_type" json:"feedback_type"`  // FeedbackType 举报类型
	FeedbackDesc string        `db:"feedback_desc" json:"feedback_desc"`  // FeedbackDesc 备注
	CreatedBy    int64         `db:"created_by" json:"created_by,string"` // CreatedBy 举报人
	Deleted      types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
	VerifyStatus int32         `db:"verify_status" json:"verify_status"`  // CheckStatus 审核状态
	VerifyDesc   string        `db:"verify_desc" json:"verify_desc"`      // CheckDesc 审核结果/原因
}
type FeedbackType struct {
	ID        int32         `db:"id" json:"id"`                 // ID ID
	Type      string        `db:"type" json:"type"`             // Type 目标类型
	Name      string        `db:"name" json:"name"`             // Name 备注
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
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

type ArticleHistory struct {
	ID          int64         `db:"id" json:"id,string"`                 // ID ID
	ArticleID   int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Content     string        `db:"content" json:"content"`              // Content 内容
	ContentText string        `db:"content_text" json:"content_text"`    // ContentText 内容
	Seq         int32         `db:"seq" json:"seq"`                      // Seq 顺序
	Diff        string        `db:"diff" json:"diff"`                    // Diff 更改内容
	UpdatedBy   int64         `db:"updated_by" json:"updated_by,string"` // UpdatedBy 更新人
	ChangeDesc  string        `db:"change_desc" json:"change_desc"`      // ChangeDesc 修订说明
	Deleted     types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
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

type Discussion struct {
	ID          int64         `db:"id" json:"id,string"`                   // ID ID
	TopicID     int64         `db:"topic_id" json:"topic_id,string"`       // TopicID 话题ID
	CategoryID  int64         `db:"category_id" json:"category_id,string"` // CategoryID 分类ID
	CreatedBy   int64         `db:"created_by" json:"created_by,string"`   // CreatedBy 创建用户ID
	Title       string        `db:"title" json:"title,omitempty"`          // Title 标题
	Content     string        `db:"content" json:"content"`                // Content 内容HTML
	ContentText string        `db:"content_text" json:"content_text"`      // ContentText 内容纯文本
	Deleted     types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
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
