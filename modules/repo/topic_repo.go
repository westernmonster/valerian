package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
	"valerian/models"

	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"

	tracerr "github.com/ztrue/tracerr"
)

type Topic struct {
	ID               int64         `db:"id" json:"id,string"`                          // ID ID
	TopicSetID       int64         `db:"topic_set_id" json:"topic_set_id,string"`      // TopicSetID 话题集合ID
	Name             string        `db:"name" json:"name"`                             // Name 话题名
	Cover            *string       `db:"cover" json:"cover"`                           // Cover 话题封面
	Bg               *string       `db:"bg" json:"bg"`                                 // Bg 背景图
	Introduction     string        `db:"introduction" json:"introduction"`             // Introduction 话题简介
	IsPrivate        types.BitBool `db:"is_private" json:"is_private"`                 // IsPrivate 是否私密
	AllowChat        types.BitBool `db:"allow_chat" json:"allow_chat"`                 // AllowChat 开启聊天
	AllowDiscuss     types.BitBool `db:"allow_discuss" json:"allow_discuss"`           // AllowChat 开启聊天
	EditPermission   string        `db:"edit_permission" json:"edit_permission"`       // EditPermission 编辑权限
	ViewPermission   string        `db:"view_permission" json:"view_permission"`       // ViewPermission 查看权限
	JoinPermission   string        `db:"join_permission" json:"join_permission"`       // JoinPermission 加入权限
	Important        types.BitBool `db:"important" json:"important"`                   // Important 重要标记
	MuteNotification types.BitBool `db:"mute_notification" json:"mute_notification"`   // MuteNotification 消息免打扰
	CategoryViewType string        `db:"category_view_type" json:"category_view_type"` // CategoryViewType 分类视图
	TopicHome        string        `db:"topic_home" json:"topic_home"`                 // TopicHome 话题首页
	TopicType        int           `db:"topic_type" json:"topic_type"`                 // TopicType 话题类型
	VersionName      string        `db:"version_name" json:"version_name"`             // VersionName 版本名称
	CreatedBy        int64         `db:"created_by" json:"created_by,string"`          // CreatedBy 创建人
	Deleted          types.BitBool `db:"deleted" json:"deleted"`                       // Deleted 是否删除
	CreatedAt        int64         `db:"created_at" json:"created_at"`                 // CreatedAt 创建时间
	UpdatedAt        int64         `db:"updated_at" json:"updated_at"`                 // UpdatedAt 更新时间
}

type TopicRepository struct{}

// GetAll get all records
func (p *TopicRepository) GetAll(ctx context.Context, node sqalx.Node) (items []*Topic, err error) {
	items = make([]*Topic, 0)
	sqlSelect := "SELECT a.* FROM topics a WHERE a.deleted=0 ORDER BY a.id DESC "

	err = node.SelectContext(ctx, &items, sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// SearchTopics
func (p *TopicRepository) SearchTopics(ctx context.Context, node sqalx.Node, cond map[string]string) (items []*models.TopicSearchResult, err error) {
	items = make([]*models.TopicSearchResult, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["query"]; ok {
		if strings.TrimSpace(val) != "" {
			clause += ` AND (a.name LIKE ?
		OR a.version_name LIKE ?)`
			condition = append(condition, val)
			condition = append(condition, val)
		}
	}

	if val, ok := cond["id"]; ok {
		clause += " AND a.id !=?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id as topic_id, a.name, a.version_name FROM topics a WHERE a.deleted=0 %s ", clause)

	err = node.SelectContext(ctx, &items, sqlSelect, condition...)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetByID get record by ID
func (p *TopicRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *Topic, exist bool, err error) {
	item = new(Topic)
	sqlSelect := "SELECT a.* FROM topics a WHERE a.id=? AND a.deleted=0"

	if e := node.GetContext(ctx, item, sqlSelect, id); e != nil {
		if e == sql.ErrNoRows {
			item = nil
			return
		}
		err = tracerr.Wrap(e)
		return
	}

	exist = true
	return

}

// GetByCondition get record by condition
func (p *TopicRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *Topic, exist bool, err error) {
	item = new(Topic)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_set_id"]; ok {
		clause += " AND a.topic_set_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["cover"]; ok {
		clause += " AND a.cover =?"
		condition = append(condition, val)
	}
	if val, ok := cond["bg"]; ok {
		clause += " AND a.bg =?"
		condition = append(condition, val)
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =?"
		condition = append(condition, val)
	}
	if val, ok := cond["is_private"]; ok {
		clause += " AND a.is_private =?"
		condition = append(condition, val)
	}
	if val, ok := cond["allow_chat"]; ok {
		clause += " AND a.allow_chat =?"
		condition = append(condition, val)
	}
	if val, ok := cond["allow_discuss"]; ok {
		clause += " AND a.allow_discuss =?"
		condition = append(condition, val)
	}
	if val, ok := cond["edit_permission"]; ok {
		clause += " AND a.edit_permission =?"
		condition = append(condition, val)
	}
	if val, ok := cond["view_permission"]; ok {
		clause += " AND a.view_permission =?"
		condition = append(condition, val)
	}
	if val, ok := cond["join_permission"]; ok {
		clause += " AND a.join_permission =?"
		condition = append(condition, val)
	}
	if val, ok := cond["important"]; ok {
		clause += " AND a.important =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mute_notification"]; ok {
		clause += " AND a.mute_notification =?"
		condition = append(condition, val)
	}
	if val, ok := cond["category_view_type"]; ok {
		clause += " AND a.category_view_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_type"]; ok {
		clause += " AND a.topic_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_home"]; ok {
		clause += " AND a.topic_home =?"
		condition = append(condition, val)
	}
	if val, ok := cond["version_name"]; ok {
		clause += " AND a.version_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topics a WHERE a.deleted=0 %s", clause)

	if e := node.GetContext(ctx, item, sqlSelect, condition...); e != nil {
		if e == sql.ErrNoRows {
			item = nil
			return
		}
		err = tracerr.Wrap(e)
		return
	}

	exist = true
	return
}

// Insert insert a new record
func (p *TopicRepository) Insert(ctx context.Context, node sqalx.Node, item *Topic) (err error) {
	sqlInsert := "INSERT INTO topics( id,topic_set_id,name,cover,bg,introduction,is_private,allow_chat,allow_discuss,edit_permission,view_permission,join_permission,important,mute_notification,category_view_type,topic_type,topic_home,version_name,created_by,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.TopicSetID, item.Name, item.Cover, item.Bg, item.Introduction, item.IsPrivate, item.AllowChat, item.AllowDiscuss, item.EditPermission, item.ViewPermission, item.JoinPermission, item.Important, item.MuteNotification, item.CategoryViewType, item.TopicType, item.TopicHome, item.VersionName, item.CreatedBy, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *TopicRepository) Update(ctx context.Context, node sqalx.Node, item *Topic) (err error) {
	sqlUpdate := "UPDATE topics SET topic_set_id=?,name=?,cover=?,bg=?,introduction=?,is_private=?,allow_chat=?,allow_discuss=?,edit_permission=?,view_permission=?,join_permission=?,important=?,mute_notification=?,category_view_type=?,topic_type=?,topic_home=?,version_name=?,created_by=?,updated_at=? WHERE id=?"

	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlUpdate, item.TopicSetID, item.Name, item.Cover, item.Bg, item.Introduction, item.IsPrivate, item.AllowChat, item.AllowDiscuss, item.EditPermission, item.ViewPermission, item.JoinPermission, item.Important, item.MuteNotification, item.CategoryViewType, item.TopicType, item.TopicHome, item.VersionName, item.CreatedBy, item.UpdatedAt, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

func (p *TopicRepository) GetTopicVersions(ctx context.Context, node sqalx.Node, topicSetID int64) (items []*models.TopicVersion, err error) {
	items = make([]*models.TopicVersion, 0)
	sqlSelect := "SELECT a.id AS topic_set_id,b.id AS topic_id,b.version_name FROM topic_sets a LEFT JOIN topics b ON a.id=b.topic_set_id WHERE a.id=?"

	err = node.SelectContext(ctx, &items, sqlSelect, topicSetID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}
