package model

type MessageItem struct {
	// 动作
	// 文字，表示actors发起的动作
	// 例如 张三 关注了你，关注了你就是web
	Verb string `json:"verb"`

	// 是否已读
	IsRead bool `json:"is_read"`

	// 合并的条数
	// 未读的通知会根据类型进行合并
	// 比如说 A B C 三个用户分别关注了你，
	// 那么就会合并成一条消息通知
	MergeCount int `json:"merge_count"`

	// 消息的类型
	// comment 评论
	// reply 回复了你
	// invite 邀请加入
	// apply 申请加入
	// like 赞
	// followed  关注了你
	// joined 加入了
	// apply_rejected 申请加入被拒绝
	Type string `json:"type"`

	ID int64 `json:"id,string" swaggertype:"string"`

	CreatedAt int64 `json:"created_at"`

	Content MessageContent `json:"content"`

	TargetType string `json:"target_type"`

	// 对象
	// 这是一个interface，包含比较全的对面具体信息，例如文章、话题等
	// 业务处理判断主要根据这个对象来
	Target interface{} `json:"target"`
}

type MessageContent struct {
	// 发起行为的人
	Actors []*Actor `json:"actors"`
	// 扩展
	// 比如用text字段承载：用户评论的文字缩略 用户回复你的内容缩略
	Extend MessageContentExtend `json:"extend"`

	// 对象
	// 目前给出了对象的类型 ID 文字 和链接 四个字段
	// 后面统一了跳转link格式后只关注 link 和 text 就好
	Target MessageContentTarget `json:"target"`
}

type MessageContentExtend struct {
	Text string `json:"text"`
}

type MessageContentTarget struct {
	// 对象的ID
	ID int64 `json:"id,string" swaggertype:"string"`
	// 对象的类型
	Type string `json:"type"`
	// 对象的标题内容
	Text string `json:"text"`

	// 对象的超链接
	Link string `json:"link"`
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

type NotificationResp struct {
	Items       []*MessageItem `json:"items"`
	Paging      *Paging        `json:"paging"`
	UnreadCount int            `json:"unread_count"`
}

type Paging struct {
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
