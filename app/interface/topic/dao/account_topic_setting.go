package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getAccountTopicSettingsSQL   = "SELECT a.* FROM account_topic_settings a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.id "
	_getAccountTopicSettingSQL    = "SELECT a.* FROM account_topic_settings a WHERE a.account_id=? AND a.topic_id=? AND a.deleted=0"
	_addAccountTopicSettingSQL    = "INSERT INTO account_topic_settings( id,account_id,topic_id,important,mute_notification,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?)"
	_updateAccountTopicSettingSQL = "UPDATE account_topic_settings SET account_id=?,topic_id=?,important=?,mute_notification=?,updated_at=? WHERE id=?"
	_delAccountTopicSettingSQL    = "UPDATE account_topic_settings SET deleted=1 WHERE id=?"
)

func (p *Dao) GetAccountTopicSettings(c context.Context, node sqalx.Node, aid int64) (items []*model.AccountTopicSetting, err error) {
	items = make([]*model.AccountTopicSetting, 0)

	if err = node.SelectContext(c, &items, _getAccountTopicSettingsSQL, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountTopicSettings error(%+v), aid(%d)", err, aid))
	}

	return
}

func (p *Dao) AddAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error) {
	if _, err = node.ExecContext(c, _addAccountTopicSettingSQL, item.ID, item.AccountID, item.TopicID, item.Important, item.MuteNotification, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountTopicSetting error(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) UpdateAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error) {
	if _, err = node.ExecContext(c, _updateAccountTopicSettingSQL, item.AccountID, item.TopicID, item.Important, item.MuteNotification, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountTopicSetting error(%+v), item(%+v)", err, item))
	}
	return
}

// Delete logic delete a exist record
func (p *Dao) DelAccountTopicSetting(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delAccountTopicSettingSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountTopicSetting error(%+v), id(%d)", err, id))
	}

	return
}

func (p *Dao) GetAccountTopicSetting(c context.Context, node sqalx.Node, aid, topicID int64) (item *model.AccountTopicSetting, err error) {
	item = new(model.AccountTopicSetting)

	if err = node.GetContext(c, item, _getAccountTopicSettingSQL, aid, topicID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountTopicSetting error(%+v) account_id(%d) topic_id(%d)", err, aid, topicID))
	}

	return
}
