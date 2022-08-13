CREATE TABLE `article`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`    datetime(3) DEFAULT NULL,
    `updated_at`    datetime(3) DEFAULT NULL,
    `deleted_at`    datetime(3) DEFAULT NULL,
    `title`         varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `cid`           bigint unsigned NOT NULL,
    `desc`          varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci          DEFAULT NULL,
    `content`       longtext CHARACTER SET utf8 COLLATE utf8_general_ci,
    `img`           varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci          DEFAULT NULL,
    `comment_count` bigint                                                  NOT NULL DEFAULT '0',
    `read_count`    bigint                                                  NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    KEY             `idx_article_deleted_at` (`deleted_at`) USING BTREE,
    KEY             `fk_article_category` (`cid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=574 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


CREATE TABLE `category`
(
    `id`   bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;