
-- +migrate Up
CREATE TABLE `articles`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `title` varchar(250) NOT NULL COMMENT '标题',
  `cover` varchar(1000) NOT NULL COMMENT '文章封面',
  `introduction` text NOT NULL COMMENT '文章简介',
  `important` int(11) NOT NULL COMMENT '重要标记',
  `created_by` bigint(20) NOT NULL COMMENT '创建人',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '文章';

-- +migrate Down
DROP TABLE articles;
