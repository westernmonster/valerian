
-- +migrate Up
CREATE TABLE `accounts`  (
  `id` bigint(0) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` bigint(255) NOT NULL,
  `updated_at` bigint(255) NOT NULL,
  PRIMARY KEY (`id`)
);
-- +migrate Down
DROP TABLE accounts;
