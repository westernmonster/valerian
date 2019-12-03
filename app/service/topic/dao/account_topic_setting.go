package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetAccountTopicSettings(c context.Context, node sqalx.Node) (items []*model.AccountTopicSetting, err error) {
	items = make([]*model.AccountTopicSetting, 0)
	sqlSelect := "SELECT a.id,a.account_id,a.topic_id,a.important,a.fav,a.mute_notification,a.deleted,a.created_at,a.updated_at FROM account_topic_settings a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountTopicSettings err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetAccountTopicSettingsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountTopicSetting, err error) {
	items = make([]*model.AccountTopicSetting, 0)
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
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["important"]; ok {
		clause += " AND a.important =?"
		condition = append(condition, val)
	}
	if val, ok := cond["fav"]; ok {
		clause += " AND a.fav =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mute_notification"]; ok {
		clause += " AND a.mute_notification =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.topic_id,a.important,a.fav,a.mute_notification,a.deleted,a.created_at,a.updated_at FROM account_topic_settings a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountTopicSettingsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetAccountTopicSettingByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountTopicSetting, err error) {
	item = new(model.AccountTopicSetting)
	sqlSelect := "SELECT a.id,a.account_id,a.topic_id,a.important,a.fav,a.mute_notification,a.deleted,a.created_at,a.updated_at FROM account_topic_settings a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountTopicSettingByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetAccountTopicSettingByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountTopicSetting, err error) {
	item = new(model.AccountTopicSetting)
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
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["important"]; ok {
		clause += " AND a.important =?"
		condition = append(condition, val)
	}
	if val, ok := cond["fav"]; ok {
		clause += " AND a.fav =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mute_notification"]; ok {
		clause += " AND a.mute_notification =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.topic_id,a.important,a.fav,a.mute_notification,a.deleted,a.created_at,a.updated_at FROM account_topic_settings a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountTopicSettingsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error) {
	sqlInsert := "INSERT INTO account_topic_settings( id,account_id,topic_id,important,fav,mute_notification,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.TopicID, item.Important, item.Fav, item.MuteNotification, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountTopicSettings err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error) {
	sqlUpdate := "UPDATE account_topic_settings SET account_id=?,topic_id=?,important=?,fav=?,mute_notification=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.TopicID, item.Important, item.Fav, item.MuteNotification, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountTopicSettings err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelAccountTopicSetting(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE account_topic_settings SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountTopicSettings err(%+v), item(%+v)", err, id))
		return
	}

	return
}
