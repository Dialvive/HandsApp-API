package models

import (
	"github.com/dgrijalva/jwt-go"
)

// WordCategory is a category in which a set of words fall into.
type WordCategory struct {
	ID       uint   `json:"ID" gorm:"type:INT AUTO_INCREMENT;primaryKey"`
	Name     string `json:"name" gorm:"type:TEXT NOT NULL"`
	Modified string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateWordCategoryInput is a category in which a set of words fall into.
type CreateWordCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

// WordCategoryClaim is  a claim that cointains WordCategory as Data.
type WordCategoryClaim struct {
	Data WordCategory `json:"data" binding:"required"`
	jwt.StandardClaims
}
