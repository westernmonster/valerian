package model

import "valerian/library/database/sqlx/types"

type TopicCatalog struct {
	ID             int64         `db:"id" json:"id,string"`                             // ID ID
	Name           string        `db:"name" json:"name"`                                // Name 名称
	Seq            int           `db:"seq" json:"seq"`                                  // Seq 顺序
	Type           string        `db:"type" json:"type"`                                // Type 类型
	ParentID       int64         `db:"parent_id" json:"parent_id,string"`               // ParentID 父ID
	RefID          *int64        `db:"ref_id" json:"ref_id,omitempty,string"`           // RefID 引用ID
	TopicID        int64         `db:"topic_id" json:"topic_id,string"`                 // TopicID 话题ID
	TopicVersionID int64         `db:"topic_version_id" json:"topic_version_id,string"` // TopicVersionID 话题类目ID
	IsPrimary      types.BitBool `db:"is_primary" json:"is_primary"`                    // IsPrimary
	Deleted        types.BitBool `db:"deleted" json:"deleted"`                          // Deleted 是否删除
	CreatedAt      int64         `db:"created_at" json:"created_at"`                    // CreatedAt 创建时间
	UpdatedAt      int64         `db:"updated_at" json:"updated_at"`                    // UpdatedAt 更新时间
}

type TopicMember struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"`     // TopicID 分类ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 成员ID
	Role      string        `db:"role" json:"role"`                    // Role 成员角色
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type TopicRelation struct {
	ID               int64         `db:"id" json:"id,string"`                                   // ID ID
	FromTopicID      int64         `db:"from_topic_id" json:"from_topic_id,string"`             // FromTopicID 话题ID
	ToTopicVersionID int64         `db:"to_topic_version_id" json:"to_topic_version_id,string"` // ToTopicVersionID 关联话题版本ID
	ToTopicID        int64         `db:"to_topic_id" json:"to_topic_id,string"`                 // ToTopicID 关联话题ID
	Relation         string        `db:"relation" json:"relation"`                              // Relation 关系
	Seq              int           `db:"seq" json:"seq"`                                        // Seq 顺序
	Deleted          types.BitBool `db:"deleted" json:"deleted"`                                // Deleted 是否删除
	CreatedAt        int64         `db:"created_at" json:"created_at"`                          // CreatedAt 创建时间
	UpdatedAt        int64         `db:"updated_at" json:"updated_at"`                          // UpdatedAt 更新时间
}

type Topic struct {
	ID              int64         `db:"id" json:"id,string"`                        // ID ID
	Name            string        `db:"name" json:"name"`                           // Name 话题名
	Cover           *string       `db:"cover" json:"cover,omitempty"`               // Cover 话题头像
	Bg              *string       `db:"bg" json:"bg,omitempty"`                     // Bg 背景图
	Introduction    string        `db:"introduction" json:"introduction"`           // Introduction 话题简介
	AllowDiscuss    types.BitBool `db:"allow_discuss" json:"allow_discuss"`         // AllowDiscuss 允许讨论
	AllowChat       types.BitBool `db:"allow_chat" json:"allow_chat"`               // AllowChat 开启群聊
	IsPrivate       types.BitBool `db:"is_private" json:"is_private"`               // IsPrivate 是否私密
	ViewPermission  string        `db:"view_permission" json:"view_permission"`     // ViewPermission 查看权限
	EditPermission  string        `db:"edit_permission" json:"edit_permission"`     // EditPermission 编辑权限
	JoinPermission  string        `db:"join_permission" json:"join_permission"`     // JoinPermission 加入权限
	CatalogViewType string        `db:"catalog_view_type" json:"catalog_view_type"` // CatalogViewType 分类视图
	TopicHome       string        `db:"topic_home" json:"topic_home"`               // TopicHome 话题首页
	CreatedBy       int64         `db:"created_by" json:"created_by,string"`        // CreatedBy 创建人
	Deleted         types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt       int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt       int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type TopicVersion struct {
	ID        int64         `db:"id" json:"id,string"`             // ID ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"` // TopicID 话题ID
	Name      string        `db:"name" json:"name"`                // Name 版本名
	Seq       int           `db:"seq" json:"seq"`                  // Seq 顺序
	Deleted   types.BitBool `db:"deleted" json:"deleted"`          // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`    // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`    // UpdatedAt 更新时间
}

type TopicType struct {
	ID        int           `db:"id" json:"id"`                 // ID ID
	Name      string        `db:"name" json:"name"`             // Name 话题类型
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type TopicFollowRequest struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 用户ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"`     // TopicID 话题ID
	Status    int           `db:"status" json:"status"`                // Status 状态
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type AccountTopicSetting struct {
	ID               int64         `db:"id" json:"id,string"`                        // ID ID
	AccountID        int64         `db:"account_id" json:"account_id,string"`        // AccountID 账户ID
	TopicID          int64         `db:"topic_id" json:"topic_id,string"`            // TopicID 话题ID
	Important        types.BitBool `db:"important" json:"important"`                 // Important 重要话题
	MuteNotification types.BitBool `db:"mute_notification" json:"mute_notification"` // MuteNotification 消息免打扰
	Deleted          types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt        int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt        int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type AccountArticleAttr struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	ArticleID int64         `db:"article_id" json:"article_id,string"` // ArticleID 文章ID
	Read      types.BitBool `db:"read" json:"read"`                    // Read 是否阅读
	Like      types.BitBool `db:"like" json:"like"`                    // Like 是否点赞
	Fav       types.BitBool `db:"fav" json:"fav"`                      // Fav 是否收藏
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type AuthTopic struct {
	ID         int64         `db:"id" json:"id,string"`                   // ID ID
	TopicID    int64         `db:"topic_id" json:"topic_id,string"`       // TopicID Topic ID
	ToTopicID  int64         `db:"to_topic_id" json:"to_topic_id,string"` // ToTopicID To Topic ID
	Permission string        `db:"permission" json:"permission"`          // Permission 权限
	Deleted    types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt  int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt  int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}
