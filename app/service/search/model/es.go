package model

type ESAccount struct {
	// ID
	ID int64 `json:"id,string,omitempty" swaggertype:"string"`

	// 手机
	Mobile *string `json:"mobile,omitempty"`

	// 邮件地址
	Email *string `json:"email,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 角色
	Role string `json:"role,omitempty"`

	// 性别
	Gender *int32 `json:"gender,omitempty"`

	// 出生年
	BirthYear *int32 `json:"birth_year,omitempty"`

	// 出生月
	BirthMonth *int32 `json:"birth_month,omitempty"`

	// 出生日
	BirthDay *int32 `json:"birth_day,omitempty"`

	// 地区
	Location *int64 `json:"location,omitempty" swaggertype:"string"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`

	// 头像
	Avatar *string `json:"avatar,omitempty"`

	// 注册来源
	Source *int32 `json:"source,omitempty"`

	// 是否身份认证
	IDCert *bool `json:"id_cert,omitempty"`

	// 是否工作认证
	WorkCert *bool `json:"work_cert,omitempty"`

	// 是否机构用户
	IsOrg *bool `json:"is_org,omitempty"`

	// 是否VIP用户
	IsVIP *bool `json:"is_vip,omitempty"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,omitempty" swaggertype:"string"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"  swaggertype:"string"`

	// 关注数
	FollowingCount int32 `json:"following_count"`
	// 粉丝数
	FansCount int32 `json:"fans_count"`
	// 是否已禁用
	IsLock *bool `json:"is_lock,omitempty"`
}

type ESCreator struct {
	// ID ID
	ID int64 `json:"id,string,omitempty"  swaggertype:"string"`
	// 用户名
	UserName *string `json:"user_name,omitempty"`
	// 头像
	Avatar *string `json:"avatar,omitempty"`

	Introduction *string `json:"introduction,omitempty"`
}

type ESTopic struct {
	// ID ID
	ID int64 `json:"id,string,omitempty"  swaggertype:"string"`
	// Name 话题名
	Name *string `json:"name,omitempty" `
	// Avatar 话题头像
	Avatar *string `json:"avatar,omitempty"`
	// Bg 背景图
	Bg *string `json:"bg,omitempty"`
	// Introduction 话题简介
	Introduction *string `json:"introduction,omitempty"`
	// AllowDiscuss 允许讨论
	AllowDiscuss *bool `json:"allow_discuss,omitempty"`
	// AllowChat 开启群聊
	AllowChat *bool `json:"allow_chat,omitempty"`
	// IsPrivate 是否私密
	IsPrivate *bool `json:"is_private,omitempty"`
	// ViewPermission 查看权限
	ViewPermission *string `json:"view_permission,omitempty"`
	// EditPermission 编辑权限
	EditPermission *string `json:"edit_permission,omitempty"`
	// JoinPermission 加入权限
	JoinPermission *string `json:"join_permission,omitempty"`
	// CatalogViewType 分类视图
	CatalogViewType *string `json:"catalog_view_type,omitempty"`
	// CreatedBy 创建人
	Creator *ESCreator `json:"creator,omitempty"`
	// CreatedAt 创建时间
	CreatedAt *int64 `json:"created_at,omitempty"`
	// UpdatedAt 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// 成员数
	MemberCount int32 `json:"member_count"`

	// 成员数
	ArticleCount int32 `json:"article_count"`

	// 讨论数
	DiscussionCount int32 `json:"discussion_count"`
}

type ESArticle struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 内容
	ContentText *string `json:"content_text,omitempty"`

	// 内容
	Excerpt *string `json:"excerpt,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`

	//  禁止补充
	DisableRevise *bool `json:"disable_revise,omitempty"`

	//  禁止评论
	DisableComment *bool `json:"disable_comment,omitempty"`

	ChangeDesc *string `json:"change_desc,omitempty"`

	Creator *ESCreator `json:"creator,omitempty"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// 喜欢数
	LikeCount int32 `json:"like_count"`
	// 反对数
	DislikeCount int32 `json:"dislike_count"`
	// 补充个数
	ReviseCount int32 `json:"revise_count"`
	// 评论数
	CommentCount int32 `json:"comment_count"`
}

type ESDiscussionTopic struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// Name 话题名
	Name *string `json:"name,omitempty" `

	// 头像
	Avatar *string `json:"avatar,omitempty"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`
}

type ESDiscussion struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`

	// 内容
	ContentText *string `json:"content_text,omitempty"`

	// 内容
	Excerpt *string `json:"excerpt,omitempty"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"  swaggertype:"string"`

	Creator *ESCreator `json:"creator,omitempty"`

	Topic *ESDiscussionTopic `json:"topic,omitempty"`

	Category *ESDiscussionCategory `json:"category,omitempty"`

	// 喜欢数
	LikeCount int32 `json:"like_count"`

	// 反对数
	DislikeCount int32 `json:"dislike_count"`

	// 评论数
	CommentCount int32 `json:"comment_count"`
}

type ESDiscussionCategory struct {
	ID int64 `json:"id,string" swaggertype:"string"`

	// Name
	Name *string `json:"name,omitempty" `

	Seq *int32 `json:"seq"`
}

type ESSearchStat struct {
	ID int64 `json:"id,string" swaggertype:"string"`
	// Keywords 关键字
	Keywords string `json:"keywords"`
	// CreatedBy 创建人
	CreatedBy int64 `json:"created_by,string"`
	// Hits 命中数量
	Hits int64 `json:"hits,string"`
	// Enterpoint 搜索入口
	Enterpoint string `json:"enterpoint"`
	// CreatedAt 创建时间(搜索时间)
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt int64 `json:"updated_at"`
}
