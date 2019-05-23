
-- +migrate Up
CREATE TABLE `drafts`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `title` varchar(250) NULL COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `category_id` int(11) NOT NULL COMMENT '分类',
  `deleted` bit(1) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '草稿';
-- +migrate Down
DROP TABLE drafts;
