-- +migrate Up
ALTER TABLE topics CHANGE COLUMN `allow_discuss` `allow_chat` bit(1) NOT NULL COMMENT '开启群聊';
ALTER TABLE topics CHANGE COLUMN `is_private` `is_private` bit(1) NOT NULL COMMENT '是否私密';
ALTER TABLE topics CHANGE COLUMN `important` `important` bit(1) NOT NULL COMMENT '重要标记';
ALTER TABLE topics CHANGE COLUMN `mute_notification` `mute_notification` bit(1) NOT NULL COMMENT '消息免打扰';
ALTER TABLE topics CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE topics CHANGE COLUMN `category_view_type` `category_view_type` varchar(20) NOT NULL COMMENT '分类视图';
ALTER TABLE topics CHANGE COLUMN `edit_permission` `edit_permission` varchar(20) NOT NULL COMMENT '编辑权限';
ALTER TABLE topics CHANGE COLUMN `view_permission` `view_permission` varchar(20) NOT NULL COMMENT '查看权限';
ALTER TABLE topics CHANGE COLUMN `join_permission` `join_permission` varchar(20) NOT NULL COMMENT '加入权限';
ALTER TABLE topics ADD COLUMN topic_type int(11) NOT NULL COMMENT '话题类型' AFTER category_view_type;
ALTER TABLE topics ADD COLUMN topic_home varchar(20) NOT NULL COMMENT '话题首页' AFTER topic_type;
ALTER TABLE topics ADD COLUMN version_name varchar(250) NOT NULL COMMENT '版本名称' AFTER topic_home;
ALTER TABLE topics ADD COLUMN version_lang varchar(50) NOT NULL COMMENT '版本语言' AFTER version_name;
ALTER TABLE topics ADD COLUMN topic_set_id bigint(20) NOT NULL COMMENT '话题集合ID' AFTER id;
CREATE TABLE `topic_types`  (
  `id` int(1) NOT NULL COMMENT 'ID',
  `name` varchar(250) NOT NULL COMMENT '话题类型',
  `deleted` bit(1) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '话题类型';
INSERT INTO topic_types (id, name, deleted, created_at, updated_at) VALUES(1, '书籍', 0 , 1553395414, 1553395414);

ALTER TABLE topic_categories CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE topic_members CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
-- +migrate Down
ALTER TABLE topics CHANGE COLUMN `allow_chat` `allow_discuss` int(11) NOT NULL COMMENT '允许讨论';
ALTER TABLE topics CHANGE COLUMN `is_private` `is_private` int(11) NOT NULL COMMENT '是否私密';
ALTER TABLE topics CHANGE COLUMN `important` `important` int(11) NOT NULL COMMENT '重要标记';
ALTER TABLE topics CHANGE COLUMN `mute_notification` `mute_notification` int(11) NOT NULL COMMENT '消息免打扰';
ALTER TABLE topics CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE topics CHANGE COLUMN `category_view_type` `category_view_type` int(11) NOT NULL COMMENT '分类视图';
ALTER TABLE topics CHANGE COLUMN `edit_permission` `edit_permission` int(11) NOT NULL COMMENT '编辑权限';
ALTER TABLE topics CHANGE COLUMN `view_permission` `view_permission` int(11) NOT NULL COMMENT '查看权限';
ALTER TABLE topics CHANGE COLUMN `join_permission` `join_permission` int(11) NOT NULL COMMENT '加入权限';
ALTER TABLE topics DROP COLUMN topic_type;
ALTER TABLE topics DROP COLUMN topic_home;
ALTER TABLE topics DROP COLUMN version_lang;
ALTER TABLE topics DROP COLUMN version_name;
ALTER TABLE topics DROP COLUMN topic_set_id;
DROP TABLE topic_types;
ALTER TABLE topic_categories CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE topic_members CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
