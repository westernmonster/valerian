package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/certification/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetIDCertifications(c context.Context, node sqalx.Node) (items []*model.IDCertification, err error) {
	items = make([]*model.IDCertification, 0)
	sqlSelect := "SELECT a.* FROM id_certifications a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetIDCertifications err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetIDCertificationsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.IDCertification, err error) {
	items = make([]*model.IDCertification, 0)
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM id_certifications a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetIDCertificationsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetIDCertificationByID(c context.Context, node sqalx.Node, id int64) (item *model.IDCertification, err error) {
	item = new(model.IDCertification)
	sqlSelect := "SELECT a.* FROM id_certifications a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetIDCertificationByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetIDCertificationByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.IDCertification, err error) {
	item = new(model.IDCertification)
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

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetIDCertificationsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddIDCertification(c context.Context, node sqalx.Node, item *model.IDCertification) (err error) {
	sqlInsert := "INSERT INTO id_certifications( id,account_id,status,audit_conclusions,name,identification_number,id_card_type,id_card_start_date,id_card_expiry,address,sex,id_card_front_pic,id_card_back_pic,face_pic,ethnic_group,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.Status, item.AuditConclusions, item.Name, item.IdentificationNumber, item.IDCardType, item.IDCardStartDate, item.IDCardExpiry, item.Address, item.Sex, item.IDCardFrontPic, item.IDCardBackPic, item.FacePic, item.EthnicGroup, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddIDCertifications err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateIDCertification(c context.Context, node sqalx.Node, item *model.IDCertification) (err error) {
	sqlUpdate := "UPDATE id_certifications SET account_id=?,status=?,audit_conclusions=?,name=?,identification_number=?,id_card_type=?,id_card_start_date=?,id_card_expiry=?,address=?,sex=?,id_card_front_pic=?,id_card_back_pic=?,face_pic=?,ethnic_group=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.Status, item.AuditConclusions, item.Name, item.IdentificationNumber, item.IDCardType, item.IDCardStartDate, item.IDCardExpiry, item.Address, item.Sex, item.IDCardFrontPic, item.IDCardBackPic, item.FacePic, item.EthnicGroup, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateIDCertifications err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelIDCertification(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE id_certifications SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelIDCertifications err(%+v), item(%+v)", err, id))
		return
	}

	return
}
