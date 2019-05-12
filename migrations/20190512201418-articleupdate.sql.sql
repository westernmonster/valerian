
-- +migrate Up
ALTER TABLE articles CHANGE COLUMN `introduction` `introduction` varchar(1000) NULL COMMENT '文章简介';
ALTER TABLE articles ADD COLUMN content text NOT NULL COMMENT '文章内容' AFTER title;
ALTER TABLE articles ADD COLUMN version_name varchar(250) NOT NULL COMMENT '版本名称' AFTER content;
ALTER TABLE articles ADD COLUMN version_lang varchar(50) NOT NULL COMMENT '版本语言' AFTER version_name;
ALTER TABLE articles ADD COLUMN article_set_id bigint(20) NOT NULL COMMENT '文章集合ID' AFTER id;

CREATE TABLE `article_sets`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `deleted` bit(1) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '文章集合';

-- +migrate Down
ALTER TABLE articles CHANGE COLUMN `introduction` `introduction` text NOT NULL COMMENT '文章简介';
ALTER TABLE articles DROP COLUMN content;
ALTER TABLE articles DROP COLUMN version_name;
ALTER TABLE articles DROP COLUMN version_lang;
ALTER TABLE articles DROP COLUMN article_set_id;
DROP TABLE article_sets;
