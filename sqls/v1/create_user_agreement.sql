CREATE TABLE `user_agreement` (
  `user_id` char(36) NOT NULL,
  `agreement_id` char(36) NOT NULL,
  `agreed_at` timestamp NOT NULL,
  UNIQUE KEY `unique_user_agreement` (`user_id`, `agreement_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;