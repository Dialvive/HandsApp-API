package models

import "github.com/dgrijalva/jwt-go"

// SignLanguage represents a Sign Language such as Mexican Sign Language.
type SignLanguage struct {
	ID           uint   `json:"id" gorm:"TINYINT AUTO_INCREMENT; primaryKey"`
	Name         string `json:"name" gorm:"TEXT NOT NULL"`
	Abbreviation string `json:"abbreviation" gorm:"VARCHAR(6) NOT NULL"`
	Modified     string `json:"modified" gorm:"TIMESTAMP"`
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
