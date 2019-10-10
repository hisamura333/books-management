# mysql -u root < init.sql
DROP DATABASE IF EXISTS `books_management`;
CREATE DATABASE `books_management` CHARACTER SET utf8mb4;

use books_management;

CREATE TABLE `books_management`.`test` (
    `id` int ,
    `name` varchar(255)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

