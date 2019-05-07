
-- +migrate Up
CREATE TABLE `topic_relations`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `from_topic_id` bigint(20) NOT NULL COMMENT '话题ID',
  `to_topic_id` bigint(20) NOT NULL COMMENT '关联话题ID',
  `relation` varchar(20) NOT NULL COMMENT '关系',
  `deleted` bit(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '关联话题';

-- +migrate Down
DROP TABLE topic_relations;
