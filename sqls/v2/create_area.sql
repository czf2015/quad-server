CREATE TABLE `area` (
  `id` char(36) NOT NULL,
  `pid` char(36) NOT NULL,
  `title` varchar(256) NOT NULL,
  `code` varchar(256) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;