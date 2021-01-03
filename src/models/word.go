package models

import "github.com/dgrijalva/jwt-go"

// Word represents a coherent spoken language set of words.
type Word struct {
	ID             uint   `json:"id" gorm:"type:INT; primaryKey"`
	LocaleID       uint   `json:"locale_ID" gorm:"type:TINYINT NOT NULL"`
	WordCategoryID uint   `json:"word_category_ID" gorm:"type:TINYINT NOT NULL"`
	Text           string `json:"name" gorm:"type:TEXT NOT NULL"`
	Context        string `json:"context" gorm:"type:TEXT"`
	Definition     string `json:"definition" gorm:"type:TEXT"`
	Modified       string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateWordInput represents a coherent spoken word.
type CreateWordInput struct {
	LocaleID       uint   `json:"locale_ID" binding:"required"`
	WordCategoryID uint   `json:"phrase_category_ID" binding:"required"`
	Text           string `json:"name" binding:"required"`
	Definition     string `json:"definition" binding:"required"`
	Context        string `json:"context" binding:"required"`
}

// WordClaim is  a claim that cointains Word as Data.
type WordClaim struct {
	Data Word `json:"data" binding:"required"`
	jwt.StandardClaims
}
