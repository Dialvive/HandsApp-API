CREATE SCHEMA IF NOT EXISTS `signa_mundi` CHARACTER SET UTF8MB4;

CREATE TABLE IF NOT EXISTS `signa_mundi`.`country`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `abbreviation` VARCHAR(2) NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`region`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `country_ID` TINYINT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`country_ID`) REFERENCES `signa_mundi`.`country`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`user`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `first_name` TEXT,
    `last_name` TEXT,
    `user_name` VARCHAR(32) NOT NULL,
    `mail` VARCHAR(254) NOT NULL,
    `password` TEXT NOT NULL,
    `biography` VARCHAR(140),
    `mailing` VARCHAR(3) DEFAULT '000',
    `privilege` VARCHAR(3) NOT NULL,
    `points` INT DEFAULT 0,
    `credits` INT DEFAULT 0,
    `region_ID` INT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`region_ID`) REFERENCES `signa_mundi`.`region`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`friendship`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`friend`(
    `user1_ID` INT NOT NULL,
    `user2_ID` INT NOT NULL,
    `friendship_ID` TINYINT NOT NULL,
    `facebook` BOOLEAN DEFAULT 0,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`user1_ID`) REFERENCES `signa_mundi`.`user`(`ID`),
    FOREIGN KEY (`user2_ID`) REFERENCES `signa_mundi`.`user`(`ID`),
    FOREIGN KEY (`friendship`) REFERENCES `signa_mundi`.`friendship`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`ad_category`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `cost` INT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`advertisement`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `user_ID` INT NOT NULL,
    `region_ID` INT NOT NULL,
    `ad_category_ID` TINYINT NOT NULL,
    `title` VARCHAR(64) NOT NULL,
    `body` TEXT,
    `media` BOOLEAN DEFAULT 0,
    `paid` INT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`user_ID`) REFERENCES `signa_mundi`.`user`(`ID`),
    FOREIGN KEY (`region_ID`) REFERENCES `signa_mundi`.`region`(`ID`),
    FOREIGN KEY (`ad_category_ID`) REFERENCES `signa_mundi`.`ad_category`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`spoken_language`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `abbreviation` VARCHAR(2) NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`sign_language`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `abbreviation` VARCHAR(2) NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`locale`(
	`ID` SMALLINT AUTO_INCREMENT NOT NULL,
    `country_ID` TINYINT NOT NULL,
    `spoken_language_ID` TINYINT NOT NULL,
    `sign_language_ID` TINYINT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`country_ID`) REFERENCES `signa_mundi`.`country`(`ID`),
    FOREIGN KEY (`spoken_language_ID`) REFERENCES `signa_mundi`.`spoken_language`(`ID`),
    FOREIGN KEY (`sign_language_ID`) REFERENCES `signa_mundi`.`sign_language`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`phrase_media`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `url` TEXT,
    `description` TEXT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`word_media`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `url` TEXT,
    `description` TEXT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`phrase_category`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`word_category`(
	`ID` TINYINT AUTO_INCREMENT NOT NULL,
    `name` TEXT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`phrase`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `locale_ID` SMALLINT NOT NULL,
    `media_ID` INT NOT NULL,
    `phrase_category_ID` TINYINT NOT NULL,
    `text` TEXT NOT NULL,
    `context` TEXT,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`locale_ID`) REFERENCES `signa_mundi`.`locale`(`ID`),
    FOREIGN KEY (`media_ID`) REFERENCES `signa_mundi`.`phrase_media`(`ID`),
    FOREIGN KEY (`phrase_category_ID`) REFERENCES `signa_mundi`.`phrase_category`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`word`(
	`ID` INT AUTO_INCREMENT NOT NULL,
    `locale_ID` SMALLINT NOT NULL,
    `media_ID` INT NOT NULL,
    `word_category_ID` TINYINT NOT NULL,
    `text` TEXT NOT NULL,
    `context` TEXT,
    `definition` TEXT,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`locale_ID`) REFERENCES `signa_mundi`.`locale`(`ID`),
    FOREIGN KEY (`media_ID`) REFERENCES `signa_mundi`.`word_media`(`ID`),
    FOREIGN KEY (`word_category_ID`) REFERENCES `signa_mundi`.`word_category`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`words_by_regions`(
    `word_ID` INT NOT NULL,
    `region_ID` INT NOT NULL,
    FOREIGN KEY (`word_ID`) REFERENCES `signa_mundi`.`word`(`ID`),
    FOREIGN KEY (`region_ID`) REFERENCES `signa_mundi`.`region`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`favorite_phrases`(
    `phrase_ID` INT NOT NULL,
    `user_ID` INT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`phrase_ID`) REFERENCES `signa_mundi`.`phrase`(`ID`),
    FOREIGN KEY (`user_ID`) REFERENCES `signa_mundi`.`user`(`ID`)
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`favorite_words`(
    `word_ID` INT NOT NULL,
    `user_ID` INT NOT NULL,
    `creation` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`word_ID`) REFERENCES `signa_mundi`.`word`(`ID`),
    FOREIGN KEY (`user_ID`) REFERENCES `signa_mundi`.`user`(`ID`)
);

