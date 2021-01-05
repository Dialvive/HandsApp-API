package models

import "github.com/dgrijalva/jwt-go"

// Word represents a coherent spoken language set of words.
type Word struct {
	ID             uint   `json:"ID" gorm:"column:ID;type:INT NOT NULL;primaryKey"`
	LocaleID       uint   `json:"locale_ID" gorm:"column:locale_ID;type:TINYINT NOT NULL"`
	WordCategoryID uint   `json:"word_category_ID" gorm:"column:word_category_ID;type:TINYINT NOT NULL"`
	Text           string `json:"text" gorm:"column:text;type:TEXT NOT NULL"`
	Context        string `json:"context" gorm:"column:context;type:TEXT"`
	Definition     string `json:"definition" gorm:"column:definition;type:TEXT"`
	Modified       string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateWordInput represents a coherent spoken word.
type CreateWordInput struct {
	LocaleID       uint   `json:"locale_ID" binding:"required"`
	WordCategoryID uint   `json:"word_category_ID" binding:"required"`
	Text           string `json:"text" binding:"required"`
	Definition     string `json:"definition" binding:"required"`
	Context        string `json:"context" binding:"required"`
}

// WordClaim is  a claim that cointains Word as Data.
type WordClaim struct {
	Data Word `json:"data" binding:"required"`
	jwt.StandardClaims
}
