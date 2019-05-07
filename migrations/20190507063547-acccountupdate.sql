
-- +migrate Up
ALTER TABLE accounts CHANGE COLUMN `deleted` `deleted` bit(1) NOT NULL COMMENT '是否删除';
ALTER TABLE accounts ADD COLUMN user_name varchar(50) NULL COMMENT '用户名' AFTER mobile;
UPDATE accounts SET user_name = cast(id as char);
-- +migrate Down
ALTER TABLE accounts CHANGE COLUMN `deleted` `deleted` int(11) NOT NULL COMMENT '是否删除';
ALTER TABLE accounts DROP COLUMN `user_name`;
