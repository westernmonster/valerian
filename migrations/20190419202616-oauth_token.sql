
-- +migrate Up
CREATE TABLE `oauth_clients`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `client_id` varchar(254) NOT NULL COMMENT 'Client ID',
    `client_secret` varchar(60) NOT NULL COMMENT 'Client Secret',
    `extra` varchar(512) NOT NULL COMMENT 'Extra',
    `redirect_uri` varchar(512) NULL COMMENT 'Redirect URI',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Clients';

CREATE TABLE `oauth_scopes`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `scope` varchar(200) NOT NULL COMMENT 'Scope',
    `description` varchar(500) NOT NULL COMMENT 'Description',
    `is_default` int(11) NOT NULL COMMENT '是否默认, 0 否，1 是',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Scope';

CREATE TABLE `oauth_roles`  (
    `id` varchar(20) NOT NULL COMMENT 'Role ID',
    `name` varchar(50) NOT NULL COMMENT 'Role Name',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Roles';

CREATE TABLE `oauth_refresh_tokens`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `client_id` varchar(254) NOT NULL COMMENT 'Client ID',
    `account_id`  bigint(20) NOT NULL COMMENT 'Account ID',
    `token` varchar(2000) NOT NULL COMMENT 'Token',
    `expires_at` bigint(20) NOT NULL COMMENT '过期时间',
    `scope` varchar(200) NOT NULL COMMENT 'Scope',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Refresh Tokens';

CREATE TABLE `oauth_access_tokens`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `client_id` varchar(254) NOT NULL COMMENT 'Client ID',
    `account_id`  bigint(20) NOT NULL COMMENT 'Account ID',
    `token` varchar(2000) NOT NULL COMMENT 'Token',
    `expires_at` bigint(20) NOT NULL COMMENT '过期时间',
    `scope` varchar(200) NOT NULL COMMENT 'Scope',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Refresh Tokens';

CREATE TABLE `oauth_authorization_codes`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `client_id` varchar(255) NOT NULL COMMENT 'Client ID',
    `account_id`  bigint(20) NOT NULL COMMENT 'Account ID',
    `code` varchar(40) NOT NULL COMMENT 'Code',
    `redirect_uri` varchar(512) NULL COMMENT 'Redirect URI',
    `expires_at` bigint(20) NOT NULL COMMENT '过期时间',
    `scope` varchar(200) NOT NULL COMMENT 'Scope',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Authorization Codes';

INSERT INTO oauth_clients(id, client_id, client_secret, extra, redirect_uri, deleted, created_at, updated_at)
VALUES(1, '532c28d5412dd75bf975fb951c740a30', '16ed0e1ab220aebf9362045ccad0664f', '','https://api.flywk.com', 0, 1555177901,1555177901);

INSERT INTO oauth_clients(id, client_id, client_secret, extra, redirect_uri, deleted, created_at, updated_at)
VALUES(2, '2567a5ec9705eb7ac2c984033e06189d', '8b17d5515cdc1939d83abe5c00d673ad', '','https://www.flywk.com', 0, 1555177901,1555177901);


-- +migrate Down
DROP TABLE oauth_clients;
DROP TABLE oauth_scopes;
DROP TABLE oauth_roles;
DROP TABLE oauth_refresh_tokens;
DROP TABLE oauth_access_tokens;
DROP TABLE oauth_authorization_codes;
