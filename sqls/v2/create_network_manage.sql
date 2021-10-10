CREATE TABLE `network_manage` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `subnet` varchar(256) NOT NULL COMMENT '子网',
  `subnet_type` varchar(36) NOT NULL COMMENT '子网类型',
  `organization` varchar(256) NOT NULL COMMENT '组织机构',
  `usage` float DEFAULT '0' COMMENT '使用率',
  `distributed` float DEFAULT '0' COMMENT '分配率',
  `create_method` varchar(36) NOT NULL COMMENT '创建方式',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;