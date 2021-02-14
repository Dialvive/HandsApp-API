package models

import "github.com/dgrijalva/jwt-go"

// FavoritePhrase is a classification of a phrase by a user.
type FavoritePhrase struct {
	PhraseID uint   `json:"phrase_ID" gorm:"column:phrase_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	UserID   uint   `json:"user_ID" gorm:"column:user_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	Modified string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// FavoritePhraseInput is a classification of a phrase by a user.
type FavoritePhraseInput struct {
	PhraseID uint `json:"phrase_ID" binding:"required"`
	UserID   uint `json:"user_ID" binding:"required"`
}

// FavoritePhraseClaim is a claim that cointains FavoritePhrases as Data.
type FavoritePhraseClaim struct {
	Data FavoritePhrase `json:"data" binding:"required"`
	jwt.StandardClaims
}
