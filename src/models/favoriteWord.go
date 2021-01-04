package models

import "github.com/dgrijalva/jwt-go"

// FavoriteWord is a classification of a word by a user.
type FavoriteWord struct {
	WordID   uint   `json:"word_ID" gorm:"column:word_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	UserID   uint   `json:"user_ID" gorm:"column:user_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	Modified string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateFavoriteWordInput is a classification of a word by a user.
type CreateFavoriteWordInput struct {
	WordID uint `json:"word_ID" binding:"required"`
	UserID uint `json:"user_ID" binding:"required"`
}

// FavoriteWordClaim is a claim that cointains FavoriteWords as Data.
type FavoriteWordClaim struct {
	Data FavoriteWord `json:"data" binding:"required"`
	jwt.StandardClaims
}
