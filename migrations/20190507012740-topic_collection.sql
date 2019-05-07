-- +migrate Up
CREATE TABLE `topic_sets`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `deleted` bit(1) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '话题集合';

-- +migrate Down
DROP TABLE topic_sets;
