DROP TABLE IF EXISTS `advertisement`;
CREATE TABLE `advertisement` (
                                 `id`           INT          NOT NULL AUTO_INCREMENT,
                                 `title`        VARCHAR(255) NOT NULL,
                                 `start_at`     DATETIME     NOT NULL,
                                 `end_at`       DATETIME     NOT NULL,
                                 `age_start` INT NOT NULL,
                                 `age_end`   INT NOT NULL,
                                 `gender`    ENUM ('F', 'M', 'B') DEFAULT 'B',
                                 `status`       BOOLEAN      NOT NULL,
                                 `created_date` TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
                                 `updated_date` TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 KEY `idx_advertisements` (gender, status, age_start, age_end, start_at, end_at)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `country`;
CREATE TABLE `country`
(
    `country_code` CHAR(2) NOT NULL,
    PRIMARY KEY (`country_code`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `platform`;
CREATE TABLE `platform`
(
    `platform_name` VARCHAR(50) NOT NULL,
    PRIMARY KEY (`platform_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `advertisement_country`;
CREATE TABLE `advertisement_country` (
                                         `advertisement_id` INT     NOT NULL,
                                         `country_code`     CHAR(2) NOT NULL,
                                         PRIMARY KEY (`advertisement_id`, `country_code`),
                                         KEY `idx_ad_country` (`country_code`, `advertisement_id`),
                                         CONSTRAINT `advertisement_country_fk_advertisement` FOREIGN KEY (`advertisement_id`) REFERENCES `advertisement` (`id`) ON DELETE CASCADE,
                                         CONSTRAINT `advertisement_country_fk_country` FOREIGN KEY (`country_code`) REFERENCES `country` (`country_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `advertisement_platform`;
CREATE TABLE `advertisement_platform` (
                                          `advertisement_id` INT         NOT NULL,
                                          `platform_name`    VARCHAR(50) NOT NULL,
                                          PRIMARY KEY (`advertisement_id`, `platform_name`),
                                          KEY `idx_ad_platform` (`platform_name`, `advertisement_id`),
                                          CONSTRAINT `advertisement_platform_fk_advertisement` FOREIGN KEY (`advertisement_id`) REFERENCES `advertisement` (`id`) ON DELETE CASCADE,
                                          CONSTRAINT `advertisement_platform_fk_platform` FOREIGN KEY (`platform_name`) REFERENCES `platform` (`platform_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `platform` (`platform_name`)
VALUES ('web'),     -- Web platform
       ('android'), -- Android platform
       ('ios'); -- IOS platform

INSERT INTO `country` (`country_code`)
VALUES ('US'), -- United States
       ('CA'), -- Canada
       ('TW'), -- Taiwan
       ('JP'); -- Japan