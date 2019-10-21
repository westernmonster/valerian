package dao

import (
	"context"
	"fmt"

	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) DelRecentPubByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE recent_pubs SET deleted=1 WHERE target_type=? AND target_id=? "

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelRecentPubs err(%+v), target_type(%s) target_id(%d)", err, targetType, targetID))
		return
	}

	return
}

func (p *Dao) DelRecentViewByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE recent_views SET deleted=1 WHERE target_type=? AND target_id=? "

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelRecentViewBy err(%+v), target_type(%s) target_id(%d)", err, targetType, targetID))
		return
	}

	return
}

func (p *Dao) DelAccountFeedByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE account_feeds SET deleted=1 WHERE target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountFeeds err(%+v),  target_type(%+v), target_id(%+v)", err, targetType, targetID))
		return
	}

	return
}

func (p *Dao) DelFeedByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE feeds SET deleted=1 WHERE  target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeed err(%+v), target_type(%+v), target_id(%+v)", err, targetType, targetID))
		return
	}

	return
}

func (p *Dao) DelTopicFeedByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE topic_feeds SET deleted=1 WHERE target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicFeeds err(%+v),  target_type(%+v), target_id(%+v)", err, targetType, targetID))
		return
	}

	return
}

func (p *Dao) DelLikeByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE likes SET deleted=1 WHERE target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelLikes err(%+v),  target_type(%+v), target_id(%+v)", err, targetType, targetID))
		return
	}

	return
}

func (p *Dao) DelDislikeByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE dislikes SET deleted=1 WHERE target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelDislikes err(%+v),  target_type(%+v), target_id(%+v)", err, targetType, targetID))
		return
	}

	return
}

func (p *Dao) DelFavByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE favs SET deleted=1 WHERE target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFavs err(%+v),  target_type(%+v), target_id(%+v)", err, targetType, targetID))
		return
	}

	return
}

func (p *Dao) DelFeedbackByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE feedbacks SET deleted=1 WHERE target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeedbacks err(%+v),  target_type(%+v), target_id(%+v)", err, targetType, targetID))
		return
	}

	return
}
