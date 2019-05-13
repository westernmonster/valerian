
-- +migrate Up
ALTER TABLE topics ADD COLUMN allow_discuss bit(1) NULL COMMENT '允许讨论' AFTER allow_chat;
UPDATE topics set allow_discuss=1;
ALTER TABLE topics CHANGE COLUMN `allow_discuss` `allow_discuss` bit(1) NOT NULL COMMENT '允许讨论';
ALTER TABLE topics ADD COLUMN bg varchar(1000) NULL COMMENT '背景图' AFTER cover;
ALTER TABLE topics CHANGE COLUMN `cover` `cover` varchar(1000)  NULL COMMENT '封面图';

ALTER TABLE accounts ADD COLUMN id_cert bit(1) NULL COMMENT '是否身份认证' AFTER ip;
ALTER TABLE accounts ADD COLUMN work_cert bit(1) NULL COMMENT '是否工作认证' AFTER id_cert;
ALTER TABLE accounts ADD COLUMN is_org bit(1) NULL COMMENT '是否机构用户' AFTER work_cert;
ALTER TABLE accounts ADD COLUMN is_vip bit(1) NULL COMMENT '是否VIP用户' AFTER is_org;

UPDATE accounts set id_cert=0, work_cert=0, is_org=0,is_vip=0;

ALTER TABLE accounts CHANGE COLUMN id_cert id_cert bit(1) NOT NULL COMMENT '是否身份认证';
ALTER TABLE accounts CHANGE COLUMN work_cert work_cert bit(1) NOT NULL COMMENT '是否工作认证';
ALTER TABLE accounts CHANGE COLUMN is_org is_org bit(1) NOT NULL COMMENT '是否机构用户';
ALTER TABLE accounts CHANGE COLUMN is_vip is_vip bit(1) NOT NULL COMMENT '是否VIP用户';
ALTER TABLE topics DROP COLUMN version_lang;

-- +migrate Down
ALTER TABLE topics DROP COLUMN allow_discuss;
ALTER TABLE topics DROP COLUMN bg;
ALTER TABLE topics CHANGE COLUMN `cover` `cover` varchar(1000)  NOT NULL COMMENT '封面图';
ALTER TABLE accounts DROP COLUMN id_cert;
ALTER TABLE accounts DROP COLUMN work_cert;
ALTER TABLE accounts DROP COLUMN is_org;
ALTER TABLE accounts DROP COLUMN is_vip;
ALTER TABLE articles ADD COLUMN version_lang varchar(50) NOT NULL COMMENT '版本语言' AFTER version_name;
