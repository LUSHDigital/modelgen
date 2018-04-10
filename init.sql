CREATE DATABASE `modelgen_tests`;

USE `modelgen_tests`;

DROP TABLE IF EXISTS `common_cases`;

CREATE TABLE `common_cases` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `not_null_string` varchar(255) NOT NULL DEFAULT '',
  `not_null_int` int(11) NOT NULL,
  `null_string` int(11) DEFAULT NULL,
  `null_int` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `complex_cases`;

CREATE TABLE `complex_cases` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `raw_json` json NOT NULL,
  `size_enum` enum('X-SMALL','SMALL','MEDIUM','LARGE','X-LARGE') DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;