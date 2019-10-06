package dao

import (
	"context"
	"fmt"
	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetFeedPaged(c context.Context, node sqalx.Node, accountID int64, limit, offset int) (items []*model.Feed, err error) {
	items = make([]*model.Feed, 0)

	sql := "SELECT a.* FROM feeds a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sql, accountID, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeedPaged error(%+v), account_id(%d) limit(%d) offset(%d)", err, accountID, limit, offset))
	}
	return
}

// GetAll get all records
func (p *Dao) GetFeeds(c context.Context, node sqalx.Node) (items []*model.Feed, err error) {
	items = make([]*model.Feed, 0)
	sqlSelect := "SELECT a.* FROM feeds a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeeds err(%+v)", err))
		return
	}
	return
}

// Insert insert a new record
func (p *Dao) AddFeed(c context.Context, node sqalx.Node, item *model.Feed) (err error) {
	sqlInsert := "INSERT INTO feeds( id,account_id,action_type,action_time,action_text,actor_id,actor_type,target_id,target_type,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.ActionType, item.ActionTime, item.ActionText, item.ActorID, item.ActorType, item.TargetID, item.TargetType, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateFeed(c context.Context, node sqalx.Node, item *model.Feed) (err error) {
	sqlUpdate := "UPDATE feeds SET account_id=?,action_type=?,action_time=?,action_text=?,actor_id=?,actor_type=?,target_id=?,target_type=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.ActionType, item.ActionTime, item.ActionText, item.ActorID, item.ActorType, item.TargetID, item.TargetType, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelFeed(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE feeds SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeeds err(%+v), item(%+v)", err, id))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelFeedByCond(c context.Context, node sqalx.Node, accountID int64, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE feeds SET deleted=1 WHERE account_id=? AND target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, accountID, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeed err(%+v), account_id(%+v), target_type(%+v), target_id(%+v)", err, accountID, targetType, targetID))
		return
	}

	return
}
