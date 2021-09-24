CREATE TABLE `approved_domain` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `domain` varchar(256) NOT NULL,
  `approved` tinyint(1) DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `modified_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_email` (`user_id`,`domain`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;