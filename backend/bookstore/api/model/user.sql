CREATE TABLE `user` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255)  NOT NULL DEFAULT '' COMMENT '用户名称',
    `password` VARCHAR(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `age` TINYINT(3) NOT NULL DEFAULT 0 COMMENT '年龄',
    `gender` CHAR(5)  NOT NULL COMMENT '男｜女｜未公开',
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name_unique` (`name`)
) ENGINE=INNODB  DEFAULT CHARSET=utf8mb4 ;