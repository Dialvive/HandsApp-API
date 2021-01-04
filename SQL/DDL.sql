CREATE SCHEMA IF NOT EXISTS `signa_mundi` CHARACTER SET UTF8MB4;

CREATE TABLE IF NOT EXISTS `signa_mundi`.`country`(
	`ID` TINYINT NOT NULL,
    `name` TEXT NOT NULL,
    `abbreviation` VARCHAR(4) NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`region`(
	`ID` INT NOT NULL,
    `name` TEXT NOT NULL,
    `country_ID` TINYINT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`user`(
	`ID` INT NOT NULL,
    `first_name` TEXT,
    `last_name` TEXT,
    `user_name` VARCHAR(32) NOT NULL,
    `mail` VARCHAR(254) NOT NULL,
    `password` TEXT NOT NULL,
    `biography` VARCHAR(142),
    `mailing` VARCHAR(5) DEFAULT '000',
    `privilege` VARCHAR(5) NOT NULL,
    `points` INT DEFAULT 0,
    `credits` INT DEFAULT 0,
    `region_ID` INT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`friendship`(
	`ID` TINYINT NOT NULL,
    `name` TEXT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`friend`(
    `user1_ID` INT NOT NULL,
    `user2_ID` INT NOT NULL,
    `friendship_ID` TINYINT NOT NULL,
    `facebook` BOOLEAN DEFAULT 0,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`ad_category`(
	`ID` TINYINT NOT NULL,
    `name` TEXT NOT NULL,
    `cost` INT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`advertisement`(
	`ID` INT NOT NULL,
    `user_ID` INT NOT NULL,
    `region_ID` INT NOT NULL,
    `ad_category_ID` TINYINT NOT NULL,
    `title` VARCHAR(64) NOT NULL,
    `body` TEXT,
    `media` BOOLEAN DEFAULT 0,
    `paid` INT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`spoken_language`(
	`ID` TINYINT NOT NULL,
    `name` TEXT NOT NULL,
    `abbreviation` VARCHAR(4) NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`sign_language`(
	`ID` TINYINT NOT NULL,
    `name` TEXT NOT NULL,
    `abbreviation` VARCHAR(8) NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`locale`(
	`ID` SMALLINT NOT NULL,
    `country_ID` TINYINT NOT NULL,
    `spoken_language_ID` TINYINT NOT NULL,
    `sign_language_ID` TINYINT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`phrase_category`(
	`ID` TINYINT NOT NULL,
    `name` TEXT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`word_category`(
	`ID` TINYINT NOT NULL,
    `name` TEXT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`phrase`(
	`ID` INT NOT NULL,
    `locale_ID` SMALLINT NOT NULL,
    `phrase_category_ID` TINYINT NOT NULL,
    `text` TEXT NOT NULL,
    `context` TEXT,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`word`(
	`ID` INT NOT NULL,
    `locale_ID` SMALLINT NOT NULL,
    `word_category_ID` TINYINT NOT NULL,
    `text` TEXT NOT NULL,
    `context` TEXT,
    `definition` TEXT,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`words_by_regions`(
    `word_ID` INT NOT NULL,
    `region_ID` INT NOT NULL
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`favorite_phrases`(
    `phrase_ID` INT NOT NULL,
    `user_ID` INT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `signa_mundi`.`favorite_words`(
    `word_ID` INT NOT NULL,
    `user_ID` INT NOT NULL,
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `signa_mundi`.`country` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`region` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`user` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`friendship` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`ad_category` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`advertisement` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`spoken_language` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`sign_language` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`locale` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`phrase_category` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`word_category` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`phrase` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`word` ADD PRIMARY KEY(`ID`);
ALTER TABLE `signa_mundi`.`friend` ADD PRIMARY KEY(`user1_ID`,`user2_ID`);
ALTER TABLE `signa_mundi`.`words_by_regions` ADD PRIMARY KEY(`word_ID`,`region_ID`);
ALTER TABLE `signa_mundi`.`favorite_phrases` ADD PRIMARY KEY(`phrase_ID`,`user_ID`);
ALTER TABLE `signa_mundi`.`favorite_words` ADD PRIMARY KEY(`word_ID`,`user_ID`);

ALTER TABLE `signa_mundi`.`country` MODIFY `ID` TINYINT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`region` MODIFY `ID` INT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`user` MODIFY `ID` INT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`friendship` MODIFY `ID` TINYINT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`ad_category` MODIFY `ID` TINYINT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`advertisement` MODIFY `ID` INT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`spoken_language` MODIFY `ID` TINYINT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`sign_language` MODIFY `ID` TINYINT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`locale` MODIFY `ID` SMALLINT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`phrase_category` MODIFY `ID` TINYINT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`word_category` MODIFY `ID` TINYINT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`phrase` MODIFY `ID` INT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`word` MODIFY `ID` INT NOT NULL AUTO_INCREMENT;
ALTER TABLE `signa_mundi`.`country` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`region` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`user` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`friendship` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`ad_category` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`advertisement` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`spoken_language` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`sign_language` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`locale` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`phrase_category` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`word_category` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`phrase` AUTO_INCREMENT=1;
ALTER TABLE `signa_mundi`.`word` AUTO_INCREMENT=1;

ALTER TABLE `signa_mundi`.`region` ADD FOREIGN KEY (`country_ID`) REFERENCES `signa_mundi`.`country`(`ID`);
ALTER TABLE `signa_mundi`.`user` ADD FOREIGN KEY (`region_ID`) REFERENCES `signa_mundi`.`region`(`ID`);
ALTER TABLE `signa_mundi`.`friend` ADD FOREIGN KEY (`user1_ID`) REFERENCES `signa_mundi`.`user`(`ID`);
ALTER TABLE `signa_mundi`.`friend` ADD FOREIGN KEY (`user2_ID`) REFERENCES `signa_mundi`.`user`(`ID`);
ALTER TABLE `signa_mundi`.`friend` ADD FOREIGN KEY (`friendship_ID`) REFERENCES `signa_mundi`.`friendship`(`ID`);
ALTER TABLE `signa_mundi`.`advertisement` ADD FOREIGN KEY (`user_ID`) REFERENCES `signa_mundi`.`user`(`ID`);
ALTER TABLE `signa_mundi`.`advertisement` ADD FOREIGN KEY (`region_ID`) REFERENCES `signa_mundi`.`region`(`ID`);
ALTER TABLE `signa_mundi`.`advertisement` ADD FOREIGN KEY (`ad_category_ID`) REFERENCES `signa_mundi`.`ad_category`(`ID`);
ALTER TABLE `signa_mundi`.`locale` ADD FOREIGN KEY (`country_ID`) REFERENCES `signa_mundi`.`country`(`ID`);
ALTER TABLE `signa_mundi`.`locale` ADD FOREIGN KEY (`spoken_language_ID`) REFERENCES `signa_mundi`.`spoken_language`(`ID`);
ALTER TABLE `signa_mundi`.`locale` ADD FOREIGN KEY (`sign_language_ID`) REFERENCES `signa_mundi`.`sign_language`(`ID`);
ALTER TABLE `signa_mundi`.`phrase` ADD FOREIGN KEY (`locale_ID`) REFERENCES `signa_mundi`.`locale`(`ID`);
ALTER TABLE `signa_mundi`.`phrase` ADD FOREIGN KEY (`phrase_category_ID`) REFERENCES `signa_mundi`.`phrase_category`(`ID`);
ALTER TABLE `signa_mundi`.`word` ADD FOREIGN KEY (`locale_ID`) REFERENCES `signa_mundi`.`locale`(`ID`);
ALTER TABLE `signa_mundi`.`word` ADD FOREIGN KEY (`word_category_ID`) REFERENCES `signa_mundi`.`word_category`(`ID`);
ALTER TABLE `signa_mundi`.`words_by_regions` ADD FOREIGN KEY (`word_ID`) REFERENCES `signa_mundi`.`word`(`ID`);
ALTER TABLE `signa_mundi`.`words_by_regions` ADD FOREIGN KEY (`region_ID`) REFERENCES `signa_mundi`.`region`(`ID`);
ALTER TABLE `signa_mundi`.`favorite_phrases` ADD FOREIGN KEY (`phrase_ID`) REFERENCES `signa_mundi`.`phrase`(`ID`);
ALTER TABLE `signa_mundi`.`favorite_phrases` ADD FOREIGN KEY (`user_ID`) REFERENCES `signa_mundi`.`user`(`ID`);
ALTER TABLE `signa_mundi`.`favorite_words` ADD FOREIGN KEY (`word_ID`) REFERENCES `signa_mundi`.`word`(`ID`);
ALTER TABLE `signa_mundi`.`favorite_words` ADD FOREIGN KEY (`user_ID`) REFERENCES `signa_mundi`.`user`(`ID`);
