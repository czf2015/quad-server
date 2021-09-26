DROP TABLE `user`;

CREATE TABLE `user` (
  `id` char(36) NOT NULL,
  `name` varchar(256) NOT NULL,
  `role_name` varchar(256) NOT NULL,
  `email` varchar(256) NOT NULL,
  `authority_level` int DEFAULT 0,
  `valid` int DEFAULT 0,
  `password` varchar(256) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `modified_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;