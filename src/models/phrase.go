package models

import "github.com/dgrijalva/jwt-go"

// Phrase represents a coherent spoken language set of words.
type Phrase struct {
	ID               uint   `json:"ID" gorm:"column:ID;type:INT NOT NULL;primaryKey"`
	LocaleID         uint   `json:"locale_ID" gorm:"column:locale_ID;type:TINYINT NOT NULL"`
	PhraseCategoryID uint   `json:"phrase_category_ID" gorm:"column:phrase_category_ID;type:TINYINT NOT NULL"`
	Text             string `json:"name" gorm:"column:text;type:TEXT NOT NULL"`
	Context          string `json:"context" gorm:"column:context;type:TEXT"`
	Modified         string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreatePhraseInput represents a coherent spoken language set of words.
type CreatePhraseInput struct {
	LocaleID         uint   `json:"locale_ID" binding:"required"`
	PhraseCategoryID uint   `json:"phrase_category_ID" binding:"required"`
	Text             string `json:"name" binding:"required"`
	Context          string `json:"context" binding:"required"`
}

// PhraseClaim is  a claim that cointains Phrase as Data.
type PhraseClaim struct {
	Data Phrase `json:"data" binding:"required"`
	jwt.StandardClaims
}
