package model

type TopicSearchResult struct {
	// 会员数据
	Data []*ESTopic `json:"data"`
	// 分页
	Page *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type ArticleSearchResult struct {
	// 会员数据
	Data []*ESArticle `json:"data"`
	// 分页
	Page *Page `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type DiscussSearchResult struct {
	// 会员数据
	Data []*ESDiscussion `json:"data"`
	// 分页
	Page *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type AllSearchParams struct {
	// 搜索关键词
	KW string `json:"kw"`
}

type AllSearchResult struct {
	Topics      []*ESTopic `json:"topics"`
	TopicsCount int        `json:"topics_count"`

	Articles      []*ESArticle `json:"articles"`
	ArticlesCount int          `json:"articles_count"`

	Accounts      []*ESAccount `json:"accounts"`
	AccountsCount int          `json:"accounts_count"`
}

type AccountSearchResult struct {
	// 会员数据
	Data []*ESAccount `json:"data"`
	// 分页
	Page *Paging `json:"paging"`
	// 调试
	Debug string `json:"debug"`
}

type Paging struct {
	// 统计数量
	Total *int `json:"total,omitempty"`
	// 是否结束
	IsEnd bool `json:"is_end"`
	// 下一页
	Next string `json:"next"`
	// 上一页
	Prev string `json:"prev"`
}
