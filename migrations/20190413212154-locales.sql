
-- +migrate Up
CREATE TABLE `locales`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `locale` varchar(50) NOT NULL COMMENT '语言编码',
  `name` varchar(250) NOT NULL COMMENT '语言名称',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '语言';

INSERT INTO locales(id, locale, name, deleted, created_at, updated_at)
VALUES( 1, 'zh-CN', '简体中文', 0, 1553395414,1553395414);

INSERT INTO locales(id, locale, name, deleted, created_at, updated_at)
VALUES( 2, 'en-US', 'English', 0, 1553395414,1553395414);

-- +migrate Down
DROP TABLE locales;
