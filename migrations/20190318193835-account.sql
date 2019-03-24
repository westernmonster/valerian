
-- +migrate Up
CREATE TABLE `accounts`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `mobile` varchar(20) NOT NULL COMMENT '手机',
  `email` varchar(320) NOT NULL COMMENT '邮件地址',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `gender` int(11) NULL COMMENT '性别',
  `birth_year` int(11) NULL COMMENT '出生年',
  `birth_month` int(11) NULL COMMENT '出生月',
  `birth_day` int(11) NULL COMMENT '出生日',
  `introduction` varchar(500) NULL COMMENT '自我介绍',
  `avatar` varchar(200) NOT NULL COMMENT '头像',
  `source` int(11) NOT NULL COMMENT '注册来源',
  `ip` bigint(20) NOT NULL COMMENT '注册IP',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_mobile` (`mobile`) COMMENT 'mobile',
  KEY `idx_email` (`email`) COMMENT 'email'
) COMMENT '账户';


-- +migrate Down
DROP TABLE accounts;
