
-- +migrate Up
CREATE TABLE `valcodes`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `code_type` int(11) NOT NULL COMMENT '类型',
  `used` int(11) NOT NULL COMMENT '类型, 0未使用，1使用',
  `code` varchar(6) NOT NULL COMMENT '验证码',
  `identity` varchar(320) NOT NULL COMMENT '用户标识，可以为邮件地址和手机号',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_identity` (`identity`) COMMENT 'index identity',
  KEY `idx_code` (`code`) COMMENT 'index code',
  KEY `idx_code_type` (`code_type`) COMMENT 'index code type',
  KEY `idx_used` (`used`) COMMENT 'index used',
  KEY `idx_created_at` (`created_at`) COMMENT 'index created_at',
  KEY `idx_deleted` (`deleted`) COMMENT 'index deleted'
) COMMENT '验证码';


-- +migrate Down
DROP TABLE valcodes;
