package model

import "valerian/library/database/sqlx/types"

type AccountStat struct {
	AccountID       int64 `db:"account_id" json:"account_id,string"`      // AccountID 用户ID
	Following       int   `db:"following" json:"following"`               // Following 关注数
	Fans            int   `db:"fans" json:"fans"`                         // Fans 粉丝数
	ArticleCount    int   `db:"article_count" json:"article_count"`       // ArticleCount 文章数
	DiscussionCount int   `db:"discussion_count" json:"discussion_count"` // DiscussionCount 讨论数
	TopicCount      int   `db:"topic_count" json:"topic_count"`           // TopicCount 讨论数
	Black           int   `db:"black" json:"black"`                       // Black 黑名单数
	CreatedAt       int64 `db:"created_at" json:"created_at"`             // CreatedAt 创建时间
	UpdatedAt       int64 `db:"updated_at" json:"updated_at"`             // UpdatedAt 更新时间
}

type AccountFollowing struct {
	ID              int64         `db:"id" json:"id,string"`                               // ID ID
	AccountID       int64         `db:"account_id" json:"account_id,string"`               // AccountID 用户ID
	TargetAccountID int64         `db:"target_account_id" json:"target_account_id,string"` // TargetAccountID 被关注者ID
	Attribute       uint32        `db:"attribute" json:"attribute"`                        // Attribute 关系
	Deleted         types.BitBool `db:"deleted" json:"deleted"`                            // Deleted 是否删除
	CreatedAt       int64         `db:"created_at" json:"created_at"`                      // CreatedAt 创建时间
	UpdatedAt       int64         `db:"updated_at" json:"updated_at"`                      // UpdatedAt 更新时间
}

type AccountFans struct {
	ID              int64         `db:"id" json:"id,string"`                               // ID ID
	AccountID       int64         `db:"account_id" json:"account_id,string"`               // AccountID 用户ID
	TargetAccountID int64         `db:"target_account_id" json:"target_account_id,string"` // TargetAccountID 粉丝ID
	Attribute       uint32        `db:"attribute" json:"attribute"`                        // Attribute 关系
	Deleted         types.BitBool `db:"deleted" json:"deleted"`                            // Deleted 是否删除
	CreatedAt       int64         `db:"created_at" json:"created_at"`                      // CreatedAt 创建时间
	UpdatedAt       int64         `db:"updated_at" json:"updated_at"`                      // UpdatedAt 更新时间
}

func (st *AccountStat) Count() int {
	return int(st.Following)
}

// BlackCount get count of black, including attr black.
func (st *AccountStat) BlackCount() int {
	return int(st.Black)
}

// Empty get if the stat is empty.
func (st *AccountStat) Empty() bool {
	return st.Following == 0 && st.Black == 0 && st.Fans == 0
}
