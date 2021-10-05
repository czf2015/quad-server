CREATE TABLE `address_plan` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `network_address` varchar(256) NOT NULL COMMENT '地址位宽',
  `subnet_type` varchar(256) NOT NULL COMMENT '子网类型',
  `organization` varchar(256) NOT NULL COMMENT '地址位宽',
  `address_list` varchar(256) NOT NULL COMMENT '地址位宽',
  `bit_width` int(11) DEFAULT '0' COMMENT '地址位宽',
  `prefix_bit_width` int(11) DEFAULT '40' COMMENT '前缀位数',
  `subnet_address_begin_value` int(11) DEFAULT '0' COMMENT '子网地址起始值',
  `address_count` int(11) DEFAULT '0' COMMENT '地址个数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;