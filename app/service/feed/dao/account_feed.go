package dao

import (
	"context"
	"fmt"
	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetAccountFeedPaged(c context.Context, node sqalx.Node, accountID int64, limit, offset int) (items []*model.AccountFeed, err error) {
	items = make([]*model.AccountFeed, 0)

	sql := "SELECT a.* FROM account_feeds a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sql, accountID, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountFeedPaged error(%+v), account_id(%d) limit(%d) offset(%d)", err, accountID, limit, offset))
	}
	return
}

// Insert insert a new record
func (p *Dao) AddAccountFeed(c context.Context, node sqalx.Node, item *model.AccountFeed) (err error) {
	sqlInsert := "INSERT INTO account_feeds( id,account_id,action_type,action_time,action_text,target_id,target_type,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.ActionType, item.ActionTime, item.ActionText, item.TargetID, item.TargetType, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccountFeed(c context.Context, node sqalx.Node, item *model.AccountFeed) (err error) {
	sqlUpdate := "UPDATE account_feeds SET account_id=?,action_type=?,action_time=?,action_text=?,target_id=?,target_type=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.ActionType, item.ActionTime, item.ActionText, item.TargetID, item.TargetType, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelAccountFeed(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE account_feeds SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountFeeds err(%+v), item(%+v)", err, id))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelAccountFeedByCond(c context.Context, node sqalx.Node, accountID int64, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE account_feeds SET deleted=1 WHERE account_id=? AND target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, accountID, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountFeeds err(%+v), account_id(%+v), target_type(%+v), target_id(%+v)", err, accountID, targetType, targetID))
		return
	}

	return
}
