package models

import "github.com/dgrijalva/jwt-go"

// SpokenLanguage represents a language such as English, Spanish, etc.
type SpokenLanguage struct {
	ID           uint   `json:"ID" gorm:"column:ID;type:TINYINT NOT NULL; primaryKey"`
	Name         string `json:"name" gorm:"column:name;type:TEXT NOT NULL"`
	Abbreviation string `json:"abbreviation" gorm:"column:abbreviation;type:VARCHAR(4) NOT NULL"`
	Modified     string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateSpokenLanguageInput type with automatic ID.
type CreateSpokenLanguageInput struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

// SpokenLanguageClaim is  a claim that cointains SpokenLanguage as Data.
type SpokenLanguageClaim struct {
	Data SpokenLanguage `json:"data" binding:"required"`
	jwt.StandardClaims
}
