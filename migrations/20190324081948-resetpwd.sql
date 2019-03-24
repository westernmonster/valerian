
-- +migrate Up
CREATE TABLE `session`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `session_type` int(11) NOT NULL COMMENT '类型',
  `used` int(11) NOT NULL COMMENT '类型, 0未使用，1使用',
  `account_id` bigint(20) NOT NULL COMMENT '账户ID',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_session_type` (`session_type`) COMMENT 'index session type',
  KEY `idx_session_used` (`used`) COMMENT 'index used'
) COMMENT 'Session';


-- +migrate Down
DROP TABLE `session`;
