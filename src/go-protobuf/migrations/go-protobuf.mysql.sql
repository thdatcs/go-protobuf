CREATE DATABASE IF NOT EXISTS `temp` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `temp`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    id BIGINT AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP,
    updated_by INT,
    username VARCHAR(100),
    password VARCHAR(100),
    fullname VARCHAR(100),
    PRIMARY KEY (id),
    INDEX (username)
)