package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is a gorm database
var DB *gorm.DB

// ConnectDatabase Initializes a conection to the desired DB through gorm and migrates the data.
func ConnectDatabase() {
	dsn := "sofrito:Diego30039!@tcp(localhost:3306)/signa_mundi?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	database.AutoMigrate(
		&Country{},
		&SpokenLanguage{},
		&SignLanguage{},
	)

	DB = database
}
