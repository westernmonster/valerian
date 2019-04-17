
-- +migrate Up
CREATE TABLE `id_certifications`  (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `account_id` bigint(20) NOT NULL COMMENT '账户ID',
  `status` int(11) NOT NULL COMMENT '状态：-1 未认证, 0 认证中,  1 认证成功, 2 认证失败',
  `audit_conclusions` varchar(250) NULL COMMENT '失败原因',
  `name` varchar(20) NULL COMMENT '姓名',
  `identification_number` varchar(50) NULL COMMENT '证件号',
  `id_card_type` varchar(50) NULL COMMENT '证件类型, identityCard代表身份证',
  `id_card_start_date` varchar(50) NULL COMMENT '证件有效期起始日期',
  `id_card_expiry` varchar(50) NULL COMMENT '证件有效期截止日期',
  `address` varchar(500) NULL COMMENT '地址',
  `sex` varchar(50) NULL COMMENT '性别',
  `id_card_front_pic` varchar(500) NULL COMMENT '证件照正面图片',
  `id_card_back_pic` varchar(500) NULL COMMENT '证件照背面图片',
  `face_pic` varchar(500) NULL COMMENT '认证过程中拍摄的人像正面照图片',
  `ethnic_group` varchar(50) NULL COMMENT '证件上的民族',
  `deleted` int(11) NOT NULL COMMENT '是否删除',
  `created_at` bigint(20) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT '身份认证';

-- +migrate Down
DROP TABLE id_certifications;
