DROP TABLE IF EXISTS `advertisement`;

CREATE TABLE `advertisement` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) NOT NULL,
  `StartAt` datetime NOT NULL,
  `EndAt` datetime NOT NULL,
  `AgeStart` int DEFAULT NULL,
  `AgeEnd` int DEFAULT NULL,
  `Gender` enum('M','F') DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `idx_advertisement_id` (`ID`),
  KEY `idx_advertisement_ageStart` (`AgeStart`),
  KEY `idx_advertisement_ageEnd` (`AgeEnd`),
  KEY `idx_advertisement_gender` (`Gender`),
  CONSTRAINT `advertisement_chk_1` CHECK (((`AgeStart` >= 1) and (`AgeStart` <= 100))),
  CONSTRAINT `advertisement_chk_2` CHECK (((`AgeEnd` >= 1) and (`AgeEnd` <= 100))),
  CONSTRAINT `chk_gender` CHECK ((`Gender` in (_utf8mb4'M',_utf8mb4'F',NULL)))
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `country`;

CREATE TABLE `country`
(
    `CountryCode` char(2) NOT NULL,
    PRIMARY KEY (`CountryCode`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `platform`;

CREATE TABLE `platform`
(
    `PlatformName` varchar(50) NOT NULL,
    PRIMARY KEY (`PlatformName`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `advertisement_country`;

CREATE TABLE `advertisement_country` (
  `advertisement_id` int NOT NULL,
  `country_code` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`advertisement_id`,`country_code`),
  KEY `CountryCode` (`country_code`),
  CONSTRAINT `advertisement_country_ibfk_1` FOREIGN KEY (`advertisement_id`) REFERENCES `advertisement` (`ID`),
  CONSTRAINT `advertisement_country_ibfk_2` FOREIGN KEY (`country_code`) REFERENCES `country` (`CountryCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `advertisement_platform`;

CREATE TABLE `advertisement_platform` (
  `advertisement_id` int NOT NULL,
  `platform_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`advertisement_id`,`platform_name`),
  KEY `PlatformName` (`platform_name`),
  CONSTRAINT `advertisement_platform_ibfk_1` FOREIGN KEY (`advertisement_id`) REFERENCES `advertisement` (`ID`),
  CONSTRAINT `advertisement_platform_ibfk_2` FOREIGN KEY (`platform_name`) REFERENCES `platform` (`PlatformName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;