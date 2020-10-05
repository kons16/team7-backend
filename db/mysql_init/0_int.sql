CREATE DATABASE IF NOT EXISTS `team7-db`;
USE `team7-db`;

CREATE TABLE IF NOT EXISTS users(
	`id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
	`name` VARCHAR(255) NOT NULL,
	`email` VARCHAR(255) NOT NULL,
	`password_hash` VARCHAR(255) NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS user_sessions(
	`id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    `user_id` INT NOT NULL,
	`token` VARCHAR(255) NULL,
    `expires_at` TIMESTAMP NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS clothes(
	`id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
	`name` VARCHAR(255) NOT NULL,
	`price` BIGINT NOT NULL,
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
