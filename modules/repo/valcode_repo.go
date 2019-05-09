package repo

import (
	"database/sql"
	"fmt"
	"time"

	"valerian/library/database/sqalx"

	packr "github.com/gobuffalo/packr"
	tracerr "github.com/ztrue/tracerr"
)

type Valcode struct {
	ID        int64  `db:"id" json:"id,string"`          // ID ID
	CodeType  int    `db:"code_type" json:"code_type"`   // CodeType 类型
	Used      int    `db:"used" json:"used"`             // Used 类型, 0未使用，1使用
	Code      string `db:"code" json:"code"`             // Code 验证码
	Identity  string `db:"identity" json:"identity"`     // Identity 用户标识，可以为邮件地址和手机号
	Deleted   int    `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type ValcodeRepository struct{}

// QueryListPaged get paged records by condition
func (p *ValcodeRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*Valcode, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*Valcode, 0)

	box := packr.NewBox("./sql/valcode")
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
func (p *ValcodeRepository) GetAll(node sqalx.Node) (items []*Valcode, err error) {
	items = make([]*Valcode, 0)
	sqlSelect := packr.NewBox("./sql/valcode").String("GET_ALL.sql")

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
func (p *ValcodeRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*Valcode, err error) {
	items = make([]*Valcode, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["code_type"]; ok {
		clause += " AND a.code_type =:code_type"
		condition["code_type"] = val
	}
	if val, ok := cond["used"]; ok {
		clause += " AND a.used =:used"
		condition["used"] = val
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =:code"
		condition["code"] = val
	}
	if val, ok := cond["identity"]; ok {
		clause += " AND a.identity =:identity"
		condition["identity"] = val
	}

	box := packr.NewBox("./sql/valcode")
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
func (p *ValcodeRepository) GetByID(node sqalx.Node, id int64) (item *Valcode, exist bool, err error) {
	item = new(Valcode)
	sqlSelect := packr.NewBox("./sql/valcode").String("GET_BY_ID.sql")

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
func (p *ValcodeRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *Valcode, exist bool, err error) {
	item = new(Valcode)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["code_type"]; ok {
		clause += " AND a.code_type =:code_type"
		condition["code_type"] = val
	}
	if val, ok := cond["used"]; ok {
		clause += " AND a.used =:used"
		condition["used"] = val
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =:code"
		condition["code"] = val
	}
	if val, ok := cond["identity"]; ok {
		clause += " AND a.identity =:identity"
		condition["identity"] = val
	}

	box := packr.NewBox("./sql/valcode")
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
func (p *ValcodeRepository) Insert(node sqalx.Node, item *Valcode) (err error) {
	sqlInsert := packr.NewBox("./sql/valcode").String("INSERT.sql")

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
func (p *ValcodeRepository) Update(node sqalx.Node, item *Valcode) (err error) {
	sqlUpdate := packr.NewBox("./sql/valcode").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *ValcodeRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/valcode").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *ValcodeRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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

// HasSentRecordsInDuration determine current identity has sent records in specified duration
func (p *ValcodeRepository) HasSentRecordsInDuration(node sqalx.Node, identity string, codeType int, duration time.Duration) (has bool, err error) {
	items := make([]*Valcode, 0)
	condition := make(map[string]interface{})
	clause := ""

	clause += " AND a.identity =:identity"
	condition["identity"] = identity

	clause += " AND a.code_type =:code_type"
	condition["code_type"] = codeType

	clause += " AND a.used =:used"
	condition["used"] = 1

	clause += " AND a.created_at >:duration"
	condition["duration"] = time.Now().Add(-duration).Unix()

	box := packr.NewBox("./sql/valcode")
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

	if len(items) > 0 {
		has = true
	}
	return
}

// IsCodeCorrect determine current code's correctness
// if used return false
// if could not found in database, return false
// if found in database and isn't used, return ture
func (p *ValcodeRepository) IsCodeCorrect(node sqalx.Node, identity string, codeType int, code string) (correct bool, item *Valcode, err error) {
	items := make([]*Valcode, 0)
	condition := make(map[string]interface{})
	clause := ""

	clause += " AND a.identity =:identity"
	condition["identity"] = identity

	clause += " AND a.code_type =:code_type"
	condition["code_type"] = codeType

	clause += " AND a.code =:code"
	condition["code"] = code

	box := packr.NewBox("./sql/valcode")
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

	if len(items) == 0 {
		correct = false
		return
	}

	// used
	if items[0].Used == 1 {
		correct = false
		return
	}

	correct = true
	item = items[0]

	return
}
