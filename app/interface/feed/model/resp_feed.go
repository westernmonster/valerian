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
	Actors []*Actor `json:"actors,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt  int64  `json:"updated_at"`
	ActionText string `json:"action_text"`
	ActionType string `json:"action_type"`
}

type Actor struct {
	// ID
	ID int64 `json:"id"`
	// 类型
	Type string `json:"type"`
	// 头像
	Avatar string `json:"avatar"`
	// 用户名
	Name string `json:"name"`
}

type FeedTarget struct {
	// ID
	ID int64 `json:"id" swaggertype:"string"`
	// 类型: article,comment,account,topic
	Type string `json:"type"`
	// 文章
	Article *TargetArticle `json:"article,omitempty"`
	// 补充
	Revise *TargetRevise `json:"revise,omitempty"`
	// 讨论
	Discuss *TargetDiscuss `json:"discuss,omitempty"`
	// 用户
	Member *MemberInfo `json:"member,omitempty"`
	// 话题
	Topic *TargetTopic `json:"topic,omitempty"`
}

type TargetArticle struct {
	// ID
	ID int64 `json:"id" swaggertype:"string"`
	// 文章标题
	Title string `json:"title"`
	// 封面
	Avatar *string `json:"avatar,omitempty"`
	// 内容
	Excerpt string `json:"excerpt"`
	// 喜欢数
	LikeCount int `json:"like_count"`
	// 补充个数
	ReviseCount int `json:"revise_count"`
	// 评论数
	CommentCount int `json:"comment_count"`
}

type TargetRevise struct {
	// ID
	ID int64 `json:"id" swaggertype:"string"`
	// 标题
	Title *string `json:"title,omitempty"`
	// 评论内容
	Excerpt string `json:"excerpt"`
	// 评论数
	CommentCount int `json:"comment_count"`
}

type TargetDiscuss struct {
	// ID
	ID int64 `json:"id" swaggertype:"string"`
	// 标题
	Title *string `json:"title,omitempty"`
	// 评论内容
	Excerpt string `json:"excerpt"`
	// 评论数
	CommentCount int `json:"comment_count"`

	// 图片
	ImageUrls []string `json:"img_urls"`
}

type TargetTopic struct {
	// ID
	ID int64 `json:"id" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`
	// 头像
	// 必须为URL
	Avatar *string `json:"avatar"`
	// 成员数
	MembersCount int `json:"members_count"`

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