package dao

import (
	"context"
	"fmt"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Delete logic delete a exist record
func (p *Dao) DelFavByCond(c context.Context, node sqalx.Node, targetID int64, targetType string) (err error) {
	sqlDelete := "UPDATE favs SET deleted=1 WHERE target_id=? AND target_type=?"

	if _, err = node.ExecContext(c, sqlDelete, targetID, targetType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFavByCond err(%+v), target_id(%+v) target_type(%d)", err, targetID, targetType))
		return
	}

	return
}

func (p *Dao) DelCommentByCond(c context.Context, node sqalx.Node, ownerID int64, ownerType string) (err error) {
	sqlDelete := "UPDATE comments SET deleted=1 WHERE owner_id=? and owner_type=? "

	if _, err = node.ExecContext(c, sqlDelete, ownerID, ownerType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelCommentByCond err(%+v), owner_id(%+v) owner_type", err, ownerID, ownerType))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelRevise(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE revises SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelRevises err(%+v), item(%+v)", err, id))
		return
	}

	return
}

func (p *Dao) DelArticleRevises(c context.Context, node sqalx.Node, articleID int64) (err error) {
	sqlDelete := "UPDATE revises SET deleted=1 WHERE article_id=? "

	if _, err = node.ExecContext(c, sqlDelete, articleID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticleRevises err(%+v), item(%+v)", err, articleID))
		return
	}

	return
}

func (p *Dao) DelArticle(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE articles SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticles err(%+v), item(%+v)", err, id))
		return
	}

	return
}

func (p *Dao) DelAccountFeedByCond(c context.Context, node sqalx.Node, targetID int64, targetType string) (err error) {
	sqlDelete := "UPDATE account_feeds SET deleted=1 WHERE target_id=? AND target_type=?"

	if _, err = node.ExecContext(c, sqlDelete, targetID, targetType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountFeedByCond err(%+v), target_id(%+v) target_type(%d)", err, targetID, targetType))
		return
	}

	return
}

func (p *Dao) DelFeedByCond(c context.Context, node sqalx.Node, targetID int64, targetType string) (err error) {
	sqlDelete := "UPDATE feeds SET deleted=1 WHERE target_id=? AND target_type=?"

	if _, err = node.ExecContext(c, sqlDelete, targetID, targetType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeedByCond err(%+v), target_id(%+v) target_type(%d)", err, targetID, targetType))
		return
	}

	return
}

func (p *Dao) DelTopicFeedByCond(c context.Context, node sqalx.Node, targetID int64, targetType string) (err error) {
	sqlDelete := "UPDATE topic_feeds SET deleted=1 WHERE target_id=? AND target_type=?"

	if _, err = node.ExecContext(c, sqlDelete, targetID, targetType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicFeedByCond err(%+v), target_id(%+v) target_type(%d)", err, targetID, targetType))
		return
	}

	return
}

func (p *Dao) DelMessageByCond(c context.Context, node sqalx.Node, targetID int64, targetType string) (err error) {
	sqlDelete := "UPDATE messages SET deleted=1 WHERE target_id=? AND target_type=?"

	if _, err = node.ExecContext(c, sqlDelete, targetID, targetType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelMessageByCond err(%+v), target_id(%+v) target_type(%d)", err, targetID, targetType))
		return
	}

	return
}

func (p *Dao) DelRecentPubByCond(c context.Context, node sqalx.Node, targetID int64, targetType string) (err error) {
	sqlDelete := "UPDATE recent_pubs SET deleted=1 WHERE target_id=? AND target_type=?"

	if _, err = node.ExecContext(c, sqlDelete, targetID, targetType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelRecentPubByCond err(%+v), target_id(%+v) target_type(%d)", err, targetID, targetType))
		return
	}

	return
}

func (p *Dao) DelFeedbacksByCond(c context.Context, node sqalx.Node, targetID int64, targetType string) (err error) {
	sqlDelete := "UPDATE feedbacks SET deleted=1 WHERE target_id=? AND target_type=?"

	if _, err = node.ExecContext(c, sqlDelete, targetID, targetType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeedbacksByCond err(%+v), target_id(%+v) target_type(%d)", err, targetID, targetType))
		return
	}

	return
}
