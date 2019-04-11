
-- +migrate Up
CREATE TABLE `topics`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `name` varchar(250) NOT NULL COMMENT '话题名',
  `cover` varchar(1000) NOT NULL COMMENT '话题封面',
  `introduction` text NOT NULL COMMENT '话题简介',
  `is_private` int(11) NOT NULL COMMENT '是否私密',
  `allow_discuss` int(11) NOT NULL COMMENT '允许讨论',
  `edit_permission` int(11) NOT NULL COMMENT '编辑权限',
  `view_permission` int(11) NOT NULL COMMENT '查看权限',
  `join_permission` int(11) NOT NULL COMMENT '加入权限',
  `important` int(11) NOT NULL COMMENT '重要标记',
  `mute_notification` int(11) NOT NULL COMMENT '消息免打扰',
  `category_view_type` int(11) NOT NULL COMMENT '分类视图',
  `created_by` bigint(20) NOT NULL COMMENT '创建人',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '话题';

CREATE TABLE `topic_categories`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `topic_id` bigint(20) NOT NULL COMMENT '分类ID',
  `name` varchar(250) NOT NULL COMMENT '分类名',
  `parent_id` bigint(20) NOT NULL COMMENT '父级ID, 一级分类的父ID为 0',
  `created_by` bigint(20) NOT NULL COMMENT '创建人n',
  `seq` int(11) NOT NULL COMMENT '顺序',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '话题分类';

CREATE TABLE `topic_members`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `topic_id` bigint(20) NOT NULL COMMENT '分类ID',
  `account_id` bigint(20) NOT NULL COMMENT '成员ID',
  `role` varchar(250) NOT NULL COMMENT '成员角色',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '话题成员';


-- +migrate Down
DROP TABLE `topics`;
DROP TABLE `topic_categories`;
DROP TABLE `topic_members`;

