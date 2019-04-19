
-- +migrate Up
CREATE TABLE `oauth_tokens`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `expired_at` bigint(20) NOT NULL COMMENT '过期时间',
  `code` varchar(512) NOT NULL COMMENT 'Authorization code ',
  `access` varchar(512) NOT NULL COMMENT 'Authorization code ',
  `refresh` varchar(512) NOT NULL COMMENT 'Authorization code ',
  `data` text NOT NULL COMMENT 'Authorization code ',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '身份认证';

-- +migrate Down
DROP TABLE oauth_tokens;
