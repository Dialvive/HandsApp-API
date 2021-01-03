package models

import (
	//"log"
	//"os"
	//"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
)

// DB is a gorm database
var DB *gorm.DB

// ConnectDatabase Initializes a conection to the desired DB through gorm and migrates the data.
func ConnectDatabase() {

	/*  f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	newLogger := logger.New(
		log.New(f, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Warn, // Log level
			Colorful:      true,        // Disable color
		},
	)
	*/
	dsn := "sofrito:Diego30039!@tcp(localhost:3306)/signa_mundi?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
	})

	if err != nil {
		panic("Failed to connect to DB. " + err.Error())
	}

	// Simple objects
	database.AutoMigrate(
		&Country{},
		&SpokenLanguage{},
		&SignLanguage{},
		&WordCategory{},
		&PhraseCategory{},
		&AdCategory{},
		&Friendship{},
		// Related objects
		&Region{},
		&User{},
		&Locale{},
		&Word{},
		&Phrase{},
		&Advertisement{},
		// Weak objects
		&WordByRegion{},
		&Friend{},
		&FavoriteWords{},
		&FavoritePhrases{},
	)

	DB = database
}
