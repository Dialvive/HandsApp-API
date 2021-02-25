package models

import "github.com/dgrijalva/jwt-go"

// WordSign is the sign of a word.
type WordSign struct {
	WordID   uint   `json:"word_ID" gorm:"column:word_ID;type:INT NOT NULL;autoIncrement:false"`
	LocaleID uint   `json:"locale_ID" gorm:"column:locale_ID;type:SMALLINT NOT NULL;autoIncrement:false"`
	Version  string `json:"version" gorm:"column:version;type:CHAR NOT NULL;"`
	RegionID uint   `json:"region_ID" gorm:"column:region_ID;type:INT;autoIncrement:false"`
	Modified string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// WordSignInput is the sign of a word.
type WordSignInput struct {
	WordID   uint   `json:"word_ID" binding:"required"`
	LocaleID uint   `json:"locale_ID" binding:"required"`
	Version  string `json:"version"`
	RegionID uint   `json:"region_ID"`
	Modified string `json:"modified"`
}

// WordSignClaim is a claim that cointains WordSign as Data.
type WordSignClaim struct {
	Data FavoriteWord `json:"data" binding:"required"`
	jwt.StandardClaims
}
