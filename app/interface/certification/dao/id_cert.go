package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/certification/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_addIDCertSQL     = "INSERT INTO id_certifications( id,account_id,status,audit_conclusions,name,identification_number,id_card_type,id_card_start_date,id_card_expiry,address,sex,id_card_front_pic,id_card_back_pic,face_pic,ethnic_group,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_updateIDCertSQL  = "UPDATE id_certifications SET account_id=?,status=?,audit_conclusions=?,name=?,identification_number=?,id_card_type=?,id_card_start_date=?,id_card_expiry=?,address=?,sex=?,id_card_front_pic=?,id_card_back_pic=?,face_pic=?,ethnic_group=?,updated_at=? WHERE id=?"
	_delIDCertSQL     = "UPDATE id_certifications SET deleted=1 WHERE id=? "
	_getUserIDCertSQL = "SELECT a.* FROM id_certifications a WHERE a.account_id=? AND a.deleted=0 LIMIT 1"
)

func (p *Dao) AddIDCertification(c context.Context, node sqalx.Node, item *model.IDCertification) (err error) {
	if _, err = node.ExecContext(c, _addIDCertSQL, item.ID, item.AccountID, item.Status, item.AuditConclusions, item.Name, item.IdentificationNumber, item.IDCardType, item.IDCardStartDate, item.IDCardExpiry, item.Address, item.Sex, item.IDCardFrontPic, item.IDCardBackPic, item.FacePic, item.EthnicGroup, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddIDCertification error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) UpdateIDCertification(c context.Context, node sqalx.Node, item *model.IDCertification) (err error) {
	if _, err = node.ExecContext(c, _updateIDCertSQL, item.AccountID, item.Status, item.AuditConclusions, item.Name, item.IdentificationNumber, item.IDCardType, item.IDCardStartDate, item.IDCardExpiry, item.Address, item.Sex, item.IDCardFrontPic, item.IDCardBackPic, item.FacePic, item.EthnicGroup, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateColor error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) DelIDCertification(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delIDCertSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelColor error(%+v), id(%d)", err, id))
	}
	return
}

func (p *Dao) GetUserIDCertification(c context.Context, node sqalx.Node, aid int64) (item *model.IDCertification, err error) {
	item = new(model.IDCertification)

	if err = node.GetContext(c, item, _getUserIDCertSQL, aid); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetUserIDCertification error(%+v), aid(%d)", err, aid))
	}

	return
}
