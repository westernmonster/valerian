package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by ID
func (p *Dao) GetAccountSettingByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountSetting, err error) {
	item = new(model.AccountSetting)
	sqlSelect := "SELECT a.* FROM account_settings a WHERE a.account_id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountSettingByID err(%+v), account_id(%+v)", err, id))
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAccountSetting(c context.Context, node sqalx.Node, item *model.AccountSetting) (err error) {
	sqlInsert := "INSERT IGNORE INTO account_settings( account_id,activity_like,activity_comment,activity_follow_topic,activity_follow_member,notify_like,notify_comment,notify_new_fans,notify_new_member,language,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.AccountID, item.ActivityLike, item.ActivityComment, item.ActivityFollowTopic, item.ActivityFollowMember, item.NotifyLike, item.NotifyComment, item.NotifyNewFans, item.NotifyNewMember, item.Language, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountSettings err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccountSetting(c context.Context, node sqalx.Node, item *model.AccountSetting) (err error) {
	sqlUpdate := "UPDATE account_settings SET activity_like=?,activity_comment=?,activity_follow_topic=?,activity_follow_member=?,notify_like=?,notify_comment=?,notify_new_fans=?,notify_new_member=?,language=?,updated_at=? WHERE account_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.ActivityLike, item.ActivityComment, item.ActivityFollowTopic, item.ActivityFollowMember, item.NotifyLike, item.NotifyComment, item.NotifyNewFans, item.NotifyNewMember, item.Language, item.UpdatedAt, item.AccountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountSettings err(%+v), item(%+v)", err, item))
		return
	}

	return
}
