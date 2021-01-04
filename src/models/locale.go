package models

import "github.com/dgrijalva/jwt-go"

// Locale is a mix of location, spoken language, and sign language classification..
type Locale struct {
	ID               uint   `json:"ID" gorm:"column:ID;type:SMALLINT NOT NULL;primaryKey"`
	CountryID        uint   `json:"country_ID" gorm:"column:country_ID;type:TINYINT NOT NULL"`
	SpokenLanguageID uint   `json:"spoken_language_ID" gorm:"column:spoken_language_ID;type:TINYINT NOT NULL"`
	SignLanguageID   uint   `json:"sign_language_ID" gorm:"column:sign_language_ID;type:TINYINT NOT NULL"`
	Modified         string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateLocaleInput is a mix of location, spoken language, and sign language classification..
type CreateLocaleInput struct {
	CountryID        uint `json:"country_ID" binding:"required"`
	SpokenLanguageID uint `json:"spoken_language_ID" binding:"required"`
	SignLanguageID   uint `json:"sign_language_ID" binding:"required"`
}

// LocaleClaim is a claim that cointains Locale as Data.
type LocaleClaim struct {
	Data Locale `json:"data" binding:"required"`
	jwt.StandardClaims
}
