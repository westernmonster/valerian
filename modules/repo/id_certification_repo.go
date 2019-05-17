package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/library/database/sqalx"

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

// GetByID get record by ID
func (p *IDCertificationRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *IDCertification, exist bool, err error) {
	item = new(IDCertification)
	sqlSelect := "SELECT a.* FROM id_certifications a WHERE a.id=? AND a.deleted=0 "

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
func (p *IDCertificationRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *IDCertification, exist bool, err error) {
	item = new(IDCertification)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["status"]; ok {
		clause += " AND a.status =?"
		condition = append(condition, val)
	}
	if val, ok := cond["audit_conclusions"]; ok {
		clause += " AND a.audit_conclusions =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["identification_number"]; ok {
		clause += " AND a.identification_number =?"
		condition = append(condition, val)
	}
	if val, ok := cond["id_card_type"]; ok {
		clause += " AND a.id_card_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["id_card_start_date"]; ok {
		clause += " AND a.id_card_start_date =?"
		condition = append(condition, val)
	}
	if val, ok := cond["id_card_expiry"]; ok {
		clause += " AND a.id_card_expiry =?"
		condition = append(condition, val)
	}
	if val, ok := cond["address"]; ok {
		clause += " AND a.address =?"
		condition = append(condition, val)
	}
	if val, ok := cond["sex"]; ok {
		clause += " AND a.sex =?"
		condition = append(condition, val)
	}
	if val, ok := cond["id_card_front_pic"]; ok {
		clause += " AND a.id_card_front_pic =?"
		condition = append(condition, val)
	}
	if val, ok := cond["id_card_back_pic"]; ok {
		clause += " AND a.id_card_back_pic =?"
		condition = append(condition, val)
	}
	if val, ok := cond["face_pic"]; ok {
		clause += " AND a.face_pic =?"
		condition = append(condition, val)
	}
	if val, ok := cond["ethnic_group"]; ok {
		clause += " AND a.ethnic_group =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM id_certifications a WHERE a.deleted=0 %s", clause)
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
func (p *IDCertificationRepository) Insert(ctx context.Context, node sqalx.Node, item *IDCertification) (err error) {
	sqlInsert := "INSERT INTO id_certifications( id, account_id, status, audit_conclusions, name, identification_number, id_card_type, id_card_start_date, id_card_expiry, address, sex, id_card_front_pic, id_card_back_pic, face_pic, ethnic_group, deleted, created_at, updated_at) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert,
		item.ID,
		item.AccountID,
		item.Status,
		item.AuditConclusions,
		item.Name,
		item.IdentificationNumber,
		item.IDCardType,
		item.IDCardStartDate,
		item.IDCardExpiry,
		item.Address,
		item.Sex,
		item.IDCardFrontPic,
		item.IDCardBackPic,
		item.FacePic,
		item.EthnicGroup,
		item.Deleted,
		item.CreatedAt,
		item.UpdatedAt,
	)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *IDCertificationRepository) Update(ctx context.Context, node sqalx.Node, item *IDCertification) (err error) {
	sqlUpdate := "UPDATE id_certifications SET account_id=?, status=?, audit_conclusions=?, name=?, identification_number=?, id_card_type=?, id_card_start_date=?, id_card_expiry=?, address=?, sex=?, id_card_front_pic=?, id_card_back_pic=?, face_pic=?, ethnic_group=?, updated_at=? WHERE id=?"

	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlUpdate,
		item.AccountID,
		item.Status,
		item.AuditConclusions,
		item.Name,
		item.IdentificationNumber,
		item.IDCardType,
		item.IDCardStartDate,
		item.IDCardExpiry,
		item.Address,
		item.Sex,
		item.IDCardFrontPic,
		item.IDCardBackPic,
		item.FacePic,
		item.EthnicGroup,
		item.UpdatedAt,
		item.ID,
	)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
