package repo

import (
	"database/sql"
	"fmt"
	packr "github.com/gobuffalo/packr"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
	"time"
)

type Topic struct {
	ID               int64  `db:"id" json:"id,string"`                          // ID ID
	Name             string `db:"name" json:"name"`                             // Name 话题名
	Description      string `db:"description" json:"description"`               // Description 描述
	IsPrivate        int    `db:"is_private" json:"is_private"`                 // IsPrivate 是否私密
	AllowDiscuss     int    `db:"allow_discuss" json:"allow_discuss"`           // AllowDiscuss 允许讨论
	EditPermission   int    `db:"edit_permission" json:"edit_permission"`       // EditPermission 编辑权限
	ViewPermission   int    `db:"view_permission" json:"view_permission"`       // ViewPermission 查看权限
	JoinPermission   int    `db:"join_permission" json:"join_permission"`       // JoinPermission 加入权限
	Important        int    `db:"important" json:"important"`                   // Important 重要标记
	MuteNotification int    `db:"mute_notification" json:"mute_notification"`   // MuteNotification 消息免打扰
	CategoryViewType int    `db:"category_view_type" json:"category_view_type"` // CategoryViewType 分类视图
	CreatedBy        int64  `db:"created_by" json:"created_by,string"`          // CreatedBy 创建人n
	Deleted          int    `db:"deleted" json:"deleted"`                       // Deleted 是否删除
	CreatedAt        int64  `db:"created_at" json:"created_at"`                 // CreatedAt 创建时间
	UpdatedAt        int64  `db:"updated_at" json:"updated_at"`                 // UpdatedAt 更新时间
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
	if val, ok := cond["description"]; ok {
		clause += " AND a.description =:description"
		condition["description"] = val
	}
	if val, ok := cond["is_private"]; ok {
		clause += " AND a.is_private =:is_private"
		condition["is_private"] = val
	}
	if val, ok := cond["allow_discuss"]; ok {
		clause += " AND a.allow_discuss =:allow_discuss"
		condition["allow_discuss"] = val
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
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}
	if val, ok := cond["description"]; ok {
		clause += " AND a.description =:description"
		condition["description"] = val
	}
	if val, ok := cond["is_private"]; ok {
		clause += " AND a.is_private =:is_private"
		condition["is_private"] = val
	}
	if val, ok := cond["allow_discuss"]; ok {
		clause += " AND a.allow_discuss =:allow_discuss"
		condition["allow_discuss"] = val
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
