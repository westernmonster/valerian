
-- +migrate Up
CREATE TABLE `areas`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `name` varchar(250) NOT NULL COMMENT '名称',
  `code` varchar(10) NOT NULL COMMENT '编码',
  `type` varchar(20) NOT NULL COMMENT '编码',
  `parent` bigint(20) NOT NULL COMMENT '父级ID',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '地区';

ALTER TABLE accounts ADD `location` bigint(20) NULL COMMENT '地区' AFTER birth_day;

-- +migrate Down
DROP TABLE areas;
ALTER TABLE accounts DROP  `location`;
