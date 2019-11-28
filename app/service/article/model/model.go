package model

import "valerian/library/database/sqlx/types"

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

type AccountArticleAttr struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	ArticleID int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Read      types.BitBool `db:"read" json:"read"`                    // Read 是否阅读
	Like      types.BitBool `db:"like" json:"like"`                    // Like 是否点赞
	Fav       types.BitBool `db:"fav" json:"fav"`                      // Fav 是否收藏
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
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

type ArticleFile struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	FileName  string        `db:"file_name" json:"file_name"`          // FileName 文件名
	FileURL   string        `db:"file_url" json:"file_url"`            // FileURL 文件地址
	Seq       int32         `db:"seq" json:"seq"`                      // Seq 文件顺序
	ArticleID int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
	FileType  string        `db:"file_type" json:"file_type"`          // FileType 文件类型
	PdfURL    string        `db:"pdf_url" json:"pdf_url"`              // PdfURL PDF文件
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

type ImageURL struct {
	ID         int64         `db:"id" json:"id,string"`               // ID ID
	TargetID   int64         `db:"target_id" json:"target_id,string"` // TargetID 目标ID
	TargetType string        `db:"target_type" json:"target_type"`    // TargetType 目标类型
	URL        string        `db:"url" json:"url"`                    // URL 图片路径
	Deleted    types.BitBool `db:"deleted" json:"deleted"`            // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`      // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`      // UpdatedAt 更新时间
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

type ReviseFile struct {
	ID        int64         `db:"id" json:"id,string"`               // ID ID
	FileName  string        `db:"file_name" json:"file_name"`        // FileName 文件名
	FileURL   string        `db:"file_url" json:"file_url"`          // FileURL 文件地址
	Seq       int32         `db:"seq" json:"seq"`                    // Seq 文件顺序
	ReviseID  int64         `db:"revise_id" json:"revise_id,string"` // ReviseID 文章ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`            // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`      // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`      // UpdatedAt 更新时间
	FileType  string        `db:"file_type" json:"file_type"`        // FileType 文件类型
	PdfURL    string        `db:"pdf_url" json:"pdf_url"`            // PdfURL PDF文件
}

type TopicStat struct {
	TopicID         int64 `db:"topic_id" json:"topic_id,string"`          // TopicID 话题ID
	MemberCount     int32 `db:"member_count" json:"member_count"`         // MemberCount 成员数
	ArticleCount    int32 `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int32 `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
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

type TopicMember struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"`     // TopicID 分类ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 成员ID
	Role      string        `db:"role" json:"role"`                    // Role 成员角色
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
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
