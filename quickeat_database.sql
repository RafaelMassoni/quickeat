-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.5.8-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=1 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
SET foreign_key_checks = 1;

-- Dumping database structure for quickeat
CREATE DATABASE IF NOT EXISTS `quickeat` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `quickeat`;

-- Dumping structure for table quickeat.categorias
CREATE TABLE IF NOT EXISTS `categorias` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table quickeat.categorias: ~0 rows (approximately)
/*!40000 ALTER TABLE `categorias` DISABLE KEYS */;
/*!40000 ALTER TABLE `categorias` ENABLE KEYS */;

-- Dumping structure for table quickeat.pratos
CREATE TABLE IF NOT EXISTS `pratos` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `id_categoria` INT(11) unsigned DEFAULT NULL,
  `nome` char(50) NOT NULL DEFAULT '0',
  `preco` decimal(10,2) unsigned NOT NULL DEFAULT 0.00,
  `tempo_de_preparo` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT 'MINUTOS',
  PRIMARY KEY (`id`),
  KEY `id_categoria` (`id_categoria`),
  CONSTRAINT `id_categoria` FOREIGN KEY (`id_categoria`) REFERENCES `categorias` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- Dumping data for table quickeat.pratos: ~0 rows (approximately)
/*!40000 ALTER TABLE `pratos` DISABLE KEYS */;
INSERT INTO `pratos` (`id`, `id_categoria`, `nome`, `preco`, `tempo_de_preparo`) VALUES
	(1, NULL, 'tilápia', 45.00, 30),
	(2, NULL, 'batata frita', 15.00, 15),
	(3, NULL, 'batata frita com queijo', 18.00, 16),
	(4, NULL, 'iscas de frango acebolada', 20.00, 20),
	(5, NULL, 'filé parmegiana', 40.00, 40);
/*!40000 ALTER TABLE `pratos` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
