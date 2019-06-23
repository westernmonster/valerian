package model

type TopicCatalogResp struct {
	ID int64 `json:"id,string" swaggertype:"string"`
	//  名称
	// required: true
	Name string `json:"name"`

	// 顺序
	// required: true
	Seq int `json:"seq"`

	// 类型
	// required: true
	Type string `json:"type"`

	// 引用ID
	RefID *int64 `json:"ref_id,string,omitempty" swaggertype:"string"`

	// 父级ID
	ParentID int64 `json:"parent_id,string" swaggertype:"string"`

	// 文章，当类型为文章时候返回
	Article *CatalogArticleItem `json:"article,omitempty"`

	// 题库，当类型为题库时候返回
	TestSet *CatalogTestSetItem `json:"test_set,omitempty"`
}

type CatalogArticleItem struct {
	// 文章 ID
	ArticleID int64 `json:"article_id,string,omitempty" swaggertype:"string"`

	// 文章标题
	Title string `json:"title"`

	// 赞
	LikeCount int `json:"like_count"`

	// 评论数
	CommentCount int `json:"comment_count"`

	// 主话题
	PrimaryTopicID int64 `json:"primary_topic_id,string" swaggertype:"string"`

	// 主话题名
	PrimaryTopicName string `json:"primary_topic_name"`

	// 创建时间
	CreatedAt int64 `json:"created_at"`

	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}

type CatalogTestSetItem struct {
	// 题库ID
	TestSetID int64 `json:"test_set_id,string,omitempty" swaggertype:"string"`

	// 题库名
	Name string `json:"name"`

	// 成员数
	MemberCount int `json:"member_count"`

	// 题目数
	TestCount int `json:"test_count"`
}
