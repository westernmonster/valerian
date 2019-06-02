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
	_getTopicSQL = "SELECT a.* FROM topics a WHERE a.id=? AND a.deleted=0"
	_addTopicSQL = "INSERT INTO topics( id,topic_set_id,name,cover,bg,introduction,is_private,allow_chat,allow_discuss,edit_permission,view_permission,join_permission,important,mute_notification,catalog_view_type,topic_type,topic_home,version_name, seq,created_by,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_delTopicSQL = "UPDATE topics SET deleted=1 WHERE id=?"
	_updateTopic = "UPDATE topics SET topic_set_id=?,name=?,cover=?,bg=?,introduction=?,is_private=?,allow_chat=?,allow_discuss=?,edit_permission=?,view_permission=?,join_permission=?,important=?,mute_notification=?,catalog_view_type=?,topic_type=?,topic_home=?,version_name=?, seq=?,created_by=?,updated_at=? WHERE id=? AND deleted=0"
)

// GetByID get record by ID
func (p *Dao) GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error) {
	item = new(model.Topic)

	if err = node.GetContext(c, item, _getTopicSQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicByID error(%+v), id(%d)", err, id))
	}

	return
}

func (p *Dao) AddTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error) {
	if _, err = node.ExecContext(c, _addTopicSQL,
		item.ID,
		item.TopicSetID,
		item.Name,
		item.Cover,
		item.Bg,
		item.Introduction,
		item.IsPrivate,
		item.AllowChat,
		item.AllowDiscuss,
		item.EditPermission,
		item.ViewPermission,
		item.JoinPermission,
		item.Important,
		item.MuteNotification,
		item.CatalogViewType,
		item.TopicType,
		item.TopicHome,
		item.VersionName,
		item.Seq,
		item.CreatedBy,
		item.Deleted,
		item.CreatedAt,
		item.UpdatedAt); err != nil {

		log.For(c).Error(fmt.Sprintf("dao.AddTopic error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) DelTopic(c context.Context, node sqalx.Node, topicID int64) (err error) {
	if _, err = node.ExecContext(c, _delTopicSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopic error(%+v), id(%d)", err, topicID))
	}
	return
}

func (p *Dao) UpdateTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error) {
	if _, err = node.ExecContext(c, _updateTopic,
		item.TopicSetID,
		item.Name,
		item.Cover,
		item.Bg,
		item.Introduction,
		item.IsPrivate,
		item.AllowChat,
		item.AllowDiscuss,
		item.EditPermission,
		item.ViewPermission,
		item.JoinPermission,
		item.Important,
		item.MuteNotification,
		item.CatalogViewType,
		item.TopicType,
		item.TopicHome,
		item.VersionName,
		item.Seq,
		item.CreatedBy,
		item.UpdatedAt,
		item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopic error(%+v), item(%+v)", err, item))
	}
	return
}
