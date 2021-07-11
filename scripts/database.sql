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
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` int unsigned NOT NULL DEFAULT 0,
    `deleted_at` int unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;