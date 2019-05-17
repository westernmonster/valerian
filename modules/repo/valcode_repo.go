package repo

import (
	"context"
	"fmt"
	"time"

	"valerian/library/database/sqalx"

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

// Insert insert a new record
func (p *ValcodeRepository) Insert(ctx context.Context, node sqalx.Node, item *Valcode) (err error) {
	sqlInsert := "INSERT INTO valcodes( id,code_type,used,code,identity,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.CodeType, item.Used, item.Code, item.Identity, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *ValcodeRepository) Update(ctx context.Context, node sqalx.Node, item *Valcode) (err error) {
	sqlUpdate := "UPDATE valcodes SET code_type=?,used=?,code=?,identity=?,updated_at=? WHERE id=?"

	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlUpdate, item.CodeType, item.Used, item.Code, item.Identity, item.UpdatedAt, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// HasSentRecordsInDuration determine current identity has sent records in specified duration
func (p *ValcodeRepository) HasSentRecordsInDuration(ctx context.Context, node sqalx.Node, identity string, codeType int, duration time.Duration) (has bool, err error) {
	items := make([]*Valcode, 0)
	condition := make([]interface{}, 0)
	clause := ""

	clause += " AND a.identity =?"
	condition = append(condition, identity)

	clause += " AND a.code_type =?"
	condition = append(condition, codeType)

	clause += " AND a.used =?"
	condition = append(condition, 1)

	clause += " AND a.created_at >?"
	condition = append(condition, time.Now().Add(-duration).Unix())

	sqlSelect := fmt.Sprintf("SELECT a.* FROM valcodes a WHERE a.deleted=0 %s ORDER BY a.id DESC ", clause)

	err = node.SelectContext(ctx, &items, sqlSelect, condition...)
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
func (p *ValcodeRepository) IsCodeCorrect(ctx context.Context, node sqalx.Node, identity string, codeType int, code string) (correct bool, item *Valcode, err error) {
	items := make([]*Valcode, 0)
	condition := make([]interface{}, 0)
	clause := ""

	clause += " AND a.identity =?"
	condition = append(condition, identity)

	clause += " AND a.code_type =?"
	condition = append(condition, codeType)

	clause += " AND a.code =?"
	condition = append(condition, code)

	sqlSelect := fmt.Sprintf("SELECT a.* FROM valcodes a WHERE a.deleted=0 %s ORDER BY a.id DESC ", clause)

	err = node.SelectContext(ctx, &items, sqlSelect, condition...)
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
