package models

import "github.com/dgrijalva/jwt-go"

// PhraseSign is the sign of a phrase.
type PhraseSign struct {
	PhraseID uint   `json:"phrase_ID" gorm:"column:phrase_ID;type:INT NOT NULL;autoIncrement:false"`
	LocaleID uint   `json:"locale_ID" gorm:"column:locale_ID;type:SMALLINT NOT NULL;autoIncrement:false"`
	RegionID uint   `json:"region_ID" gorm:"column:region_ID;type:INT;autoIncrement:false"`
	Modified string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// PhraseSignInput is the sign of a phrase.
type PhraseSignInput struct {
	PhraseID uint   `json:"phrase_ID" binding:"required"`
	LocaleID uint   `json:"locale_ID" binding:"required"`
	RegionID uint   `json:"region_ID"`
	Modified string `json:"modified"`
}

// PhraseSignClaim is a claim that cointains PhraseSign as Data.
type PhraseSignClaim struct {
	Data FavoritePhrase `json:"data" binding:"required"`
	jwt.StandardClaims
}
