package models

import "github.com/dgrijalva/jwt-go"

// FavoriteWords is a classification of a word by a user.
type FavoriteWords struct {
	WordID   uint   `json:"word_ID" gorm:"column:word_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	UserID   uint   `json:"user_ID" gorm:"column:user_ID;type:INT NOT NULL;primaryKey;autoIncrement:false"`
	Modified string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateFavoriteWordsInput is a classification of a word by a user.
type CreateFavoriteWordsInput struct {
	WordID uint `json:"word_ID" binding:"required"`
	UserID uint `json:"user_ID" binding:"required"`
}

// FavoriteWordsClaim is a claim that cointains FavoriteWords as Data.
type FavoriteWordsClaim struct {
	Data FavoriteWords `json:"data" binding:"required"`
	jwt.StandardClaims
}
