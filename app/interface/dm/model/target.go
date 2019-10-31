package model

type Creator struct {
	// 用户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 用户名
	UserName string `json:"user_name"`
	// 头像
	Avatar string `json:"avatar"`

	// 自我介绍
	Introduction string `json:"introduction,omitempty"`
}

type TargetTopic struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`
	// 头像
	// 必须为URL
	Avatar string `json:"avatar"`

	// 成员数
	MemberCount int `json:"member_count"`

	// 文章数
	ArticleCount int `json:"article_count"`

	// 讨论数
	DiscussionCount int `json:"discussion_count"`

	// 简介
	Introduction string `json:"introduction"`

	Creator *Creator `json:"creator,omitempty"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type TargetReplyComment struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 内容
	Excerpt string `json:"excerpt"`

	Creator *Creator `json:"creator,omitempty"`

	// 评论资源类型 article,revise,discussion
	TargetType string `json:"target_type"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`
}
type TargetComment struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 内容
	Excerpt string `json:"excerpt"`

	// 评论目标类型 article,revise,discussion,comment
	TargetType string `json:"target_type"`

	Creator *Creator `json:"creator,omitempty"`

	// 被回复的人
	ReplyTo *Creator `json:"reply_to,omitempty"`

	// 父级评论
	ParentComment *TargetReplyComment `json:"parent_comment,omitempty"`

	// 所属资源信息 文章/补充/评论
	Owner interface{} `json:"owner"`

	// 所属资源类型 article,revise,discussion
	OwnerType string `json:"owner_type"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`
}

type MemberInfo struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64"`
	// 用户名
	UserName string `json:"user_name" format:"user_name"`
	// 性别 1为男， 2为女
	Gender int `json:"gender,omitempty"`

	// 所在地区值
	Location int64 `json:"location,string,omitempty"`

	// 所在地区名，地区是层级结构，这里将国家、州/省、市、区全部获取出来
	LocationString string `json:"location_string,omitempty"`

	// 自我介绍
	Introduction string `json:"introduction,omitempty"`

	// 头像
	Avatar string `json:"avatar"`

	// 是否身份认证
	IDCert bool `json:"id_cert"`

	// 是否工作认证
	WorkCert bool `json:"work_cert"`

	// 是否机构用户
	IsOrg bool `json:"is_org"`

	// 是否VIP
	IsVIP bool `json:"is_vip"`

	// 状态
	Stat *MemberInfoStat `json:"stat"`
}

type MemberInfoStat struct {
	// 是否关注
	IsFollow bool `json:"is_follow"`

	// 关注数
	FollowingCount int `json:"following_count"`

	// 粉丝数
	FansCount int `json:"fans_count"  db:"-"`

	// 话题数
	TopicCount int `json:"topic_count"`

	// 文章数
	ArticleCount int `json:"article_count"`

	// 讨论数
	DiscussionCount int `json:"discussion_count"`
}

type TargetArticle struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 文章标题
	Title string `json:"title"`
	// 内容
	Excerpt string `json:"excerpt"`

	ChangeDesc string `json:"change_desc"`

	// 图片
	ImageUrls []string `json:"images"`

	// 补充个数
	ReviseCount int `json:"revise_count"`

	// 喜欢数
	LikeCount int `json:"like_count"`

	// 反对数
	DislikeCount int `json:"dislike_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	Creator *Creator `json:"creator,omitempty"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type TargetRevise struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 文章标题
	Title string `json:"title,omitempty"`
	// 评论内容
	Excerpt string `json:"excerpt"`

	// 喜欢数
	LikeCount int `json:"like_count"`

	// 反对数
	DislikeCount int `json:"dislike_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	// 图片
	ImageUrls []string `json:"images"`

	Creator *Creator `json:"creator,omitempty"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type TargetDiscuss struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 标题
	Title string `json:"title,omitempty"`
	// 评论内容
	Excerpt string `json:"excerpt"`

	// 喜欢数
	LikeCount int `json:"like_count"`

	// 反对数
	DislikeCount int `json:"dislike_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	// 图片
	ImageUrls []string `json:"images"`

	Creator *Creator `json:"creator,omitempty"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type TargetTopicInviteRequest struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 邀请人
	Creator *Creator `json:"creator,omitempty"`

	// 邀请加入的话题
	Topic *TargetTopic `json:"topic,omitempty"`

	// 状态
	// 1 请求已经发送
	// 2 已经加入
	// 3 已经拒绝
	Status int `json:"status"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`
}

type TargetTopicFollowRequest struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 申请加入的话题
	Topic *TargetTopic `json:"topic,omitempty"`

	// 申请人
	Member *Creator `json:"member"`

	// 申请理由
	Reason string `json:"reason"`

	// 状态
	// 0 请求已经发送
	// 1 审批通过
	// 2 已经拒绝
	Status int32 `json:"status"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`

	AllowViewCert bool `json:"allow_view_cert"`
}
