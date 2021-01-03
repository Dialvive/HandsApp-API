package models

import "github.com/dgrijalva/jwt-go"

// WordByRegion is a classification of a word by region.
type WordByRegion struct {
	WordID   uint   `json:"word_ID" gorm:"type:INT NOT NULL"`
	RegionID string `json:"region_ID" gorm:"type:INT NOT NULL"`
	Modified string `json:"modified" gorm:"type:TIMESTAMP"`
}

// CreateWordByRegionInput is a classification of a word by region.
type CreateWordByRegionInput struct {
	WordID   uint   `json:"word_ID" binding:"required"`
	RegionID string `json:"region_ID" binding:"required"`
}

// WordByRegionClaim is  a claim that cointains WordByRegion as Data.
type WordByRegionClaim struct {
	Data WordByRegion `json:"data" binding:"required"`
	jwt.StandardClaims
}
