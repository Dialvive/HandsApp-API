package models

import "github.com/dgrijalva/jwt-go"

// PhraseCategory is a category in which a set of phrases fall into.
type PhraseCategory struct {
	ID       uint   `json:"ID" gorm:"type:TINYINT NOT NULL;primaryKey"`
	Name     string `json:"name" gorm:"type:TEXT NOT NULL"`
	Modified string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreatePhraseCategoryInput is a category in which a set of phrases fall into.
type CreatePhraseCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

// PhraseCategoryClaim is  a claim that cointains PhraseCategory as Data.
type PhraseCategoryClaim struct {
	Data PhraseCategory `json:"data" binding:"required"`
	jwt.StandardClaims
}
