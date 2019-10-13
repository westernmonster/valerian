package model

type Creator struct {
	// 用户ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 用户名
	UserName string `json:"user_name"`
	// 头像
	Avatar string `json:"avatar"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`
}

type TargetTopic struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`
	// 头像
	// 必须为URL
	Avatar *string `json:"avatar"`

	// 成员数
	MemberCount int `json:"member_count"`

	// 成员数
	ArticleCount int `json:"article_count"`

	// 讨论数
	DiscussionCount int `json:"discussion_count"`

	// 简介
	Introduction string `json:"introduction"`

	Creator *Creator `json:"creator,omitempty"`
}

type TargetComment struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 内容
	Excerpt string `json:"excerpt"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`
}

type MemberInfo struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string" format:"int64"`
	// 用户名
	UserName string `json:"user_name" format:"user_name"`
	// 性别 1为男， 2为女
	Gender *int `json:"gender,omitempty"`

	// 所在地区值
	Location *int64 `json:"location,string,omitempty"`

	// 所在地区名，地区是层级结构，这里将国家、州/省、市、区全部获取出来
	LocationString *string `json:"location_string,omitempty"`

	// 自我介绍
	Introduction *string `json:"introduction,omitempty"`

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
}

type TargetRevise struct {
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
}

type TargetDiscuss struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 标题
	Title *string `json:"title,omitempty"`

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
}