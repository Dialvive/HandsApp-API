package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB is a gorm database
var DB *gorm.DB

// ConnectDatabase Initializes a conection to the desired DB through gorm and migrates the data.
func ConnectDatabase() {

	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	newLogger := logger.New(
		log.New(f, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)

	dsn := "sofrito:Diego30039!@tcp(localhost:3306)/signa_mundi?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic("Failed to connect to DB. " + err.Error())
	}

	// Simple objects
	database.AutoMigrate(
		&AdCategory{},
		&Country{},
		&Friendship{},
		&PhraseCategory{},
		&SignLanguage{},
		&SpokenLanguage{},
		&WordCategory{},
		// Related objects
		&Advertisement{},
		&Region{},
		&Locale{},
		&Phrase{},
		&User{},
		&Word{},
		// Weak objects
		&Friend{},
		&FavoritePhrase{},
		&FavoriteWord{},
	)

	DB = database
}
