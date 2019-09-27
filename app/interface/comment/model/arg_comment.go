package model

type ArgAddComment struct {
	// 回复的评论ID
	// 留空代表在资源下评论，而不是在某个评论下追加
	CommentID *int64 `json:"comment_id,string" swaggertype:"string"`
	// 类型
	// revise 补充
	// article 文章
	// discuss 话题讨论
	Type string `json:"string"`

	// 内容
	Content string `json:"string"`
	// 资源ID
	// 表示话题、文章、讨论的ID
	ResourceID int64 `json:"resource_id,string" swaggertype:"string"`
}

type ArgDelComment struct {
	CommentID int64 `json:"comment_id,string" swaggertype:"string"`
}
