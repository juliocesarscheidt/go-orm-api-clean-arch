create schema go_orm_api;

drop schema go_orm_api;

use go_orm_api;

show tables;

CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`),
  KEY `idx_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

select * from users;

select COUNT(id) as counter from users;

truncate users;

show create table users;
