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

type FeedItem struct {
	// 来源
	Source *FeedSource `json:"source"`
	// 目标类型
	// article, topic, member, comment
	TargetType string `json:"target_type"`
	// 目标
	Target *FeedTarget `json:"target"`

	Deleted bool `json:"deleted"`
}

type FeedSource struct {
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

	// 自我介绍
	Introduction string `json:"introduction,omitempty"`
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
	ReviseCount int32 `json:"revise_count"`

	// 喜欢数
	LikeCount int32 `json:"like_count"`

	// 反对数
	DislikeCount int32 `json:"dislike_count"`

	// 评论数
	CommentCount int32 `json:"comment_count"`

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
	LikeCount int32 `json:"like_count"`

	// 反对数
	DislikeCount int32 `json:"dislike_count"`

	// 评论数
	CommentCount int32 `json:"comment_count"`

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
	LikeCount int32 `json:"like_count"`

	// 反对数
	DislikeCount int32 `json:"dislike_count"`

	// 评论数
	CommentCount int32 `json:"comment_count"`

	// 图片
	ImageUrls []string `json:"images"`

	Creator *Creator `json:"creator,omitempty"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type TargetComment struct {
	// ID
	ID int64 `json:"id,string" swaggertype:"string"`

	// 类型
	// revise 补充
	// article 文章
	// discussion 话题讨论
	// comment 评论
	Type string `json:"type"`

	// 内容
	Excerpt string `json:"excerpt"`

	// 发布日期
	CreatedAt int64 `json:"created_at"`

	// 资源ID
	// 表示话题、文章、讨论、评论的ID
	ResourceID int64 `json:"resource_id,string" swaggertype:"string"`

	// 对象
	// 这是一个interface，包含比较全的对面具体信息，例如文章、话题等
	// 业务处理判断主要根据这个对象来
	Target interface{} `json:"target"`

	// 喜欢数
	LikeCount int32 `json:"like_count"`

	// 子评论数
	ChildrenCount int32 `json:"children_count"`
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
	MemberCount int32 `json:"member_count"`

	// 文章数
	ArticleCount int32 `json:"article_count"`

	// 讨论数
	DiscussionCount int32 `json:"discussion_count"`

	// 简介
	Introduction string `json:"introduction"`

	Creator *Creator `json:"creator,omitempty"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`
}

type FeedResp struct {
	Items  []*FeedItem `json:"items"`
	Paging *Paging     `json:"paging"`
}
