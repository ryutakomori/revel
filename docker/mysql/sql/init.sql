CREATE DATABASE IF NOT EXISTS test CHARACTER SET utf8mb4;

---- drop ----
DROP TABLE IF EXISTS `test.test_table`;

---- create ----
create table IF not exists `test.test_table`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;