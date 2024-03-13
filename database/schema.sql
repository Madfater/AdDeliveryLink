-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: DcardAssignment
-- ------------------------------------------------------
-- Server version	8.0.36-0ubuntu0.22.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `advertisement`
--

DROP TABLE IF EXISTS `advertisement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `advertisement_country`
--

DROP TABLE IF EXISTS `advertisement_country`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `advertisement_country` (
  `advertisement_id` int NOT NULL,
  `country_code` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`advertisement_id`,`country_code`),
  KEY `CountryCode` (`country_code`),
  CONSTRAINT `advertisement_country_ibfk_1` FOREIGN KEY (`advertisement_id`) REFERENCES `advertisement` (`ID`),
  CONSTRAINT `advertisement_country_ibfk_2` FOREIGN KEY (`country_code`) REFERENCES `country` (`CountryCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `advertisement_platform`
--

DROP TABLE IF EXISTS `advertisement_platform`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `advertisement_platform` (
  `advertisement_id` int NOT NULL,
  `platform_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`advertisement_id`,`platform_name`),
  KEY `PlatformName` (`platform_name`),
  CONSTRAINT `advertisement_platform_ibfk_1` FOREIGN KEY (`advertisement_id`) REFERENCES `advertisement` (`ID`),
  CONSTRAINT `advertisement_platform_ibfk_2` FOREIGN KEY (`platform_name`) REFERENCES `platform` (`PlatformName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `country`
--

DROP TABLE IF EXISTS `country`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `country` (
  `CountryCode` char(2) NOT NULL,
  PRIMARY KEY (`CountryCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `platform`
--

DROP TABLE IF EXISTS `platform`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `platform` (
  `PlatformName` varchar(50) NOT NULL,
  PRIMARY KEY (`PlatformName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-03-11 20:01:47
