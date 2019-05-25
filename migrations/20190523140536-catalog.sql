
-- +migrate Up
CREATE TABLE `topic_catalogs`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `name` varchar(250) NOT NULL COMMENT '名称',
  `seq` int(11) NOT NULL COMMENT '顺序',
  `type` varchar(20) NOT NULL COMMENT '类型',
  `parent_id` bigint(20) NOT NULL COMMENT '父ID',
  `ref_id` bigint(20) NULL COMMENT '引用ID',
  `topic_id` bigint(20) NOT NULL COMMENT '话题ID',
  `deleted` bit(1) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '文章类目';


ALTER TABLE topics CHANGE COLUMN `category_view_type` `catalog_view_type` varchar(20) NOT NULL COMMENT '分类视图';
-- +migrate Down
ALTER TABLE topics CHANGE COLUMN `catalog_view_type` `category_view_type` varchar(20) NOT NULL COMMENT '分类视图';
DROP TABLE topic_catalogs;
