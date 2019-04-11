
-- +migrate Up
CREATE TABLE `account_followers`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `account_id` bigint(20) NOT NULL COMMENT '用户ID',
  `followers_id` bigint(20) NOT NULL COMMENT '关注者ID',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '用户关注者';

CREATE TABLE `topic_followers`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `topic_id` bigint(20) NOT NULL COMMENT '话题ID',
  `followers_id` bigint(20) NOT NULL COMMENT '关注者ID',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '话题关注者';

-- +migrate Down
DROP TABLE account_followers;
DROP TABLE topic_followers;
