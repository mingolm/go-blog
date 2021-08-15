SET NAMES utf8mb4;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username`   varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
    `password`   char(40) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
    `status`     tinyint(3) NOT NULL DEFAULT '0',
    `ip`         varbinary(16) NOT NULL DEFAULT '',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` int unsigned NOT NULL DEFAULT 0,
    `deleted_at` int unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `articles`;

CREATE TABLE `articles`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int(10) unsigned NOT NULL,
    `type`       tinyint(3) NOT NULL DEFAULT '0',
    `title`      varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `content`    text COLLATE utf8mb4_unicode_ci NOT NULL,
    `status`     tinyint(3) NOT NULL DEFAULT '0',
    `ip`         varbinary(16) NOT NULL DEFAULT '',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `articles` (`user_id`,`type`,`title`,`content`,`status`,`ip`,`created_at`)
VALUES (1,1,'我是标题1','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题2','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题3','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题4','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题5','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题6','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题7','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题8','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题9','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题10','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题11','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题12','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题13','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题14','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题15','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题16','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题17','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题18','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题19','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题20','我是内容1','1','','2021-08-15 16:24:22'),
(1,1,'我是标题21','我是内容1','1','','2021-08-15 16:24:22');