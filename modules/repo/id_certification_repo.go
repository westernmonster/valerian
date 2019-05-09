package repo

import (
	"database/sql"
	"fmt"
	"time"
	"valerian/library/database/sqalx"

	packr "github.com/gobuffalo/packr"
	tracerr "github.com/ztrue/tracerr"
)

type IDCertification struct {
	ID                   int64   `db:"id" json:"id,string"`                                          // ID ID
	AccountID            int64   `db:"account_id" json:"account_id,string"`                          // AccountID 账户ID
	Status               int     `db:"status" json:"status"`                                         // Status 状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败
	AuditConclusions     *string `db:"audit_conclusions" json:"audit_conclusions,omitempty"`         // AuditConclusions 失败原因
	Name                 *string `db:"name" json:"name,omitempty"`                                   // Name 姓名
	IdentificationNumber *string `db:"identification_number" json:"identification_number,omitempty"` // IdentificationNumber 证件号
	IDCardType           *string `db:"id_card_type" json:"id_card_type,omitempty"`                   // IDCardType 证件类型, identityCard代表身份证
	IDCardStartDate      *string `db:"id_card_start_date" json:"id_card_start_date,omitempty"`       // IDCardStartDate 证件有效期起始日期
	IDCardExpiry         *string `db:"id_card_expiry" json:"id_card_expiry,omitempty"`               // IDCardExpiry 证件有效期截止日期
	Address              *string `db:"address" json:"address,omitempty"`                             // Address 地址
	Sex                  *string `db:"sex" json:"sex,omitempty"`                                     // Sex 性别
	IDCardFrontPic       *string `db:"id_card_front_pic" json:"id_card_front_pic,omitempty"`         // IDCardFrontPic 证件照正面图片
	IDCardBackPic        *string `db:"id_card_back_pic" json:"id_card_back_pic,omitempty"`           // IDCardBackPic 证件照背面图片
	FacePic              *string `db:"face_pic" json:"face_pic,omitempty"`                           // FacePic 认证过程中拍摄的人像正面照图片
	EthnicGroup          *string `db:"ethnic_group" json:"ethnic_group,omitempty"`                   // EthnicGroup 证件上的民族
	Deleted              int     `db:"deleted" json:"deleted"`                                       // Deleted 是否删除
	CreatedAt            int64   `db:"created_at" json:"created_at"`                                 // CreatedAt 创建时间
	UpdatedAt            int64   `db:"updated_at" json:"updated_at"`                                 // UpdatedAt 更新时间
}

type IDCertificationRepository struct{}

// QueryListPaged get paged records by condition
func (p *IDCertificationRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*IDCertification, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*IDCertification, 0)

	box := packr.NewBox("./sql/id_certification")
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
func (p *IDCertificationRepository) GetAll(node sqalx.Node) (items []*IDCertification, err error) {
	items = make([]*IDCertification, 0)
	sqlSelect := packr.NewBox("./sql/id_certification").String("GET_ALL.sql")

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
func (p *IDCertificationRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*IDCertification, err error) {
	items = make([]*IDCertification, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =:account_id"
		condition["account_id"] = val
	}
	if val, ok := cond["status"]; ok {
		clause += " AND a.status =:status"
		condition["status"] = val
	}
	if val, ok := cond["audit_conclusions"]; ok {
		clause += " AND a.audit_conclusions =:audit_conclusions"
		condition["audit_conclusions"] = val
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}
	if val, ok := cond["identification_number"]; ok {
		clause += " AND a.identification_number =:identification_number"
		condition["identification_number"] = val
	}
	if val, ok := cond["id_card_type"]; ok {
		clause += " AND a.id_card_type =:id_card_type"
		condition["id_card_type"] = val
	}
	if val, ok := cond["id_card_start_date"]; ok {
		clause += " AND a.id_card_start_date =:id_card_start_date"
		condition["id_card_start_date"] = val
	}
	if val, ok := cond["id_card_expiry"]; ok {
		clause += " AND a.id_card_expiry =:id_card_expiry"
		condition["id_card_expiry"] = val
	}
	if val, ok := cond["address"]; ok {
		clause += " AND a.address =:address"
		condition["address"] = val
	}
	if val, ok := cond["sex"]; ok {
		clause += " AND a.sex =:sex"
		condition["sex"] = val
	}
	if val, ok := cond["id_card_front_pic"]; ok {
		clause += " AND a.id_card_front_pic =:id_card_front_pic"
		condition["id_card_front_pic"] = val
	}
	if val, ok := cond["id_card_back_pic"]; ok {
		clause += " AND a.id_card_back_pic =:id_card_back_pic"
		condition["id_card_back_pic"] = val
	}
	if val, ok := cond["face_pic"]; ok {
		clause += " AND a.face_pic =:face_pic"
		condition["face_pic"] = val
	}
	if val, ok := cond["ethnic_group"]; ok {
		clause += " AND a.ethnic_group =:ethnic_group"
		condition["ethnic_group"] = val
	}

	box := packr.NewBox("./sql/id_certification")
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
func (p *IDCertificationRepository) GetByID(node sqalx.Node, id int64) (item *IDCertification, exist bool, err error) {
	item = new(IDCertification)
	sqlSelect := packr.NewBox("./sql/id_certification").String("GET_BY_ID.sql")

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
func (p *IDCertificationRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *IDCertification, exist bool, err error) {
	item = new(IDCertification)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =:account_id"
		condition["account_id"] = val
	}
	if val, ok := cond["status"]; ok {
		clause += " AND a.status =:status"
		condition["status"] = val
	}
	if val, ok := cond["audit_conclusions"]; ok {
		clause += " AND a.audit_conclusions =:audit_conclusions"
		condition["audit_conclusions"] = val
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}
	if val, ok := cond["identification_number"]; ok {
		clause += " AND a.identification_number =:identification_number"
		condition["identification_number"] = val
	}
	if val, ok := cond["id_card_type"]; ok {
		clause += " AND a.id_card_type =:id_card_type"
		condition["id_card_type"] = val
	}
	if val, ok := cond["id_card_start_date"]; ok {
		clause += " AND a.id_card_start_date =:id_card_start_date"
		condition["id_card_start_date"] = val
	}
	if val, ok := cond["id_card_expiry"]; ok {
		clause += " AND a.id_card_expiry =:id_card_expiry"
		condition["id_card_expiry"] = val
	}
	if val, ok := cond["address"]; ok {
		clause += " AND a.address =:address"
		condition["address"] = val
	}
	if val, ok := cond["sex"]; ok {
		clause += " AND a.sex =:sex"
		condition["sex"] = val
	}
	if val, ok := cond["id_card_front_pic"]; ok {
		clause += " AND a.id_card_front_pic =:id_card_front_pic"
		condition["id_card_front_pic"] = val
	}
	if val, ok := cond["id_card_back_pic"]; ok {
		clause += " AND a.id_card_back_pic =:id_card_back_pic"
		condition["id_card_back_pic"] = val
	}
	if val, ok := cond["face_pic"]; ok {
		clause += " AND a.face_pic =:face_pic"
		condition["face_pic"] = val
	}
	if val, ok := cond["ethnic_group"]; ok {
		clause += " AND a.ethnic_group =:ethnic_group"
		condition["ethnic_group"] = val
	}

	box := packr.NewBox("./sql/id_certification")
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
func (p *IDCertificationRepository) Insert(node sqalx.Node, item *IDCertification) (err error) {
	sqlInsert := packr.NewBox("./sql/id_certification").String("INSERT.sql")

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
func (p *IDCertificationRepository) Update(node sqalx.Node, item *IDCertification) (err error) {
	sqlUpdate := packr.NewBox("./sql/id_certification").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *IDCertificationRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/id_certification").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *IDCertificationRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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
