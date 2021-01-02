package models

import "github.com/dgrijalva/jwt-go"

// FavoritePhrase is a classification of a phrase by a user.
type FavoritePhrase struct {
	PhraseID uint   `json:"phrase_ID" gorm:"INT NOT NULL;primaryKey;autoIncrement:false"`
	UserID   string `json:"user_ID" gorm:"INT NOT NULL;primaryKey;autoIncrement:false"`
	Modified string `json:"modified" gorm:"TIMESTAMP"`
}

// CreateFavoritePhraseInput is a classification of a phrase by a user.
type CreateFavoritePhraseInput struct {
	PhraseID uint   `json:"phrase_ID" binding:"required"`
	UserID   string `json:"user_ID" binding:"required"`
}

// FavoritePhraseClaim is a claim that cointains FavoritePhrase as Data.
type FavoritePhraseClaim struct {
	Data FavoritePhrase `json:"data" binding:"required"`
	jwt.StandardClaims
}
