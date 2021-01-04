package models

import "github.com/dgrijalva/jwt-go"

// SignLanguage represents a Sign Language such as Mexican Sign Language.
type SignLanguage struct {
	ID           uint   `json:"ID" gorm:"column:ID;type:TINYINT NOT NULL;primaryKey"`
	Name         string `json:"name" gorm:"column:name;type:TEXT NOT NULL"`
	Abbreviation string `json:"abbreviation" gorm:"column:abbreviation;type:VARCHAR(8) NOT NULL"`
	Modified     string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateSignLanguageInput type with automatic ID.
type CreateSignLanguageInput struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}

// SignLanguageClaim is  a claim that cointains SignLanguage as Data.
type SignLanguageClaim struct {
	Data SignLanguage `json:"data" binding:"required"`
	jwt.StandardClaims
}
