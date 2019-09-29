package model

type PublishItem struct {
	// 类型
	// article, revise, discussion
	Type string `json:"type"`

	// 文章
	Article *TargetArticle `json:"article,omitempty"`

	// 文章补充
	Revise *TargetRevise `json:"revise,omitempty"`

	// 讨论
	Discussion *TargetDiscuss `json:"discussion,omitempty"`
}

type RecentPublishResp struct {
	Items  []*PublishItem `json:"items"`
	Paging *Paging        `json:"paging"`
}
