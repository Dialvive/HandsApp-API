CREATE SCHEMA IF NOT EXISTS `signa_mundi` CHARACTER SET utf8;

CREATE TABLE IF NOT EXISTS `signa_mundi`.`word_category`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`word`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `text` TEXT NOT NULL,
    `media` TEXT,
    `definition` TEXT,
    `context` TEXT,
    `word_category_ID` TINYINT NOT NULL,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`word_category_ID`) REFERENCES `signa_mundi`.`word_category`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`phrase_category`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`phrase`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `text` TEXT NOT NULL,
    `media` TEXT,
    `context` TEXT,
    `phrase_category_ID` TINYINT NOT NULL,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`phrase_category_ID`) REFERENCES `signa_mundi`.`phrase_category`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`spoken_language`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `abbreviation` VARCHAR(3) NOT NULL,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`sign_language`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `abbreviation` VARCHAR(3) NOT NULL,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`country`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `abbreviation` VARCHAR(3) NOT NULL,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`region`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `country_ID` TINYINT NOT NULL,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`country_ID`) REFERENCES `signa_mundi`.`country`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`user`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `first_name` TEXT NOT NULL,
    `last_name` TEXT,
    `user_name` VARCHAR(36) NOT NULL,
    `mail` VARCHAR(254) NOT NULL,
    `password` TEXT NOT NULL,
    `biography` TEXT,
    `profile_picture` TEXT,
    `mailing` VARCHAR(3),
    `type` VARCHAR(3) NOT NULL,
    `privilege` VARCHAR(2),
    `country_ID` TINYINT,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`word_category_ID`) REFERENCES `signa_mundi`.`word_category`(`ID`)
); 

CREATE TABLE IF NOT EXISTS `signa_mundi`.`words_by_regions`(
	`word_ID` TINYINT NOT NULL,
    `region_ID` TINYINT NOT NULL,
    FOREIGN KEY (`word_ID`) REFERENCES `signa_mundi`.`word`(`ID`)
    FOREIGN KEY (`region_ID`) REFERENCES `signa_mundi`.`region`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`phrases_by_sign_languages`(
	`phrase_ID` TINYINT NOT NULL,
    `sign_language_ID` TINYINT NOT NULL,
    FOREIGN KEY (`phrase_ID`) REFERENCES `signa_mundi`.`phrase`(`ID`)
    FOREIGN KEY (`sign_language_ID`) REFERENCES `signa_mundi`.`sign_language`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`spoken_languages_by_countries`(
	`spoken_language_ID` TINYINT NOT NULL,
    `country_ID` TINYINT NOT NULL,
    FOREIGN KEY (`spoken_language_ID`) REFERENCES `signa_mundi`.`spoken_language`(`ID`)
    FOREIGN KEY (`country_ID`) REFERENCES `signa_mundi`.`country`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`sign_languages_by_countries`(
	`sign_language_ID` TINYINT NOT NULL,
    `country_ID` TINYINT NOT NULL,
    FOREIGN KEY (`sign_language_ID`) REFERENCES `signa_mundi`.`sign_language`(`ID`)
    FOREIGN KEY (`country_ID`) REFERENCES `signa_mundi`.`country`(`ID`)
);