package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getAllTopicMembersSQL   = "SELECT a.* FROM topic_members a WHERE a.topic_id=? ORDER BY a.id DESC"
	_getTopicMembersCountSQL = "SELECT COUNT(1) as count FROM topic_members a WHERE a.topic_id=?"
	_getTopicMembersPagedSQL = "SELECT a.* FROM topic_members a WHERE a.topic_id=? ORDER BY a.role,a.id DESC limit ?,?"
	_addTopicMemberSQL       = "INSERT INTO topic_members( id,topic_id,account_id,role,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"
	_updateTopicMemberSQL    = "UPDATE topic_members SET topic_id=?,account_id=?,role=?,updated_at=? WHERE id=?"
	_deleteTopicMemberSQL    = "UPDATE topic_members SET deleted=1 WHERE id=? "
)

func (p *Dao) GetAllTopicMembers(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicMember, err error) {
	items = make([]*model.TopicMember, 0)

	if err = node.SelectContext(c, &items, _getAllTopicMembersSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllTopicMembers error(%+v), topic id(%d)", err, topicID))
	}
	return
}

// GetTopicMembersCount
func (p *Dao) GetTopicMembersCount(c context.Context, node sqalx.Node, topicID int64) (count int, err error) {
	if err = node.GetContext(c, &count, _getTopicMembersCountSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersCount error(%+v), topic id(%d)", err, topicID))
	}
	return
}

// GetTopicMembersPaged
func (p *Dao) GetTopicMembersPaged(c context.Context, node sqalx.Node, topicID int64, page, pageSize int) (count int, items []*model.TopicMember, err error) {
	items = make([]*model.TopicMember, 0)
	offset := (page - 1) * pageSize

	if err = node.GetContext(c, &count, _getTopicMembersCountSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersPaged error(%+v), topic id(%d)", err, topicID))
	}

	if err = node.SelectContext(c, &items, _getTopicMembersPagedSQL, topicID, offset, pageSize); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersPaged error(%+v), topic id(%d) page(%d) pageSize(%d)", err, topicID, page, pageSize))
	}
	return
}

// Insert insert a new record
func (p *Dao) AddTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error) {
	if _, err = node.ExecContext(c, _addTopicMemberSQL, item.ID, item.TopicID, item.AccountID, item.Role, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.addTopicMember error(%+v), item(%+v)", err, item))
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error) {
	if _, err = node.ExecContext(c, _updateTopicMemberSQL, item.TopicID, item.AccountID, item.Role, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicMember error(%+v), item(%+v)", err, item))
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DeleteTopicMember(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _deleteTopicMemberSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersPaged error(%+v), topic member id(%d)", err, id))
	}

	return
}
