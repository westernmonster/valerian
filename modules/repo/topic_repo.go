package repo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
	"valerian/models"

	packr "github.com/gobuffalo/packr"
	"github.com/jmoiron/sqlx/types"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
)

type Topic struct {
	ID               int64         `db:"id" json:"id,string"`                          // ID ID
	TopicSetID       int64         `db:"topic_set_id" json:"topic_set_id,string"`      // TopicSetID 话题集合ID
	Name             string        `db:"name" json:"name"`                             // Name 话题名
	Cover            string        `db:"cover" json:"cover"`                           // Cover 话题封面
	Introduction     string        `db:"introduction" json:"introduction"`             // Introduction 话题简介
	IsPrivate        types.BitBool `db:"is_private" json:"is_private"`                 // IsPrivate 是否私密
	AllowChat        types.BitBool `db:"allow_chat" json:"allow_chat"`                 // AllowChat 开启聊天
	EditPermission   string        `db:"edit_permission" json:"edit_permission"`       // EditPermission 编辑权限
	ViewPermission   string        `db:"view_permission" json:"view_permission"`       // ViewPermission 查看权限
	JoinPermission   string        `db:"join_permission" json:"join_permission"`       // JoinPermission 加入权限
	Important        types.BitBool `db:"important" json:"important"`                   // Important 重要标记
	MuteNotification types.BitBool `db:"mute_notification" json:"mute_notification"`   // MuteNotification 消息免打扰
	CategoryViewType string        `db:"category_view_type" json:"category_view_type"` // CategoryViewType 分类视图
	TopicHome        string        `db:"topic_home" json:"topic_home"`                 // TopicHome 话题首页
	TopicType        int           `db:"topic_type" json:"topic_type"`                 // TopicType 话题类型
	VersionName      string        `db:"version_name" json:"version_name"`             // VersionName 版本名称
	VersionLanguage  string        `db:"version_lang" json:"version_lang"`             // VersionLanguage 版本语言
	CreatedBy        int64         `db:"created_by" json:"created_by,string"`          // CreatedBy 创建人
	Deleted          types.BitBool `db:"deleted" json:"deleted"`                       // Deleted 是否删除
	CreatedAt        int64         `db:"created_at" json:"created_at"`                 // CreatedAt 创建时间
	UpdatedAt        int64         `db:"updated_at" json:"updated_at"`                 // UpdatedAt 更新时间
}

type TopicRepository struct{}

// QueryListPaged get paged records by condition
func (p *TopicRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*Topic, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*Topic, 0)

	box := packr.NewBox("./sql/topic")
	sqlCount := fmt.Sprintf(box.String("QUERY_LIST_PAGED_COUNT.sql"), clause)
	sqlSelect := fmt.Sprintf(box.String("QUERY_LIST_PAGED_DATA.sql"), clause)

	stmtCount, err := node.PrepareNamed(sqlCount)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtCount.Get(&total, condition)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	condition["limit"] = pageSize
	condition["offset"] = offset

	stmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtSelect.Select(&items, condition)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetAll get all records
func (p *TopicRepository) GetAll(node sqalx.Node) (items []*Topic, err error) {
	items = make([]*Topic, 0)
	sqlSelect := packr.NewBox("./sql/topic").String("GET_ALL.sql")

	stmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtSelect.Select(&items, map[string]interface{}{})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// SearchTopics
func (p *TopicRepository) SearchTopics(node sqalx.Node, cond map[string]string) (items []*models.TopicSearchResult, err error) {
	items = make([]*models.TopicSearchResult, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["query"]; ok {
		if strings.TrimSpace(val) != "" {
			clause += ` AND (a.name LIKE :query
		OR a.version_name LIKE :query
		OR a.version_lang LIKE :query)`
			condition["query"] = "%" + val + "%"
		}
	}

	if val, ok := cond["id"]; ok {
		clause += " AND a.id !=:id"
		condition["id"] = val
	}

	box := packr.NewBox("./sql/topic")
	sqlSelect := fmt.Sprintf(box.String("SEARCH_TOPICS.sql"), clause)

	fmt.Println(sqlSelect)

	stmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtSelect.Select(&items, condition)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *TopicRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*Topic, err error) {
	items = make([]*Topic, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}
	if val, ok := cond["cover"]; ok {
		clause += " AND a.cover =:cover"
		condition["cover"] = val
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =:introduction"
		condition["introduction"] = val
	}
	if val, ok := cond["is_private"]; ok {
		clause += " AND a.is_private =:is_private"
		condition["is_private"] = val
	}
	if val, ok := cond["allow_chat"]; ok {
		clause += " AND a.allow_chat =:allow_chat"
		condition["allow_chat"] = val
	}
	if val, ok := cond["edit_permission"]; ok {
		clause += " AND a.edit_permission =:edit_permission"
		condition["edit_permission"] = val
	}
	if val, ok := cond["view_permission"]; ok {
		clause += " AND a.view_permission =:view_permission"
		condition["view_permission"] = val
	}
	if val, ok := cond["join_permission"]; ok {
		clause += " AND a.join_permission =:join_permission"
		condition["join_permission"] = val
	}
	if val, ok := cond["important"]; ok {
		clause += " AND a.important =:important"
		condition["important"] = val
	}
	if val, ok := cond["mute_notification"]; ok {
		clause += " AND a.mute_notification =:mute_notification"
		condition["mute_notification"] = val
	}
	if val, ok := cond["category_view_type"]; ok {
		clause += " AND a.category_view_type =:category_view_type"
		condition["category_view_type"] = val
	}

	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =:created_by"
		condition["created_by"] = val
	}

	box := packr.NewBox("./sql/topic")
	sqlSelect := fmt.Sprintf(box.String("GET_ALL_BY_CONDITION.sql"), clause)

	stmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtSelect.Select(&items, condition)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetByID get record by ID
func (p *TopicRepository) GetByID(node sqalx.Node, id int64) (item *Topic, exist bool, err error) {
	item = new(Topic)
	sqlSelect := packr.NewBox("./sql/topic").String("GET_BY_ID.sql")

	tmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if e := tmtSelect.Get(item, map[string]interface{}{"id": id}); e != nil {
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
func (p *TopicRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *Topic, exist bool, err error) {
	item = new(Topic)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}

	if val, ok := cond["topic_set_id"]; ok {
		clause += " AND a.topic_set_id =:topic_set_id"
		condition["topic_set_id"] = val
	}

	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}
	if val, ok := cond["cover"]; ok {
		clause += " AND a.cover =:cover"
		condition["cover"] = val
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =:introduction"
		condition["introduction"] = val
	}
	if val, ok := cond["is_private"]; ok {
		clause += " AND a.is_private =:is_private"
		condition["is_private"] = val
	}
	if val, ok := cond["allow_chat"]; ok {
		clause += " AND a.allow_chat =:allow_chat"
		condition["allow_chat"] = val
	}
	if val, ok := cond["edit_permission"]; ok {
		clause += " AND a.edit_permission =:edit_permission"
		condition["edit_permission"] = val
	}
	if val, ok := cond["view_permission"]; ok {
		clause += " AND a.view_permission =:view_permission"
		condition["view_permission"] = val
	}
	if val, ok := cond["join_permission"]; ok {
		clause += " AND a.join_permission =:join_permission"
		condition["join_permission"] = val
	}
	if val, ok := cond["important"]; ok {
		clause += " AND a.important =:important"
		condition["important"] = val
	}
	if val, ok := cond["mute_notification"]; ok {
		clause += " AND a.mute_notification =:mute_notification"
		condition["mute_notification"] = val
	}
	if val, ok := cond["category_view_type"]; ok {
		clause += " AND a.category_view_type =:category_view_type"
		condition["category_view_type"] = val
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =:created_by"
		condition["created_by"] = val
	}

	box := packr.NewBox("./sql/topic")
	sqlSelect := fmt.Sprintf(box.String("GET_BY_CONDITION.sql"), clause)

	tmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if e := tmtSelect.Get(item, condition); e != nil {
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
func (p *TopicRepository) Insert(node sqalx.Node, item *Topic) (err error) {
	sqlInsert := packr.NewBox("./sql/topic").String("INSERT.sql")

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlInsert, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *TopicRepository) Update(node sqalx.Node, item *Topic) (err error) {
	sqlUpdate := packr.NewBox("./sql/topic").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/topic").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *TopicRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
	tx, err := node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	defer tx.Rollback()
	for _, id := range ids {
		errDelete := p.Delete(tx, id)
		if errDelete != nil {
			err = tracerr.Wrap(err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
