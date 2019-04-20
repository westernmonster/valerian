
-- +migrate Up
CREATE TABLE `auth_clients`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `client_id` varchar(255) NOT NULL COMMENT 'Client ID',
    `client_secret` varchar(255) NOT NULL COMMENT 'Client Secret',
    `extra` varchar(512) NOT NULL COMMENT 'Extra',
    `redirect_uri` varchar(512) NOT NULL COMMENT 'Redirect URI',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Clients';

CREATE TABLE `auth_authorize`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `client_id` varchar(255) NOT NULL COMMENT 'Client ID',
    `code` varchar(255) NOT NULL COMMENT 'Code ',
    `expired_in` bigint(20) NOT NULL COMMENT '过期时间',
    `scope` varchar(255) NOT NULL COMMENT 'Scope',
    `redirect_uri` varchar(512) NOT NULL COMMENT 'Redirect URI',
    `state` varchar(255) NOT NULL COMMENT 'State',
    `extra` varchar(512) NOT NULL COMMENT 'Extra',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Authorize';

CREATE TABLE `auth_access`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `client_id` varchar(255) NOT NULL COMMENT 'Client ID',
    `authorize` varchar(255) NOT NULL COMMENT 'Authorize',
    `previous` varchar(255) NOT NULL COMMENT 'Previous',
    `access_token` varchar(255) NOT NULL COMMENT 'Access Token',
    `refresh_token` varchar(255) NOT NULL COMMENT 'Refresh Token',
    `expired_in` bigint(20) NOT NULL COMMENT '过期时间',
    `scope` varchar(255) NOT NULL COMMENT 'Scope',
    `redirect_uri` varchar(512) NOT NULL COMMENT 'Redirect URI',
    `extra` varchar(512) NOT NULL COMMENT 'Extra',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Access';

CREATE TABLE `auth_refresh`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `token` varchar(512) NOT NULL COMMENT 'Token',
    `access` varchar(512) NOT NULL COMMENT 'Access',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) COMMENT 'OAUTH2 Fresh';

CREATE TABLE `auth_expires`  (
    `id` bigint(20) NOT NULL COMMENT 'ID',
    `token` varchar(512) NOT NULL COMMENT 'Token',
    `expires_at` bigint(20) NOT NULL COMMENT '过期时间',
    `deleted` int(11) NOT NULL COMMENT '是否删除',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX expires_index (expires_at),
    INDEX token_expires_index (token)
) COMMENT 'OAUTH2 Expires';



-- +migrate Down
DROP TABLE auth_clients;
DROP TABLE auth_authorize;
DROP TABLE auth_access;
DROP TABLE auth_refresh;
DROP TABLE auth_expires;
