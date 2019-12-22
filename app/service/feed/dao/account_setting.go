package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by ID
func (p *Dao) GetAccountSettingByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountSetting, err error) {
	item = new(model.AccountSetting)
	sqlSelect := "SELECT a.account_id,a.activity_like,a.activity_comment,a.activity_follow_topic,a.activity_follow_member,a.notify_like,a.notify_comment,a.notify_new_fans,a.notify_new_member,a.language,a.created_at,a.updated_at FROM account_settings a WHERE a.account_id=?"

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
