package models

import "github.com/dgrijalva/jwt-go"

// FavoritePhrases is a classification of a phrase by a user.
type FavoritePhrases struct {
	PhraseID uint   `json:"phrase_ID" gorm:"column:phrase_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	UserID   uint   `json:"user_ID" gorm:"column:user_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	Modified string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateFavoritePhrasesInput is a classification of a phrase by a user.
type CreateFavoritePhrasesInput struct {
	PhraseID uint `json:"phrase_ID" binding:"required"`
	UserID   uint `json:"user_ID" binding:"required"`
}

// FavoritePhrasesClaim is a claim that cointains FavoritePhrases as Data.
type FavoritePhrasesClaim struct {
	Data FavoritePhrases `json:"data" binding:"required"`
	jwt.StandardClaims
}
