package model

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
	// MEMBER_CREATE_ARTICLE
	// MEMBER_EDIT_ARTICLE
	// MEMBER_FOLLOW_TOPIC
	// MEMBER_FOLLOW_MEMBER
	// MEMBER_CREATE_COMMENT
	// MEMBER_CREATE_DISCUSS
	// MEMBER_UPDATE_DISCUSS
	// MEMBER_DELETE_DISCUSS
	// MEMBER_LIKE_ARTICLE
	// MEMBER_FAV_ARTICLE
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

	Introduction *string `json:"introduction"`
}

type FeedTarget struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 类型: article,comment,account,topic
	Type string `json:"type"`
	// 文章
	Article *TargetArticle `json:"article,omitempty"`
	// 评论
	Comment *TargetComment `json:"comment,omitempty"`
	// 用户
	Member *TargetMember `json:"member,omitempty"`
	// 话题
	Topic *TargetTopic `json:"topic,omitempty"`
	// 讨论
	Discussion *TargetDiscussion `json:"discussion,omitempty"`
}

type TargetArticle struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
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

type TargetDiscussion struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 标题
	Title *string `json:"title"`
	// 内容
	Excerpt string `json:"excerpt"`

	Images []string `json:"images"`
	// 喜欢数
	LikeCount int `json:"like_count"`
	// 评论数
	CommentCount int `json:"comment_count"`
}

type TargetComment struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 文章标题
	Title string `json:"title"`
	// 评论内容
	Excerpt string `json:"excerpt"`
	// 子评论数
	CommentCount int `json:"comment_count"`
}

type TargetTopic struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 话题名
	Name string `json:"name"`

	// 简介
	Introduction string `json:"introduction"`

	// 封面
	Avatar *string `json:"avatar,omitempty"`

	// 资源数量
	ResourcesCount int `json:"resources_count"`

	// 成员数
	MembersCount int `json:"members_count"`
}

type TargetMember struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 关注数
	FollowCount int `json:"follow_count"`
	// 粉丝数
	FansCount int `json:"fans_count"`
	// 话题数
	TopicCount int `json:"topic_count"`
	// 文章数
	ArticleCount int `json:"article_count"`
	// 简介
	Introduction string
	// 头像
	Avatar string `json:"avatar"`
	// 用户名
	Name string `json:"name"`

	// 身份认证
	IDCert bool `json:"id_cert"`
	// 工作认证
	WorkCert bool `json:"work_cert"`
	// 是否VIP
	IsVIP bool `json:"is_vip"`
	// 是否机构
	IsOrg bool `json:"is_org"`
}

type FeedResp struct {
	Items  []*FeedItem `json:"items"`
	Paging *FeedPaging `json:"paging"`
}

type FeedPaging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
