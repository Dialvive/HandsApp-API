package models

import (
	"github.com/jinzhu/gorm"
)

// DB is a gorm database
var DB *gorm.DB

// ConnectDataBase Initializes a conection to the desired DB through gorm and migrates the data.
func ConnectDatabase() {
	database, err := gorm.Open("mysql", "signaMundi.db")

	if err != nil {
		panic("Failed to connect to DB")
	}

	database.AutoMigrate(
		&SignLanguage{},
		&SpokenLanguage{},
		&Country{},
		&Region{},
		&WordCategory{},
		&PhraseCategory{},
		&Word{},
		&Phrase{})

	DB = database
}
