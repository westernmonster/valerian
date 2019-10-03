package model

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
	FollowCount int `json:"follow_count" db:"-"`

	// 粉丝数
	FansCount int `json:"fans_count"  db:"-"`

	// 话题数
	TopicCount int `json:"topic_count"`

	// 文章数
	ArticleCount int `json:"article_count"`

	// 讨论数
	DiscussionCount int `json:"discussion_count"`
}

type FeedItem struct {
	// 来源
	Source *FeedSource `json:"source"`
	// 目标类型
	// article, topic, member, comment
	TargetType string `json:"target_type"`
	// 目标
	Target *FeedTarget `json:"target"`
}

type FeedSource struct {
	// 动作发起者
	Actor *Actor `json:"actor,omitempty"`
	// 时间
	ActionTime int64 `json:"action_time"`
	// 文字内容
	// 例如：编辑了文章，收藏了文章，关注了用户
	ActionText string `json:"action_text"`
	// 类型
	// MEMBER_CREATE_TOPIC  创建话题
	// MEMBER_CREATE_DISCUSS 发布讨论
	// MEMBER_CREATE_ARTICLE 发布文章
	// MEMBER_CREATE_REVISE  发布补充
	// MEMBER_FOLLOW_TOPIC  关注话题
	// MEMBER_FOLLOW_MEMBER 关注用户
	// MEMBER_LIKE_ARTICLE  点赞文章
	// MEMBER_LIKE_REVISE 点赞补充
	// MEMBER_LIKE_DISCUSS 点赞讨论
	// MEMBER_FAV_ARTICLE  收藏文章
	// MEMBER_FAV_TOPIC 收藏话题
	ActionType string `json:"action_type"`
}

type Actor struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 类型
	// user, org
	Type string `json:"type"`
	// 头像
	Avatar string `json:"avatar"`
	// 用户名
	Name string `json:"name"`
}

type FeedTarget struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 类型: article,comment,account,topic
	Type string `json:"type"`
	// 文章
	Article *TargetArticle `json:"article,omitempty"`
	// 补充
	Revise *TargetRevise `json:"revise,omitempty"`
	// 讨论
	Discussion *TargetDiscuss `json:"discussion,omitempty"`
	// 用户
	Member *MemberInfo `json:"member,omitempty"`
	// 话题
	Topic *TargetTopic `json:"topic,omitempty"`
	// 评论
	Comment *TargetComment `json:"comment,omitempty"`
}

type TargetComment struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 内容
	Excerpt string `json:"excerpt"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`
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
}

type TargetRevise struct {
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
}

type FeedResp struct {
	Items  []*FeedItem `json:"items"`
	Paging *Paging     `json:"paging"`
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
