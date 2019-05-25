
-- +migrate Up
ALTER TABLE account_followers CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE areas CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE articles CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE country_codes CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE id_certifications CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE locales CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_access_tokens CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_authorization_codes CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_clients CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_refresh_tokens CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_roles CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_scopes CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE topic_members CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE valcodes CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
DROP TABLE `topic_categories`;
DROP TABLE `topic_followers`;

-- +migrate Down
ALTER TABLE account_followers CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE areas CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE articles CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE country_codes CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE id_certifications CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE locales CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_access_tokens CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_authorization_codes CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_clients CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_refresh_tokens CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_roles CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE oauth_scopes CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE topic_members CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE valcodes CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';

CREATE TABLE `topic_categories`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `topic_id` bigint(20) NOT NULL COMMENT '分类ID',
  `name` varchar(250) NOT NULL COMMENT '分类名',
  `parent_id` bigint(20) NOT NULL COMMENT '父级ID, 一级分类的父ID为 0',
  `created_by` bigint(20) NOT NULL COMMENT '创建人n',
  `seq` int(11) NOT NULL COMMENT '顺序',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '话题分类';

CREATE TABLE `topic_followers` (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `topic_id` bigint(20) NOT NULL COMMENT '话题ID',
  `followers_id` bigint(20) NOT NULL COMMENT '关注者ID',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
)  COMMENT='话题关注者';
