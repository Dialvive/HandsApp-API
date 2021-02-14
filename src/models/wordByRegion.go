package models

import "github.com/dgrijalva/jwt-go"

// WordByRegion is a classification of a word by region.
type WordByRegion struct {
	WordID   uint `json:"word_ID" gorm:"column:word_ID;type:INT NOT NULL"`
	RegionID uint `json:"region_ID" gorm:"column:region_ID;type:INT NOT NULL"`
}

// WordByRegionInput is a classification of a word by region.
type WordByRegionInput struct {
	WordID   uint `json:"word_ID" binding:"required"`
	RegionID uint `json:"region_ID" binding:"required"`
}

// WordByRegionClaim is  a claim that cointains WordByRegion as Data.
type WordByRegionClaim struct {
	Data WordByRegion `json:"data" binding:"required"`
	jwt.StandardClaims
}
