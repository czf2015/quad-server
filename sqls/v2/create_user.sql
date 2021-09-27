CREATE TABLE `user` (
  `id` char(36) NOT NULL,
  `name` varchar(256) NOT NULL,
  `role_name` varchar(256) NOT NULL,
  `email` varchar(256) NOT NULL,
  `authority_level` int(11) DEFAULT 0,
  `valid` int(11) DEFAULT 0,
  `password` varchar(256) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`name`),
  UNIQUE (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;