# Create TestUser
CREATE USER IF NOT EXISTS 'url'@'localhost' IDENTIFIED BY 'password';
GRANT SELECT,INSERT,UPDATE,DELETE,CREATE,DROP ON *.* to 'url'@'localhost';

# Create DB
CREATE DATABASE IF NOT EXISTS `urlshortener` DEFAULT CHARACTER SET utf8;
USE `urlshortener`;

# Ceate table
CREATE TABLE IF NOT EXISTS `mappings` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `original_url` varchar(300) NOT NULL,
  `shortened_url` varchar(100) NOT NULL,
  `single_use` tinyint(1) NOT NULL,
  `expired` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8;
